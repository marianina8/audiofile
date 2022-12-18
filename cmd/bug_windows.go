//go:build windows && (free || pro)

package cmd

import (
	"os/exec"
)

func openBrowser(targetURL string) bool {
	exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", targetURL).Start() == nil
}
