package datetimeclient

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	backoff "github.com/cenkalti/backoff/v4"
)

// Client is the class used to make requests to the dateTime-server.
type Client struct {
	url    string
	client *http.Client
}

// NewClient creates a server to the Client to use.
// input is the url of site hitting to.
// return instance of the client
func NewClient(url string, timeout time.Duration) *Client {
	client := &Client{
		url: url,
		client: &http.Client{
			Timeout: timeout,
		},
	}
	return client
}

// NewClient creates a server to the Client to use.
// input is the key in Enviroment variables to get url of site hitting to.
// return instance of the client and err default is nil
func NewClientUsingEnv(key string, timeout time.Duration) (*Client, error) {
	url, exist := os.LookupEnv(key)
	if !exist {
		return &Client{}, fmt.Errorf("key isn't found %v", key)
	}
	client := &Client{
		url: url,
		client: &http.Client{
			Timeout: timeout,
		},
	}
	return client, nil
}

type currDateRes struct {
	Date string `json:"date"`
}

// GetDateTime makes a request to the dateTime-server to get the current date and time.
func (c *Client) GetCurrentDate() (currDateRes, error) {
	operation := func() (currDateRes, error) {
		req, err := http.NewRequest(http.MethodGet, c.url+"/datetime", nil)
		req.Header.Add("Accept", "text/plain;charset=UTF-8, application/json")

		if err != nil {
			return currDateRes{}, err
		}
		resp, err := c.client.Do(req)
		if err != nil {
			return currDateRes{}, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return currDateRes{}, fmt.Errorf("unexpected status code")
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return currDateRes{}, err
		}

		return currDateRes{string(body)}, nil
	}

	resp, err := backoff.RetryWithData(operation, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 3))
	if err != nil {
		return currDateRes{}, err
	}
	return resp, nil

}
