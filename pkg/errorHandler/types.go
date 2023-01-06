package errorHandler

import (
	"fmt"

	"github.com/siruspen/logrus"
)

type internalError struct {
	message string
}

func NewInternalServerError(message string) *internalError {
	return &internalError{
		message: message,
	}
}

func (e *internalError) Error() string {
	return fmt.Sprintf("Internal server error: %s", e.message)
}

type clientError struct {
	message string
}

func NewClientError(message string) *clientError {
	return &clientError{
		message: message,
	}
}

func (e *clientError) Error() string {
	return fmt.Sprintf("Client error: %s", e.message)
}

// TODO: сделать возможность определять, админ ли пользователь, из-за которого произошла ошибка.
// Если админ - отдавать полную информацию об ошибке на клиент.
type internalSecretError struct {
	message string
}

func NewInternalSecretError(message string) *internalSecretError {
	return &internalSecretError{
		message: message,
	}
}

func (e *internalSecretError) Error() string {
	logrus.Error(e.message)
	return fmt.Sprintf("Internal error")
}
