package responses

import (
	"net/http"
	"time"
)

type Response struct {
	Timestamp  time.Time
	StatusCode int32
	Message    string
	Content    any
}

func CreateResponse(data any) Response {
	res := Response{
		Timestamp:  time.Now(),
		Message:    "null",
		StatusCode: http.StatusFound,
		Content:    data,
	}

	return res
}
