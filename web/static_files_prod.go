//go:build prod
// +build prod

package web

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed frontend-app/index.html
//go:embed frontend-app/favicon_plain.png
var rootFiles embed.FS
var RootFiles fs.FS

//go:embed frontend-app/static/css
//go:embed frontend-app/static/js
//go:embed frontend-app/static/img
var staticFiles embed.FS
var StaticFiles fs.FS

func init() {
	var err error
	if RootFiles, err = fs.Sub(rootFiles, "frontend-app"); err != nil {
		panic(fmt.Sprintf("An error occured when processing the static files: %v", err))
	}

	if StaticFiles, err = fs.Sub(staticFiles, "frontend-app/static"); err != nil {
		panic(fmt.Sprintf("An error occured when processing the static files: %v", err))
	}
}
