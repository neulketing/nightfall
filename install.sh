#!/bin/bash
set -e

echo "Installing Nightfall dependencies..."

# Check for Go
if ! command -v go &> /dev/null; then
    echo "Go not found. Installing..."
    brew install go
fi

# Check for Python
if ! command -v python3 &> /dev/null; then
    echo "Python not found. Installing..."
    brew install python3
fi

# Check for Rust
if ! command -v rustc &> /dev/null; then
    echo "Rust not found. Installing..."
    curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
fi

echo "Building Nightfall..."
make build

echo "Installing Nightfall..."
make install

echo "Installation complete!"
