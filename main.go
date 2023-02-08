package main

import "fmt"
import "github.com/zeebo/xxh3"

func main() {
	data := "https://github.com/popov-andrei"
	hash := convertStringToHash(data)
	fmt.Printf("XXH3 hash of '%s': %x\n", data, hash)
}

func convertStringToHash(s string) []byte {
	h := xxh3.New()
	data := []byte(s)
	_, _ = h.Write(data)

	return h.Sum(nil)
}
