package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Приветствую!")
		fmt.Println("Запрос к корневому пути")
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format(time.RFC1123)
		fmt.Fprintf(w, currentTime)
		fmt.Println("Запрос к пути /time")
	})

	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
		fmt.Println("Обработка 404 ошибки")
	})

	http.ListenAndServe(":8080", nil)
}
