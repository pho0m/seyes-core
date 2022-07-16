.PHONY:
test:
	go test ./...

.PHONY:
build:
	echo "Should be building something"

.PHONY:
dev:
	docker-compose up

.PHONY:
down:
	docker-compose down -v

.PHONY:
psql:
	docker-compose exec db psql -U mns_local
