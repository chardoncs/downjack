package gitignore

import "embed"

//go:embed files/*.gitignore
var root embed.FS
