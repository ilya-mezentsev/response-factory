package responses

import (
	"net/http"
)

type ServerErrorResponse struct {
	DefaultResponse
}

func (r ServerErrorResponse) ApplicationStatus() string {
	return StatusError
}

func (r ServerErrorResponse) HttpStatus() int {
	return http.StatusInternalServerError
}
