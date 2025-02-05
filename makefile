SHELL = /bin/bash

chat-run:
	go run web/test.go

tidy:
	go mod tidy
	go mod vendor