package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := http.Get("http://127.0.0.1:9999/")
		log.Println(resp)
		resp.Body.Close()
		w.Write([]byte("hello"))

	})
	http.ListenAndServe(":9999", nil)
}
