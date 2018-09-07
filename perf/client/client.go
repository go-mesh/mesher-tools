package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(os.Getenv("TARGET"))
		if err != nil {
			w.WriteHeader(500)
		}
		log.Println(resp)
		b := make([]byte, 0)
		_, _ = resp.Body.Read(b)
		resp.Body.Close()
		w.Write(b)

	})
	http.ListenAndServe(":9000", nil)
}
