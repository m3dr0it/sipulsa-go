package responses

import (
	"net/http"
	"time"
)

type Response struct {
	Timestamp  time.Time `json:"timestamp"`
	StatusCode int32     `json:"status_code"`
	Message    string    `json:"message"`
	Content    any       `json:"content"`
}

func PaginationResponse() {

}

func OkResponse(data any) Response {
	res := Response{
		Timestamp:  time.Now(),
		StatusCode: http.StatusOK,
		Content:    data,
	}

	return res
}

func CreatedResponse(data any) Response {
	res := Response{
		Timestamp:  time.Now(),
		Message:    "created",
		StatusCode: http.StatusCreated,
		Content:    data,
	}

	return res
}

func ErrorResponse(err error) Response {
	res := Response{
		Timestamp:  time.Now(),
		Message:    err.Error(),
		StatusCode: http.StatusInternalServerError,
		Content:    nil,
	}

	return res
}

func ConflictResponse(err error) Response {
	res := Response{
		Timestamp:  time.Now(),
		Message:    err.Error(),
		StatusCode: http.StatusInternalServerError,
		Content:    nil,
	}

	return res
}
