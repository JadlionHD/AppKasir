package resources

import (
	"embed"

	"github.com/gin-contrib/static"
)

//go:embed all:website/build
var assets embed.FS

func Assets() static.ServeFileSystem {
	// return fs.Sub(assets, "website/build")
	return static.EmbedFolder(assets, "website/build")
}
