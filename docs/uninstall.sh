#!/usr/bin/env bash

set -e

BINARY="gshortcuts"
INSTALL_PATH="$HOME/.local/bin/$BINARY"

if [ -f "$INSTALL_PATH" ]; then
  echo "🗑️ Removing $INSTALL_PATH..."
  rm "$INSTALL_PATH"
  echo "✅ Uninstalled $BINARY"
else
  echo "⚠️  $BINARY not found in $INSTALL_PATH"
fi
