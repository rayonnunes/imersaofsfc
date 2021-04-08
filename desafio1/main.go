package main

import (
	"fmt"
	"net/http"
	"log"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
	}

	if r.Method != "GET" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
	}


	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./pages"))
	http.Handle("/", fileServer)
	// http.HandleFunc("/", homeHandler)

	fmt.Printf("Starting server at port 8000\n")
	err := http.ListenAndServe(":8000", nil);
	if err != nil {
		log.Fatal(err)
	}
}