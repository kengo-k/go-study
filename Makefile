target = 

image:
	docker-compose build --no-cache

air:
	docker-compose run --rm app

run:
	docker-compose run --rm app go run $(target)

test:
	docker-compose run --rm app go test -v $(target)
