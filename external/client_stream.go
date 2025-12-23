package external

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type StreamWriter struct {
	file   *os.File
	writer *bufio.Writer
}

func NewStreamWriter(filepath string) (*StreamWriter, error) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("opening file error: %w", err)
	}

	return &StreamWriter{
		file:   file,
		writer: bufio.NewWriter(file),
	}, nil
}

func (sw *StreamWriter) WriteLine(line string) error {
	if _, err := sw.writer.WriteString(line + "\n"); err != nil {
		return fmt.Errorf("error writing line: %w", err)
	}
	return nil
}

func (sw *StreamWriter) Flush() error {
	if err := sw.writer.Flush(); err != nil {
		return fmt.Errorf("error flushing buffer: %w", err)
	}
	if err := sw.file.Sync(); err != nil {
		return fmt.Errorf("error syncing file: %w", err)
	}
	return nil
}

func (sw *StreamWriter) Close() error {
	if err := sw.Flush(); err != nil {
		sw.file.Close()
		return err
	}
	return sw.file.Close()
}

func ReadStreamCorrectly() error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://stream.upfluence.co/stream", nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	writer, err := NewStreamWriter("./events.jsonl")
	if err != nil {
		return fmt.Errorf("error creating writer: %w", err)
	}
	defer writer.Close()

	scanner := bufio.NewScanner(resp.Body)

	// Increase buffer size if needed (following the number of events)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore empty lines (between two events)
		if line == "" {
			continue
		}

		if err := writer.WriteLine(line); err != nil {
			return fmt.Errorf("écriture ligne %d: %w", lineCount, err)
		}

		lineCount++

		// TODO: be more granular to deal with tiny duration
		// Flush every 100 events
		if lineCount%100 == 0 {
			if err := writer.Flush(); err != nil {
				return fmt.Errorf("flush after %d lines: %w", lineCount, err)
			}
			fmt.Printf("Total of received events: %d\n", lineCount)
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("flush final: %w", err)
	}

	if err := scanner.Err(); err != nil {
		// Check if timeout context error happens
		if ctx.Err() != nil {
			return fmt.Errorf("timeout after %d lines: %w", lineCount, ctx.Err())
		}
		return fmt.Errorf("reading error after %d lines: %w", lineCount, err)
	}

	fmt.Printf("End of stream. %d received events.\n", lineCount)
	return nil
}
