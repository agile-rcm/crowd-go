package crowd

import "fmt"

func (api *API) AddUserToGroup(userName, groupName string) error {

	body := GroupInfo{Name: groupName}

	url := fmt.Sprintf(
		"/rest/usermanagement/1/user/group/direct?username=%s",
		urlEscape(userName),
	)

	status, err := api.doPostRequest(url, "json", body)

	if err != nil {
		return err
	}

	switch status {
	case 201:
		return nil
	case 400:
		return ErrorGroupNotFound
	case 403:
		return ErrorNoPermissions
	case 404:
		return ErrorUserNotFound
	case 409:
		return ErrorUserAlreadyInGroup
	default:
		return unknownResponse(status)
	}
}