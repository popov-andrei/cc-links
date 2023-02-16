package storage

import (
	"context"
	"github.com/zeebo/xxh3"
	"strings"
)

type Storage interface {
	Create(ctx context.Context, link *Link) error
	Get(ctx context.Context, link *Link) error
}

const (
	alphabet        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	seed            = uint64(63) // len(alphabet)
	shortLinkLength = 10
)

type Link struct {
	URL      string
	ShortURL string
}

func (link *Link) Hash() {
	// использование имени переменной "link" вместо "l"
	b := []byte(link.URL)
	h := xxh3.Hash(b)

	var sb strings.Builder
	for i := 0; i < shortLinkLength; i++ {
		sb.WriteByte(alphabet[h%seed])
		h /= seed
	}

	link.ShortURL = sb.String()
}
