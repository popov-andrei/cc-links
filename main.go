package main

import (
	"fmt"
	"my/transport/rest"
	"net/http"
)

func main() {
	// Делаем локальное хранилище в виде мапы
	var storage = make(map[string]string)

	// Тестовое заполнение мапы
	storage["42f706038c113da3"] = "https://vk.com/coding21"
	storage["86ce2ea1c7eec5be"] = "https://github.com/popov-andrei"
	storage["d4dbe9f1ccdea6f5"] = "https://www.apple.com/iphone-14-pro"

	// Задаем обработчик для HTTP-запросов, путь "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rest.HandlePostRequest(w, r, &storage)
	})

	// Запускаем сервер
	fmt.Println("Server started")
	http.ListenAndServe(":8080", nil)
}
