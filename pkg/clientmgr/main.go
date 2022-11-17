package clientmgr

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/go-crypto/openpgp/armor"
	"github.com/corentindeboisset/neutrino/pkg/transfermgr"
	"github.com/google/uuid"
)

type ClientInstance struct {
	Id         int
	Name       string
	PublicKey  *openpgp.Entity
	Expiration time.Time
	Online     bool

	NewPeers     chan int
	NewTransfers chan *transfermgr.ContentTransfer
}

// TODO: store the public keys in a key-value store, such as Redis
type ClientStore struct {
	Clients map[string]*ClientInstance
	Mtx     sync.RWMutex
}

func (c *ClientInstance) SerializePubKey() (string, error) {
	b := new(bytes.Buffer)
	w, _ := armor.Encode(b, openpgp.PublicKeyType, nil)
	if err := c.PublicKey.Serialize(w); err != nil {
		// Handle the error
		return "", fmt.Errorf("failed to serialize the public key for the client: %w", err)
	}
	_ = w.Close()

	return b.String(), nil
}

func (c *ClientInstance) IsOnline() bool {
	return c.Online
}

func InitStore() *ClientStore {
	store := &ClientStore{
		Clients: make(map[string]*ClientInstance),
	}

	return store
}

// CleanupExpired removes client instances where client.Exp is before the current time
func (store *ClientStore) CleanupExpired() {
	store.Mtx.Lock()
	defer store.Mtx.Unlock()

	now := time.Now()
	for sessionId, client := range store.Clients {
		if now.After(client.Expiration) {
			delete(store.Clients, sessionId)
		}
	}
}

// NewClient creates a new client instance and stores it in the store
func (store *ClientStore) NewClient(name string, pKey *openpgp.Entity, exp time.Time) (string, error) {
	store.Mtx.Lock()
	defer store.Mtx.Unlock()

	var sessionId string
	for {
		candidate, err := uuid.NewRandom()
		if err != nil {
			return "", fmt.Errorf("failed to create a UUID for the new client: %w", err)
		}
		if _, ok := store.Clients[candidate.String()]; !ok {
			// Stop the loop as soon as there is an unknown UUID
			sessionId = candidate.String()
			break
		}
	}

	// find a new unique id
	existingIds := make(map[int]bool, len(store.Clients))
	for _, client := range store.Clients {
		existingIds[client.Id] = true
	}

	newId := -1
	for tries := 0; tries < 50; tries++ {
		candidate := rand.Intn(8999) + 1000 // random int between 1000 and 9999
		if _, ok := existingIds[candidate]; !ok {
			newId = candidate
			break
		}
	}
	if newId == -1 {
		return "", fmt.Errorf("failed to find a unique numeric id for the client after 50 tries")
	}

	store.Clients[sessionId] = &ClientInstance{
		Id:         newId,
		Name:       name,
		PublicKey:  pKey,
		Expiration: exp,
	}

	return sessionId, nil
}

// RemoveClient deletes the client instance indexed on the given sessionId
func (store *ClientStore) RemoveClient(sessionId string) error {
	store.Mtx.Lock()
	defer store.Mtx.Unlock()

	delete(store.Clients, sessionId)

	return nil
}

// GetClient tries to find a client instance from the given uuid.UUID. If no one can be found, it returns an error
func (store *ClientStore) GetClient(sessionId string) (*ClientInstance, error) {
	store.Mtx.RLock()
	defer store.Mtx.RUnlock()

	client, ok := store.Clients[sessionId]
	if !ok {
		return nil, fmt.Errorf("no client could be found for the identifier \"%s\"", sessionId)
	}

	if time.Now().After(client.Expiration) {
		return nil, fmt.Errorf("the client with the given session identifier (%s) is expired", sessionId)
	}

	return client, nil
}

// GetPublicKey returns a serialized representation of the public key for the user with the given id
func (store *ClientStore) GetClientFromId(id int) (*ClientInstance, error) {
	store.Mtx.RLock()
	defer store.Mtx.RUnlock()

	var foundClient *ClientInstance
	for _, client := range store.Clients {
		if client.Id == id {
			foundClient = client
		}
	}
	if foundClient == nil {
		return nil, fmt.Errorf("no client could be found with id %d", id)
	}

	if time.Now().After(foundClient.Expiration) {
		return nil, fmt.Errorf("the client with the given id (%d) is expired", id)
	}

	return foundClient, nil
}
