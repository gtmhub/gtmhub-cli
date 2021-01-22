#!/usr/bin/env bash

: ${BINARY_NAME:="gtmhub"}
: ${GTMHUB_INSTALL_DIR:="/usr/local/bin"}

runAsRoot() {
  local CMD="$*"

  if [ $EUID -ne 0 ]; then
    CMD="sudo $CMD"
  fi

  $CMD
}

downloadZip() {
  DOWNLOAD_URL="https://github.com/gtmhub/gtmhub-cli/releases/latest/download/gtmhub_linux.tar.gz"
  GTMHUB_TMP_ROOT="$(mktemp -dt gtmhub-XXXXXX)"
  GTMHUB_TMP_FILE="$GTMHUB_TMP_ROOT/gtmhub_linux.zip"
  echo "Downloading $DOWNLOAD_URL"
  curl -SsL "$DOWNLOAD_URL" -o "$GTMHUB_TMP_FILE"

}

installFile() {
  GTMHUB_TMP="GTMHUB_TMP_ROOT/$BINARY_NAME"
  mkdir -p "$GTMHUB_TMP"
  tar xf "$GTMHUB_TMP_FILE" -C "$GTMHUB_TMP"
  GTMHUB_TMP_BIN="$GTMHUB_TMP/gtmhub"
  echo "Preparing to install $BINARY_NAME into ${GTMHUB_INSTALL_DIR}"
  runAsRoot cp "$GTMHUB_TMP_BIN" "$GTMHUB_INSTALL_DIR/$BINARY_NAME"
  runAsRoot chmod +x "$GTMHUB_INSTALL_DIR/$BINARY_NAME"
  echo "$BINARY_NAME installed into $GTMHUB_INSTALL_DIR/$BINARY_NAME"
}

downloadZip
installFile