include config/.env
export

run: i18n swag
	go run cmd/api/main.go

i18n:
	go run cmd/i18n/gen.go --csv config/errors.csv --out pkg/core/i18ngen.go --pkg core

swag:
	cd internal && swag fmt && swag init -g ../cmd/api/main.go