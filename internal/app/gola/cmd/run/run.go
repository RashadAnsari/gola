package run

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/RashadAnsari/gola/internal/app/gola/config"
	"github.com/RashadAnsari/gola/internal/app/gola/log"
)

func main(name string) {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal("Failed to load configuration.")
	}

	command := findCommand(cfg.Commands, name)
	if command == nil {
		log.Fatalf("Command %s not found in configuration file.", name)
	}

	shellCmd := shellCommand(*command)

	if err := shellCmd.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

// findCommand finds a command from configuration file.
func findCommand(commands []config.Command, name string) *config.Cmd {
	for _, command := range commands {
		if command.Name == name {
			return &command.Cmd
		}
	}

	return nil
}

// shellCommand generates a new shell command from configuration file.
func shellCommand(command config.Cmd) *exec.Cmd {
	cmd := exec.Command(command.Path, command.Args...)

	cmd.Env = command.Env
	cmd.Dir = command.Dir

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd
}

// Register registers run command for gola binary.
func Register(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run specific gola command",
		Run: func(_ *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatal("Not enough arguments.")
			}

			if len(args) > 1 {
				log.Fatal("Too many arguments.")
			}

			main(args[0])
		},
	}

	cmd.SetUsageTemplate(
		`Usage:
  gola run [command]

Example:
  gola run echo

Flags:
  -h, --help   help for run
`,
	)

	root.AddCommand(cmd)
}
