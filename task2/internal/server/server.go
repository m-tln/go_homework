package server

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type requestBody struct {
	Str string `json:"inputString"`
}

type OutputData struct {
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
	b, _ := io.ReadAll(r.Body)

	var rb requestBody
	_ = json.Unmarshal(b, &rb)

	decodedBytes, _ := base64.StdEncoding.DecodeString(rb.Str)

	decodedString := string(decodedBytes)
	output := OutputData{OutputString: decodedString}
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

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", HandleVersion)
	mux.HandleFunc("/decode", HandleDecode)
	mux.HandleFunc("/hard-op", HandleHardOp)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go httpServer.ListenAndServe()
}
