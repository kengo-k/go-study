target =

bin = main

image:
	docker-compose build --no-cache

start:
	docker-compose run -w /app/src --rm app /bin/ash

air:
	docker-compose run --rm app

run:
	docker-compose run --rm app go run $(target)

build:
	docker-compose run --rm app go build -o /app/bin/$(bin) $(target)

clean:
	docker-compose run --rm app rm -rf /app/bin

test:
	docker-compose run --rm app go test -v $(target)