//go:build !free && pro

package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/marianina8/audiofile/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Command to search for audiofiles by string",
	Long:    `Command to search for audiofiles by search string within the metadata file.  Search string is not case sensitive`,
	Example: `./bin/audiofile search --value electronic`,
	RunE: func(cmd *cobra.Command, args []string) error {
		verbose, _ := cmd.Flags().GetBool("verbose")
		var err error
		value, _ := cmd.Flags().GetString("value")
		if value == "" {
			value, err = utils.AskForValue()
			if err != nil {
				return utils.Error("\n  %v\n  try again and enter a value", err, verbose)
			}
		}
		params := "searchFor=" + url.QueryEscape(value)
		path := fmt.Sprintf("http://%s:%d/search?%s", viper.Get("cli.hostname"), viper.GetInt("cli.port"), params)
		payload := &bytes.Buffer{}
		utils.Verbose.Info(fmt.Sprintf("sending request: %s %s %s...\n", http.MethodGet, path, payload))
		req, err := http.NewRequest(http.MethodGet, path, payload)
		if err != nil {
			return utils.Error("\n  %v\n  check configuration to ensure properly configured hostname and port", err, verbose)
		}
		utils.LogRequest(verbose, http.MethodGet, path, payload.String())
		resp, err := getClient.Do(req)
		if err != nil {
			return utils.Error("\n  %v\n  check configuration to ensure properly configured hostname and port\n  or check that api is running", err, verbose)
		}
		defer resp.Body.Close()
		err = utils.CheckResponse(resp)
		if err != nil {
			return utils.Error("\n  checking response: %v", err, verbose)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return utils.Error("\n  reading response: %v\n  ", err, verbose)
		}
		utils.LogHTTPResponse(verbose, resp, b)
		jsonFormat, err := cmd.Flags().GetBool("json")
		_, err = utils.Print(b, jsonFormat)
		if err != nil {
			return utils.Error("\n printing result: %v", err, verbose)
		}
		fmt.Fprintf(cmd.OutOrStdout(), string(b))
		return nil
	},
}

func init() {
	searchCmd.Flags().String("value", "", "string to search for in metadata")
	searchCmd.Flags().Bool("json", false, "return json format")
	rootCmd.AddCommand(searchCmd)
}
