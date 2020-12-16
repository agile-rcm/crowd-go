package crowd

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/url"
	"runtime"
	"time"
)

const (
	VERSION 	= "3.0.1"
	NAME 		= "crowd-go"
)

type API struct {
	Client		*fasthttp.Client
	Url			string
	BasicAuth 	string
}

func NewAPI(url, application, applicationPassword string) (*API, error) {

	switch {
	case url == "":
		return nil, NewApiEmptyURL
	case application == "":
		return nil, NewApiEmptyApplication
	case applicationPassword == "":
		return nil, NewApiEmptyPassword
	}

	return &API{
		Client: &fasthttp.Client{
			Name:                generateUserAgent(),
			MaxIdleConnDuration: 5 * time.Second,
			ReadTimeout:         3 * time.Second,
			WriteTimeout:        3 * time.Second,
			MaxConnsPerHost:     150,
		},

		Url:       url,
		BasicAuth: generateBasicAuthString(application, applicationPassword),
	}, nil
}

func (api *API) requestPost(uri string, contentType string) *fasthttp.Request {

	r := fasthttp.AcquireRequest()
	r.SetRequestURI(api.Url + uri)
	r.Header.Add("Authorization", "Basic "+api.BasicAuth)

	switch contentType {
	case "json":
		r.Header.Set("Content-Type", "application/json")
		r.Header.Add("Accept", "application/json")
	case "xml":
		r.Header.Set("Content-Type", "application/xml")
		r.Header.Add("Accept", "application/xml")
	default:
		r.Header.Set("Content-Type", "application/json")
		r.Header.Add("Accept", "application/json")
	}

	return r

}

func (api *API) requestGet(uri string) *fasthttp.Request {

	r := fasthttp.AcquireRequest()
	r.SetRequestURI(api.Url + uri)
	r.Header.Add("Authorization", "Basic "+api.BasicAuth)

	return r

}

func (api *API) doPostRequest(uri string, contentType string, body interface{}) (int, error) {
	request := api.requestPost(uri, contentType)
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	switch contentType {
	case "json":
		bodyContent, err := json.Marshal(body)

		if err != nil {
			return 0, err
		}

		request.SetBody(bodyContent)
	case "xml":
		bodyContent, err := xml.Marshal(body)

		if err != nil {
			return 0, err
		}

		request.SetBody(bodyContent)
	default:
		bodyContent, err := json.Marshal(body)

		if err != nil {
			return 0, err
		}

		request.SetBody(bodyContent)
	}

	err := api.Client.Do(request, response)

	if err != nil {
		return 0, err
	}

	status := response.StatusCode()

	if !(status >= 200 && status <= 204) && status < 500 {
		return status, getCrowdErrorMessage(response.Body())
	}

	return status, nil

}

func getCrowdErrorMessage(data []byte) error {
	crowdErrorMessage := &crowdErrorMessage{}
	err := json.Unmarshal(data, crowdErrorMessage)

	if err != nil {
		return nil
	}

	return errors.New(crowdErrorMessage.Message)
}

func urlEscape(s string) string {
	return url.QueryEscape(s)
}

func unknownResponse(status int) error {
	return fmt.Errorf("Unknown response: %d)", status)
}

func generateBasicAuthString(username string, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

func generateUserAgent() string {

	userAgent := fmt.Sprintf(
		"%s/%s (go; %s; %s-%s)",
		NAME, VERSION, runtime.Version(),
		runtime.GOARCH, runtime.GOOS,
	)

	return userAgent
}