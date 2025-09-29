package main

import (
	"log"
	"net/http"
)

func SetupRoutes() {

}

func main() {

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
