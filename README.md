# downjack

A simple ~down jacket~ developer helper tool.

It can set up `.gitignore` and licenses in your projects with damn-fast speed.

## Install 
Use Golang package manager to install downjack, here's the command
```bash
go install github.com/chardoncs/downjack@latest
```

## Quick Start

For example, if you want to set up `.gitignore` for a Go project:

```bash
downjack gitignore go
# OR simply
downjack g go
```

then create a license file with `MIT` license:

```bash
downjack license mit
# OR simply
downjack l mit
```

and that's it, your project is now ready to work with!

## Usage

`downjack [command]`

### Available commands 
- completion - Generate the autocompletion script for the specified shell
- gitignore - Create or append a `.gitignore` file in the project (aliases: g/git/i/ignore)
- license - Add an open source license (aliases: l)
- help - Get help about any command

### Flags 
- `-h, --help` - get help for downjack
- `-v, --version` - check the installed version 

Use `downjack [command] --help` for more information about a command.

## Install - manual way
Go to [releases](https://github.com/Fynjirby/gmfi/releases/) and download latest binary for your OS, then move it to `/usr/local/bin/` and enjoy with simple `gmfi` in terminal!

## Building
- Install [Go](https://go.dev/) and make sure it's working with `go version`
- Clone repo
- Run `go build` in repo directory, then move it to `/usr/local/bin/` or any other directory in your `$PATH`

## Tips

It may be useful to add an alias for `dj` to your shell config (pun intended 📀🤘)

```bash
alias dj='downjack'
```
