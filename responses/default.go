package responses

import (
	"net/http"
)

type DefaultResponse struct {
	D interface{}
}

func (r DefaultResponse) HttpStatus() int {
	return http.StatusNoContent
}

func (r DefaultResponse) ApplicationStatus() string {
	return StatusOk
}

func (r DefaultResponse) HasData() bool {
	return r.D != nil
}

func (r DefaultResponse) Data() interface{} {
	return r.D
}
