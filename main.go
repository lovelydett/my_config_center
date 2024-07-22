package main

import (
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	router := initAppRouter()
	http.Handle("/", router)
	log.Println("Server started at http://localhost:8889")
	if err := http.ListenAndServe(":8889", nil); err != nil {
		log.Fatal(err)
	}
}
