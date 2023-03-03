package handler

import "myapp/storage"

// Pull Функция получает значение длинной ссылки из базы по короткой
func Pull(s storage.Storage, data string) (string, error) {
	return s.GetOriginalLink(data)
}

// Store Функция сохраняет значения длинной и короткой ссылки в базе
func Store(s storage.Storage, data string) (string, error) {
	return s.CreateShortLink(data)
}
