package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ответ сервера:", string(body))
}
