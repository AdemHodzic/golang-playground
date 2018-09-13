package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World!\n")
	fmt.Fprintf(w, "Go is amazing!\n")

}

func users(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	json.NewEncoder(w).Encode(articles)

}

func main() {
	fmt.Println("Running server at port 8080")
	http.HandleFunc("/", index)
	http.HandleFunc("/users", users)
	http.ListenAndServe(":8080", nil)

}
