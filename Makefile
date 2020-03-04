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

dev:
	cd ./vue-app && npm run start
