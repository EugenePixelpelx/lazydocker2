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

`/` is contextual: it filters the focused side panel or searches the focused main
tab. Press `Enter` to confirm a search and move to the next result, `n`/`N` to move
forward/backward, and `Esc` to clear it.

## Requirements

- Docker >= 29.0.0 (API >= 1.24)
- Docker Compose >= 1.23.2 (optional)

## Installation

Release binaries are not published yet. To build and run the project locally,
install Go and run:

```bash
GOFLAGS=-mod=vendor go build -o lazydocker2 .
./lazydocker2
```

Published release binaries and installation instructions will be added here.

## Documentation

- [Configuration](docs/Config.md)
- [Keybindings](docs/keybindings/Keybindings_en.md)
- [Upstream repository](https://github.com/jesseduffield/lazydocker)

## License

This project is licensed under the [MIT License](LICENSE). It includes work
copyrighted by the original lazydocker contributors.
