package errors

const (
	CodeMalformedRequest = "konnect:error:malformed_request"
	CodeValidationError  = "konnect:error:validation"
	CodeInternalError    = "konnect:error:internal"
	CodeMissingField     = "konnect:error:missing_field"
	CodeNotFoundError    = "konnect:error:not_found"
)

type ValidationError struct {
	GenericError

	Details map[string]interface{}
}

func NewMalformedRequestError(reason string) ValidationError {
	return ValidationError{
		GenericError: GenericError{
			Code:         CodeMalformedRequest,
			Message:      "Payload is wrong or malformed",
			MessageTitle: "Malformed Request",
		},
		Details: map[string]interface{}{
			"reason": reason,
		},
	}
}

func NewInvalidValueError(field string, reason string) ValidationError {
	return ValidationError{
		GenericError: GenericError{
			Code:         CodeValidationError,
			Message:      "A parameter has invalid value",
			MessageTitle: "Validation Error",
		},
		Details: map[string]interface{}{
			"parameter": field,
			"reason":    reason,
		},
	}
}

func NewMissingFieldError(field string) ValidationError {
	return ValidationError{
		GenericError: GenericError{
			Code:         CodeMissingField,
			Message:      "At least one required parameter is missing",
			MessageTitle: "Validation Error",
		},
		Details: map[string]interface{}{
			"parameter": field,
		},
	}
}

func (err ValidationError) Error() string {
	return err.Message
}

func (err ValidationError) ErrorCode() string {
	if _, found := err.Details["parameter"]; found {
		return err.Code + ":" + err.Details["parameter"].(string)
	}
	return err.Code
}

type GenericError struct {
	Code         string      `json:"code"`
	Message      string      `json:"message"`
	MessageTitle string      `json:"message_title"`
	Details      interface{} `json:"details,omitempty"`
}
