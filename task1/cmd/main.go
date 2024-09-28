package main

import (
	. "task1/storage"
)

func main() {
	var library ILibrary = &LibraryMap{}
	library.AddBook("The Whale")
}
