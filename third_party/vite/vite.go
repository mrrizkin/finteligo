package vite

import (
	goVite "github.com/mrrizkin/go-vite-parser"
)

var vite = goVite.Parse(goVite.Config{
	OutDir:       "/build/",
	ManifestPath: "public/build/manifest.json",
	HotFilePath:  "public/hot",
})

func Entry(entries ...string) string {
	if vite.IsDev() {
		return vite.RenderDevEntriesTag(entries...)
	}

	return vite.RenderEntriesTag(entries...)
}

func ReactRefresh() string {
	if vite.IsDev() {
		return vite.RenderReactRefreshTag()
	}

	return ""
}
