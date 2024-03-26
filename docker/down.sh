#!/bin/bash

docker stop gateway
docker rm gateway

docker build -f docker/Dockerfile -t gateway .

docker run -it --name gateway \
-p 8081:8081 \
-v /usr/local/www/gateway/logs:/app/logs
--log-opt max-size=10m --log-opt max-file=3 \
--restart always \
-d gateway