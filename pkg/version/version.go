package version

import (
	"fmt"
)

// Following variables are set via -ldflags.
var (
	// AppVersion represents the semantic version of the app.
	AppVersion string
	// VCSRef represents name of branch at build time.
	VCSRef string
	// Version represents git SHA at build time.
	BuildVersion string
	// Date represents the time of build.
	Date string
)

// String generates a string for representing the version of application.
func String() string {
	return fmt.Sprintf(
		"AppVersion = %s\n"+
			"VCSRef = %s\n"+
			"BuildVersion = %s\n"+
			"BuildDate = %s",
		AppVersion, VCSRef, BuildVersion, Date,
	)
}
