package main

import (
	"fmt"

	"github.com/google/uuid"
)

type IStorage interface {
	SearchByID(id string) (any, bool)
	Add(elem any) string
}

type Storage struct {
	units map[string]any
}

func (s *Storage) SearchByID(id string) (any, bool) {
	res, ok := s.units[id]
	return res, ok
}

func (s *Storage) Add(elem any) string {
	if len(s.units) == 0 {
		s.units = make(map[string]any)
	}
	key := uuid.New().String()
	s.units[key] = elem
	return key
}

type Book struct {
	name string
}

type ILibrary interface {
	IStorage
	SearchByName(name string) (Book, bool)
	AddBook(name string)
}

type Library struct {
	Storage
	GetID map[string]string
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

func main() {
	var l ILibrary = &Library{}
	l.AddBook("Pronin")
	fmt.Print(l.SearchByName("Pronin"))
}
