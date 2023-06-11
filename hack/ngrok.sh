#!/bin/bash

# Check if curl is installed
if! command -v curl &> /dev/null
then
    echo "Error: curl is not installed. Please install it first."
    exit 1
fi

# Check if ngrok is installed
if! command -v ngrok &> /dev/null
then
    # Download ngrok binary
    curl -LO https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip
    unzip ngrok-stable-linux-amd64.zip
    sudo mv ngrok /usr/local/bin/
fi

# Start ngrok on port 8555
ngrok http 8555