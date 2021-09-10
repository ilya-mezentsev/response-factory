package response_factory

import (
	"net/http"
)

type Response interface {
	HttpStatus() int
	IsOk() bool
	HasData() bool
	Data() interface{}
}

type response struct {
	httpStatus int
	isOk       bool
	data       interface{}
}

func (r response) HttpStatus() int {
	return r.httpStatus
}

func (r response) IsOk() bool {
	return r.isOk
}

func (r response) HasData() bool {
	return r.data != nil
}

func (r response) Data() interface{} {
	return r.data
}

func DefaultResponse() Response {
	return response{
		httpStatus: http.StatusNoContent,
		isOk:       true,
	}
}

func SuccessResponse(data interface{}) Response {
	return response{
		httpStatus: http.StatusOK,
		isOk:       true,
		data:       data,
	}
}

func ServerError(data interface{}) Response {
	return response{
		httpStatus: http.StatusInternalServerError,
		data:       data,
	}
}

func EmptyServerError() Response {
	return ServerError(nil)
}

func ClientError(data interface{}) Response {
	return response{
		httpStatus: http.StatusBadRequest,
		data:       data,
	}
}

func EmptyClientError() Response {
	return ClientError(nil)
}

func ForbiddenError(data interface{}) Response {
	return response{
		httpStatus: http.StatusForbidden,
		data:       data,
	}
}

func UnauthorizedError(data interface{}) Response {
	return response{
		httpStatus: http.StatusUnauthorized,
		data:       data,
	}
}
