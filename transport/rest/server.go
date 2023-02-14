package rest

import (
	"fmt"
	"github.com/zeebo/xxh3"
	"io"
	"net/http"
	"strings"
)

const URL = "pw4.su/"

// Конвертор url-строки в уникальный хэш
func convertStringToHash(s string) []byte {
	h := xxh3.New()
	data := []byte(s)
	_, _ = h.Write(data)

	return h.Sum(nil)
}

// Достраивать ссылки
func linkCompleter(hash []byte) string {
	return URL + fmt.Sprintf("%x", hash)
}

// HandlePostRequest реализация обработки запросов
func HandlePostRequest(w http.ResponseWriter, r *http.Request, storage *map[string]string) {

	if r.Method == "GET" {
		// Получаем ссылку из URL
		link := r.URL.Query().Get("link")
		hash := strings.Split(link, "/")[1]

		// Проверяем ссылку в хранилище
		if value, ok := (*storage)[hash]; ok {
			w.Write([]byte(value))
		} else {
			w.Write([]byte("Link not found"))
		}

	} else if r.Method == "POST" {
		// Читаем тело запроса
		body, _ := io.ReadAll(r.Body)
		link := string(body)
		fmt.Println("Link:", link)

		// Делаем Hash из Link
		hash := convertStringToHash(link)
		// Компонуем ответ
		answer := linkCompleter(hash)

		// Сохранение результатов в мапе
		key := fmt.Sprintf("%x", hash)
		(*storage)[key] = link

		// Отправляем ответ
		w.Write([]byte(answer))

	} else {

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	// Печатаем хранилище
	fmt.Println(storage)
}
