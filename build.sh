#!/bin/bash

# Define the source and destination directories
SRC_DIR="./src"
DIST_DIR="./dist"
ASSETS_DIR="./dist/assets"

rm -rf $DIST_DIR

# Loop through each file in the source directory
for file in $(find "$SRC_DIR" -type f ! -name "*.go"); do
    # Get the relative path of the file
    rel_path="${file#$SRC_DIR/}"
    # Create the destination directory if it doesn't exist
    mkdir -p "$(dirname "$ASSETS_DIR/$rel_path")"
    # Copy the file to the destination directory
    cp "$file" "$ASSETS_DIR/$rel_path"
done

go build -o $DIST_DIR/bhakti-buana-api
