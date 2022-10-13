package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/marianina8/audiofile/models"
	"github.com/marianina8/audiofile/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get audio metadata",
	Long:  `Get audio metadata by audiofile id.  Metadata includes available tags and transcript.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := getAudioByID(cmd)
		if err != nil {
			return err
		}
		jsonFormat, _ := cmd.Flags().GetBool("json")
		if jsonFormat {
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

func getAudioByID(cmd *cobra.Command) ([]byte, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	var err error
	id, _ := cmd.Flags().GetString("id")
	if id == "" {
		id, err = utils.AskForID()
		if err != nil {
			return nil, err
		}
	}
	params := "id=" + url.QueryEscape(id)
	path := fmt.Sprintf("http://%s:%d/request?%s", viper.Get("cli.hostname"), int(viper.Get("cli.port").(float64)), params)
	payload := &bytes.Buffer{}

	req, err := http.NewRequest(http.MethodGet, path, payload)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Sending request: %s %s %s...\n", http.MethodGet, path, payload)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = utils.CheckResponse(resp)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
