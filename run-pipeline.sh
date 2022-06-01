#!/bin/bash

source .env
source .env.secrets
source commons.sh

MODE=dev
if [[ $1 != '' ]]; then
    MODE=$1
fi

runPipelineContainer() {
    containerRunning=$(docker ps | grep pipeline-container)
    if [[ $containerRunning ]]; then
        echo "$(dateTimeNow) [INFO] - Stopping pipeline-container container"
        docker container stop pipeline-container
    fi

    nodeContainer=$(docker container ls --all | grep pipeline-container)
    if [[ $nodeContainer ]]; then
        docker container rm pipeline-container
    fi

    echo "$(dateTimeNow) [INFO] - Building new pipeline-container container"
    docker build --platform linux/amd64 \
        -f data_pipeline/Dockerfile \
        --build-arg MODE="$MODE" \
        --build-arg DATA_URL="$DATA_URL" \
        --build-arg MONGO_PASS="$MONGO_PASS" \
        -t pipeline-container . || exit

    echo "$(dateTimeNow) [INFO] - Running pipeline-container container"
    docker run -d \
        --platform linux/amd64 \
        --name=pipeline-container \
        -v "$PWD"/data:/app/data \
        -v "$PWD"/data/info.log:/app/data/info.log \
        pipeline-container || exit

    docker logs -f pipeline-container
}

runPipelineContainer
