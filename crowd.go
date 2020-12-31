package crowd

import (
	"encoding/json"
	"fmt"
)

// User management

// Get details of a crowd user.
func (api *API) GetUser(userName string) (*User, error) {

	user := &User{}

	url := fmt.Sprintf(
		"/rest/usermanagement/1/user?username=%s", urlEscape(userName),
	)

	status, result, err := api.doGetRequest(url)

	if status == 200 {

		err = json.Unmarshal(result, user)

		if err != nil {
			return nil, err
		}

	}

	switch status {
	case 200:
		return user, nil
	case 404:
		return nil, ErrorUserNotFound
	default:
		return nil, unknownResponse(status)
	}

}

// Add a new crowd user.
func (api *API) AddUser(userName, userPassword, userFirstName, userLastName, userDisplayName, userEmail string, isActive bool) error {

	if userDisplayName == "" {
		userDisplayName = userFirstName + userLastName
	}

	body := User{
		Name:        userName,
		FirstName:   userFirstName,
		LastName:    userLastName,
		DisplayName: userDisplayName,
		Email:       userEmail,
		IsActive:    isActive,
		Password:    PasswordValue{Value: userPassword},
	}

	url := "/rest/usermanagement/1/user"

	status, err := api.doPostRequest(url, body)

	if err != nil {
		return err
	}

	switch status {
	case 201:
		return nil
	case 400:
		return ErrorInvalidUserDataOrUserExists
	case 403:
		return ErrorGeneralNoPermissions
	default:
		return unknownResponse(status)
	}

}

// Remove a crowd user.
func (api *API) RemoveUser(userName string) error {

	url := fmt.Sprintf("/rest/usermanagement/1/user?username=%s", urlEscape(userName))

	status, err := api.doDeleteRequest(url)

	if err != nil {
		return err
	}

	switch status {
	case 204:
		return nil
	case 403:
		return ErrorGeneralNoPermissions
	case 404:
		return ErrorUserNotFound
	default:
		return unknownResponse(status)
	}

}

// Update details of a crowd user.
func (api *API) UpdateUser(userName, userFirstName, userLastName, userDisplayName, userEmail string, isActive bool) error {

	existingUser := &User{}

	existingUser, err := api.GetUser(userName)

	if err != nil {
		return err
	}

	if userName 		== "" {userName = existingUser.Name}
	if userFirstName 	== "" {userFirstName = existingUser.FirstName}
	if userLastName		== "" {userLastName = existingUser.LastName}
	if userDisplayName	== "" {userDisplayName = existingUser.DisplayName}
	if userEmail		== "" {userEmail = existingUser.Email}

	body := &User{
		Name:        userName,
		FirstName:   userFirstName,
		LastName:    userLastName,
		DisplayName: userDisplayName,
		Email:       userEmail,
		IsActive:    isActive,
		Password: 	 PasswordValue{},
	}

	url := fmt.Sprintf("/rest/usermanagement/1/user?username=%s", urlEscape(userName))

	status, err := api.doPutRequest(url, body)

	if err != nil {
		return err
	}

	switch status {
	case 204:
		return nil
	case 400:
		return ErrorInvalidUserDataOrMismatch
	case 403:
		return ErrorGeneralNoPermissions
	case 404:
		return ErrorUserNotFound
	default:
		return unknownResponse(status)
	}

}

// Get the attributes of a crowd user.
func (api *API) GetUserAttributes(userName string) (*Attributes, error) {

	attributes := &Attributes{}

	url := fmt.Sprintf(
		"/rest/usermanagement/1/user/attribute?username=%s", urlEscape(userName),
	)

	status, result, err := api.doGetRequest(url)

	if status == 200 {

		err = json.Unmarshal(result, attributes)

		if err != nil {
			return nil, err
		}

	}

	switch status {
	case 200:
		return attributes, nil
	case 404:
		return nil, ErrorUserNotFound
	default:
		return nil, unknownResponse(status)
	}

}

// Store (new) attributes for a crowd user.
func (api *API) StoreUserAttributes(userName string, attributes *Attributes) error {

	body := attributes

	url := fmt.Sprintf("/rest/usermanagement/1/user/attribute?username=%s", urlEscape(userName))

	status, err := api.doPostRequest(url, body)

	if err != nil {
		return err
	}

	switch status {
	case 204:
		return nil
	case 403:
		return ErrorGeneralNoPermissions
	case 404:
		return ErrorGroupNotFound
	default:
		return unknownResponse(status)
	}

}

// Remove attributes from a crowd user.
func (api *API) RemoveUserAttribute(userName, attributeName string) error {

	url := fmt.Sprintf("/rest/usermanagement/1/user/attribute?username=%s&attributename=%s", urlEscape(userName), urlEscape(attributeName))

	status, err := api.doDeleteRequest(url)

	if err != nil {
		return err
	}

	switch status {
	case 204:
		return nil
	case 403:
		return ErrorGeneralNoPermissions
	case 404:
		return ErrorUserNotFound
	default:
		return unknownResponse(status)
	}

}

// Add a user to an existing group.
func (api *API) AddUserToGroup(userName, groupName string) error {

	body := GroupName{Name: groupName}

	url := fmt.Sprintf(
		"/rest/usermanagement/1/user/group/direct?username=%s",
		urlEscape(userName),
	)

	status, err := api.doPostRequest(url, body)

	if err != nil {
		return err
	}

	switch status {
	case 201:
		return nil
	case 400:
		return ErrorGroupNotFound
	case 403:
		return ErrorGeneralNoPermissions
	case 404:
		return ErrorUserNotFound
	case 409:
		return ErrorUserAlreadyInGroup
	default:
		return unknownResponse(status)
	}
}

// Remove a user from a group.
func (api *API) RemoveUserFromGroup(userName, groupName string) error {

	url := fmt.Sprintf("/rest/usermanagement/1/user/group/direct?username=%s&groupname=%s", urlEscape(userName), urlEscape(groupName))

	status, err := api.doDeleteRequest(url)

	if err != nil {
		return err
	}

	switch status {
	case 204:
		return nil
	case 403:
		return ErrorGeneralNoPermissions
	case 404:
		return ErrorUserNotFound
	default:
		return unknownResponse(status)
	}

}

// Group management

// Create a new group.
func (api *API) CreateGroup(groupName, description string, isActive bool) error {

	body := Group{
		Name:        	groupName,
		Description: 	description,
		Type:        	"GROUP",
		Active: 		isActive,
	}

	url := "/rest/usermanagement/1/group"

	status, err := api.doPostRequest(url, body)

	if err != nil {
		return err
	}

	switch status {
	case 201:
		return nil
	case 400:
		return ErrorGroupAlreadyExists
	case 403:
		return ErrorGeneralNoPermissions
	default:
		return unknownResponse(status)
	}
}

// Remove a group.
func (api *API) RemoveGroup(groupName string) error {

	url := fmt.Sprintf("/rest/usermanagement/1/group?groupname=%s", urlEscape(groupName))

	status, err := api.doDeleteRequest(url)

	if err != nil {
		return err
	}

	switch status {
	case 204:
		return nil
	case 404:
		return ErrorGroupNotFound
	default:
		return unknownResponse(status)
	}
}

// Add a new child group membership.
func (api *API) AddChildGroupMembership(parentGroupName, childGroupName string) error {

	body := GroupName{
		Name:	childGroupName,
	}

	url := fmt.Sprintf("/rest/usermanagement/1/group/child-group/direct?groupname=%s", urlEscape(parentGroupName))

	status, err := api.doPostRequest(url, body)

	if err != nil {
		return err
	}

	switch status {
	case 201:
		return nil
	case 400:
		return ErrorGroupNotFoundOrCircularDependency
	case 404:
		return ErrorGroupNotFound
	default:
		return unknownResponse(status)
	}

}

// Add a new parent group membership.
func (api *API) AddParentGroupMembership(parentGroupName, childGroupName string) error {

	body := GroupName{
		Name:	parentGroupName,
	}

	url := fmt.Sprintf("/rest/usermanagement/1/group/parent-group/direct?groupname=%s", urlEscape(childGroupName))

	status, err := api.doPostRequest(url, body)

	if err != nil {
		return err
	}

	switch status {
	case 201:
		return nil
	case 400:
		return ErrorGroupNotFoundOrCircularDependency
	case 404:
		return ErrorGroupNotFound
	default:
		return unknownResponse(status)
	}

}