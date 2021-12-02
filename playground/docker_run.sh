#!/usr/bin/env bash

docker run -p 9090:9090 --name prom -dt --user $(id -u) --mount type=bind,source="$(pwd)/tsdb_out/",target=/prometheus prom/prometheus:latest