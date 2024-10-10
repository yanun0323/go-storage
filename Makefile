.PHONY:

run:
	go run ./main.go

build:	## Build backend Docker image
	docker build . \
		-t go-storage

docker.run:
	docker run -d \
	-p 8001:8001 \
	--name go-storage go-storage

docker.up:
	docker container start go-storage

docker.down:
	docker container stop go-storage

docker.rm:
	docker rm -f go-storage