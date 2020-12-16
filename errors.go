package crowd

import "errors"

var (
	ErrorUserAlreadyInGroup	= errors.New("User is already a direct member of the group")
	ErrorNoPermissions     	= errors.New("Your application has no permission to access crowd")
	ErrorUserNotFound      	= errors.New("User could not be found")
	ErrorGroupNotFound      = errors.New("Group could not be found")
	NewApiEmptyURL      	= errors.New("You must set the crowd base URL")
	NewApiEmptyApplication 	= errors.New("You must set the crowd application name")
	NewApiEmptyPassword 	= errors.New("You must set a password to access the crowd application")
)