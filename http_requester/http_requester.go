package http_requester

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetRequest(url string, headers map[string]string, timeout int8) (string, error) {
	if timeout < 0 {
		return "", errors.New("timeout must be greater than 0, got: " + fmt.Sprint(timeout) + " instead")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	http.DefaultClient.Timeout = time.Duration(timeout) * time.Second

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func PostRequest(url string, headers map[string]string, body io.Reader, timeout int8) (string, error) {
	if timeout < 0 {
		return "", errors.New("timeout must be greater than 0, got: " + fmt.Sprint(timeout) + " instead")
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return "", err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	http.DefaultClient.Timeout = time.Duration(timeout) * time.Second

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}
