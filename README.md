# gshortcuts

[![Go Version](https://img.shields.io/badge/go-1.24-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`gshortcuts` is a command-line tool to manage your GNOME custom shortcuts. Easily import and reset your custom shortcuts from a configuration file.

## Overview

This tool helps you manage your custom keyboard shortcuts in GNOME-based desktop environments. You can define your shortcuts in a YAML or JSON file and import them using `gshortcuts`. You can find sample configuration files in the `docs/` directory: [sample-shortcuts.yaml](docs/sample-shortcuts.yaml) and [sample-shortcuts.json](docs/sample-shortcuts.json).

## Features

- Import custom shortcuts from a YAML or JSON file.
- Reset existing custom shortcuts.
- Command-line interface for easy integration with scripts.

## Installation

The recommended way to install `gshortcuts` is using `go install`:

```bash
go install github.com/jpinilloslr/gshortcuts/cmd/gshortcuts@latest
```

Ensure your `$GOPATH/bin` or `$HOME/go/bin` directory is in your system's `PATH` to run the installed binary.

### From source (for development or specific versions)

If you prefer to build from source or want to install a specific version from a clone:

1.  Clone the repository:
    ```bash
    git clone https://github.com/jpinilloslr/gshortcuts.git
    cd gshortcuts
    ```
2.  Build and install using Make:
    ```bash
    make install
    ```
    This will build the binary and install it to your Go bin path (usually `$GOPATH/bin` or `$HOME/go/bin`).
    Alternatively, to just build the binary into the `./bin/` directory:
    ```bash
    make build
    ```

## Usage

The main command is `gshortcuts`.

```
A command line tool to manage your custom shortcuts in Gnome
```

### Commands

- `import`: Imports shortcuts from a specified file.

  ```bash
  gshortcuts import /path/to/your/shortcuts.yaml
  gshortcuts import /path/to/your/shortcuts.json
  ```

  Sample files can be found in the `docs/` directory:

  - [`docs/sample-shortcuts.yaml`](docs/sample-shortcuts.yaml)
  - [`docs/sample-shortcuts.json`](docs/sample-shortcuts.json)

- `reset`: Resets all custom shortcuts managed by this tool or specific ones.

For more information on each command, use the `-h` or `--help` flag:

```bash
gshortcuts import --help
gshortcuts reset --help
```

## Building from Source

If you want to build the project from source:

```bash
make build
```

This command will:

1.  Run `go vet` to check for suspicious constructs.
2.  Build the `gshortcuts` binary into the `bin/` directory.

Other `make` commands available:

- `make run`: Runs the application.
- `make clean`: Removes the `bin/` directory.
- `make install`: Builds and installs the application to your Go bin path.

## Configuration File Format

Shortcuts should be defined in a YAML or JSON file. Here's an example structure (refer to [docs/sample-shortcuts.yaml](docs/sample-shortcuts.yaml) or [docs/sample-shortcuts.json](docs/sample-shortcuts.json) for more):

**YAML Example (`shortcuts.yaml`):**

```yaml
- id: "gnome-terminal"
  name: "Open Terminal"
  command: "gnome-terminal"
  binding: "<Super>t"
- id: "firefox"
  name: "Open Browser"
  command: "firefox"
  binding: "<Super>b"
```

**JSON Example (`shortcuts.json`):**

```json
[
  {
    "id": "gnome-terminal",
    "name": "Open Terminal",
    "command": "gnome-terminal",
    "binding": "<Super>t"
  },
  {
    "id": "firefox",
    "name": "Open Browser",
    "command": "firefox",
    "binding": "<Super>b"
  }
]
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
