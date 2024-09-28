package main

import (
	"fmt"
	. "task1/storage"

	"github.com/google/uuid"
)

func main() {
	var l ILibrary = &Library{}
	l.SetIDfunc(uuid.New().String)
	l.AddBook("Pronin")
	fmt.Print(l.SearchByName("Pronin"))
}
