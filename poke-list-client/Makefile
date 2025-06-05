ifneq (,$(wildcard ./.env))
	include .env
	export
endif

.PHONY: deps dev
.DEFAULT_GOAL := deps

deps:
	pnpm i

dev:
	pnpm run dev
