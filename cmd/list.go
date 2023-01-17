package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/marianina8/audiofile/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all audio files",
	Long: `List audio file metadata in JSON format.  Data includes id, tags, 
and transcript if available.`,
	Example: `audiofile list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		verbose, _ := cmd.Flags().GetBool("verbose")
		client := &http.Client{
			Timeout: 15 * time.Second,
		}
		path := fmt.Sprintf("http://%s:%d/list", viper.Get("cli.hostname"), viper.GetInt("cli.port"))
		payload := &bytes.Buffer{}
		req, err := http.NewRequest(http.MethodGet, path, payload)
		if err != nil {
			return utils.Error("\n  %v\n  check configuration to ensure properly configured hostname and port", err, verbose)
		}
		utils.LogRequest(verbose, http.MethodGet, path, payload.String())
		resp, err := client.Do(req)
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
		jsonFormat, _ := cmd.Flags().GetBool("json")
		err = print(b, jsonFormat)
		return err
	},
}

func init() {
	listCmd.Flags().Bool("json", false, "return json format")
	rootCmd.AddCommand(listCmd)
}
