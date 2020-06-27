package controller

import (
	"fmt"
	"net/http"
)

type Error struct {
	error
	Status int
	Code   *int
	Msg    string
}

func (e Error) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "error"
}

// http status
func errorHttpStatus(err error) int {
	e, ok := err.(Error)
	if !ok {
		return http.StatusInternalServerError
	}

	if *e.Code >= 400 && *e.Code <= 499 {
		return *e.Code
	}
	return 403
}

func errorBody(err error) interface{} {
	r := map[string]interface{}{"message": err.Error()}
	if e, ok := err.(Error); ok && e.Code != nil {
		r["code"] = e.Code
	}
	return r
}

func errorf(status int, code *int, format string, a ...interface{}) error {
	return Error{
		fmt.Errorf(format, a...),
		status,
		code,
		fmt.Sprintf(format, a...),
	}
}

func newInvalidInputErrorf(format string, a ...interface{}) error {
	code := 0
	return errorf(http.StatusBadRequest, &code, format, a...)
}

func newUnauthorizedErrorf(format string, a ...interface{}) error {
	code := 0
	return errorf(http.StatusUnauthorized, &code, format, a...)
}

func newNotFoundError(name, id interface{}) error {
	code := 0
	return errorf(http.StatusNotFound, &code, "%s not found, id: %v", name, id)
}

func newInvalidParameterError(name, value interface{}) error {
	return newInvalidInputErrorf("Invalid parameter [%s]='%v'", name, value)
}

func newCannotProcessErrorf(format string, a ...interface{}) error {
	code := 0
	return errorf(http.StatusUnprocessableEntity, &code, format, a...)
}

func newCannotProcessErrorfWithCode(code int, format string, a ...interface{}) error {
	return errorf(http.StatusUnprocessableEntity, &code, format, a...)
}
