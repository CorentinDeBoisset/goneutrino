//go:build !prod
// +build !prod

package web

import (
	"io/fs"
	"os"
)

type dummyFS struct{}

func (f dummyFS) Open(name string) (fs.File, error) {
	return nil, &fs.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

// Dummy,empty filesystems (in dev mode, the static files are handled by vue-cli)
var RootFiles dummyFS
var StaticFiles dummyFS
