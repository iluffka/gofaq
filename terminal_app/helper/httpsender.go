package helper

import (
	"net/http"
	"bytes"
	"errors"
	"fmt"
	"gofaq/terminal_app/model"
	"io/ioutil"
	"encoding/json"
)

func GetAllRepos(urlFromConfig, userName string, gitHubResponse *[]model.GitHubResponse) error {
	url := fmt.Sprintf(urlFromConfig, userName)
	resp, err := makeRequest(url, "GET")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, gitHubResponse); err != nil {
		return err
	}
	return nil
}

func makeRequest(url, method string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, errors.New("Status: " + resp.Status)
	}

	return resp, err
}