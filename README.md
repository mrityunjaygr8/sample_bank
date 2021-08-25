### Sample Bank

Created in golang, follows the tutorial in [this youtube playlist](https://www.youtube.com/playlist?list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE).

For running, requires the following environment variables in a file called ".env" at the root level of the repo:
- DB_PASS
- DB_USER
- DB_PORT
- DB_NAME
- DB_HOST

Requires the following binaries to be installed:
- [docker](https://www.docker.com/get-started)
- [migrate](https://github.com/golang-migrate/migrate)
- [sqlc](https://github.com/kyleconroy/sqlc)

There are the following `make` commands for ease of use:
- dbstart  
Create a docker based instance of PostgreSQL 13, alpine flavor
- dbstop  
Destroy the docker based instance of PostgreSQL 13, alpine flavor
- dbcreate  
Create a Database in the created PSQL container, name is defined in .env by DB_NAME
- dbdrop  
Drop the Database in the created PSQL container, name is defined in .env by DB_NAME
- migrateup  
Apply all upward migrations
- migratedown  
Apply all downward migrations, reset the DB to blank, except the schema table
- sqlc  
Compile GO code based upon the specified queries using sqlc