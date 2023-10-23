package main

import (
	"fmt"
	"os"

	"github.com/drillby/codium-extension-manager/cli"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: cem <command> [argument]")
		os.Exit(1)
	}

	if len(args) > 2 {
		fmt.Println("Usage: cem <command> [argument]")
		os.Exit(1)
	}

	app := cli.Cli{}

	app.AddComand(cli.ListExtensions)
	app.AddComand(cli.DownloadExtension)
	app.AddComand(cli.UninstallExtension)
	app.AddComand(cli.UpdateExtensions)

	app.Run(args)

}
