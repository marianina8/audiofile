//go:build windows

package cmd

import (
	"fmt"
	"os/exec"
)

func play(audiofilePath string) error {
	cmd := exec.Command("start", audiofilePath)
	if err := cmd.Start(); err != nil {
		return err
	}
	spinnerInfo := &pterm.SpinnerPrinter{}
	if utils.IsAtty() {
		spinnerInfo, _ = pterm.DefaultSpinner.Start("Enjoy the music...")
	}	err := cmd.Wait()
	if err != nil {
		return err
	}
	if utils.IsAtty() {
		spinnerInfo.Stop()
	}
	return nil
}
