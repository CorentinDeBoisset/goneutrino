package keymgr

import (
	"fmt"
)

type KeyStore map[int]string

// TODO: store the public keys in a key-value store, such as Redis

func InitKeyStore() KeyStore {
	keyStore := make(KeyStore)

	return keyStore
}

func (m *KeyStore) StoreKey(id int, pKey string) error {
	if _, ok := (*m)[id]; ok {
		return fmt.Errorf("there is already a public key for the given identifier (%d)", id)
	}

	(*m)[id] = pKey

	return nil
}

func (m *KeyStore) GetKey(id int) (string, error) {
	if _, ok := (*m)[id]; !ok {
		return "", fmt.Errorf("there are no public key for the given identifier (%d)", id)
	}
	return (*m)[id], nil
}
