package http_requester

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Extension struct {
	Results []struct {
		Extensions []struct {
			Publisher struct {
				PublisherName string `json:"displayName"`
			} `json:"publisher"`
			ExtensionName string `json:"displayName"`
			Versions      []struct {
				Version string `json:"version"`
				Files   []struct {
					AssetType string `json:"assetType"`
					Source    string `json:"source"`
				} `json:"files"`
			} `json:"versions"`
		} `json:"extensions"`
	} `json:"results"`
}

func GetRequest(url string, headers map[string]string, timeout int8) (Extension, error) {
	if timeout < 0 {
		return Extension{}, errors.New("timeout must be greater than 0, got: " + fmt.Sprint(timeout) + " instead")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Extension{}, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	http.DefaultClient.Timeout = time.Duration(timeout) * time.Second

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Extension{}, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Extension{}, err
	}

	var extension Extension
	err = json.Unmarshal(respBody, &extension)
	if err != nil {
		return Extension{}, err
	}

	return extension, nil
}

func PostRequest(url string, headers map[string]string, body io.Reader, timeout int8) (Extension, error) {
	if timeout < 0 {
		return Extension{}, errors.New("timeout must be greater than 0, got: " + fmt.Sprint(timeout) + " instead")
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return Extension{}, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	http.DefaultClient.Timeout = time.Duration(timeout) * time.Second

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Extension{}, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Extension{}, err
	}

	var extension Extension
	err = json.Unmarshal(respBody, &extension)
	if err != nil {
		return Extension{}, err
	}

	const errResponse = "{\"results\":[{\"extensions\":[],\"pagingToken\":null,\"resultMetadata\":[{\"metadataType\":\"ResultCount\",\"metadataItems\":[{\"name\":\"TotalCount\",\"count\":0}]}]}]}"

	if string(respBody) == errResponse {
		return Extension{}, errors.New("extension not found")
	}

	return extension, nil
}
