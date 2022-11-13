//go:build darwin

package cmd

import (
	"os/exec"

	"github.com/marianina8/audiofile/utils"
	"github.com/pterm/pterm"
)

func play(audiofilePath string, verbose bool) error {
	cmd := exec.Command("afplay", audiofilePath)
	if err := cmd.Start(); err != nil {
		return utils.Error("\n  starting afplay command: %v", err, verbose)
	}
	spinnerInfo := &pterm.SpinnerPrinter{}
	if utils.IsAtty() {
		spinnerInfo, _ = pterm.DefaultSpinner.Start("Enjoy the music...")
	}
	err := cmd.Wait()
	if err != nil {
		return utils.Error("\n  running afplay command: %v", err, verbose)
	}
	if utils.IsAtty() {
		spinnerInfo.Stop()
	}
	return nil
}
