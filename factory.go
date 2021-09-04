package response_factory

import "github.com/ilya-mezentsev/response-factory/responses"

type Response interface {
	HttpStatus() int
	ApplicationStatus() string
	HasData() bool
	Data() interface{}
}

func DefaultResponse() Response {
	return responses.DefaultResponse{}
}

func SuccessResponse(data interface{}) Response {
	return responses.SuccessResponse{DefaultResponse: responses.DefaultResponse{D: data}}
}

func ServerError(data interface{}) Response {
	return responses.ServerErrorResponse{DefaultResponse: responses.DefaultResponse{D: data}}
}

func EmptyServerError() Response {
	return ServerError(nil)
}

func ClientError(data interface{}) Response {
	return responses.ClientErrorResponse{DefaultResponse: responses.DefaultResponse{D: data}}
}

func EmptyClientError() Response {
	return ClientError(nil)
}

func ForbiddenError(data interface{}) Response {
	return responses.ForbiddenErrorResponse{DefaultResponse: responses.DefaultResponse{D: data}}
}

func UnauthorizedError(data interface{}) Response {
	return responses.UnauthorizedErrorResponse{DefaultResponse: responses.DefaultResponse{D: data}}
}
