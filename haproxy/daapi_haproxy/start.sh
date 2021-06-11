#!/bin/bash
docker run --name=haapi --net=host -v /etc/timezone:/etc/timezone -v /etc/localtime:/etc/localtime -d app-haproxy:2.2.13
