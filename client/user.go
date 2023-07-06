package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"kmipn-2023/config"
	"kmipn-2023/model"
	"net/http"
)

type UserClient interface {
	Login(email, password string) (respCode int, err error)
	Register(username, email, password string) (respCode int, err error)
	GetUserProductCategory(token string) (*[]model.UserProductCategory, error)
}

type userClient struct {
}

func NewUserClient() *userClient {
	return &userClient{}
}

func (u *userClient) Login(email, password string) (respCode int, err error) {
	datajson := map[string]string{
		"email":    email,
		"password": password,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("POST", config.SetUrl("/api/v1/user/login"), bytes.NewBuffer(data))
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

func (u *userClient) Register(username, email, password string) (respCode int, err error) {
	datajson := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
	}

	data, err := json.Marshal(datajson)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("POST", config.SetUrl("/api/v1/user/register"), bytes.NewBuffer(data))
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

func (u *userClient) GetUserProductCategory(token string) (*[]model.UserProductCategory, error) {
	client, err := GetClientWithCookie(token)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", config.SetUrl("/api/v1/user/product"), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("status code not 200")
	}

	var userProduct []model.UserProductCategory
	err = json.Unmarshal(b, &userProduct)
	if err != nil {
		return nil, err
	}

	return &userProduct, nil
}
