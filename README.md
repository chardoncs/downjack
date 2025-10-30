# downjack

A simple ~down jacket~ developer helper tool.

![Demo screenshot](./docs/demo.png)

It can set up `.gitignore` and licenses in your projects with damn-fast speed.

## Usage

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

## Install 

### Binary

Go to the [release page](https://github.com/chardoncs/downjack/releases) and find the binary for your OS.

### Arch Linux (btw)

We provide AURs
([`downjack`](https://aur.archlinux.org/packages/downjack)
and
[`downjack-bin`](https://aur.archlinux.org/packages/downjack-bin))
for Arch users.

For convenience, you may use an AUR helper:

```bash
yay -S downjack-bin
# OR
paru -S downjack-bin
```

### Go

```bash
go install github.com/chardoncs/downjack@latest
```

## Tips

It may be useful to add an alias for `dj` to your shell config (pun intended 📀🤘)

```bash
alias dj='downjack'
```
