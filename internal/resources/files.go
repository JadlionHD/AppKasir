package resources

import (
	"embed"
	"html/template"
	"io/fs"

	"github.com/gin-contrib/static"
)

//go:embed all:website/build
var assets embed.FS

var Tmpl *template.Template
var Assets fs.FS
var StaticAssets static.ServeFileSystem

func GetAssets() {
	asset, _ := fs.Sub(assets, "website/build")
	Assets = asset

	tmpl := template.Must(template.ParseFS(Assets, "*.html"))
	Tmpl = tmpl

	StaticAssets = static.EmbedFolder(assets, "website/build")
}
