package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type requestBody struct {
	Str string `json:"inputString"`
}

type Output struct {
	OutputString string `json:"outputString"`
}

type OutputVersion struct {
	Version string `json:"version"`
}

func HandleVersion(w http.ResponseWriter, r *http.Request) {
	output := OutputVersion{Version: "v1.0.0"}
	json.NewEncoder(w).Encode(output)
}

func HandleDecode(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var rb requestBody
	err = json.Unmarshal(b, &rb)
	if err != nil {
		fmt.Println(err)
		return
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(rb.Str)
	if err != nil {
		fmt.Println(err)
		return
	}

	decodedString := string(decodedBytes)
	output := Output{OutputString: decodedString}
	json.NewEncoder(w).Encode(output)
}

func HandleHardOp(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(11)+10) * time.Second)
	if rand.Intn(200)%2 == 1 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

const shutdownTimeout = 15 * time.Second

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", HandleVersion)
	mux.HandleFunc("/decode", HandleDecode)
	mux.HandleFunc("/hard-op", HandleHardOp)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("err in listen: %s\n", err)
			return fmt.Errorf("failed to serve http server: %w", err)
		}
		fmt.Println("after listener")

		return nil
	})

	group.Go(func() error {
		fmt.Println("before ctx done")
		<-ctx.Done()
		fmt.Println("after ctx done")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			return err
		}
		fmt.Println("after server shutdown")

		return nil
	})

	err := group.Wait()
	if err != nil {
		fmt.Printf("after wait: %s\n", err)
		return
	}
}
