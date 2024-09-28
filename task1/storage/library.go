package storage

type Book struct {
	Name string
}

type ILibrary interface {
	IStorage
	SearchByName(name string) (Book, bool)
	AddBook(name string)
	SetIDfunc(NewID func() string)
}

type LibraryMap struct {
	StorageMap
	GetID map[string]string
}

func (l *LibraryMap) AddBook(name string) {
	b := Book{Name: name}
	key := l.Add(b)
	if len(l.GetID) == 0 {
		l.GetID = make(map[string]string)
	}
	l.GetID[name] = key
}

func (l *LibraryMap) SearchByName(name string) (Book, bool) {
	book, ok := l.SearchByID(l.GetID[name])
	return book.(Book), ok
}

func (l *LibraryMap) SetIDfunc(NewID func() string) {
	l.StorageMap.SetIDfunc(NewID)
}

type LibrarySlice struct {
	StorageSlice
	GetID map[string]string
}

func (l *LibrarySlice) AddBook(name string) {
	b := Book{Name: name}
	key := l.Add(b)
	if len(l.GetID) == 0 {
		l.GetID = make(map[string]string)
	}
	l.GetID[name] = key
}

func (l *LibrarySlice) SearchByName(name string) (Book, bool) {
	book, ok := l.SearchByID(l.GetID[name])
	return book.(Book), ok
}

func (l *LibrarySlice) SetIDfunc(NewID func() string) {
	l.StorageSlice.SetIDfunc(NewID)
}
