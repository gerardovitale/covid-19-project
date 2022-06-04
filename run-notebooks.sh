#!/bin/bash

source .env
source commons.sh

SERVICE_NAME=jupyter-container

runJupyterContainer() {
    cleanUpDocker

    echo "$(dateTimeNow) - [INFO] - Building new $SERVICE_NAME"
    docker build --platform linux/amd64 \
        -f notebooks/Dockerfile \
        --build-arg DATA_URL="$DATA_URL" \
        --build-arg JUPYTER_PORT="$JUPYTER_PORT" \
        -t "$SERVICE_NAME" . || exit

    echo "$(dateTimeNow) - [INFO] - Running $SERVICE_NAME"
    docker run -d \
        --platform linux/amd64 \
        --name="$SERVICE_NAME" \
        -p "$JUPYTER_PORT":"$JUPYTER_PORT" \
        -v "$PWD"/notebooks/src:/app/src \
        -v "$PWD"/data:/app/data \
        --restart unless-stopped \
        "$SERVICE_NAME" || exit

    docker logs -f "$SERVICE_NAME"
}

runJupyterContainer
