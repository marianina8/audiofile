package command

import (
	"audiofile/internal/interfaces"
	"flag"
	"fmt"
)

func NewListCommand(client interfaces.Client) *ListCommand {
	gc := &ListCommand{
		fs:     flag.NewFlagSet("list", flag.ContinueOnError),
		client: client,
	}
	return gc
}

type ListCommand struct {
	fs     *flag.FlagSet
	client interfaces.Client
}

func (l *ListCommand) Name() string {
	return l.fs.Name()
}

func (l *ListCommand) ParseFlags(flags []string) error {
	return l.fs.Parse(flags)
}

func (l *ListCommand) Run() error {
	fmt.Println("listing audiofiles...")
	return nil
}
