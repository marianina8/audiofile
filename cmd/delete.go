/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/marianina8/audiofile/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete audiofile by id",
	Long:  `Delete audiofile by id. This command removes the entire folder containing all stored metadata`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{
			Timeout: 15 * time.Second,
		}
		var err error
		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			id, err = utils.AskForID()
			if err != nil {
				return err
			}
		}
		params := "id=" + url.QueryEscape(id)
		path := fmt.Sprintf("http://%s:%d/delete?%s", viper.Get("cli.hostname"), int(viper.Get("cli.port").(float64)), params)
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
		if strings.Contains(string(b), "success") {
			fmt.Printf("\U00002705 Successfully deleted audiofile (%s)!\n", id)
		} else {
			fmt.Printf("\U0000274C Unsuccessful delete of audiofile (%s): %s\n", id, string(b))
		}
		return nil
	},
}

func init() {
	deleteCmd.Flags().String("id", "", "audiofile id")
	rootCmd.AddCommand(deleteCmd)
}
