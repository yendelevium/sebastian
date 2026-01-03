build:
	go build -o sebastian cmd/sebastian/main.go

compose:
	docker compose up --build