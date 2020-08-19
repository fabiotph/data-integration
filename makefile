start:
	docker-compose up -d

clean:
	docker-compose stop
	docker-compose rm -fsv

test:
	docker-compose stop
	docker-compose rm -fsv
	docker-compose up -d
	docker-compose exec api go test server/tests
	docker-compose stop
	docker-compose rm -fsv