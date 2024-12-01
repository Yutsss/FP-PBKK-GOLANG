APP_NAME=backend

run:
	go run main.go

dev:
	CompileDaemon -command="go run main.go" -build="go build -o $(APP_NAME)"

build:
	go build -o $(APP_NAME)

migrate:
	go run main.go --migrate

run-build:
	./$(APP_NAME)

run-start-dev:
	GO_ENV=development ./$(APP_NAME)

run-start-prod:
	GO_ENV=production ./$(APP_NAME)
deploy-dev:
	pm2 start ecosystem.config.js --env development

deploy-prod:
	pm2 start ecosystem.config.js --env production

clean:
	rm -f $(APP_NAME)