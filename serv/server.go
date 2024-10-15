package serv

import (
	asciiart "asciiart/src"
	"html/template"
	"net/http"
	"strings"
)

// global variable containing the template
var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	// if method not "GET", will return Only GET method is allowed
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// execute the template with nil data
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AsciiWeb(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	// if method not "POST", will return Only POST method is allowed
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	//get the text from the request
	textInput := asciiart.CheckInput(r.FormValue("text"))
	textLines := strings.Split(textInput, "\r\n")
	//get the banner from the request
	banner := r.FormValue("banner")
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "Invalid banner", http.StatusBadRequest)
		return
	}
	maps, err := asciiart.MapBanner(banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// draw the ascii art and execute the template with the result
	asciiArt := asciiart.Draw(maps, textLines)
	err = tmpl.Execute(w, asciiArt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
