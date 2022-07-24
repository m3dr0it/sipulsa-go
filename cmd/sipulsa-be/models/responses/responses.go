package responses

import (
	"time"
)

type Response struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Content   any       `json:"content"`
}

func PaginationResponse() {

}

func OkResponse(data any) Response {
	res := Response{
		Timestamp: time.Now(),
		Content:   data,
	}

	return res
}

func CreatedResponse(m string) Response {
	res := Response{
		Timestamp: time.Now(),
		Message:   "created " + m,
		Content:   nil,
	}

	return res
}

func ErrorResponse(err error) Response {
	res := Response{
		Timestamp: time.Now(),
		Message:   err.Error(),
		Content:   nil,
	}

	return res
}

func ConflictResponse(err error) Response {
	res := Response{
		Timestamp: time.Now(),
		Message:   err.Error(),
		Content:   nil,
	}

	return res
}
