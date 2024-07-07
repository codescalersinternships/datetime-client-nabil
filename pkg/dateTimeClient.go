package datetimeclient

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	url    string
	client *http.Client
}

func NewClient(url string, timeout time.Duration) *Client {
	client := &Client{
		url: url,
		client: &http.Client{
			Timeout: timeout,
		},
	}
	return client
}

func (c *Client) GetCurrentDate() (string, error) {
	req, err := http.NewRequest(http.MethodGet, c.url+"/datetime", nil)
	req.Header.Add("Accept", "text/plain;charset=UTF-8, application/json")

	if err != nil {
		return "", err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
