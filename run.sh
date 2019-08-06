#!/bin/bash
if [ "$1" != "" ] && [ "$2" != "" ]; then
    directory_name="$2"
    cd "$directory_name"
    if [ "$1" = "build" ]; then
        sudo docker build -t "$directory_name" .
    fi
    if [ "$1" = "run" ]; then
        sudo docker-compose up
    fi
else
    echo "Please specify whether you want to build or run."
fi
