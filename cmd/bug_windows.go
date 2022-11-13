//go:build windows

package cmd

import (
	"fmt"
	"os/exec"
)

func openBrowser(targetURL string) bool {
	exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", targetURL).Start() == nil
}
