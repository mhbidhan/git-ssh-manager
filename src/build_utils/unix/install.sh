#!/bin/bash

go build -o bin/gsm ./src
echo "✅ Build"

if [ ! -d ~/.local/bin/ ]; then
    read -p "Directory ~/.local/bin/ does not exist. Create it? (y/n): " answer
    if [ "$answer" = "y" ] || [ "$answer" = "Y" ]; then
        mkdir -p ~/.local/bin/
        echo "✅ Create ~/.local/bin/ folder"
    else
        echo "Directory not created. Exiting."
        exit 1
    fi
fi

cp bin/gsm ~/.local/bin/
echo "✅ Update gsm"

chmod +x ~/.local/bin/gsm