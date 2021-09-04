package responses

import (
	"net/http"
)

type ForbiddenErrorResponse struct {
	DefaultResponse
}

func (r ForbiddenErrorResponse) ApplicationStatus() string {
	return StatusError
}

func (r ForbiddenErrorResponse) HttpStatus() int {
	return http.StatusForbidden
}
