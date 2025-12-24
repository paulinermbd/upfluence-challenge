package service

import (
	"bufio"
	"challenge/internal/model"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// DataFileReader représente un lecteur de fichier de données
type DataFileReader struct {
	filepath string
}

// NewDataFileReader crée une nouvelle instance du lecteur
func NewDataFileReader(filepath string) *DataFileReader {
	return &DataFileReader{
		filepath: filepath,
	}
}

// Read lit et parse les données depuis le fichier
func (r *DataFileReader) Read() ([]model.ProcessedData, error) {
	file, err := r.openFile()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return r.parseFile(file)
}

func (r *DataFileReader) openFile() (*os.File, error) {
	file, err := os.Open(r.filepath)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le fichier %s: %w", r.filepath, err)
	}
	return file, nil
}

func (r *DataFileReader) parseFile(file *os.File) ([]model.ProcessedData, error) {
	var allData []model.ProcessedData
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		data, err := r.parseLine(line, lineNumber)
		if err != nil {
			return nil, err
		}

		allData = append(allData, data)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("erreur de lecture du fichier: %w", err)
	}

	return allData, nil
}

func (r *DataFileReader) parseLine(line string, lineNumber int) (model.ProcessedData, error) {
	var event model.EventData

	if err := json.Unmarshal([]byte(line), &event); err != nil {
		return model.ProcessedData{}, fmt.Errorf("erreur de parsing JSON à la ligne %d: %w", lineNumber, err)
	}

	return event.Data, nil
}

func (r *DataFileReader) ExtractDimension(dimensionType string, data []model.ProcessedData) []int {
	dimension := make([]int, 0, len(data))

	for _, d := range data {
		var value int
		switch dimensionType {
		case "likes":
			value = d.Likes
		case "comments":
			value = d.Comments
		case "favorites":
			value = d.Favorites
		case "retweets":
			value = d.Retweets
		default:
			return nil
		}
		dimension = append(dimension, value)
	}

	return dimension
}
