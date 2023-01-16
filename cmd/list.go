package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/marianina8/audiofile/models"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{
			Timeout: 15 * time.Second,
		}
		path := fmt.Sprintf("http://%s:%d/list", viper.Get("cli.hostname"), int(viper.Get("cli.port").(float64)))
		payload := &bytes.Buffer{}
		req, err := http.NewRequest(http.MethodGet, path, payload)
		if err != nil {
			return err
		}
		fmt.Printf("Sending request: %s %s %s...\n", http.MethodGet, path, payload)
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		err = utils.CheckResponse(resp)
		if err != nil {
			return err
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		jsonFormat, _ := cmd.Flags().GetBool("json")
		plainFormat, _ := cmd.Flags().GetBool("plain")
		if jsonFormat {
			if utils.IsaTTY() {
				err = utils.Pager(string(b))
				if err != nil {
					return err
				}
			} else {
				fmt.Println(string(b))
			}
		} else if plainFormat {
			var audios models.AudioList
			json.Unmarshal(b, &audios)
			if utils.IsaTTY() {
				err = utils.Pager(audios.Plain())
				if err != nil {
					return err
				}
			} else {
				fmt.Println(audios.Plain())
			}
		} else {
			var audios models.AudioList
			json.Unmarshal(b, &audios)
			tableData, err := audios.Table()
			if err != nil {
				return err
			}
			if utils.IsaTTY() {
				err = utils.Pager(tableData)
				if err != nil {
					return err
				}
			} else {
				fmt.Println(tableData)
			}
		}
		return nil
	},
}

func init() {
	listCmd.Flags().Bool("json", false, "return json format")
	listCmd.Flags().Bool("plain", false, "return plain format")
	rootCmd.AddCommand(listCmd)
}
