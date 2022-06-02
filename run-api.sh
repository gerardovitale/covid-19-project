#!/bin/bash

source .env
source .env.secrets
source commons.sh

SERVICE_NAME=api-container

runAPIContainer() {
    cleanUpDocker $SERVICE_NAME

    echo "$(dateTimeNow) [INFO] - Building new $SERVICE_NAME"
    docker build --platform linux/amd64 \
        -f api/Dockerfile \
        --build-arg MONGO_PASS="$MONGO_PASS" \
        -t $SERVICE_NAME . || exit

    echo "$(dateTimeNow) [INFO] - Running $SERVICE_NAME"
    docker run -d \
        --platform linux/amd64 \
        --name=$SERVICE_NAME \
        $SERVICE_NAME || exit

    docker logs -f $SERVICE_NAME
}

runAPIContainer
