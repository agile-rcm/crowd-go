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
		return ErrGroupNoFound
	case 403:
		return ErrNoPerms
	case 404:
		return ErrUserNoFound
	case 409:
		return ErrUserAlreadyInGroup
	default:
		return unknownResponse(status)
	}
}