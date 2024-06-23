package resources

import (
	"embed"
	"io/fs"
)

//go:embed all:website/build
var assets embed.FS

func Assets() (fs.FS, error) {
	return fs.Sub(assets, "website/build")
}
