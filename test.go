package main

import (
	"cc-links/storage"
	"cc-links/storage/memory"
	"context"
	"fmt"
)

func main() {
	s := memory.New()
	//s, err := postgres.New()
	//s.Db["key"] = "value"
	//fmt.Println(s)

	link := storage.Link{
		URL: "https://example.com",
	}
	link.Hash()
	//fmt.Println(link)
	ctx := context.Background()

	if err := s.Create(ctx, &link); err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	//var st Storage = &MemoryStorage{}
	//fmt.Println(st)
}
