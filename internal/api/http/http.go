package http

import (
	"bytes"
	"encoding/json"
	error2 "go-microservices/internal/error"
	"io"
	"net/http"
)

func DoRequest(request *http.Request) (*string, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		var data map[string]string
		err = json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			return nil, err
		}
		return nil, error2.NewAppError(data["error"])
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	result := string(responseBody)

	defer response.Body.Close()
	return &result, nil

}

func MakeRequest(uri, method string, requestBody map[string]string, header map[string]string) (*http.Request, error) {
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, uri, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	for key, value := range header {
		request.Header.Set(key, value)
	}
	return request, nil

}
