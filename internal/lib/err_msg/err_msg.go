package err_msg

import "errors"

var (
	RequestBodyIsEmpty = errors.New("request body is empty")
	NoRowsAffected     = errors.New("no rows affected")
	MethodNotAllowed   = errors.New("method not allowed")
)
