package version

import "runtime/debug"

var Version = "unknown"

func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	mainVersion := info.Main.Version
	if mainVersion == "" || mainVersion == "(devel)" {
		// non-standard build
		return
	}

	Version = mainVersion
}
