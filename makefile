SHELL = /bin/bash

chat-run:
	go run web/test.go

tidy:
	go mod tidy
	go mod vendor

chat-test:
	go run tests/test_web.go

chat-curl:
	curl -H "X-API-Key: secret" http://localhost:8080