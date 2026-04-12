package main

import (
	"encoding/json"
	"net/http"
)

// Note structure
type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// In-memory storage
var notes = []Note{}
var idCounter = 1

// Handler function
func handler(w http.ResponseWriter, r *http.Request) {

	// 👉 POST (create note)
	if r.Method == "POST" {
		var note Note
		json.NewDecoder(r.Body).Decode(&note)

		note.ID = idCounter
		idCounter++

		notes = append(notes, note)
		json.NewEncoder(w).Encode(note)
		return
	}

	// 👉 DELETE
	if r.Method == "DELETE" {
		notes = []Note{}
		w.Write([]byte("Deleted all notes"))
		return
	}

	// 👉 FORCE HTML for browser
	w.Header().Set("Content-Type", "text/html")

	html := `
	<html>
	<head>
		<title>Notes API</title>
		<style>
			body { font-family: Arial; background: #111; color: #fff; padding: 20px; }
			h1 { color: #00ffcc; }
			li { margin: 10px 0; }
		</style>
	</head>
	<body>
		<h1>🚀 Notes API Running</h1>
		<ul>
	`

	for _, note := range notes {
		html += "<li>📝 " + note.Text + "</li>"
	}

	html += `
		</ul>
	</body>
	</html>
	`

	w.Write([]byte(html))
}

// Main function
func main() {
	http.HandleFunc("/notes", handler)

	println("Server running on http://localhost:8080/notes 🚀")

	http.ListenAndServe(":8080", nil)
}