# downjack

A simple ~down jacket~ developer helper tool.

![Demo screenshot](./docs/demo.png)

It can set up `.gitignore` and licenses in your projects with damn-fast speed.

## Install 

### From source

```bash
go install github.com/chardoncs/downjack@latest
```

## Quick Start

For example, if you want to set up `.gitignore` for a Go project:

```bash
downjack gitignore go
# OR simply
downjack g go
# OR fuzzy find
downjack g
```

then create a license file with `MIT` license:

```bash
downjack license mit
# OR simply
downjack l mit
# OR again fuzzy find
downjack l
```

and that's it, your project is now ready to work with!

## Usage

`downjack [command]`

### Commands
- completion - Generate the autocompletion script for the specified shell
- gitignore - Create or append a `.gitignore` file in the project (aliases: g/git/i/ignore)
- license - Add an open source license (aliases: l)
- help - Get help about any command

### Flags
- `-h, --help` - get help for downjack
- `-v, --version` - check the installed version 

Use `downjack [command] --help` for more information about a command.

## Tips

It may be useful to add an alias for `dj` to your shell config (pun intended 📀🤘)

```bash
alias dj='downjack'
```
