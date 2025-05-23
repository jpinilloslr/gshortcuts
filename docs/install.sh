#!/usr/bin/env bash

set -e

REPO="jpinilloslr/gshortcuts"
BINARY="gshortcuts"
INSTALL_DIR="$HOME/.local/bin"

if [ "$(uname)" != "Linux" ]; then
  echo "❌ This tool only works on Linux."
  exit 1
fi

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
  x86_64) ARCH=amd64 ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

VERSION=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" \
    | grep tag_name | cut -d '"' -f4)

echo "Installing $BINARY $VERSION for $OS/$ARCH..."
URL="https://github.com/$REPO/releases/download/$VERSION/${BINARY}-${OS}-${ARCH}"

mkdir -p "$INSTALL_DIR"

curl -L -o "$INSTALL_DIR/$BINARY" "$URL"
chmod +x "$INSTALL_DIR/$BINARY"

if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
  echo "⚠️  $INSTALL_DIR is not in your PATH."
  echo "Add this to your shell config, e.g.:"
  echo "  export PATH=\"\$HOME/.local/bin:\$PATH\""
else
  echo "✅ Installed: $INSTALL_DIR/$BINARY"
fi
