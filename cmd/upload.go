//go:build free || pro

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/marianina8/audiofile/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	checkMark = "\U00002705"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload an audio file",
	Long: `Upload an audio file by passing in the --filename or -f flag followed by the 
filepath of the audiofile.`,
	Example:    `audiofile upload --filename ./audio/beatdoctor.mp3`,
	SuggestFor: []string{"add"},
	RunE: func(cmd *cobra.Command, args []string) error {
		verbose, _ := cmd.Flags().GetBool("verbose")
		var err error
		var p = &pterm.ProgressbarPrinter{}
		if utils.IsAtty() {
			p, _ = pterm.DefaultProgressbar.WithTotal(4).WithTitle("Initiating upload...").Start()
		}
		filename, _ := cmd.Flags().GetString("filename")
		if filename == "" {
			filename, err = utils.AskForFilename()
			if err != nil {
				return utils.Error("\n  %v\n  try again and enter a filename", err, verbose)
			}
		}
		path := fmt.Sprintf("http://%s:%d/upload", viper.Get("cli.hostname"), int(viper.Get("cli.port").(float64)))
		payload := &bytes.Buffer{}
		multipartWriter := multipart.NewWriter(payload)
		file, err := os.Open(filename)
		if err != nil {
			return utils.Error("\n  unable to open file, "+filename+": %v", err, verbose)
		}
		defer file.Close()
		partWriter, err := multipartWriter.CreateFormFile("file", filepath.Base(filename))
		if err != nil {
			return utils.Error("\n  %v\n  problems preparing file for upload", err, verbose)
		}

		_, err = io.Copy(partWriter, file)
		if err != nil {
			return utils.Error("\n  %v\n  problems preparing file for upload", err, verbose)
		}
		if utils.IsAtty() && runtime.GOOS != "windows" {
			p.UpdateTitle("Creating multipart writer...")
		} else {
			fmt.Println("Creating multipart writer...")
		}
		err = multipartWriter.Close()
		if err != nil {
			return utils.Error("\n  %v\n  problems preparing file for upload", err, verbose)
		}
		if utils.IsAtty() && runtime.GOOS != "windows" {
			pterm.Success.Println("Created multipart writer")
			p.Increment()
			p.UpdateTitle("Sending request...")
		} else {
			fmt.Println("Sending request...")
		}
		req, err := http.NewRequest(http.MethodPost, path, payload)
		if err != nil {
			return utils.Error("\n  %v\n  check configuration to ensure properly configured hostname and port", err, verbose)
		}
		utils.LogRequest(verbose, http.MethodPost, path, payload.String())
		req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
		if utils.IsAtty() && runtime.GOOS != "windows" {
			pterm.Success.Printf("Sending request: %s %s...", http.MethodPost, path)
			p.Increment()
		} else {
			fmt.Printf("Sending request: %s %s...\n", http.MethodPost, path)
		}
		resp, err := getClient.Do(req)
		if err != nil {
			return utils.Error("\n  %v\n  check configuration to ensure properly configured hostname and port\n  or check that api is running", err, verbose)
		}
		defer resp.Body.Close()
		if utils.IsAtty() && runtime.GOOS != "windows" {
			p.UpdateTitle("Receive response...")
			pterm.Success.Println("Receive response...")
			p.Increment()
		} else {
			fmt.Println("Receive response...")
		}
		err = utils.CheckResponse(resp)
		if err != nil {
			return utils.Error("\n  checking response: %v", err, verbose)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return utils.Error("\n  reading response: %v\n  ", err, verbose)
		}
		utils.LogHTTPResponse(verbose, resp, b)
		if utils.IsAtty() && runtime.GOOS != "windows" {
			p.UpdateTitle("Process response...")
			pterm.Success.Println("Process response...")
			p.Increment()
		} else {
			fmt.Println("Receive response...")
		}
		if utils.IsAtty() && runtime.GOOS != "windows" {
			fmt.Fprintf(cmd.OutOrStdout(), fmt.Sprintf(" Successfully uploaded!\n Audiofile ID: %s", string(b)))
			fmt.Println(checkMark, " Successfully uploaded!")
			fmt.Println(checkMark, " Audiofile ID: ", string(b))
		} else {
			//fmt.Fprintf(cmd.OutOrStdout(), string(b))
			fmt.Println("ID: ", string(b))
		}
		return nil
	},
}

func init() {
	uploadCmd.Flags().StringP("filename", "f", "", "Filepath of filename to be uploaded")
	rootCmd.AddCommand(uploadCmd)
}
