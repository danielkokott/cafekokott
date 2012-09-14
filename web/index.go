//https://developers.google.com/appengine/docs/python/tools/uploadinganapp

package web

import (
    "net/http"
    "html/template"
)

func init() {
    http.HandleFunc("/", workinprogress)
    http.HandleFunc("/test", index)
}

type data struct {
	Title string
}

func workinprogress(w http.ResponseWriter, r *http.Request) {
    ob := &data{Title:"Cafe Kokott"}
    //renderTemplate(w, "index", ob)
    renderTemplate(w, "workinprogress", ob)
}
func index(w http.ResponseWriter, r *http.Request) {
    ob := &data{Title:"Cafe Kokott"}
    //renderTemplate(w, "index", ob)
    renderTemplate(w, "index", ob)
}

func renderTemplate(w http.ResponseWriter, name string, ob *data) {
	t, err := template.ParseFiles("web/" + name + ".html")
	if t == nil {
		http.Error(w, "File "+ name + "not found", http.StatusInternalServerError)
	}
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
    t.Execute(w,ob)
}