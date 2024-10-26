#!/bin/bash

APP_NAME="bebra"

PLATFORMS=("linux/amd64" "windows/amd64")

mkdir -p bin

for PLATFORM in "${PLATFORMS[@]}"; do
    IFS="/" read -r OS ARCH <<< "$PLATFORM"

    OUTPUT_DIR="bin/$APP_NAME-$OS-$ARCH"

    mkdir -p "$OUTPUT_DIR"

    # Determine the appropriate file extension
    if [[ "$OS" == "windows" ]]; then
        EXT=".exe"
    else
        EXT=""
    fi

    GOOS=$OS GOARCH=$ARCH go build -o "$OUTPUT_DIR/$APP_NAME$EXT"

    echo "Built $APP_NAME for $OS/$ARCH in $OUTPUT_DIR"
done

echo "All builds completed."
