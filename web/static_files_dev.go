//go:build !prod
// +build !prod

package web

import "io/fs"

// Dummy, empty filesystems (in dev mode, the static files are handled by vue-cli)
var RootFiles fs.FS
var StaticFiles fs.FS
