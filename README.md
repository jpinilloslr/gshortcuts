# gshortcuts

[![Go Version](https://img.shields.io/badge/go-1.24-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://github.com/jpinilloslr/gshortcuts/actions/workflows/build.yaml/badge.svg?branch=master)](https://github.com/jpinilloslr/gshortcuts/actions/workflows/build.yaml)

## Table of Contents

* [Overview](#overview)
* [Features](#features)
* [Installation](#installation)
* [Usage](#usage)
* [Configuration File Format](#configuration-file-format)
* [License](#license)

`gshortcuts` is a command-line tool to manage your GNOME custom shortcuts. Easily import and export your custom shortcuts using YAML or JSON files.

## Overview

This tool helps you manage your custom keyboard shortcuts in GNOME-based desktop environments. You can define your shortcuts in a YAML or JSON file and import them using `gshortcuts`. You can find sample configuration files in the `docs/` directory: [sample-shortcuts.yaml](docs/sample-shortcuts.yaml) and [sample-shortcuts.json](docs/sample-shortcuts.json).

## Features

* Consistent, idempotent management of GNOME custom shortcuts via declarative config files.
* Native, direct GIO Settings integration for true GNOME compatibility (no external binaries).
* Supports both YAML and JSON for flexible workflows.
* Git-friendly: track and version your shortcut definitions alongside dotfiles.
* Automatable: integrate into provisioning scripts.
* Minimal external dependencies, easy to install.
* Configurable import strategies (replace or merge).

## Installation

### Standalone binary

```bash
curl -sSL https://gshortcuts.jpinillos.dev/install.sh | bash
```

### Alternative: Install via Go

If you have Go installed and prefer to install from source or use a pinned version:

```bash
go install github.com/jpinilloslr/gshortcuts/cmd/gshortcuts@latest
```

### Alternative: Install from source

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
A command line tool to manage your custom shortcuts in GNOME
```

### Commands

- `import`: Imports custom shortcuts from a specified file.

  ```bash
  gshortcuts import /path/to/your/shortcuts.yaml
  gshortcuts import /path/to/your/shortcuts.json
  ```

  Sample files can be found in the `docs/` directory:

  - [`docs/sample-shortcuts.yaml`](docs/sample-shortcuts.yaml)
  - [`docs/sample-shortcuts.json`](docs/sample-shortcuts.json)

- `export`: Exports custom shortcuts to a specified file.

  ```bash
  gshortcuts export /path/to/your/shortcuts.yaml
  gshortcuts export /path/to/your/shortcuts.json
  ```

- `reset`: Resets all custom shortcuts.

For more information on each command, use the `-h` or `--help` flag:

```bash
gshortcuts import --help
gshortcuts export --help
gshortcuts reset --help
```

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
