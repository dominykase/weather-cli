#!/bin/bash

# Set the source and binary directories
src_dir="src"
bin_dir="bin"
output_name="weather"

if [ ! -d "$src_dir" ]; then
    echo "Error: Source directory '$src_dir' not found."
    exit 1
fi

if [ ! -d "$bin_dir" ]; then
    mkdir -p "$bin_dir"
fi

go build -o "$bin_dir/$output_name" "$src_dir/main.go"

if [ $? -eq 0 ]; then
    echo "Compilation successful. Binary saved to '$bin_dir/$output_name'."
else
    echo "Compilation failed."
fi
