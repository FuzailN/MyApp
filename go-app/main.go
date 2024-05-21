package main

import (
    "html/template"
    "net/http"
    "github.com/gorilla/mux"
)

var templates = template.Must(template.ParseFiles("templates/layout.html", "templates/home.html", "templates/about.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := templates.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "layout.html", map[string]interface{}{
        "Title": "Home",
        "ContentTemplate": "home.html",
    })
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "layout.html", map[string]interface{}{
        "Title": "About",
        "ContentTemplate": "about.html",
    })
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler)
    r.HandleFunc("/about", aboutHandler)
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

    http.ListenAndServe(":8080", r)
}

