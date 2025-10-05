package licenses

import "embed"

//go:embed files/*.LICENSE
var Root embed.FS

const DirPrefix = "files"
