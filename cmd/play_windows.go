//go:build windows

package cmd

import (
	"fmt"
	"os/exec"
)

func play(audiofilePath string, verbose bool) error {
	cmd := exec.Command("start", audiofilePath)
	if err := cmd.Start(); err != nil {
		return utils.Error("\n  starting start command: %v", err, verbose)
	}
	spinnerInfo := &pterm.SpinnerPrinter{}
	if utils.IsAtty() {
		spinnerInfo, _ = pterm.DefaultSpinner.Start("Enjoy the music...")
	}	err := cmd.Wait()
	if err != nil {
		return utils.Error("\n  running start command: %v", err, verbose)
	}
	if utils.IsAtty() {
		spinnerInfo.Stop()
	}
	return nil
}
