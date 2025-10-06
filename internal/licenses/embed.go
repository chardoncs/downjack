package licenses

import "embed"

//go:embed files/*.txt files/*.tmpl
var Root embed.FS

const DirPrefix = "files"
