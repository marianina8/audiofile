/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
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

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Command to search for audiofiles by string",
	Long:  `Command to search for audiofiles by search string within the metadata file.  Search string is not case sensitive`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{
			Timeout: 15 * time.Second,
		}
		var err error
		value, _ := cmd.Flags().GetString("value")
		if value == "" {
			value, err = utils.AskForValue()
			if err != nil {
				return err
			}
		}
		params := "searchFor=" + url.QueryEscape(value)
		path := fmt.Sprintf("http://%s:%d/search?%s", viper.Get("cli.hostname"), int(viper.Get("cli.port").(float64)), params)
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
	searchCmd.Flags().String("value", "", "string to search for in metadata")
	searchCmd.Flags().Bool("json", false, "return json format")
	searchCmd.Flags().Bool("plain", false, "return plain format")
	rootCmd.AddCommand(searchCmd)
}
