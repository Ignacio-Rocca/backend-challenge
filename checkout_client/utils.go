package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func isOkRequest(err error) {
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
}

func isOkResponse(resp *http.Response, err error) {
	if resp.StatusCode > 299 {
		log.Println("Status code error: " + resp.Status )
		os.Exit(1)
	}

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println(fmt.Sprintf("OK: %v", resp.StatusCode))
}

func parseResponse(resp *http.Response, model interface{}) {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(data, model)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func buildBody(model interface{}) io.Reader {
	bodyBytes, err := json.Marshal(model)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	return strings.NewReader(string(bodyBytes))
}