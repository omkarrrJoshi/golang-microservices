package github_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/omkarrrJoshi/golang-microservices/github-microservices/src/api/client/rest_client"
	"github.com/omkarrrJoshi/golang-microservices/github-microservices/src/api/model/github"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest_client.StartMockups()
	os.Exit(m.Run())
}

func TestGetAuthorizationHeader(t *testing.T) {
	accessToken := "abcd"
	authorizationHeader := getAuthorizationHeader(accessToken)

	assert.EqualValues(t, authorizationHeader, "token abcd")
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	rest_client.FlushMocks()
	rest_client.AddMockUp(&rest_client.Mock{
		Url:        CreateRepoURL,
		MethodType: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})

	resp, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message, "invalid restclient response")
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	rest_client.FlushMocks()

	invalidCloser, _ := os.Open("random non json string")
	rest_client.AddMockUp(&rest_client.Mock{
		Url:        CreateRepoURL,
		MethodType: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})

	resp, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message, "invalid response body")
}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	rest_client.FlushMocks()

	rest_client.AddMockUp(&rest_client.Mock{
		Url:        CreateRepoURL,
		MethodType: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
	})

	resp, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message, "invalid json response body of error")
}

func TestCreateRepoInvalidUnauthorized(t *testing.T) {
	rest_client.FlushMocks()

	rest_client.AddMockUp(&rest_client.Mock{
		Url:        CreateRepoURL,
		MethodType: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(
				`{"message": "Requires authentication", "documnetation_url":"xyz/url"}`,
			)),
		},
	})

	resp, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message, "Requires authentication")
}

func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	rest_client.FlushMocks()

	rest_client.AddMockUp(&rest_client.Mock{
		Url:        CreateRepoURL,
		MethodType: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(
				`{"id": "string val"}`,
			)),
		},
	})

	resp, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message, "error when trying to unmarshal create repo successfull response")
}

func TestCreateRepoSuccess(t *testing.T) {
	rest_client.FlushMocks()

	rest_client.AddMockUp(&rest_client.Mock{
		Url:        CreateRepoURL,
		MethodType: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(
				`{"id": 1, "name": "omkar", "full_name":"omkars full name"}`,
			)),
		},
	})

	resp, err := CreateRepo("", github.CreateRepoRequest{})

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, resp.Id, 1)
	assert.EqualValues(t, resp.Name, "omkar")
	assert.EqualValues(t, resp.FullName, "omkars full name")
}
