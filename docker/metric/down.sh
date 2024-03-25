#!/bin/bash
###
 # @Author: yujiajie
 # @Date: 2024-03-25 16:18:27
 # @LastEditors: yujiajie
 # @LastEditTime: 2024-03-25 16:24:59
 # @FilePath: /gateway/docker/metric/down.sh
 # @Description: 
### 

docker stop prometheus
docker rm prometheus

docker build -f prometheus/Dockerfile -t prometheus .

docker run -it --name prometheus \
-p 9090:9090 \
-v /usr/local/www/metric/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml
-v /usr/local/www/metric/prometheus/alertmanager.yml:/etc/prometheus/alertmanager.yml
-v /usr/local/www/metric/prometheus/node:/etc/prometheus/node
-v /usr/local/www/metric/prometheus/rules:/etc/prometheus/rules
-v /usr/local/www/metric/prometheus/data:/etc/prometheus/data
--log-opt max-size=10m --log-opt max-file=3 \
--restart always \
-d prometheus

docker stop grafana
docker rm grafana

docker build -f grafana/Dockerfile -t grafana .

docker run -it --name grafana \
-p 3011:3000 \
-v /usr/local/www/metric/grafana:/var/lib/grafana
--log-opt max-size=10m --log-opt max-file=3 \
--restart always \
-d grafana