# lazydocker2

`lazydocker2` is a fork of [lazydocker](https://github.com/jesseduffield/lazydocker),
the terminal UI for managing Docker and Docker Compose environments.

It preserves the familiar lazydocker workflow while making selected usability
fixes available independently of upstream releases.

## Changes from upstream

- Main-panel tabs remain clickable and receive focus when selected with the mouse.
- The selected resource remains highlighted while viewing its details or using
  search and filters.
- Search in the main panel scans all loaded content, not only the text currently
  visible in the terminal.
- Search and filter prompts identify their target, and the corresponding panel
  remains visually active.
- Clicking another panel or switching tabs clears the current main-panel search.
- The main panel has a predictable fullscreen toggle for focused viewing and
  easier terminal text selection.

`/` is contextual: it filters the focused side panel or searches the focused main
tab. Press `Enter` to confirm a search and move to the next result, `n`/`N` to move
forward/backward, and `Esc` to clear it.

Press `z` while the main panel is focused to enter fullscreen, and `z` again to
restore the previous layout.

## Requirements

- Docker >= 29.0.0 (API >= 1.24)
- Docker Compose >= 1.23.2 (optional)

## Installation

### Quick install (Linux, macOS)

1. Install (or update to the latest version):

   ```bash
   curl -fsSL https://raw.githubusercontent.com/EugenePixelpelx/lazydocker2/master/scripts/install.sh | sh
   ```

2. Run:

   ```bash
   lazydocker2
   ```

That's it. The script detects your OS and CPU, downloads the latest release and
puts `lazydocker2` into `/usr/local/bin` (asking for your `sudo` password if
needed). Run the same command again to update.

### Other platforms

Binaries for all platforms, including Windows, are available on the
[Releases](https://github.com/EugenePixelpelx/lazydocker2/releases) page.

`lazydocker2` reads the same config file as lazydocker, and both can be installed
side by side.

### Building from source

Install Go and run:

```bash
GOFLAGS=-mod=vendor go build -o lazydocker2 .
./lazydocker2
```

## Documentation

- [Configuration](docs/Config.md)
- [Keybindings](docs/keybindings/Keybindings_en.md)
- [Upstream repository](https://github.com/jesseduffield/lazydocker)

## License

This project is licensed under the [MIT License](LICENSE). It includes work
copyrighted by the original lazydocker contributors.
