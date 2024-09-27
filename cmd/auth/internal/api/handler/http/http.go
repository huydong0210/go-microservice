package http

import (
	"encoding/json"
	"github.com/joho/godotenv"
	http2 "go-microservices/internal/api/http"
	"go-microservices/internal/api/request"
	"go-microservices/internal/api/response"
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
	request, err := http2.MakeRequest("http://localhost"+h.address.UserServiceAddress+"/api/user/"+username, http.MethodGet, nil, nil)
	if err != nil {
		return nil, err
	}
	res, err := http2.DoRequest(request)
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

	request, err := http2.MakeRequest("http://localhost"+h.address.UserServiceAddress+"/api/user", http.MethodPost, requestBody, nil)
	if err != nil {
		return err
	}
	res, err := http2.DoRequest(request)
	_ = res
	if err != nil {
		return err
	}
	return nil
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
