#!/usr/bin/env bash

docker build . -t gin-server
docker run -i -t -p 8080:8080 gin-server