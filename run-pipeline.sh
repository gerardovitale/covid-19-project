#!/bin/bash

source .env
source .env.secrets
source commons.sh

SERVICE_NAME=pipeline-container
MODE=dev
if [[ $1 != '' ]]; then
    MODE=$1
fi

runPipelineContainer() {
    cleanUpDocker $SERVICE_NAME

    echo "$(dateTimeNow) - [INFO] - Building new $SERVICE_NAME"
    docker build --platform linux/amd64 \
        -f data_pipeline/Dockerfile \
        --build-arg MODE="$MODE" \
        --build-arg DATA_URL="$DATA_URL" \
        --build-arg MONGO_PASS="$MONGO_PASS" \
        -t "$SERVICE_NAME" . || exit

    echo "$(dateTimeNow) - [INFO] - Running $SERVICE_NAME"
    docker run -d \
        --platform linux/amd64 \
        --name="$SERVICE_NAME" \
        -v "$PWD"/data:/app/data \
        -v "$PWD"/data/info.log:/app/data/info.log \
        "$SERVICE_NAME" || exit

    docker logs -f "$SERVICE_NAME"
}

runPipelineContainer
