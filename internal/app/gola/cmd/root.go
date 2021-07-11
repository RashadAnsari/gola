package cmd

import (
	"github.com/spf13/cobra"

	"github.com/RashadAnsari/gola/internal/app/gola/cmd/version"
	"github.com/RashadAnsari/gola/internal/app/gola/config"
	"github.com/RashadAnsari/gola/internal/app/gola/log"
)

// NewRootCommand creates a new gola root command.
func NewRootCommand() *cobra.Command {
	var root = &cobra.Command{
		Use:   "gola",
		Short: "Gola is a tool for automated scripting purpose.",
	}

	root.CompletionOptions.DisableDefaultCmd = true

	// Set Cmd global variable for logging.
	log.Cmd = root

	_, err := config.Init()
	if err != nil {
		log.PrintErr(err.Error())
	}

	version.Register(root)

	return root
}
