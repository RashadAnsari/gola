package cmd

import (
	"github.com/spf13/cobra"

	"github.com/RashadAnsari/gola/internal/app/gola/cmd/commands"

	"github.com/RashadAnsari/gola/internal/app/gola/cmd/run"

	"github.com/RashadAnsari/gola/internal/app/gola/cmd/version"
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

	version.Register(root)
	run.Register(root)
	commands.Register(root)

	return root
}
