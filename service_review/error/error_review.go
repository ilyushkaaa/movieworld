package errorreview

import "errors"

var (
	ErrorNoReview = errors.New("user has not got review with such id")
)
