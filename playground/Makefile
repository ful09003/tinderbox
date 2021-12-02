clean:
	-docker stop prom 
	-docker rm prom
	-rm -rf tsdb_out/*
	>hm

build:
	go build .

run: clean build
	pwd
	./docker_run.sh
	./playground > hm
	promtool tsdb create-blocks-from openmetrics hm tsdb_out/
	docker restart prom