#!/bin/bash
set -x
IMAGE_NAME=quay.io/agustinlare/test-image
TAG=golang120
PORT=8082
UNIQUEID=$RANDOM

main (){
    docker build -t $IMAGE_NAME:$TAG .
    test
}
test (){
    echo "===> DEBUG: $IMAGE_NAME:$TAG $PORT"
    docker run -p $PORT:$PORT --name $UNIQUEID -itd $IMAGE_NAME:$TAG
    if [ $? -ne 0 ]; then
        echo "Failed to run the Docker container."
        exit 1
    fi

    sleep 3 && response=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:$PORT/ping)

    if [ "$response" -eq 200 ]; then
        echo "Endpoint is transmitting data correctly."
        docker stop $UNIQUEID
        push
    else
        echo "Endpoint is not transmitting data correctly. Response code: $response"
    fi
}
push (){
    docker push $IMAGE_NAME:$TAG
}
main