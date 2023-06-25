package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	url  string
	http *http.Client
}

func NewClient(url string) Client {
	return Client{
		url:  url,
		http: &http.Client{},
	}
}

func (a *Client) Get() (string, error) {
	res, err := a.http.Get(a.url)
	if err != nil {
		return "", fmt.Errorf("error: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %s", body)
	}

	return string(body), nil
}

func (a *Client) Post(bz []byte) (string, error) {
	buf := bytes.NewBuffer(bz)

	res, err := a.http.Post(a.url, "text/turtle", buf)
	if err != nil {
		return "", fmt.Errorf("error: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: %s", body)
	}

	return string(body), nil
}
