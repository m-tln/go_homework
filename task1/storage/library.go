package storage

type Book struct {
	Name string
}

type ILibrary interface {
	SearchByName(name string) (Book, bool)
	AddBook(name string)
	SetIDfunc(NewID func() string)
}

type Library struct {
	storage IStorage
	GetID   map[string]string
}

func New(s IStorage) *Library {
	return &Library{
		storage: s,
	}
}

func (l *Library) AddBook(name string) {
	b := Book{Name: name}
	key := l.storage.Add(b)
	if len(l.GetID) == 0 {
		l.GetID = make(map[string]string)
	}
	l.GetID[name] = key
}

func (l *Library) SearchByName(name string) (Book, bool) {
	book, ok := l.storage.SearchByID(l.GetID[name])
	return book.(Book), ok
}

func (l *Library) SetIDfunc(NewID func() string) {
	l.storage.SetIDfunc(NewID)
}
