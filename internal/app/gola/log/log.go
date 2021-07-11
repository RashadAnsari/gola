package log

import (
	"os"

	"github.com/spf13/cobra"
)

// Cmd is a pointer to root cobra command.
var Cmd *cobra.Command

// Println represents a logging output function.
func Println(i ...interface{}) {
	Cmd.Println(i...)
}

// Fatal represents a logging output function.
func Fatal(i ...interface{}) {
	Cmd.PrintErrln(i...)
	os.Exit(1)
}

// Fatalf represents a logging output function.
func Fatalf(format string, i ...interface{}) {
	Cmd.PrintErrf(format+"\n", i...)
	os.Exit(1)
}
