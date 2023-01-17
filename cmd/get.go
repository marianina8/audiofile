//go:build free || pro

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/marianina8/audiofile/models"
	"github.com/marianina8/audiofile/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get audio metadata",
	Long:    `Get audio metadata by audiofile id.  Metadata includes available tags and transcript.`,
	Example: `audiofile get --id 45705eba-9342-4952-8cd4-baa2acc25188`,
	RunE: func(cmd *cobra.Command, args []string) error {
		verbose, _ := cmd.Flags().GetBool("verbose")
		b, err := getAudioByID(cmd, verbose)
		if err != nil {
			return err
		}
		jsonFormat, _ := cmd.Flags().GetBool("json")
		if jsonFormat {
			fmt.Fprintf(cmd.OutOrStdout(), string(b))
			fmt.Println(string(b))
		} else {
			var audio models.Audio
			json.Unmarshal(b, &audio)
			tableData, err := audio.Table() // could use another flag here to show more or less detail
			if err != nil {
				return fmt.Errorf("\n  printing table: %v\n  ", err)
			}
			fmt.Println(tableData)
		}
		return nil
	},
}

func init() {
	getCmd.Flags().String("id", "", "audiofile id")
	getCmd.Flags().Bool("json", false, "return json format")
	rootCmd.AddCommand(getCmd)
}

func getAudioByID(cmd *cobra.Command, verbose bool) ([]byte, error) {
	var err error
	id, _ := cmd.Flags().GetString("id")
	if id == "" {
		id, err = utils.AskForID()
		if err != nil {
			return nil, utils.Error("\n  %v\n  try again and enter an id", err, verbose)
		}
	}
	params := "id=" + url.QueryEscape(id)
	path := fmt.Sprintf("http://%s:%d/request?%s", viper.Get("cli.hostname"), viper.Get("cli.port").(int), params)
	payload := &bytes.Buffer{}
	req, err := http.NewRequest(http.MethodGet, path, payload)
	if err != nil {
		return nil, utils.Error("\n  %v\n  check configuration to ensure properly configured hostname and port", err, verbose)
	}
	utils.LogRequest(verbose, http.MethodGet, path, payload.String())
	resp, err := getClient.Do(req)
	if err != nil {
		return nil, utils.Error("\n  %v\n  check configuration to ensure properly configured hostname and port\n  or check that api is running", err, verbose)
	}
	defer resp.Body.Close()
	err = utils.CheckResponse(resp)
	if err != nil {
		return nil, utils.Error("\n  checking response: %v", err, verbose)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, utils.Error("\n  reading response: %v\n  ", err, verbose)
	}
	utils.LogHTTPResponse(verbose, resp, b)
	return b, nil
}
