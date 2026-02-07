# Makefile

# Load .env nếu có (tự động đọc biến từ .env)
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DB_URL = mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?charset=utf8mb4&parseTime=True&loc=Local

MIGRATE = migrate -path migrations -database "$(DB_URL)"

up:
	$(MIGRATE) up

down:
	$(MIGRATE) down

create-%:
	$(MIGRATE) create -ext sql -dir migrations -seq $*

status:
	$(MIGRATE) version

force-%:
	$(MIGRATE) force $*