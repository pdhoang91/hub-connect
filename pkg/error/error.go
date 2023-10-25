package error

type ErrorHandler struct {
	Message string
	Code    int
}

func (err ErrorHandler) Error() string {
	return err.Message
}

func NewErrorHandler(code int, message string) ErrorHandler {
	return ErrorHandler{
		Code:    code,
		Message: message,
	}
}

func NewError(message string) ErrorHandler {
	return ErrorHandler{
		Message: message,
	}
}
