package app

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type App struct {
	url  string
	http *http.Client
}

func NewApp(url string) App {
	return App{
		url:  url,
		http: &http.Client{},
	}
}

func (c App) Get() (string, error) {
	res, err := c.http.Get(c.url)
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

func (c App) Post(bz []byte) (string, error) {
	buf := bytes.NewBuffer(bz)

	res, err := c.http.Post(c.url, "text/turtle", buf)
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
