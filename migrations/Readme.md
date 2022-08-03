## Setup sql-migrate library
- To install run this command: `go get -v github.com/rubenv/sql-migrate/...`
And then Setup up it as a *Global* application on your machine. 

- Create a config file
- Create new migration file `sql-migrate new -config=./migrations/dbconfig.yaml -env="dev" ddl_create_permission_table`