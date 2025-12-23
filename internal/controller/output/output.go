package output

import "time"

type StatsResponse struct {
	TotalPosts       int       `json:"total_posts"`
	MaximumTimestamp time.Time `json:"maximum_timestamp"`
	MinimumTimestamp time.Time `json:"minimum_timestamp"`
	P50              int64     `json:"p50"`
	P90              int64     `json:"p90"`
	P99              int64     `json:"p99"`
}
