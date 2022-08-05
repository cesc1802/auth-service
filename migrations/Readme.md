## Setup sql-migrate library
- To install run this command: `go get -v github.com/rubenv/sql-migrate/...`
And then Setup up it as a *Global* application on your machine. 

- Create a config file
- Create new migration file `sql-migrate new -config=./migrations/dbconfig.yaml -env="dev" ddl_create_permission_table`
- Migrate up `sql-migrate up -config=./migrations/dbconfig.yaml -env="local"`
- Migrate down `sql-migrate down -config=./migrations/dbconfig.yaml -env="local"`

## Before you want to run migration on Production ENV
- export all variable you define in `dbconfig.yaml` file (e.g: `export MYSQL_USER=admin`, `export MYSQL_PASSWORD=admin@123`)
- Run migrate up 