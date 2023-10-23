package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/drillby/codium-extension-manager/cli"
	"github.com/drillby/codium-extension-manager/conf"
	"github.com/drillby/codium-extension-manager/http_requester"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: cem <command> [arguments]")
		os.Exit(1)
	}

	if len(args) > 2 {
		fmt.Println("Usage: cem <command> [arguments]")
		os.Exit(1)
	}

	app := cli.Cli{}



	wantedCommand := os.Args[1]
	commandArgs := os.Args[2:]

	url := conf.Url
	headers := conf.Headers
	payload := conf.Payload

	extension_name := "aaron-bond.better-comments"

	filters := payload["filters"].([]map[string]interface{})
	filters[0]["criteria"].([]map[string]interface{})[0]["value"] = extension_name

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http_requester.PostRequest(url, headers, bytes.NewReader(payloadBytes), 10)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(res.Results[0].Extensions[0].Versions[0].Files[0].Source)

}
