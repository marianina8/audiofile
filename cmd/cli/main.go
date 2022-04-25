package main

import (
	"audiofile/cmd/cli/command"
	"audiofile/internal/interfaces"
	"fmt"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	cmds := []interfaces.Command{
		command.NewGetCommand(client),
		command.NewUploadCommand(client),
		command.NewListCommand(client),
	}
	parser := command.NewParser(cmds)
	if err := parser.Parse(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
