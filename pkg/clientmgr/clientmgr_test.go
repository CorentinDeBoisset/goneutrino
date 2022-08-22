package clientmgr

import (
	"strings"
	"testing"
	"time"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/google/uuid"
)

const testingPgpKey = `
-----BEGIN PGP PUBLIC KEY BLOCK-----

mDMEYwTU8xYJKwYBBAHaRw8BAQdAkf98L/SB3EwcKZhl6yyfuk3vpHdarTX/Ze+x
jSYmGcK0FlRlc3QgPHRlc3RAc2FuY2FyZS5mcj6ImQQTFgoAQRYhBM/FiXXZdkkz
cEHiyuvFTsrXmCfbBQJjBNTzAhsDBQkDwmcABQsJCAcCAiICBhUKCQgLAgQWAgMB
Ah4HAheAAAoJEOvFTsrXmCfb+TIA/RpMO1+scSl9kE738+Bsq5sdbiLR21rPc55h
7hHNP+RRAQDGcNt3d92ld7HKAlIn+aHwaxUdSikzQIqbQ8xuYmRLCbg4BGME1PMS
CisGAQQBl1UBBQEBB0BfWxXSLfbp4WXRwXmrSWBIvhs0v3LqC1lIPO+J7N/0HwMB
CAeIfgQYFgoAJhYhBM/FiXXZdkkzcEHiyuvFTsrXmCfbBQJjBNTzAhsMBQkDwmcA
AAoJEOvFTsrXmCfbDysA/1IJEzxJhqkCBYpxZ6L/YSTBygezRu1aWtr7QIR0dE+B
APwPaeB7Vj0amuytg0N0G1gcx3L1BDH0zTScav+DidJADQ==
=ImSo
-----END PGP PUBLIC KEY BLOCK-----
`

func getPgpKey(t *testing.T) *openpgp.Entity {
	entityList, err := openpgp.ReadArmoredKeyRing(strings.NewReader(testingPgpKey))
	if err != nil {
		t.Fatalf("failed to load the PGP key: %s", err.Error())
		return nil
	}
	if len(entityList) == 0 {
		t.Fatalf("no PGP key was found in the keyring: %s", err.Error())
		return nil
	}

	return entityList[0]
}

func TestInit(t *testing.T) {
	store := InitStore()
	if store == nil {
		t.Fatal("After initialization, the store should not be nil")
	}
}

func TestNewClient(t *testing.T) {
	store := InitStore()
	clientId, err := store.NewClient("superclient", getPgpKey(t), time.Now().Add(10*time.Minute))
	if err != nil {
		t.Fatalf("There was an error creating a new client: %s", err.Error())
	}
	if clientId == uuid.Nil.String() {
		t.Fatalf("The generated client id is only made of null bytes")
	}

	// Check the client
	client, err := store.GetClient(clientId)
	if err != nil {
		t.Fatalf("failed to retrieve the client after it was registerd: %s", err.Error())
	}
	if client.Name != "superclient" {
		t.Fatalf("the new client does not have the supplied name: (expected: %s, received: %s)", "superclient", client.Name)
	}

	_, err = store.GetClient("some-invalid-id")
	if err == nil {
		t.Fatal("no error were thrown when trying to find a client with an invalid id")
	}
}

func TestExpiredClient(t *testing.T) {
	store := InitStore()
	clientId, err := store.NewClient("superclient", getPgpKey(t), time.Now().Add(-10*time.Minute))
	if err != nil {
		t.Fatalf("There was an error creating a new client: %s", err.Error())
	}
	if clientId == uuid.Nil.String() {
		t.Fatal("The generated client id is only made of null bytes")
	}

	store.CleanupExpired()

	client, _ := store.GetClient(clientId)
	if client != nil {
		t.Fatal("the expired client was not cleaned up")
	}
}
