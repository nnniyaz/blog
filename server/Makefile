include .env
export

run: i18n swag
	go run cmd/api/main.go

i18n:
	go run cmd/i18n/gen.go --csv errors.csv --out pkg/core/i18ngen.go --pkg core

swag:
	cd cmd/api && swag fmt && swag init -g main.go