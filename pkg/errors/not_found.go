package errors

type NotFoundError struct {
	GenericError

	Details map[string]interface{}
}

func NewNotFoundError(reason string) NotFoundError {
	return NotFoundError{
		GenericError: GenericError{
			Code:         CodeNotFoundError,
			Message:      "Payload is wrong or malformed",
			MessageTitle: "Malformed Request",
		},
		Details: map[string]interface{}{
			"reason": reason,
		},
	}
}

func (err NotFoundError) Error() string {
	return err.Message
}

func (err NotFoundError) ErrorCode() string {
	return err.Code
}
