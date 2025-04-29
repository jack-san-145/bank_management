package main

import (
	"net/http"
	"path"
)

func LoadAsserts(w http.ResponseWriter, r *http.Request) {
	image := path.Base(r.URL.Path)
	im := "asserts/" + image
	http.ServeFile(w, r, im)
}
