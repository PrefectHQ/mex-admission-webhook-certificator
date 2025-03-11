package version

import (
	"fmt"
)

// Version represents main version number being run right now.
var Version = "v0.1.0"

// ReleasePhase represents pre-release marker for the version. If this is an empty string,
// then the release is a final release. Otherwise this is a pre-release
// version e.g. "dev", "alpha", etc.
var ReleasePhase = ""

// String prints the version of the dha CLI.
func String() string {
	if ReleasePhase != "" {
		return fmt.Sprintf("%s-%s", Version, ReleasePhase)
	}
	return Version
}
