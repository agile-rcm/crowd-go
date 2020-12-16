package crowd

import "errors"

var (
	ErrInitEmptyURL      	= errors.New("URL can't be empty")
	ErrInitEmptyApp      	= errors.New("App can't be empty")
	ErrInitEmptyPassword 	= errors.New("Password can't be empty")
	ErrNoPerms           	= errors.New("Application does not have permission to use Crowd")
	ErrUserNoFound      	= errors.New("User could not be found")
	ErrGroupNoFound      	= errors.New("Group could not be found")
	ErrUserAlreadyInGroup	= errors.New("User is already a direct member of the group")
)