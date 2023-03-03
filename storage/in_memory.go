package storage

import (
	"errors"
	"myapp/internal/handler"
	"sync"
)

type InMemoryStorage struct {
	data sync.Map
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{}
}

func (in *InMemoryStorage) GetOriginalLink(shortLink string) (string, error) {
	originalLink, ok := in.data.Load(shortLink)
	if !ok {
		return "", errors.New("link not found")
	}

	return originalLink.(string), nil
}

func (in *InMemoryStorage) CreateShortLink(originalLink string) (string, error) {
	shortLink, err := handler.Hash(originalLink)
	if err != nil {
		return "", err
	}

	_, ok := in.data.Load(shortLink)
	if ok {
		return "", errors.New("link already exists")
	}
	in.data.Store(shortLink, originalLink)

	return shortLink, nil
}
