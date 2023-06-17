## build: will build go to a binary
build:
	@echo "Building go binary"
	go mod tidy
	cd cmd && go build -o ../bin/kubeomatic
	@echo "Build Done!"

## run: this will run the go binary
run:
	@echo "Running Kube-o-Matic..."
	cd bin && chmod +x kubeomatic && ./kubeomatic
	@echo "Done!"
	

## BuildAndRun: builds and runs the docker image
BuildAndRun:
	@echo "Building Docker Image"
	docker build -t kube-o-matic:latest .
	@echo "Done Building Image"
	@echo "Running docker image"
	docker run -p 8555:8555 -d --name kube-o-matic kube-o-matic:latest
	@echo "Docker Run complete"


## PullAndRun: pulls the docker image from Github Image Registry ghcr.io
PullAndRun:
	@echo "Pulling Docker image..."
	docker pull docker pull ghcr.io/miladhzzzz/kube-o-matic
	@echo "Running the Cotainer..."
	docker run -p 8555:8555 -d --name kube-o-matic ghcr.io/miladhzzzz/kube-o-matic
	@echo "Done!"

## down: stops the docker container
down:
	@echo "Stoping container..."
	docker stop kube-o-matic
	@echo "Deleting Cotainer..."
	docker rm kube-o-matic
	@echo "Done!"

## CleanDocker: cleans the docker of any image or container
CleanDocker:
	@echo "Cleaning Docker..."
	@echo "Stoping container..."
	docker stop kube-o-matic
	@echo "Deleting Cotainer..."
	docker rm kube-o-matic
	@echo "Deleting image"
	docker image rm kube-o-matic:latest
	@echo "Done!"

## swag: makes docs/ swagger ui
swag:
	@echo "Making swagger UI"
	swag init
	@echo "Done!"