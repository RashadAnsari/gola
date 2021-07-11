package main

import (
	"os"

	"github.com/RashadAnsari/gola/internal/app/gola/cmd"
)

func main() {
	root := cmd.NewRootCommand()
	if root != nil {
		if err := root.Execute(); err != nil {
			os.Exit(1)
		}
	}
}
