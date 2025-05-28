package errs

import "errors"

var (
	NotFoundError         = errors.New("not found error")
	ApplicationError      = errors.New("application error")
	ValidationError       = errors.New("validation error")
	NotAuthenticatedError = errors.New("not authenticated error")
	NotAuthorizedError    = errors.New("not authorized error")
	AlreadyExistsError    = errors.New("already exists error")
	CacheError            = errors.New("cache error")
	DatabaseError         = errors.New("database error")
	QueueError            = errors.New("queue error")
)
