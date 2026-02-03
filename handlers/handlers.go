package handlers

import (
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

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Get input values
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// Validate input
	if text == "" || banner == "" {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Check if text contains only printable ASCII characters (32-126)
	for _, char := range text {
		if char < 32 || char > 126 {
			if char != '\n' && char != '\r' {
				errorHandler(w, r, http.StatusBadRequest)
				return
			}
		}
	}

	// Generate ASCII art
	asciiArt, err := asciiart.Generate(text, banner)
	if err != nil {
		if strings.Contains(err.Error(), "banner file not found") {
			errorHandler(w, r, http.StatusNotFound)
		} else {
			errorHandler(w, r, http.StatusInternalServerError)
		}
		return
	}

	// Parse template
	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Prepare data for template
	data := struct {
		Input    string
		Banner   string
		AsciiArt string
	}{
		Input:    text,
		Banner:   banner,
		AsciiArt: asciiArt,
	}

	// Execute template
	err = tmpl.Execute(w, data)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		// If error template doesn't exist, send basic error message
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