package command

import (
	"audiofile/internal/interfaces"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func NewGetCommand(client interfaces.Client) *GetCommand {
	gc := &GetCommand{
		fs:     flag.NewFlagSet("get", flag.ContinueOnError),
		client: client,
	}

	gc.fs.StringVar(&gc.id, "id", "", "id of audiofile requested")

	return gc
}

type GetCommand struct {
	fs     *flag.FlagSet
	client interfaces.Client
	id     string
}

func (cmd *GetCommand) Name() string {
	return cmd.fs.Name()
}

func (cmd *GetCommand) ParseFlags(flags []string) error {
	return cmd.fs.Parse(flags)
}

func (cmd *GetCommand) Run() error {
	params := "id=" + url.QueryEscape(cmd.id)
	path := fmt.Sprintf("http://localhost/request?%s", params)
	payload := &bytes.Buffer{}
	method := "GET"
	client := cmd.client

	req, err := http.NewRequest(method, path, payload)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
