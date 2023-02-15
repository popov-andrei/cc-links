package storage

import (
	"github.com/zeebo/xxh3"
	"strings"
)

type Storage interface {
	// Create(ctx context.Context, link *Link) (*Link, error)
	// Get(ctx context.Context, link *Link) (*Link, error)
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const seed = uint64(63) // len(alphabet)

const ShortLinkLength = 10

type Link struct {
	URL      []byte
	ShortURL string
}

func (link *Link) Hash() {
	// использование имени переменной "link" вместо "l"
	h := xxh3.Hash(link.URL)

	var sb strings.Builder
	for i := 0; i < ShortLinkLength; i++ {
		sb.WriteByte(alphabet[h%seed])
		h /= seed
	}
	link.ShortURL = sb.String()
}
