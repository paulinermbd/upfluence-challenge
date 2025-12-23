package internal

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func ListenStream(body io.Reader) string {
	fmt.Println("start")

	client := http.Client{Timeout: 0 * time.Second}
	req, err := http.NewRequest(http.MethodGet, "https://stream.upfluence.co/stream", body)
	//resp, err := http.Get("https://stream.upfluence.co/stream")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	var buf bytes.Buffer
	written, err := io.Copy(&buf, resp.Body)
	if err != nil {
		return ""
	}
	fmt.Println(written)
	//jsonData := json.NewDecoder(resp.Body).Decode(&data)

	fmt.Printf("%v", buf.String())

	return "client"
}

func ReadStreamCorrectly() {
	fmt.Println("start")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://stream.upfluence.co/stream", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{
		Timeout: 0,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Status code: %d", resp.StatusCode)
	}

	fmt.Println("Connecté, lecture du stream...")

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Reçu: %s\n", line)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Erreur de lecture: %v", err)
	}

	fmt.Println("Stream terminé")
}
