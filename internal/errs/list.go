package errs

import "errors"

var (
	ErrUserIDNotFoundInContext = errors.New("user id not found in context")
	ErrNotFound                = errors.New("not found")
	ErrArticleNotFound         = errors.New("article not found")

	ErrUsernameAlreadyExists       = errors.New("username already exists")
	ErrIncorrectUsernameOrPassword = errors.New("incorrect username or password")

	ErrFillRequiredFields = errors.New("fill required fields")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidPathParam   = errors.New("invalid path param")
)
