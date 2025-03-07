SHELL = /bin/bash

chat-run:
	go run web/test.go

tidy:
	go mod tidy
	go mod vendor

chat-run:
	go run chat/api/services/main.go

chat-curl:
	curl -H "X-API-Key: secret" http://localhost:3000/test