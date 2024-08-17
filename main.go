package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	Data string `json:"data"`
}

func main() {
	data := Data{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/save-snippet", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &data)
	})

	http.HandleFunc("/load-snippet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, data.Data)
	})

	fmt.Println("Server is running on port 38642")
	http.ListenAndServe("0.0.0.0:38642", nil)
}
