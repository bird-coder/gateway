#!/bin/bash
###
 # @Author: yujiajie
 # @Date: 2024-03-25 16:18:27
 # @LastEditors: yujiajie
 # @LastEditTime: 2024-03-28 16:04:16
 # @FilePath: /Gateway/docker/metric/down.sh
 # @Description: 
### 

docker stop prometheus
docker rm prometheus

docker build -f docker/metric/prometheus/Dockerfile -t prometheus .

docker run -it --name prometheus \
-p 9090:9090 \
-v /usr/local/www/Gateway/docker/metric/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml \
-v /usr/local/www/Gateway/docker/metric/prometheus/node:/etc/prometheus/node \
-v /usr/local/www/Gateway/docker/metric/prometheus/rules:/etc/prometheus/rules \
-v /usr/local/www/Gateway/docker/metric/prometheus/data:/etc/prometheus/data \
--log-opt max-size=10m --log-opt max-file=3 \
--restart always \
-d prometheus

docker stop grafana
docker rm grafana

docker build -f docker/metric/grafana/Dockerfile -t grafana .

docker run -it --name grafana \
-p 3011:3000 \
-v /usr/local/www/Gateway/docker/metric/grafana:/var/lib/grafana \
-e "GF_SECURITY_ADMIN_PASSWORD=123456" \
-e "GF_USERS_ALLOW_SIGN_UP=false" \
--log-opt max-size=10m --log-opt max-file=3 \
--restart always \
-d grafana
