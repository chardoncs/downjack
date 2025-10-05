package gitignore

import "embed"

//go:embed files/*.gitignore
var Root embed.FS

const DirPrefix = "files"
