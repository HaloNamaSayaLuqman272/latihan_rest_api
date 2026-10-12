dev-api:
	air

dev-worker:
	air -c .air.worker.toml

make rabbit:
	docker start rabbitmq