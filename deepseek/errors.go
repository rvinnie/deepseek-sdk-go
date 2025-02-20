package deepseek

import (
	"encoding/json"
	"fmt"
	"io"
)

type DeepseekError struct {
	InternalError *Error `json:"error"`
}

type Error struct {
	Code    string `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Param   any    `json:"param"`
}

func (e *DeepseekError) Error() string {
	return fmt.Sprintf("code=%s type=%s message=[%v] param=%v", e.InternalError.Code, e.InternalError.Type, e.InternalError.Message, e.InternalError.Param)
}

func getError(r io.Reader) *DeepseekError {
	responseBytes, err := io.ReadAll(r)
	if err != nil {
		return createError("unknown", "unexpected_error", "Unable to read response body")
	}

	errResp := &DeepseekError{}
	err = json.Unmarshal(responseBytes, errResp)
	if err != nil {
		return createError("unknown", "unexpected_error", "Unexpected error occurs")
	}

	return errResp
}

func createError(code, errType, message string) *DeepseekError {
	return &DeepseekError{
		InternalError: &Error{
			Code:    code,
			Type:    errType,
			Message: message,
			Param:   nil,
		},
	}
}
