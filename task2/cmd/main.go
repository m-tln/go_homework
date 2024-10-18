package main

import (
	"errors"
	"fmt"
	"net/http"
)

func HandleStart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Start our process"))
	fmt.Println("process was started")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/start", HandleStart)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := httpServer.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		fmt.Print(err)
	}
}
