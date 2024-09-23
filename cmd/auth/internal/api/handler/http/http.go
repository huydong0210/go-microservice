package http

import (
	"fmt"
	"github.com/joho/godotenv"
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
func (h *HttpHandler) GetUserInfoByUsername(username string) {
	log.Println("Calling user service : GetUserInfoByUsername()")
	uri := buildUrI(h.address.userServiceAddress) + username
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)
	_ = responseBody

	if error != nil {
		fmt.Println(error)
	}

	//formattedData := formatJSON(responseBody)
	//fmt.Println("Status: ", response.Status)
	//fmt.Println("Response body: ", formattedData)

	// clean up memory after execution
	defer response.Body.Close()
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
