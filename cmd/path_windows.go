package cmd

import (
	"os"
	"path/filepath"
)

func Path() string {
	return filepath.Join(os.Getenv("APPDATA"), "tic_toc_time_stack")
}
