package response_factory

import (
	"net/http"
)

const (
	statusOk    = "ok"
	statusError = "error"
)

type Response interface {
	HttpStatus() int
	ApplicationStatus() string
	HasData() bool
	Data() interface{}
}

type defaultResponse struct {
	httpStatus        int
	applicationStatus string
	data              interface{}
}

func (r defaultResponse) HttpStatus() int {
	return r.httpStatus
}

func (r defaultResponse) ApplicationStatus() string {
	return r.applicationStatus
}

func (r defaultResponse) HasData() bool {
	return r.data != nil
}

func (r defaultResponse) Data() interface{} {
	return r.data
}

func DefaultResponse() Response {
	return defaultResponse{
		httpStatus:        http.StatusNoContent,
		applicationStatus: statusOk,
	}
}

func SuccessResponse(data interface{}) Response {
	return defaultResponse{
		httpStatus:        http.StatusOK,
		applicationStatus: statusOk,
		data:              data,
	}
}

func ServerError(data interface{}) Response {
	return defaultResponse{
		httpStatus:        http.StatusInternalServerError,
		applicationStatus: statusError,
		data:              data,
	}
}

func EmptyServerError() Response {
	return ServerError(nil)
}

func ClientError(data interface{}) Response {
	return defaultResponse{
		httpStatus:        http.StatusBadRequest,
		applicationStatus: statusError,
		data:              data,
	}
}

func EmptyClientError() Response {
	return ClientError(nil)
}

func ForbiddenError(data interface{}) Response {
	return defaultResponse{
		httpStatus:        http.StatusForbidden,
		applicationStatus: statusError,
		data:              data,
	}
}

func UnauthorizedError(data interface{}) Response {
	return defaultResponse{
		httpStatus:        http.StatusUnauthorized,
		applicationStatus: statusError,
		data:              data,
	}
}
