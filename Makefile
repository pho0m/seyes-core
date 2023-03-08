.PHONY:
dev:
	docker-compose up

.PHONY:
down:
	docker-compose down -v

.PHONY:
psql:
	docker-compose exec db psql -U seyes_core
