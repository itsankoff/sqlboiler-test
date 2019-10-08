export DATABASE_HOST      ?= localhost
export DATABASE_PORT      ?= 5432
export DATABASE_USER      ?= ${USER}
export DATABASE_NAME      ?= sqlboiler-test
export DATABASE_SSLMODE   ?= disable
export DATABASE_FILES     ?= ./schema
export SCHEMA_FINGERPRINT := ./schema-postgresql-$(shell psql --version | grep -o '[0-9]*\.[0-9]').sum

db.setup:
	@createdb -h ${DATABASE_HOST} -p ${DATABASE_PORT} -U ${DATABASE_USER} ${DATABASE_NAME}
	@psql -h ${DATABASE_HOST} -p ${DATABASE_PORT} -U ${DATABASE_USER} ${DATABASE_NAME} < schema.sql

db.drop:
	@dropdb --if-exists -h ${DATABASE_HOST} -p ${DATABASE_PORT} -U ${DATABASE_USER} ${DATABASE_NAME}

db.reset: db.drop db.setup clean.schema generate

clean.schema:
	@rm -f schema/*.go

clean: clean.schema
	@rm -f *.sum

generate:
	@(PSQL_USER=${DATABASE_USER} \
		PSQL_HOST=${DATABASE_HOST} \
		PSQL_DBNAME=${DATABASE_NAME} \
		PSQL_PASS=${DATABASE_PASSWORD} \
		PSQL_SSLMODE=${DATABASE_SSLMODE} \
	sqlboiler --no-hooks --no-context --pkgname schema --output ./schema psql)
	@ls -d ${DATABASE_FILES}/* | sort | xargs -n 1 shasum -a 256 > ${SCHEMA_FINGERPRINT}
