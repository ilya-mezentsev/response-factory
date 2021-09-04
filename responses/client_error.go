package responses

import (
	"net/http"
)

type ClientErrorResponse struct {
	DefaultResponse
}

func (r ClientErrorResponse) ApplicationStatus() string {
	return StatusError
}

func (r ClientErrorResponse) HttpStatus() int {
	return http.StatusBadRequest
}
