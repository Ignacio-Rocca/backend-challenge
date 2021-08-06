package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	fmt.Println(req.URL.String())
}
