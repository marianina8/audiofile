//go:build windows

package cmd

import (
	"fmt"
	"github.com/marianina8/audiofile/utils"
	"github.com/pterm/pterm"
	"os/exec"
	"runtime"
)

func play(audiofilePath string, verbose, disableOutput bool) (int, error) {
	cmd := exec.Command("cmd", "/C", "start", audiofilePath)
	if err := cmd.Start(); err != nil {
		return 0, utils.Error("\n  starting start command: %v", err, verbose)
	}
	if !disableOutput {
		spinnerInfo := &pterm.SpinnerPrinter{}
		if utils.IsaTTY() && runtime.GOOS != "windows" {
			spinnerInfo, _ = pterm.DefaultSpinner.Start("Enjoy the music...")
		}
		if runtime.GOOS == "windows" {
			fmt.Println("Enjoy the music...")
		}
		err := cmd.Wait()
		if err != nil {
			return 0, utils.Error("\n  running start command: %v", err, verbose)
		}
		if utils.IsaTTY() && runtime.GOOS != "windows" {
			spinnerInfo.Stop()
		}
		return 0, nil
	} else {
		return cmd.Process.Pid, nil
	}
}
