ENV_LOCAL_FILE := env.local
ENV_LOCAL = $(shell cat $(ENV_LOCAL_FILE))

up:
	$(ENV_LOCAL) docker-compose up
down:
	$(ENV_LOCAL) docker-compose down
build:
	$(ENV_LOCAL) docker-compose build
