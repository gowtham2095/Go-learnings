package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	counterCookie, err := r.Cookie("counter")
	if err == http.ErrNoCookie {
		counterCookie = &http.Cookie{
			Name:  "counter",
			Value: "0",
		}
	}
	count, _ := strconv.Atoi(counterCookie.Value)
	count++
	counterCookie.Value = strconv.Itoa(count)
	http.SetCookie(w, counterCookie)
	io.WriteString(w, counterCookie.Value)
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hari")
}

func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
