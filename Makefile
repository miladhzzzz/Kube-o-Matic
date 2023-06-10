## build: will build go to a binary
build:
	@echo "Building go binary"
	go mod tidy
	cd cmd && go build -o ../bin/kubeomatic
	@echo "Build Done!"

## swag: makes docs/ swagger ui
swag:
	@echo "Making swagger UI"
	swag init
	@echo "Done!"

## docker: builds and runs the docker image
docker:
	@echo "Building Docker Image"
	docker build -t kube-o-matic:latest .
	@echo "Done Building Image"
	@echo "Running docker image"
	docker run -p 8555:8555 -d --name kube-o-matic kube-o-matic:latest
	@echo "Docker Run complete"
	
## down: stops the docker container
down:
	@echo "Stoping container..."
	docker stop kube-o-matic
	@echo "Done!"