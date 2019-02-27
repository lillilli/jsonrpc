package jsonrpc

import "fmt"

const (
	// errorCodeParse is parse error code.
	errorCodeParse errorCode = -32700
	// errorCodeInvalidRequest is invalid request error code.
	errorCodeInvalidRequest errorCode = -32600
	// errorCodeMethodNotFound is method not found error code.
	errorCodeMethodNotFound errorCode = -32601
	// errorCodeInvalidParams is invalid params error code.
	errorCodeInvalidParams errorCode = -32602
	// errorCodeInternal is internal error code.
	errorCodeInternal errorCode = -32603
)

type (
	// A errorCode by JSON-RPC v2.0.
	errorCode int

	// respError is a wrapper for a JSON interface value.
	respError struct {
		Code    errorCode   `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
)

// Error implements error interface.
func (e *respError) Error() string {
	return fmt.Sprintf("jsonrpc: code: %d, message: %s, data: %+v", e.Code, e.Message, e.Data)
}

// errParse returns parse error.
func errParse() *respError {
	return &respError{
		Code:    errorCodeParse,
		Message: "Parse error",
	}
}

// errInvalidRequest returns invalid request error.
func errInvalidRequest() *respError {
	return &respError{
		Code:    errorCodeInvalidRequest,
		Message: "Invalid Request",
	}
}

// errMethodNotFound returns method not found error.
func errMethodNotFound() *respError {
	return &respError{
		Code:    errorCodeMethodNotFound,
		Message: "Method not found",
	}
}

// errInvalidParams returns invalid params error.
func errInvalidParams() *respError {
	return &respError{
		Code:    errorCodeInvalidParams,
		Message: "Invalid params",
	}
}

// errInternal returns internal error.
func errInternal() *respError {
	return &respError{
		Code:    errorCodeInternal,
		Message: "Internal error",
	}
}
