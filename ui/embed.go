package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:generate pnpm build

//go:embed all:build/**
var distFS embed.FS

func Handler() http.Handler {
	f, err := fs.Sub(distFS, "build")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(f))
}
