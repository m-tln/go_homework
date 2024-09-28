package storage

type Book struct {
	name string
}

type ILibrary interface {
	IStorage
	SearchByName(name string) (Book, bool)
	AddBook(name string)
}

type Library struct {
	StorageMap
	GetID map[string]string
	NewID func() string
}

func (l *Library) AddBook(name string) {
	b := Book{name: name}
	key := l.Add(b)
	if len(l.GetID) == 0 {
		l.GetID = make(map[string]string)
	}
	l.GetID[name] = key
}

func (l *Library) SearchByName(name string) (Book, bool) {
	book, ok := l.SearchByID(l.GetID[name])
	return book.(Book), ok
}
