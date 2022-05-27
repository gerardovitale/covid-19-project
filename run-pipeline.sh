#!/bin/bash

source .env
source commons.sh

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
        --build-arg DATA_URL=$DATA_URL \
        -t pipeline-container . || exit

    echo "$(dateTimeNow) [INFO] - Running pipeline-container container"
    docker run -d \
        --platform linux/amd64 \
        --name=pipeline-container \
        -v $PWD/data:/app/data \
        pipeline-container || exit

    docker logs -f pipeline-container
}

runPipelineContainer