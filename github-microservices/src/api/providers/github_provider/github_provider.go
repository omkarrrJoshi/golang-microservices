package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/omkarrrJoshi/golang-microservices/github-microservices/src/api/client/rest_client"
	"github.com/omkarrrJoshi/golang-microservices/github-microservices/src/api/model/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	CreateRepoURL             = "https://github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	header := http.Header{}
	header.Set(headerAuthorization, getAuthorizationHeader(accessToken))
	response, err := rest_client.Post(CreateRepoURL, request, header)

	if err != nil {
		log.Println("error when trying to create new repo in github", err)
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errReponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errReponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid json response body of error",
			}
		}
		errReponse.StatusCode = response.StatusCode
		return nil, &errReponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println("error when trying to unmarshal create repo successfull response", err)
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error when trying to unmarshal create repo successfull response",
		}
	}

	return &result, nil
}
