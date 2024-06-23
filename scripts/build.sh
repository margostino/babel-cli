#!/bin/bash

# Ensure the script exits if any command fails
set -e

# Clean up any previous builds
rm -rf ./bin
mkdir -p ./bin

# Build the application
go build -v -o ./bin/babel
cp ./bin/babel $GOBIN

echo "Build completed successfully. The binary is located at ./bin/babel-agent"