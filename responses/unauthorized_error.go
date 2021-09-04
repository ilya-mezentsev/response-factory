package responses

import (
	"net/http"
)

type UnauthorizedErrorResponse struct {
	DefaultResponse
}

func (r UnauthorizedErrorResponse) ApplicationStatus() string {
	return StatusError
}

func (r UnauthorizedErrorResponse) HttpStatus() int {
	return http.StatusUnauthorized
}
