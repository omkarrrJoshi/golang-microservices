package rest_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var (
	enableMocks = false
	mocks       = make(map[string]*Mock)
)

type Mock struct {
	MethodType string
	Url        string
	Response   *http.Response
	Err        error
}

func getMockId(methodType, url string) string {
	return methodType + "_" + url
}

func StartMockups() {
	enableMocks = true
}

func StopMockups() {
	enableMocks = false
}

func FlushMocks() {
	mocks = make(map[string]*Mock)
}

func AddMockUp(mock *Mock) {
	mocks[mock.MethodType+"_"+mock.Url] = mock
}

func Post(url string, body interface{}, header http.Header) (*http.Response, error) {
	if enableMocks {
		if mock := mocks[getMockId(http.MethodPost, url)]; mock != nil {
			return mock.Response, mock.Err
		}
		return nil, errors.New("no mocup found for given request")
	}
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		log.Println("error while marshalling json string")
		return nil, err
	}

	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = header
	client := http.Client{}
	return client.Do(request)
}
