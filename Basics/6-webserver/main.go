package main

import (
	"fmt"
	"log"
	"net/http"
)

//handler is kind of a path to function mapping
func main() {
	http.HandleFunc("/", index_handler) //handler  
	log.Fatal(http.ListenAndServe(":8000",nil))
}

func index_handler(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprintf(w,"whoa, Go is neat!")  //format based on writer 
}

