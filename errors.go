package crowd

import "errors"

var (
	ErrorGeneralEmptyURL      		= errors.New("You must set the crowd base URL")
	ErrorGeneralEmptyApplication 	= errors.New("You must set the crowd application name")
	ErrorGeneralEmptyPassword 		= errors.New("You must set a password to access the crowd application")
	ErrorGeneralNoPermissions     	= errors.New("Your application has no permission to perform the desired request")
)

var (
	ErrorUserAlreadyInGroup				= errors.New("User is already a direct member of the group")
	ErrorUserNotFound      				= errors.New("User could not be found")
	ErrorInvalidUserDataOrUserExists	= errors.New("Invalid user data, for example missing password or the user already exists")
	ErrorInvalidUserDataOrMismatch		= errors.New("Invalid user data, for example the usernames in the body and the uri don't match")
)

var (
	ErrorGroupNotFound      				= errors.New("Group could not be found")
	ErrorGroupAlreadyExists 				= errors.New("Group already exists")
	ErrorGroupNotFoundOrCircularDependency	= errors.New("Child group could not be found, or adding the membership would result in a circular dependency.")
)