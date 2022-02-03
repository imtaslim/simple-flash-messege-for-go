package main

import (
	"flash/fls"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)
	http.HandleFunc("/", he)
	fmt.Println("Listening...")
	http.ListenAndServe(":8000", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	fm := "This is a flashed message!"
	flash.SetFlash(w, "message", fm)
	http.Redirect(w, r, "/get", http.StatusTemporaryRedirect)
}

func get(w http.ResponseWriter, r *http.Request) {
	fm, err := flash.GetFlash(w, r, "message")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", fm)
}


func he(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now())
}