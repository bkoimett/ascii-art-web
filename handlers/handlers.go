package handlers

import (
	// "bytes"
	"fmt"
	"test/asciiart"
	"html/template"
	"net/http"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
	}
}


func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	fmt.Printf("Received text: %q\n", text)
	fmt.Printf("Text length: %d\n", len(text))

	if text == "" {
		tmpl, err := template.ParseFiles("templates/result.html")
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
			return
		}
		
		data := struct {
			Input    string
			Banner   string
			AsciiArt string
		}{
			Input:    text,
			Banner:   banner,
			AsciiArt: "",
		}
		
		err = tmpl.Execute(w, data)
		if err != nil {
			errorHandler(w, r, http.StatusInternalServerError)
		}
		return
	}

	if banner == "" {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}

	for i, char := range text {
		if char < 32 || char > 126 {
			if char != '\n' && char != '\r' {
				fmt.Printf("Invalid character at position %d: %q (code: %d)\n", i, char, char)
				errorHandler(w, r, http.StatusBadRequest)
				return
			}
		}
	}

	asciiArt, err := asciiart.Generate(text, banner)
	if err != nil {
		fmt.Printf("Error generating ASCII art: %v\n", err)
		if strings.Contains(err.Error(), "banner file not found") || 
		   strings.Contains(err.Error(), "cannot read banner file") {
			errorHandler(w, r, http.StatusNotFound)
		} else if strings.Contains(err.Error(), "invalid banner") {
			errorHandler(w, r, http.StatusBadRequest)
		} else {
			errorHandler(w, r, http.StatusInternalServerError)
		}
		return
	}

	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

	data := struct {
		Input    string
		Banner   string
		AsciiArt string
	}{
		Input:    text,
		Banner:   banner,
		AsciiArt: asciiArt,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		switch status {
		case http.StatusNotFound:
			http.Error(w, "404 Page Not Found", http.StatusNotFound)
		case http.StatusBadRequest:
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
		case http.StatusInternalServerError:
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		case http.StatusMethodNotAllowed:
			http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		default:
			http.Error(w, "Error", status)
		}
		return
	}

	data := struct{ Status int }{Status: status}
	tmpl.Execute(w, data)
}