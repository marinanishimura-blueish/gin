up:
	docker-compose -p gin up -d
serve:
	docker-compose -p gin exec gin-app go run api.go
	# docker-compose -p gin exec gin-app go run main.go
	# docker-compose -p gin exec gin-app go run gorm.go
app:
	docker-compose -p gin exec gin-app sh
db:
	docker-compose -p gin exec gin-db sh
migrate:
	docker-compose -p gin exec gin-migrate sh
stop:
	docker-compose -p gin stop
down:
	docker-compose -p gin down