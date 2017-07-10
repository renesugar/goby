package vm

import (
	"fmt"
)

// Nothing to describe here, just error classes
var (
	// ArgumentErrorClass ...
	ArgumentErrorClass *RClass
	// InternalErrorClass ...
	InternalErrorClass *RClass
	// NameErrorClass ...
	NameErrorClass *RClass
	// TypeErrorClass ...
	TypeErrorClass *RClass
	// UndefinedMethodErrorClass ...
	UndefinedMethodErrorClass *RClass
	// UnsupportedMethodClass ...
	UnsupportedMethodClass *RClass
)

const (
	// ArgumentError describes the error type in string
	ArgumentError = "ArgumentError"
	// InternalError is the default error type
	InternalError = "InternalError"
	// NameError describes constant related errors
	NameError = "NameError"
	// TypeError describes the error type in string
	TypeError = "TypeError"
	// UndefinedMethodError describes the error type in string
	UndefinedMethodError = "UndefinedMethodError"
	// UnsupportedMethodError describes the error type in string
	UnsupportedMethodError = "UnsupportedMethodError"
)

// Error class is just for special usages that holds internal error types.
// Goby maintainers should consider using the appropriate error type.
// Cannot create instances of Error class, or inherit Error class.
//
// The type of internal errors:
//
// * `InternalError`: default error type
// * `ArgumentError`: indicates an argument-related error
// * `NameError: indicates a constant-related error
// * `TypeError`: indicates a type-related error
// * `UndefinedMethodError`: indicates "undefined-method error"
// * `UnsupportedMethodError`: indicates "unsupported method error"
//
type Error struct {
	*baseObj
	Message string
}

func initErrorObject(errorType *RClass, format string, args ...interface{}) *Error {
	return &Error{
		baseObj: &baseObj{class: errorType},
		Message: fmt.Sprintf(errorType.Name+": "+format, args...),
	}
}

func (vm *VM) initErrorClasses() {
	ArgumentErrorClass = vm.initializeClass(ArgumentError, false)
	InternalErrorClass = vm.initializeClass(InternalError, false)
	NameErrorClass = vm.initializeClass(NameError, false)
	TypeErrorClass = vm.initializeClass(TypeError, false)
	UndefinedMethodErrorClass = vm.initializeClass(UndefinedMethodError, false)
	UnsupportedMethodClass = vm.initializeClass(UnsupportedMethodError, false)
}

// Polymorphic helper functions -----------------------------------------

// toString converts error messages into string.
func (e *Error) toString() string {
	return "ERROR: " + e.Message
}

// toJSON converts the receiver into JSON string.
func (e *Error) toJSON() string {
	return e.toString()
}
