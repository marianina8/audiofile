//go:build free || pro

package utils

import (
	"os"
	"os/exec"
	"strings"
)

func Pager(data string) error {
	lessCmd := exec.Command("cmd", "/C", "more")
	lessCmd.Stdin = strings.NewReader(data)
	lessCmd.Stdout = os.Stdout
	lessCmd.Stderr = os.Stderr
	err := lessCmd.Run()
	if err != nil {
		return err
	}
	return nil
}
