ifneq (,$(wildcard ./.env))
	include .env
	export
endif

.PHONY: deps dev
.DEFAULT_GOAL := deps

deps:
	go mod tidy
	go mod download

dev:
	go run cmd/main.go
