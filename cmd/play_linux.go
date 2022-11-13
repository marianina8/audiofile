//go:build linux

package cmd

import (
	"fmt"
	"os/exec"
)

func play(audiofilePath string, verbose bool) error {
	cmd := exec.Command("aplay", audiofilePath)
	if err := cmd.Start(); err != nil {
		return utils.Error("\n  starting aplay command: %v", err, verbose)
	}
	spinnerInfo := &pterm.SpinnerPrinter{}
	if utils.IsAtty() {
		spinnerInfo, _ = pterm.DefaultSpinner.Start("Enjoy the music...")
	}	err := cmd.Wait()
	if err != nil {
		return utils.Error("\n  running aplay command: %v", err, verbose)
	}
	if utils.IsAtty() {
		spinnerInfo.Stop()
	}
	return nil
}
