#!/bin/bash

source .env
source commons.sh

runJupyterContainer() {
    containerRunning=$(docker ps | grep jupyter-container)
    if [[ $containerRunning ]]; then
        echo "$(dateTimeNow) [INFO] - Stopping jupyter-container container"
        docker container stop jupyter-container
    fi

    nodeContainer=$(docker container ls --all | grep jupyter-container)
    if [[ $nodeContainer ]]; then
        docker container rm jupyter-container
    fi

    echo "$(dateTimeNow) [INFO] - Building new jupyter-container container"
    docker build --platform linux/amd64 \
        -f notebooks/Dockerfile \
        --build-arg DATA_URL="$DATA_URL" \
        --build-arg JUPYTER_PORT="$JUPYTER_PORT" \
        -t jupyter-container . || exit

    echo "$(dateTimeNow) [INFO] - Running jupyter-container container"
    docker run -d \
        --platform linux/amd64 \
        --name=jupyter-container \
        -p "$JUPYTER_PORT":"$JUPYTER_PORT" \
        -v "$PWD"/notebooks/src:/app/src \
        -v "$PWD"/data:/app/data \
        --restart unless-stopped \
        jupyter-container || exit

    docker logs -f jupyter-container
}

runJupyterContainer
