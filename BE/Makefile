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

clean:
	rm -f $(APP_NAME)