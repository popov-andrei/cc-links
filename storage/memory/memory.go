package memory

import (
	"cc-links/storage"
	"context"
	"errors"
)

type Storage struct {
	Db map[string]string
}

func New() *Storage {
	return &Storage{Db: make(map[string]string)}
}

func (ms *Storage) Create(_ context.Context, link *storage.Link) error {

	if _, ok := ms.Db[link.ShortURL]; ok {
		return errors.New("link already exists")
	}

	link.Hash()
	ms.Db[link.ShortURL] = link.URL

	return nil
}

func (ms *Storage) Get(_ context.Context, link *storage.Link) error {

	if url, ok := ms.Db[link.ShortURL]; ok {
		link.URL = url
		return nil
	}

	return errors.New("link not found")
}
