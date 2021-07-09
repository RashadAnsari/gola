package version

import (
	"github.com/RshadAnsari/gola/internal/app/gola/log"
	"github.com/RshadAnsari/gola/pkg/version"
	"github.com/spf13/cobra"
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
