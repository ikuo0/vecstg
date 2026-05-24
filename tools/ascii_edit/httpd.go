package main

import (
	"log"
	"net/http"
)

func main() {
	addr := ":8000"
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Printf("listening on http://localhost%s/ascii_edit.html", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

/*
cd tools/ascii_edit
go run httpd.go
*/
