package memory

type Storage struct {
	basePath string
}

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

// Create(ctx context.Context, link *Link) (*Link, error)
//func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
//	defer func() { err = e.WrapIfErr("can't pick random page", err) }()
//
//	path := filepath.Join(s.basePath, userName)
//
//	// 1. check user folder
//	// 2. create folder
//
//	files, err := os.ReadDir(path)
//	if err != nil {
//		return nil, err
//	}
//
//	if len(files) == 0 {
//		return nil, storage.ErrNoSavedPages
//	}
//
//	rand.Seed(time.Now().UnixNano())
//	n := rand.Intn(len(files))
//
//	file := files[n]
//
//	return s.decodePage(filepath.Join(path, file.Name()))
//}
