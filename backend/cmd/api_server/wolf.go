package main

import (
	"log"
	"net/http"
	"wolf/api"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	router := api.GetRouter()
	http.Handle("/", router)
	log.Println("Server started at http://localhost:8889")
	if err := http.ListenAndServe(":8889", nil); err != nil {
		log.Fatal(err)
	}
}
