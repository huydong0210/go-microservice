package http

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"go-microservices/internal/api/request"
	"go-microservices/internal/api/response"
	error2 "go-microservices/internal/error"
	"io"
	"log"
	"net/http"
	"os"
)

type HttpHandler struct {
	address *AddressServiceConfig
}

type AddressServiceConfig struct {
	GatewayServiceAddress string
	UserServiceAddress    string
	AuthServiceAddress    string
	TodoServiceAddress    string
}

func NewHttpHandler() *HttpHandler {
	address, err := loadAddressConfig()
	if err != nil {
		panic(err)
	}
	return &HttpHandler{address: address}
}
func (h *HttpHandler) GetUserInfoByUsername(username string) (*response.UserInfoResponse, error) {
	log.Println("Calling user service : GetUserInfoByUsername")
	request, err := h.makeRequest(http.MethodGet, "/api/user/"+username, nil)
	if err != nil {
		return nil, err
	}
	res, err := h.doRequest(request)
	if err != nil {
		return nil, err
	}

	var tmp map[string]response.UserInfoResponse
	err = json.Unmarshal([]byte(*res), &tmp)
	if err != nil {
		return nil, err
	}
	user := tmp["data"]
	return &user, nil

}

func (h *HttpHandler) CreateUser(creationRequest *request.UserCreationRequest) error {
	log.Println("Calling user service : CreateUser")

	requestBody := make(map[string]string)
	requestBody["username"] = creationRequest.Username
	requestBody["password"] = creationRequest.Password
	requestBody["email"] = creationRequest.Email

	request, err := h.makeRequest(http.MethodPost, "/api/user", requestBody)
	if err != nil {
		return err
	}
	res, err := h.doRequest(request)
	_ = res
	if err != nil {
		return err
	}
	return nil
}

func (h *HttpHandler) doRequest(request *http.Request) (*string, error) {
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

func (h *HttpHandler) makeRequest(method, path string, requestBody map[string]string) (*http.Request, error) {
	uri := buildUrI(h.address.UserServiceAddress) + path
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, uri, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	return request, nil

}

func loadAddressConfig() (*AddressServiceConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	config := &AddressServiceConfig{
		GatewayServiceAddress: os.Getenv("GATEWAY_SERVICE_ADDRESS"),
		UserServiceAddress:    os.Getenv("USER_SERVICE_ADDRESS"),
		AuthServiceAddress:    os.Getenv("AUTH_SERVICE_ADDRESS"),
		TodoServiceAddress:    os.Getenv("TODO_SERVICE_ADDRESS"),
	}

	return config, nil
}

func buildUrI(address string) string {
	return "http://" + "localhost" + address
}
