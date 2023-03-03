package handler

import (
	"errors"
	"github.com/zeebo/xxh3"
	"net/url"
	"strings"
)

const (
	code          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	seed          = uint64(63) // len(alphabet)
	shortLinkSize = 10
)

// Hash create a unique encrypted string from URL.
func Hash(originalLink string) (string, error) {
	// Validate the input string.
	if originalLink == "" {
		return "", errors.New("URL is empty")
	}

	_, err := url.Parse(originalLink)
	if err != nil {
		return "", errors.New("URL is invalid")
	}

	h := xxh3.Hash([]byte(originalLink))

	var sb strings.Builder
	// We use double precision to reduce the number of collisions.
	for i := 0; i < shortLinkSize*2; i++ {
		sb.WriteByte(code[h%seed])
		h /= seed
	}

	return sb.String()[:shortLinkSize], nil
}
