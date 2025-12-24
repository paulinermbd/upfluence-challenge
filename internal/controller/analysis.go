package controller

import (
	"challenge/external"
	"challenge/internal/controller/output"
	"challenge/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetAnalysisHandler(w http.ResponseWriter, r *http.Request) {
	durationStr := r.URL.Query().Get("duration")
	dimension := r.URL.Query().Get("dimension")

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid duration given: %v", err), http.StatusBadRequest)
	}

	external.ReadStreamAndWriteData(duration)

	dataFileReader := service.NewDataFileReader("events.jsonl")
	data, err := dataFileReader.Read()
	if err != nil {
		http.Error(w, fmt.Sprintf("error reading data file: %v", err), http.StatusInternalServerError)
		return
	}

	dimensions := dataFileReader.ExtractDimension(dimension, data)

	percentiles, err := service.ComputePercentiles(dimensions)
	if err != nil {
		http.Error(w, fmt.Sprintf("error calculating percentiles: %v", err), http.StatusInternalServerError)
		return
	}

	response := output.StatsResponse{
		TotalPosts:       len(data),
		MaximumTimestamp: data[0].Timestamp.Time,
		MinimumTimestamp: data[len(data)-1].Timestamp.Time,
		P50:              int64(percentiles[50]),
		P90:              int64(percentiles[90]),
		P99:              int64(percentiles[99]),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}
