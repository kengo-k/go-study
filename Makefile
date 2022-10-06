build-image:
	docker build -t study-go:v1.0.0 .

connect:
	docker run -it --rm --name study-go -v "${PWD}/:/app" study-go:v1.0.0

test:
	docker run -it --rm --name study-go -v "${PWD}/:/app" study-go:v1.0.0 go test ./...

# ホスト上にgoが直接インストールされている場合はこちらを使ってよい
test-local:
	go test ./...
