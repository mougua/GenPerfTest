package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	port := flag.String("p", "8028", "port")
	flag.Parse()
	err := http.ListenAndServe(":"+*port, http.FileServer(http.Dir("./")))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
