ifneq (,$(wildcard ./.env))
	include .env
	export
endif
ifneq (,$(wildcard ./.env.local))
	include .env.local
	export
endif

dev:
	pnpm run dev
