package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Input struct {
	InputString string `json:"inputString"`
}

type Output struct {
	OutputString string `json:"outputString"`
}

type OutputVersion struct {
	Version string `json:"version"`
}

func MakeRequest(ctx context.Context, client *http.Client, method, url string, body io.Reader) ([]byte, int, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}, 0, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return []byte{}, 0, err
	}

	return bodyBytes, resp.StatusCode, nil
}

func VersionRequest(ctx context.Context, client *http.Client) {
	versionURL := "http://localhost:8080/version"
	versionRequest, _, err := MakeRequest(ctx, client, http.MethodGet, versionURL, bytes.NewReader([]byte("")))
	if err != nil {
		fmt.Println(err)
	}
	var decodeVersion OutputVersion
	err = json.Unmarshal(versionRequest, &decodeVersion)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(decodeVersion.Version)
}

func DecodeRequest(ctx context.Context, client *http.Client) {
	decodeURL := "http://localhost:8080/decode"
	inputData := Input{InputString: "SGVsbG8gV29ybGQh"}
	jsonData, err := json.Marshal(inputData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	decodeResp, _, err := MakeRequest(ctx, client, http.MethodPost, decodeURL, bytes.NewReader(jsonData))
	if err != nil {
		fmt.Println(err)
		return
	}

	var decodeOutput Output
	err = json.Unmarshal(decodeResp, &decodeOutput)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(decodeOutput.OutputString)
}

func HardOp(ctx context.Context, client *http.Client) {
	hardopURL := "http://localhost:8080/hard-op"
	_, StatusCode, err := MakeRequest(ctx, client, http.MethodPost, hardopURL, bytes.NewReader([]byte("")))
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("false")
		} else {
			fmt.Println("Error getting /version:", err)
		}
		return
	}
	fmt.Println("true", StatusCode)
}

func RunClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := &http.Client{}

	VersionRequest(ctx, client)
	DecodeRequest(ctx, client)
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	HardOp(ctx, client)
}
