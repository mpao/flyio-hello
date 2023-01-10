package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		out := []byte("hello fly.io\n")
		_, _ = w.Write(out)
	})

	fmt.Println("starting app...")
	e := http.ListenAndServe(":8080", nil)
	if e != nil {
		log.Fatalln(e)
	}

}
