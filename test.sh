#!/bin/bash

SRC_DIR="./src"

for file in $(find "$SRC_DIR" -type f -name "*_test.go"); do
    go test -cover "$(dirname "$file")"
done
