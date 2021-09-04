package responses

import "net/http"

type SuccessResponse struct {
	DefaultResponse
}

func (r SuccessResponse) HttpStatus() int {
	return http.StatusOK
}
