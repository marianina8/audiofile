package command

import (
	"audiofile/internal/interfaces"
	"fmt"
)

type Parser struct {
	commands []interfaces.Command
}

func NewParser(commands []interfaces.Command) *Parser {
	return &Parser{commands: commands}
}

func (p *Parser) Parse(args []string) error {
	if len(args) < 1 {
		help()
		return nil
	}

	subcommand := args[0]
	for _, cmd := range p.commands {
		if cmd.Name() == subcommand {
			cmd.ParseFlags(args[1:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

func help() {
	help := `usage: ./audiofile-cli <command> [<flags>]

These are a few Audiofile commands:
    get      Get metadata for a particular audio file by id
    list     List all metadata
    upload   Upload audio file
	`
	fmt.Println(help)
}
