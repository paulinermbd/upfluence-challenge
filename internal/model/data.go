package model

import "time"

type ContentType string

const (
	InstagramMedia ContentType = "instagram_media"
	Pinterest      ContentType = "pin"
	Youtube        ContentType = "youtube_video"
	Article        ContentType = "article"
	Tweet          ContentType = "tweet"
	FacebookStatus ContentType = "facebook_status"
)

type ProcessedData struct {
	Type      ContentType
	Likes     int
	Comments  int
	Favorites int
	Retweets  int
	Timestamp time.Time
}
