package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/objects"
	"gateway/utils"
	"io/ioutil"
	"net/http"
)

type PrivilegesM struct {
	client *http.Client
}

func NewPrivilegesM(client *http.Client) *PrivilegesM {
	return &PrivilegesM{client: client}
}

func (model *PrivilegesM) Fetch(authHeader string) *objects.PrivilegeInfoResponse {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/privilege", utils.Config.PrivilegesEndpoint), nil)
	req.Header.Add("Authorization", authHeader)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("client: error making http request\n")
	}

	data := &objects.PrivilegeInfoResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, data)
	return data
}

func (model *PrivilegesM) AddTicket(authHeader string, request *objects.AddHistoryRequest) (*objects.AddHistoryResponce, error) {
	req_body, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/history", utils.Config.PrivilegesEndpoint), bytes.NewBuffer(req_body))
	req.Header.Add("Authorization", authHeader)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	data := &objects.AddHistoryResponce{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, data)
	return data, nil
}

func (model *PrivilegesM) DeleteTicket(authHeader string, ticket_uid string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/history/%s", utils.Config.PrivilegesEndpoint, ticket_uid), nil)
	req.Header.Add("Authorization", authHeader)
	client := &http.Client{}
	_, err := client.Do(req)
	return err
}
