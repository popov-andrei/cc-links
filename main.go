package main

import (
	"myapp/server"
	"myapp/storage"
)

func main() {
	s := storage.NewInMemoryStorage()

	srv := server.NewHTTPServer(s)
	srv
}
