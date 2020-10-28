### CS130 Backend Server

## Dependencies
1. Golang 1.9 or higher
2. Postgres

## Setup
Initialize a postgres database and update username, password, and database title variables in app-env accordingly.

## Instructions to run server
1. Update database variables
Go into "app.go" and update variables APP_DB_USERNAME, APP_DB_PASSWORD, and APP_DB_NAME.

2. Build the project
```bash
make
```

3. Run the server
```bash
make run
```