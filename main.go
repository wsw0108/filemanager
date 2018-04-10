package main

import (
	_ "github.com/filebrowser/filebrowser/caddy/filemanager"
	_ "github.com/hacdias/caddy-webdav"
	"github.com/mholt/caddy/caddy/caddymain"
)

func main() {
	caddymain.Run()
}
