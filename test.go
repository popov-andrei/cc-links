package main

import (
	"fmt"
	"my/storage"
)

func main() {
	link := storage.Link{URL: []byte("a.io")}
	fmt.Println(string(link.ShortURL))

	link.Hash()
	fmt.Println(string(link.ShortURL))
}
