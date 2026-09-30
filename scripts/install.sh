#!/bin/sh
# Installs or updates lazydocker2 from the latest GitHub release.
#
#   curl -fsSL https://raw.githubusercontent.com/EugenePixelpelx/lazydocker2/master/scripts/install.sh | sh
#
# Set DIR to install somewhere else, e.g. DIR="$HOME/.local/bin".
set -eu

REPO="EugenePixelpelx/lazydocker2"
DIR="${DIR:-/usr/local/bin}"

OS=$(uname -s)
case "$OS" in
    Linux|Darwin) ;;
    *) echo "Unsupported OS: $OS" >&2; exit 1 ;;
esac

# map different architecture variations to the available binaries
ARCH=$(uname -m)
case "$ARCH" in
    x86_64|amd64) ARCH=x86_64 ;;
    i386|i686) ARCH=x86 ;;
    aarch64*|arm64) ARCH=arm64 ;;
    armv6*) ARCH=armv6 ;;
    armv7*) ARCH=armv7 ;;
    *) echo "Unsupported architecture: $ARCH" >&2; exit 1 ;;
esac

FILE="lazydocker2_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/latest/download/${FILE}"

TMP=$(mktemp -d)
trap 'rm -rf "$TMP"' EXIT

echo "Downloading $URL"
curl -fsSL -o "$TMP/$FILE" "$URL"
tar -xzf "$TMP/$FILE" -C "$TMP" lazydocker2

# use sudo only when the destination is not writable by the current user
SUDO=""
if ! mkdir -p "$DIR" 2>/dev/null || [ ! -w "$DIR" ]; then
    SUDO="sudo"
fi
$SUDO mkdir -p "$DIR"
$SUDO install -m 755 "$TMP/lazydocker2" "$DIR/lazydocker2"

echo "Installed: $("$DIR/lazydocker2" --version | head -n 1) -> $DIR/lazydocker2"
