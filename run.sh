#!/usr/bin/env bash

docker build . -t wwd990810/gin-todo-server
docker run -i -t -p 8080:8080 wwd990810/gin-todo-server