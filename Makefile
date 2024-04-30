
SQLDEF=./db/sqldef/psqldef
SCHEMADEF=./db/schema/schema.sql
DOWNLOAD_SQLDEF=psqldef_linux_arm64
IMPORT_DATA_SH=./db/testdata/load_csv.sh
IMPORT_DATA_DIR=./db/testdata

DB_HOST=db
DB_PORT=5432
DB_NAME=dbname
DB_USER=user
DB_PASSWORD=password

.PHONY: sqldef-setup
sqldef-setup:
	curl -L https://github.com/k0kubun/sqldef/releases/download/v0.17.7/$(DOWNLOAD_SQLDEF).tar.gz -o db/sqldef/$(DOWNLOAD_SQLDEF).tar.gz
	tar -xf db/sqldef/$(DOWNLOAD_SQLDEF).tar.gz -C db/sqldef
	rm db/sqldef/$(DOWNLOAD_SQLDEF).tar.gz

.PHONY: db-dry-run
db-dry-run:
	$(SQLDEF) -U $(DB_USER) -W $(DB_PASSWORD) -h $(DB_HOST) -p $(DB_PORT) $(DB_NAME) --dry-run < $(SCHEMADEF)

.PHONY: db-apply
db-apply:
	$(SQLDEF) -U $(DB_USER) -W $(DB_PASSWORD) -h $(DB_HOST) -p $(DB_PORT) $(DB_NAME) < $(SCHEMADEF)
	

# drop table and create table
.PHONY: db-apply-force
db-apply-force:
	$(SQLDEF) -U $(DB_USER) -W $(DB_PASSWORD) -h $(DB_HOST) -p $(DB_PORT) $(DB_NAME) --enable-drop-table < $(SCHEMADEF)

.PHONY: db-connect
db-connect:
	psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME)
