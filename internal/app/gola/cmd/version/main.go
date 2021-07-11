package version

import (
	"github.com/spf13/cobra"

	"github.com/RashadAnsari/gola/internal/app/gola/log"
	"github.com/RashadAnsari/gola/pkg/version"
)

// Register registers version command for gola binary.
func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "version",
			Short: "Print gola version",
			Run: func(_ *cobra.Command, _ []string) {
				log.Println(version.String())
			},
		},
	)
}
