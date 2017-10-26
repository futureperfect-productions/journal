package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"gopkg.in/mgo.v2/bson"
)

type Entry struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Title string        `bson:"title" json:"title"`
}

func AddEntry(w http.ResponseWriter, r *http.Request) {
	session := Connect()
	defer session.Close()

	collection := session.DB("fp-journal").C("entries")

	title := r.FormValue("Title")
	e := Entry{}
	e.Title = title

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&e)

	// Write the entry to mongo
	collection.Insert(e)

	ej, _ := json.Marshal(e)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", ej)
}

func GetEntries(w http.ResponseWriter, r *http.Request) {
	session := Connect()
	defer session.Close()

	collection := session.DB("fp-journal").C("entries")
	var results []Entry

	if err := collection.Find(nil).All(&results); err != nil {
		w.WriteHeader(404)
		return
	}

	e, _ := json.Marshal(results)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", e)
}

func main() {
	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/save", AddEntry)
	http.HandleFunc("/entries", GetEntries)
	http.ListenAndServe(":8080", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")

	tmpl, _ := template.ParseFiles(lp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}
