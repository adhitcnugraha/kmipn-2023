package client

import (
	"bytes"
	"encoding/json"
	"kmipn-2023/config"
	"net/http"
)

type SellerClient interface {
	Login(email, password string) (respCode int, err error)
	Register(username, email, password string) (respCode int, err error)
}

type sellerClient struct {
}

func NewSellerClient() *sellerClient {
	return &sellerClient{}
}

func (u *sellerClient) Login(email, password string) (respCode int, err error) {
	datajson := map[string]string{
		"email":    email,
		"password": password,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("POST", config.SetUrl("/api/v1/seller/login"), bytes.NewBuffer(data))
	if err != nil {
		return -1, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if err != nil {
		return -1, err
	} else {
		return resp.StatusCode, nil
	}
}

func (u *sellerClient) Register(username, email, password string) (respCode int, err error) {
	datajson := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("POST", config.SetUrl("/api/v1/seller/register"), bytes.NewBuffer(data))
	if err != nil {
		return -1, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if err != nil {
		return -1, err
	} else {
		return resp.StatusCode, nil
	}
}
