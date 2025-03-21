.PHONY:docker
docker:
	@rm redbook || true
	@GOOS=linux GOARCH=arm go build -tags=k8s -o redbook .
	@docker rmi -f redbook:v0.0.1
	@docker build -t redbook:v0.0.1 .