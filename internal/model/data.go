package model

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type ContentType string

const (
	InstagramMedia ContentType = "instagram_media"
	Pinterest      ContentType = "pin"
	Youtube        ContentType = "youtube_video"
	Article        ContentType = "article"
	Tweet          ContentType = "tweet"
	FacebookStatus ContentType = "facebook_status"
)

type CustomTime struct {
	time.Time
}

type ProcessedData struct {
	Type      ContentType `json:"-"`
	Likes     int         `json:"likes"`
	Comments  int         `json:"comments"`
	Favorites int         `json:"favorites"`
	Retweets  int         `json:"retweets"`
	Timestamp CustomTime  `json:"timestamp"`
}

type EventData struct {
	Data ProcessedData
}

func (e *EventData) UnmarshalJSON(data []byte) error {
	var raw map[ContentType]ProcessedData
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	for key, processedData := range raw {
		processedData.Type = key // Inject key into the structure
		e.Data = processedData
	}

	return nil
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")

	if s == "null" || s == "" {
		ct.Time = time.Time{}
		return nil
	}

	timestamp, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	ct.Time = time.Unix(timestamp, 0)
	return nil
}
