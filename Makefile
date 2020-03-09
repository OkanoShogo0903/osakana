log:
	heroku logs -t

ps:
	heroku ps

restart:
	heroku restart

cancel:
	heroku builds:cancel

local:
	heroku local web

run:
	go run cmd/api/main.go

build:
	go build cmd/api/main.go

get:
	curl -v -X POST \
	  http://localhost:8080/user/aquarium \
	  -H 'content-type: application/json' \
	  -d '{ "background": 0, "plant_ids": [0, 3, 6], "creature_ids": [1, 2, 3, 4, 5]}'

post:
	curl -v -X POST \
	  http://localhost:8080/user/signup \
	  -H 'content-type: application/json' \
	  -d '{ "background": 0, "plant_ids": [0, 3, 6], "creature_ids": [1, 2, 3, 4, 5]}'

put:
	curl -v -X PUT \
	  http://localhost:8080/user/update \
	  -H 'content-type: application/json' \
	  -d '{ "background": 0, "plant_ids": [1], "creature_ids": [1, 2, 3, 4, 5]}'

dev:
	cd database && make __migrate/down && make migrate/up
