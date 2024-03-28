#!/bin/bash
###
 # @Author: yujiajie
 # @Date: 2024-03-28 15:57:56
 # @LastEditors: yujiajie
 # @LastEditTime: 2024-03-28 16:03:47
 # @FilePath: /Gateway/docker/down.sh
 # @Description: 
### 

docker stop gateway
docker rm gateway

docker build -f docker/Dockerfile -t gateway .

docker run -it --name gateway \
-p 8081:8081 \
-v /usr/local/www/Gateway/logs:/app/logs \
--log-opt max-size=10m --log-opt max-file=3 \
--restart always \
-d gateway