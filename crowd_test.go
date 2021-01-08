package crowd

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Usermanagement

func TestAPI_GetUser(t *testing.T){

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/rest/usermanagement/1/user?username=testuser", r.RequestURI)

		resp := User{}
		respBytes, err := json.Marshal(resp)

		if err != nil {
			http.Error(w, string(respBytes), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)

	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password")

	assert.Nil(t, err)

	res, err := api.GetUser("testuser")

	assert.Nil(t, err)
	assert.Equal(t, &User{}, res)

}

func TestAPI_AddUser(t *testing.T){

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/rest/usermanagement/1/user", r.RequestURI)

		body := new(bytes.Buffer)
		_, err := body.ReadFrom(r.Body)
		content := &User{}
		err = json.Unmarshal(body.Bytes(), content)

		assert.Nil(t, err)
		assert.Equal(t, &User{}, content)

		respBytes, err := json.Marshal(content)

		if err != nil {
			http.Error(w, string(respBytes), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(respBytes)
	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password")

	assert.Nil(t, err)

	err = api.AddUser("", "", "", "", "", "", false)

	assert.Nil(t, err)

}

func TestAPI_RemoveUser(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "/rest/usermanagement/1/user?username=testuser", r.RequestURI)

		w.WriteHeader(http.StatusNoContent)

	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password")

	assert.Nil(t, err)

	err = api.RemoveUser("testuser")

	assert.Nil(t, err)

}

func TestAPI_UpdateUser(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			assert.Equal(t, "/rest/usermanagement/1/user?username=testuser", r.RequestURI)

			resp := User{}
			respBytes, err := json.Marshal(resp)

			if err != nil {
				http.Error(w, string(respBytes), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(respBytes)
		case "PUT":
			assert.Equal(t, "/rest/usermanagement/1/user?username=testuser", r.RequestURI)

			body := new(bytes.Buffer)
			body.ReadFrom(r.Body)
			content := &User{}
			json.Unmarshal(body.Bytes(), content)

			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password")

	assert.Nil(t, err)

	err = api.UpdateUser("testuser", "", "", "", "", false)

	assert.Nil(t, err)

}

func TestAPI_GetUserAttributes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/rest/usermanagement/1/user/attribute?username=testuser", r.RequestURI)

		resp := Attributes{}
		respBytes, err := json.Marshal(resp)

		if err != nil {
			http.Error(w, string(respBytes), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)

	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password")

	assert.Nil(t, err)

	res, err := api.GetUserAttributes("testuser")

	assert.Nil(t, err)
	assert.Equal(t, &Attributes{}, res)
}

func TestAPI_StoreUserAttributes(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/rest/usermanagement/1/user/attribute?username=testuser", r.RequestURI)

		body := new(bytes.Buffer)
		_, err := body.ReadFrom(r.Body)
		content := &Attributes{}
		err = json.Unmarshal(body.Bytes(), content)

		assert.Nil(t, err)
		assert.Equal(t, &Attributes{}, content)

		w.WriteHeader(http.StatusNoContent)

	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password")

	assert.Nil(t, err)

	err = api.StoreUserAttributes("testuser", &Attributes{})
	assert.Nil(t, err)

}

func TestAPI_RemoveUserAttribute(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "/rest/usermanagement/1/user/attribute?username=testuser&attributename=testattribute", r.RequestURI)

		w.WriteHeader(http.StatusNoContent)

	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password")

	assert.Nil(t, err)

	err = api.RemoveUserAttribute("testuser", "testattribute")

	assert.Nil(t, err)

}

func TestAPI_AddUserToGroup(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/rest/usermanagement/1/user/group/direct?username=testuser", r.RequestURI)

		body := new(bytes.Buffer)
		_, err := body.ReadFrom(r.Body)
		content := &GroupName{}
		err = json.Unmarshal(body.Bytes(), content)

		assert.Nil(t, err)
		assert.Equal(t, &GroupName{}, content)

		w.WriteHeader(http.StatusCreated)

	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password")

	assert.Nil(t, err)

	err = api.AddUserToGroup("testuser", "")

	assert.Nil(t, err)

}

func TestAPI_RemoveUserFromGroup(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "/rest/usermanagement/1/user/group/direct?username=testuser&groupname=testgroup", r.RequestURI)

		w.WriteHeader(http.StatusNoContent)

	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password")

	assert.Nil(t, err)

	err = api.RemoveUserFromGroup("testuser", "testgroup")

	assert.Nil(t, err)

}

// Groupmanagement

func TestAPI_CreateGroup(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/rest/usermanagement/1/group", r.RequestURI)

		body := new(bytes.Buffer)
		_, err := body.ReadFrom(r.Body)
		content := &Group{}
		err = json.Unmarshal(body.Bytes(), content)

		assert.Nil(t, err)
		assert.Equal(t, &Group{Name: "testgroup", Type: "GROUP"}, content)

		respBytes, err := json.Marshal(content)

		if err != nil {
			http.Error(w, string(respBytes), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(respBytes)
	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password")

	assert.Nil(t, err)

	err = api.CreateGroup("testgroup", "", false)

	assert.Nil(t, err)

}

func TestAPI_RemoveGroup(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "/rest/usermanagement/1/group?groupname=testgroup", r.RequestURI)

		w.WriteHeader(http.StatusNoContent)

	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password")

	assert.Nil(t, err)

	err = api.RemoveGroup("testgroup")

	assert.Nil(t, err)

}

func TestAPI_AddChildGroupMembership(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/rest/usermanagement/1/group/child-group/direct?groupname=parentgroup", r.RequestURI)

		body := new(bytes.Buffer)
		_, err := body.ReadFrom(r.Body)
		content := &GroupName{}
		err = json.Unmarshal(body.Bytes(), content)

		assert.Nil(t, err)
		assert.Equal(t, &GroupName{Name: "childgroup"}, content)

		respBytes, err := json.Marshal(content)

		if err != nil {
			http.Error(w, string(respBytes), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(respBytes)
	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password")

	assert.Nil(t, err)

	err = api.AddChildGroupMembership("parentgroup", "childgroup")

	assert.Nil(t, err)

}

func TestAPI_AddParentGroupMembership(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/rest/usermanagement/1/group/parent-group/direct?groupname=childgroup", r.RequestURI)

		body := new(bytes.Buffer)
		_, err := body.ReadFrom(r.Body)
		content := &GroupName{}
		err = json.Unmarshal(body.Bytes(), content)

		assert.Nil(t, err)
		assert.Equal(t, &GroupName{Name: "parentgroup"}, content)

		respBytes, err := json.Marshal(content)

		if err != nil {
			http.Error(w, string(respBytes), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(respBytes)
	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password")

	assert.Nil(t, err)

	err = api.AddParentGroupMembership("parentgroup", "childgroup")

	assert.Nil(t, err)

}