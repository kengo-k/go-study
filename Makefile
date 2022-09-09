target =

bin = main

image:
	docker-compose build --no-cache

air:
	docker-compose run --rm app

run:
	docker-compose run --rm app go run $(target)

build:
	docker-compose run --rm app go build -o /app/bin/$(bin) $(target)

test:
	docker-compose run --rm app go test -v $(target)

