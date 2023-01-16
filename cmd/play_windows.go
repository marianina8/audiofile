//go:build windows

package cmd

import (
	"github.com/marianina8/audiofile/utils"
	"github.com/pterm/pterm"
	"os/exec"
)

func play(audiofilePath string) error {
	cmd := exec.Command("cmd", "/C", "start", audiofilePath)
	if err := cmd.Start(); err != nil {
		return err
	}
	spinnerInfo := &pterm.SpinnerPrinter{}
	if utils.IsaTTY() {
		spinnerInfo, _ = pterm.DefaultSpinner.Start("Enjoy the music...")
	}
	err := cmd.Wait()
	if err != nil {
		return err
	}
	if utils.IsaTTY() {
		spinnerInfo.Stop()
	}
	return nil
}
