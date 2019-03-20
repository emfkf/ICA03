package main

import (
	"encoding/json"
	"net/http"
)

var info = message{"Fredrik Meltveit", "fredm18@uia.no"}

type message struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func giveInfo(w http.ResponseWriter, r *http.Request) {
	// m := message{"Fredrik", "fredm18@uia.no"}
	b, err := json.Marshal(info)

	if err != nil {
		panic(err)
	}

	w.Write([]byte(b))
}

func main() {
	http.HandleFunc("/", giveInfo)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
