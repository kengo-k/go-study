build-image:
	docker build -t study-go:v1.0.0 .

connect:
	docker run -it --rm --name study-go -v "${PWD}/:/app" study-go:v1.0.0

test:
	go test ./...