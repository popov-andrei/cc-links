package main

import "fmt"
import "github.com/zeebo/xxh3"

const URL = "pw4.su/"

func main() {
	data := "https://github.com/popov-andrei"
	hash := convertStringToHash(data)
	fmt.Printf("XXH3 hash of '%s': %x\n", data, hash)

	res := linkCompleter(hash)
	fmt.Println("Short url:", res)
}

// конвертор url-строки в уникальный хэш
func convertStringToHash(s string) []byte {
	h := xxh3.New()
	data := []byte(s)
	_, _ = h.Write(data)

	return h.Sum(nil)
}

// достраиватель ссылки
func linkCompleter(hash []byte) string {
	return URL + fmt.Sprintf("%x", hash)
}
