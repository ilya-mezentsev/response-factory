package response_factory

import "github.com/ilya-mezentsev/response-factory/responses"

func DefaultResponse() responses.Response {
	return responses.DefaultResponse{}
}

func SuccessResponse(data interface{}) responses.Response {
	return responses.SuccessResponse{DefaultResponse: responses.DefaultResponse{D: data}}
}

func ServerError(data interface{}) responses.Response {
	return responses.ServerErrorResponse{DefaultResponse: responses.DefaultResponse{D: data}}
}

func EmptyServerError() responses.Response {
	return ServerError(nil)
}

func ClientError(data interface{}) responses.Response {
	return responses.ClientErrorResponse{DefaultResponse: responses.DefaultResponse{D: data}}
}

func EmptyClientError() responses.Response {
	return ClientError(nil)
}

func ForbiddenError(data interface{}) responses.Response {
	return responses.ForbiddenErrorResponse{DefaultResponse: responses.DefaultResponse{D: data}}
}

func UnauthorizedError(data interface{}) responses.Response {
	return responses.UnauthorizedErrorResponse{DefaultResponse: responses.DefaultResponse{D: data}}
}
