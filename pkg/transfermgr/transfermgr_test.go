package transfermgr

import (
	"mime/multipart"
	"testing"

	"github.com/google/uuid"
)

func TestInitManager(t *testing.T) {
	store := InitTransferStore()
	if store == nil {
		t.Fatal("The returned store after initialization is nil")
	}

	someId, err := store.newUuid()
	if err != nil {
		t.Fatalf("There was an error when generating an id in the store: %s", err.Error())
	}
	if someId == uuid.Nil.String() {
		t.Fatal("The generated uuid contains only empty bytes")
	}
}

func TestNewStringTransfer(t *testing.T) {
	store := InitTransferStore()

	transfer, err := store.NewStringTransfer(12, 21, "some secret")
	if err != nil {
		t.Fatalf("There was an error creating a new string transfer: %s", err.Error())
	}
	if transfer == nil {
		t.Fatalf("The new transfer is nil")
	}

	if transfer.Id == uuid.Nil.String() {
		t.Fatal("The generated uuid contains only empty bytes")
	}
	if transfer.From != 12 || transfer.To != 21 || transfer.TransferedString != "some secret" || transfer.TransferedFile != nil {
		t.Fatalf("The contents of the new transfer is not coherent with the submitted values")
	}
}
func TestNewFileTransfer(t *testing.T) {
	store := InitTransferStore()

	sampleFile := &multipart.FileHeader{}
	transfer, err := store.NewFileTransfer(12, 21, sampleFile)
	if err != nil {
		t.Fatalf("There was an error creating a new string transfer: %s", err.Error())
	}
	if transfer == nil {
		t.Fatalf("The new transfer is nil")
	}

	if transfer.Id == uuid.Nil.String() {
		t.Fatal("The generated uuid contains only empty bytes")
	}
	if transfer.From != 12 || transfer.To != 21 || transfer.TransferedString != "" || transfer.TransferedFile != sampleFile {
		t.Fatalf("The contents of the new transfer is not coherent with the submitted values")
	}
}
