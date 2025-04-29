package main

import (
	"html/template"
	"net/http"
	"path"
)

func LoadhtmlPage(w http.ResponseWriter, file string) *template.Template {
	fileToLoad := "HTML/" + file
	templ, err := template.ParseFiles(fileToLoad)
	if err != nil {
		http.Error(w, "unable to load the HTML page", http.StatusInternalServerError)
		return nil
	}
	return templ
}

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	file := path.Base(r.URL.Path)
	templ := LoadhtmlPage(w, file)
	templ.Execute(w, nil)

}
