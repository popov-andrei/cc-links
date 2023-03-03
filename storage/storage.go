package storage

type Storage interface {
	CreateShortLink(originalLink string) (string, error)
	GetOriginalLink(shortLink string) (string, error)
}
