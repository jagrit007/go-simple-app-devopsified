package main

import (
	"html/template"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Unable to get hostname", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("static/home.html")
	if err != nil {
		http.Error(w, "Unable to parse template", http.StatusInternalServerError)
		return
	}

	data := struct {
		Hostname string
	}{
		Hostname: hostname,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
