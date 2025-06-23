package errs

import (
	"errors"
	"net/http"
)

var (
	NotFoundError         = errors.New("not found error")
	ApplicationError      = errors.New("application error")
	LogicError            = errors.New("application error")
	ExistsAlreadyError    = errors.New("exists already error")
	ValidationError       = errors.New("validation error")
	NotAuthenticatedError = errors.New("not authenticated error")
	NotAuthorizedError    = errors.New("not authorized error")
	CacheError            = errors.New("cache error")
	DatabaseError         = errors.New("database error")
	QueueError            = errors.New("queue error")
)

func GetHttpStatusCode(err error) int {
	switch err {
	case NotFoundError:
		return http.StatusNotFound
	case ExistsAlreadyError, ValidationError, LogicError:
		return http.StatusBadRequest
	case NotAuthenticatedError:
		return http.StatusUnauthorized
	case NotAuthorizedError:
		return http.StatusForbidden
	}

	return http.StatusInternalServerError
}
