package commands

import (
	"github.com/spf13/cobra"

	"github.com/RashadAnsari/gola/internal/app/gola/config"
	"github.com/RashadAnsari/gola/internal/app/gola/log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal("Failed to load configuration.")
	}

	for _, command := range cfg.Commands {
		log.Println(command.Name)
	}
}

// Register registers commands command for gola binary.
func Register(root *cobra.Command) {
	root.AddCommand(
		&cobra.Command{
			Use:   "commands",
			Short: "Print all available gola commands",
			Run: func(_ *cobra.Command, args []string) {
				main()
			},
		},
	)
}
