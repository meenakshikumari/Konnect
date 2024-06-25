package errors

import "strings"

func NewInternalError(err error, msg ...string) InternalError {
	ierr := InternalError{
		GenericError: GenericError{
			Code:         CodeInternalError,
			Message:      "Something went wrong",
			MessageTitle: "Internal Error",
		},
		Cause: err,
	}

	if len(msg) > 0 {
		ierr.Message = strings.TrimSpace(strings.Join(msg, ": "))
	}

	return ierr
}

type InternalError struct {
	GenericError
	Cause error
}

func (ierr InternalError) Error() string {
	return ierr.Message
}

func (ierr InternalError) ErrorCode() string {
	return ierr.Code
}
