#!/bin/bash

#停止服务
docker stop austin-web-api
docker stop austin-web-rpc
docker stop austin-job

#删除容器
docker rm austin-web-api
docker rm austin-web-rpc
docker rm austin-job

#删除镜像
docker rmi austin-web-api:v1
docker rmi austin-web-rpc:v1
docker rmi austin-job:v1

#删除none镜像
docker rmi $(docker images | grep "none" | awk '{print $3}')

#构建服务
docker build -t austin-job:v1 -f app/austin-job/Dockerfile .
docker build -t austin-web-rpc:v1 -f app/austin-web/rpc/Dockerfile .
docker build -t austin-web-api:v1 -f app/austin-web/api/Dockerfile .

#启动服务
docker run -itd --net=host --name=austin-job austin-job:v1
docker run -itd --net=host --name=austin-web-rpc austin-web-rpc:v1
docker run -itd --net=host --name=austin-web-api austin-web-api:v1

