package main

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/drillby/codium-local-extension-manager/http_requester"
)

func main() {
	url := "https://marketplace.visualstudio.com/_apis/public/gallery/extensionquery?api-version=7.2-preview.1"
	headers := map[string]string{
		"Content-Type": "application/json",
		"Cookie":       "VstsSession=%7B%22PersistentSessionId%22%3A%224c0e54ac-75de-4d3a-bdd0-9c6f219a7a5c%22%2C%22PendingAuthenticationSessionId%22%3A%2200000000-0000-0000-0000-000000000000%22%2C%22CurrentAuthenticationSessionId%22%3A%2200000000-0000-0000-0000-000000000000%22%2C%22SignInState%22%3A%7B%7D%7D",
	}

	extension_name := "aaron-bond.better-comments"
	payload := map[string]interface{}{
		"assetTypes": "",
		"filters": []map[string]interface{}{
			{
				"criteria": []map[string]interface{}{
					{
						"filterType": 7,
						"value":      extension_name,
					},
				},
				"direction":   2,
				"pageSize":    100,
				"pageNumber":  1,
				"sortBy":      0,
				"sortOrder":   0,
				"pagingToken": "",
			},
		},
		"flags": 2151,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http_requester.PostRequest(url, headers, bytes.NewReader(payloadBytes), 10)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

}