package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Dog")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hari")
}

func main() {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
