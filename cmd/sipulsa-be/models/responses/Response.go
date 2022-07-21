package models

import "time"

type Response struct {
	Timestamp  time.Time
	StatusCode int32
	Message    string
	Content    any
}
