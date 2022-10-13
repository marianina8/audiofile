package utils

import (
	"os"

	isatty "github.com/mattn/go-isatty"
)

func IsAtty() bool {
	if isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		return true
	}
	return false
}
