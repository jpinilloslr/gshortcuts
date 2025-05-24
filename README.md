# gshortcuts

[![Go Version](https://img.shields.io/badge/go-1.24-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://github.com/jpinilloslr/gshortcuts/actions/workflows/build.yaml/badge.svg?branch=master)](https://github.com/jpinilloslr/gshortcuts/actions/workflows/build.yaml)

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Dependencies](#dependencies)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration File Format](#configuration-file-format)
- [License](#license)

`gshortcuts` is a command-line tool to import, export, and declaratively manage GNOME built-in and custom keyboard shortcuts.

## Overview

This tool helps you manage your keyboard shortcuts in GNOME-based desktop environments. You can define your shortcuts in a YAML or JSON file and import them using `gshortcuts`.

## Features

- Reliable and idempotent.
- Manage both, built-in and custom GNOME shortcuts.
- Native GIO Settings integration (no external processes).
- No need to restart your GNOME session, changes take effect immediately.
- Declarative config format: track changes easily with Git.
- Supports both YAML and JSON.
- Compatible with provisioning scripts or system setup automation.
- Avoids manual DConf branch management, no need to export and merge multiple DConf paths when syncing shortcut settings.

## Dependencies

To build `gshortcuts` from source, you need GLib/GIO development headers installed.

Debian/Ubuntu:

```bash
sudo apt install libglib2.0-dev pkg-config
```

Fedora:

```bash
sudo dnf install glib2-devel pkgconf-pkg-config
```

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

> ⚠️ Requires development headers for GLib and GIO.  
> On Debian/Ubuntu: `sudo apt install libglib2.0-dev pkg-config`
> On Fedora: sudo dnf install glib2-devel pkgconf-pkg-config

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

### Subcommands

- `import`: Imports shortcuts from a specified file.

  ```bash
  gshortcuts import /path/to/your/shortcuts.yaml
  gshortcuts import /path/to/your/shortcuts.json
  ```

- `export`: Exports shortcuts to a specified file.

  ```bash
  gshortcuts export /path/to/your/shortcuts.yaml
  gshortcuts export /path/to/your/shortcuts.json
  ```

- `reset-custom`: Resets all custom shortcuts.

For more information on each command, use the `-h` or `--help` flag:

```bash
gshortcuts --help
gshortcuts <subcommand> --help
```

## Configuration File Format

Shortcuts should be defined in a YAML or JSON file. Here's an example structure (refer to [docs/sample-shortcuts.yaml](docs/sample-shortcuts.yaml) or [docs/sample-shortcuts.json](docs/sample-shortcuts.json) for more):

**YAML Example:**

```yaml
builtin:
  org.gnome.desktop.wm.keybindings:
    - key: switch-to-workspace-1
      bindings:
        - <Super>1
    - key: switch-to-workspace-2
      bindings:
        - <Super>2
    - key: switch-to-workspace-3
      bindings:
        - <Super>3
    - key: switch-to-workspace-4
      bindings:
        - <Super>4
  org.gnome.shell.keybindings:
    - key: switch-to-application-1
      bindings: []
    - key: switch-to-application-2
      bindings: []
    - key: switch-to-application-3
      bindings: []
    - key: switch-to-application-4
      bindings: []
custom:
  - id: gnome-terminal
    name: Open Terminal
    binding: <Super>t
    command: gnome-terminal
  - id: firefox
    name: Open Browser
    binding: <Super>b
    command: firefox
```

**JSON Example:**

```json
{
  "builtIn": {
    "org.gnome.desktop.wm.keybindings": [
      { "Key": "switch-to-workspace-1", "Bindings": ["<Super>1"] },
      { "Key": "switch-to-workspace-2", "Bindings": ["<Super>2"] },
      { "Key": "switch-to-workspace-3", "Bindings": ["<Super>3"] },
      { "Key": "switch-to-workspace-4", "Bindings": ["<Super>4"] }
    ],
    "org.gnome.shell.keybindings": [
      { "Key": "switch-to-application-1", "Bindings": null },
      { "Key": "switch-to-application-2", "Bindings": null },
      { "Key": "switch-to-application-3", "Bindings": null },
      { "Key": "switch-to-application-4", "Bindings": null }
    ]
  },
  "custom": [
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
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
