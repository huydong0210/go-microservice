package http

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"go-microservices/internal/api/response"
	"io"
	"log"
	"net/http"
	"os"
)

type HttpHandler struct {
	address *AddressServiceConfig
}

type AddressServiceConfig struct {
	gatewayServiceAddress string
	userServiceAddress    string
	authServiceAddress    string
	todoServiceAddress    string
}

func NewHttpHandler() *HttpHandler {
	address, err := loadAddressConfig()
	if err != nil {
		panic(err)
	}
	return &HttpHandler{address: address}
}
func (h *HttpHandler) GetUserInfoByUsername(username string) (*response.UserInfoResponse, error) {

	var result response.UserInfoResponse
	log.Println("Calling user service : GetUserInfoByUsername()")
	uri := buildUrI(h.address.userServiceAddress) + "api/user/user-login/" + username
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)
	_ = responseBody
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return nil, err
	}

	// clean up memory after execution
	defer response.Body.Close()
	return &result, nil
}

func loadAddressConfig() (*AddressServiceConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	config := &AddressServiceConfig{
		gatewayServiceAddress: os.Getenv("GATEWAY_SERVICE_ADDRESS"),
		userServiceAddress:    os.Getenv("USER_SERVICE_ADDRESS"),
		authServiceAddress:    os.Getenv("AUTH_SERVICE_ADDRESS"),
		todoServiceAddress:    os.Getenv("TODO_SERVICE_ADDRESS"),
	}

	return config, nil
}
func buildUrI(address string) string {
	return "http://" + "localhost" + address + "/"
}
