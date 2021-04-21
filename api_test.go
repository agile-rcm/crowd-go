package crowd

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"runtime"
	"testing"
)

type testmessage struct {
	Test string `json:"test"`
}

func TestGetCrowdErrorMessage(t *testing.T){

	message := &crowdErrorMessage{
		Message: "Some Error",
	}
	bytes, _ := json.Marshal(message)

	err := getCrowdErrorMessage(bytes)

	assert.Equal(t, errors.New(message.Message), err)

}

func TestUrlEscape(t *testing.T){

	testString := "some test string with char's & to be escaped"

	escapedString := urlEscape(testString)

	assert.Equal(t, url.QueryEscape(testString), escapedString)

}

func TestUnknownResponse(t *testing.T){

	status := 123

	unknownResponse := unknownResponse(status)

	assert.Equal(t, fmt.Errorf("Unknown response: %d", status), unknownResponse)

}

func TestGenerateBasicAuthString(t *testing.T){

	username := "test"
	password := "password"

	basicAuthString := generateBasicAuthString(username, password)

	assert.Equal(t, base64.StdEncoding.EncodeToString([]byte(username + ":" + password)), basicAuthString)
}

func TestGenerateUserAgent(t *testing.T){

	userAgent := fmt.Sprintf(
		"%s/%s (go; %s; %s-%s)",
		NAME, VERSION, runtime.Version(),
		runtime.GOARCH, runtime.GOOS,
	)

	assert.Equal(t, userAgent, generateUserAgent())

}

func TestNewAPI(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)

	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password", true)

	assert.Nil(t, err)
	assert.ObjectsAreEqual(API{}, api)

}

func TestRequestDelete(t *testing.T){

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password", true)

	assert.Nil(t, err)
	assert.ObjectsAreEqual(API{}, api)

	request := api.requestDelete("/testuri")

	assert.ObjectsAreEqual(request, fasthttp.Request{})
	assert.Equal(t, server.URL+"/testuri", request.URI().String())
	assert.Equal(t, []byte("DELETE"), request.Header.Method())

}

func TestRequestPost(t *testing.T){

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password", true)

	assert.Nil(t, err)

	request := api.requestPost("/testuri")

	assert.ObjectsAreEqual(request, fasthttp.Request{})
	assert.Equal(t, server.URL+"/testuri", request.URI().String())
	assert.Equal(t, []byte("POST"), request.Header.Method())
	assert.Equal(t, "application/json", string(request.Header.ContentType()))

}

func TestRequestPut(t *testing.T){

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password", true)

	assert.Nil(t, err)

	request := api.requestPut("/testuri")

	assert.ObjectsAreEqual(request, fasthttp.Request{})
	assert.Equal(t, server.URL+"/testuri", request.URI().String())
	assert.Equal(t, []byte("PUT"), request.Header.Method())
	assert.Equal(t, "application/json", string(request.Header.ContentType()))

}

func TestRequestGet(t *testing.T){

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password", true)

	assert.Nil(t, err)

	request := api.requestGet("/testuri")

	assert.ObjectsAreEqual(request, fasthttp.Request{})
	assert.Equal(t, server.URL+"/testuri", request.URI().String())
	assert.Equal(t, []byte("GET"), request.Header.Method())

}

func TestDoDeleteRequest(t *testing.T){

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)

		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "/testuri", r.RequestURI)

	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password", true)

	assert.Nil(t, err)

	res, err := api.doDeleteRequest("/testuri")

	assert.Equal(t, 200, res)

}

func TestDoGetRequest(t *testing.T){

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/testuri", r.RequestURI)

		respBytes, err := json.Marshal(testmessage{Test: "message"})

		assert.Nil(t, err)

		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)

	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password", true)

	assert.Nil(t, err)

	status, res, err := api.doGetRequest("/testuri")
	expectedResponse, err := json.Marshal(testmessage{Test: "message"})

	assert.Nil(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, expectedResponse , res)

}

func TestDoPostRequest(t *testing.T){

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.RequestURI == "/testuri400" {
			w.WriteHeader(http.StatusBadRequest)
			testError, err := json.Marshal(crowdErrorMessage{Message: "Test Error"})

			assert.Nil(t, err)

			w.Write(testError)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		expectedBody, err := json.Marshal(testmessage{Test: "message"})

		assert.Nil(t, err)
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/testuri", r.RequestURI)
		assert.Equal(t, expectedBody, body)

		w.WriteHeader(http.StatusOK)

	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password", true)

	assert.Nil(t, err)

	body := testmessage{Test: "message"}

	status, err := api.doPostRequest("/testuri", body)

	assert.Nil(t, err)
	assert.Equal(t, 200, status)

	status400, err := api.doPostRequest("/testuri400", body)

	assert.Equal(t, 400, status400)
	assert.Equal(t, errors.New("Test Error"), err)

}

func TestDoPutRequest(t *testing.T){

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.RequestURI == "/testuri400" {
			w.WriteHeader(http.StatusBadRequest)
			testError, err := json.Marshal(crowdErrorMessage{Message: "Test Error"})

			assert.Nil(t, err)

			w.Write(testError)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		expectedBody, err := json.Marshal(testmessage{Test: "message"})

		assert.Nil(t, err)
		assert.Equal(t, "PUT", r.Method)
		assert.Equal(t, "/testuri", r.RequestURI)
		assert.Equal(t, expectedBody, body)

		w.WriteHeader(http.StatusOK)

	}))
	defer server.Close()

	api, err := NewAPI(server.URL, "testapp", "password", true)

	assert.Nil(t, err)

	body := testmessage{Test: "message"}

	status, err := api.doPutRequest("/testuri", body)

	assert.Nil(t, err)
	assert.Equal(t, 200, status)

	status400, err := api.doPostRequest("/testuri400", body)

	assert.Equal(t, 400, status400)
	assert.Equal(t, errors.New("Test Error"), err)

}