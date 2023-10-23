package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/drillby/codium-extension-manager/conf"
	"github.com/drillby/codium-extension-manager/http_requester"
	// "os"
)

var DownloadExtension = Comand{
	Name: "download",
	Help: "Download extension",
	Run: func(args []string) error {
		url := conf.Url
		headers := conf.Headers
		payload := conf.Payload

		if len(args) != 1 {
			fmt.Println("Usage: cem download <extension_name@version>")
			os.Exit(1)
		}

		extensionToDwnld := strings.Split(args[0], "@")
		if len(extensionToDwnld) != 2 {
			fmt.Println("Usage: cem download <extension_nameextension_name@version>")
			os.Exit(1)
		}

		wantedExtensionName := extensionToDwnld[0]
		wantedExtensionVersion := extensionToDwnld[1]

		filters := payload["filters"].([]map[string]interface{})
		filters[0]["criteria"].([]map[string]interface{})[0]["value"] = wantedExtensionName
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		res, err := http_requester.PostRequest(url, headers, bytes.NewReader(payloadBytes), 10)

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		versionList := res.Results[0].Extensions[0].Versions
		var extensionVersion struct {
			Version string "json:\"version\""
			Files   []struct {
				AssetType string "json:\"assetType\""
				Source    string "json:\"source\""
			} "json:\"files\""
		}
		if wantedExtensionVersion == "latest" {
			extensionVersion = versionList[0]
		} else {
			for _, version := range versionList {
				if version.Version == wantedExtensionVersion {
					extensionVersion = version
					break
				}
			}
		}

		if extensionVersion.Version == "" {
			fmt.Println("Extension not found")
			os.Exit(1)
		}

		fmt.Println("Found " + wantedExtensionName + "@" + extensionVersion.Version)
		return nil
	},
}

var UninstallExtension = Comand{
	Name: "uninstall",
	Help: "Uninstall extension",
	Run: func(args []string) error {
		// TODO: implement
		return nil
	},
}

var ListExtensions = Comand{
	Name: "list",
	Help: "Lists installed extensions",
	Run: func(args []string) error {
		// TODO: implement
		if len(args) > 0 {
			fmt.Println("list command does not take any arguments")
			return errors.New("list command does not take any arguments")
		}
		list, err := os.ReadDir(conf.ExtensionDir)
		if err != nil {
			fmt.Println(err)
			return err
		}
		for _, file := range list {
			fmt.Println(file.Name())
		}

		return nil
	},
}

var UpdateExtensions = Comand{
	Name: "update",
	Help: "Update extensions",
	Run: func(args []string) error {
		// TODO: implement
		return nil
	},
}
