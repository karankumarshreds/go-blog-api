package custom_errors

type CustomError struct {
	Status  int
	Message string
}

func NotFoundError(message string) *CustomError {
	return &CustomError{Status: 404, Message: message}
}

func BadRequestError(message string) *CustomError {
	return &CustomError{Status: 400, Message: message}
}

func NotAuthorizedError(message string) *CustomError {
	return &CustomError{Status: 401, Message: message}
}
func InternalServerError(message string) *CustomError {
	return &CustomError{Status: 500, Message: message}
}
