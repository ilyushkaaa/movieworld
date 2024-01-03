package errorapp

import "errors"

var (
	ErrorNoUser      = errors.New("no user with such login")
	ErrorBadPassword = errors.New("wrong password for such user")
	ErrorUserExists  = errors.New("user with such username already exist")
	ErrorNoFilm      = errors.New("film with such id does not exist")
)
