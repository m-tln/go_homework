package main

import (
	"fmt"
	. "task1/storage"
	"time"

	"github.com/google/uuid"
)

func StringID() string {
	return uuid.New().String()
}

func main() {
	books := []string{"The Whale", "1984", "Frankenstein"}
	var library ILibrary = &LibraryMap{}
	library.SetIDfunc(StringID)

	for _, book := range books {
		library.AddBook(book)
	}

	fmt.Println(library.SearchByName("The Whale"))

	library.SetIDfunc(time.UTC.String)
	fmt.Println(library.SearchByName("1984"))

	library = &LibrarySlice{}
	library.SetIDfunc(StringID)
	library.AddBook("The Great Gatsby")
	library.AddBook("To Kill a Mockingbird")

	fmt.Println(library.SearchByName("The Great Gatsby"))
}
