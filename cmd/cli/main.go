package main

import (
	"audiofile/cmd/cli/command"
	"audiofile/internal/interfaces"
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
		os.Exit(1)
	}
}
