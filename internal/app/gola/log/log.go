package log

import "github.com/spf13/cobra"

// Cmd is a pointer to root cobra command.
var Cmd *cobra.Command

// Print is a convenience method to Print to the defined output, fallback to Stderr if not set.
func Print(i ...interface{}) {
	Cmd.Print(i...)
}

// Println is a convenience method to Println to the defined output, fallback to Stderr if not set.
func Println(i ...interface{}) {
	Cmd.Println(i...)
}

// Printf is a convenience method to Printf to the defined output, fallback to Stderr if not set.
func Printf(format string, i ...interface{}) {
	Cmd.Printf(format, i...)
}

// PrintErr is a convenience method to Print to the defined Err output, fallback to Stderr if not set.
func PrintErr(i ...interface{}) {
	Cmd.PrintErr(i...)
}

// PrintErrln is a convenience method to Println to the defined Err output, fallback to Stderr if not set.
func PrintErrln(i ...interface{}) {
	Cmd.PrintErrln(i...)
}

// PrintErrf is a convenience method to Printf to the defined Err output, fallback to Stderr if not set.
func PrintErrf(format string, i ...interface{}) {
	Cmd.PrintErrf(format, i...)
}
