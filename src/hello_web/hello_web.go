package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	// fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
	title := "aaa"
	body := "adfjjf"
	fmt.Fprintf(w, "<h1>%s<h1><div>%s</div>", title, body)
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
