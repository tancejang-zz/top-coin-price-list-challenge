#!/bin/bash
PORT=8080
DOCKER_PORT=8081
if [ "$1" != "" ] && [ "$2" != "" ]; then
    directory_name="$2"
    cd "$directory_name"
    if [ "$1" = "build" ]; then
        sudo docker build -t "$directory_name" .
    fi
    if [ "$1" = "run" ]; then
        sudo docker run -e COINAPI_API_KEY="A1269882-C37A-49A8-9413-270B3E2FDC5B" -p "$PORT":"$DOCKER_PORT" -it "$directory_name"
    fi
else
    echo "Please specify whether you want to build or run."
fi