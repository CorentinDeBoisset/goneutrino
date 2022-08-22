package transfermgr

import (
	"fmt"
	"mime/multipart"
	"sync"

	"github.com/google/uuid"
)

type ContentTransfer struct {
	Id   string
	From int
	To   int

	TransferedFile   *multipart.FileHeader
	TransferedString string
}

type TransferStore struct {
	store map[string]*ContentTransfer

	mtx sync.Mutex
}

func InitTransferStore() *TransferStore {
	return &TransferStore{
		store: make(map[string]*ContentTransfer),
	}
}

func (t *TransferStore) newUuid() (string, error) {
	for i := 0; i < 50; i++ {
		candidate, err := uuid.NewRandom()
		if err != nil {
			return "", fmt.Errorf("there was an error generating an id for the file transfer: %w", err)
		}
		if _, ok := t.store[candidate.String()]; !ok {
			return candidate.String(), nil
		}
	}
	return "", fmt.Errorf("failed to find a new unique id for the file transfer after 50 tries")
}

func (t *TransferStore) NewFileTransfer(from, to int, file *multipart.FileHeader) (*ContentTransfer, error) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	transferId, err := t.newUuid()
	if err != nil {
		return nil, fmt.Errorf("there was an error generating an id for the file transfer: %w", err)
	}

	newTransfer := &ContentTransfer{
		Id:             transferId,
		From:           from,
		To:             to,
		TransferedFile: file,
	}
	t.store[transferId] = newTransfer

	return newTransfer, nil
}

func (t *TransferStore) NewStringTransfer(from, to int, content string) (*ContentTransfer, error) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	transferId, err := t.newUuid()
	if err != nil {
		return nil, fmt.Errorf("there was an error generating an id for the file transfer: %w", err)
	}

	newTransfer := &ContentTransfer{
		Id:               transferId,
		From:             from,
		To:               to,
		TransferedString: content,
	}
	t.store[transferId] = newTransfer

	return newTransfer, nil
}

func (t *TransferStore) GetTransfer(id string) (*ContentTransfer, error) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	if transfer, ok := t.store[id]; ok {
		return transfer, nil
	}

	return nil, fmt.Errorf("there is no transfer associated with this id")
}
