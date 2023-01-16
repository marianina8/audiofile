//go:build windows

package cmd

import (
	"fmt"
	"os/exec"
)

func play(audiofilePath string, verbose, disableOutput bool) (int, error) {
	cmd := exec.Command("cmd", "/C", "start", audiofilePath)
	if err := cmd.Start(); err != nil {
		return 0, utils.Error("\n  starting start command: %v", err, verbose)
	}
	if !disableOutput {
		spinnerInfo := &pterm.SpinnerPrinter{}
		if utils.IsAtty() {
			spinnerInfo, _ = pterm.DefaultSpinner.Start("Enjoy the music...")
		}	err := cmd.Wait()
		if err != nil {
			return 0, utils.Error("\n  running start command: %v", err, verbose)
		}
		if utils.IsAtty() {
			spinnerInfo.Stop()
		}
		return 0, nil
	} else {
		return cmd.Process.Pid, nil
	}
}
