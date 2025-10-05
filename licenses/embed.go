package licenses

import "embed"

//go:embed files/*.LICENSE
var root embed.FS

const DirPrefix = "files"
