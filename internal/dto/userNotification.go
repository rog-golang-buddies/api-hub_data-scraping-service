package dto

import "fmt"

//UserNotification represents basic DTO notification to user if requested
//Initially it supposed to be simple - if err != nil => error happens, else all is ok
type UserNotification struct {
	Error *ProcessingError
}

func NewUserNotification(procErr *ProcessingError) UserNotification {
	return UserNotification{Error: procErr}
}

//ProcessingError represents basic DTO to provide information about the error
//when the processing request contains a notification request
type ProcessingError struct {
	Cause error

	Message string
}

func (pe *ProcessingError) Error() string {
	return fmt.Sprintf("%s: %v", pe.Message, pe.Cause)
}

func NewProcessingError(message string, err error) ProcessingError {
	return ProcessingError{
		Cause:   err,
		Message: message,
	}
}
