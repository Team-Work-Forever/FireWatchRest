package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Api struct {
	baseUrl string
	client  *http.Client
}

func New(url string) *Api {
	client := &http.Client{}

	return &Api{
		baseUrl: url,
		client:  client,
	}
}

func (api *Api) getJson(path string, response interface{}) error {
	body, err := api.get(path)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, response); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return nil
}

func (api *Api) getXml(path string, response interface{}) error {
	body, err := api.get(path)

	if err != nil {
		return err
	}

	if err := xml.Unmarshal(body, response); err != nil {
		return fmt.Errorf("failed to unmarshal XML: %v", err)
	}

	return nil
}

func (api *Api) get(path string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", api.baseUrl, path)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "HTTPie")

	resp, err := api.client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}
