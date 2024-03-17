package main

import (
	"fmt"
	"go-todo/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("starting server on the port 9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}
