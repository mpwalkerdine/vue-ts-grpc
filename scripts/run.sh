#!/usr/bin/env bash

set -e

UIPORT=8080
SERVERPORT=8081
GRPCPORT=8082

# Go "up one" to root of the repository
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$DIR"

killPort() {
    local pid=$(lsof -t -i ":$1")
    if [ -n "$pid" ]; then
        echo killing process $pid on port $1
        kill -SIGKILL $pid
    fi
}

# Kill off any old processes
killPort $UIPORT
killPort $SERVERPORT
killPort $GRPCPORT

# Spawn backend server and exit when frontend dev server is killed with this script
trap "killPort $SERVERPORT" EXIT
bin/example-server -webport $SERVERPORT -port $GRPCPORT &

# Start the frontend dev server
cd $DIR/client
npm run serve -- --port $UIPORT
