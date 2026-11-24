package main

import (
	"fmt"
	"log"
	"net/http"
)

var Url string = "https://dinofyi.com/api/v1/"

func main() {
	mux := http.NewServeMux()
	log.Fatal(http.ListenAndServe(":6001", mux))
	fmt.Println()
}

