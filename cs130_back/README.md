### CS130 Backend Server

## Dependencies
1. Golang 1.9 or higher
2. Postgres

## Setup
Initialize a postgres database and update username, password, and database title variables in app-env accordingly.

## Instructions to run server
1. Source environment variables
```bash
source app-env
```

2. Build the project
```bash
make
```

3. Run the server
```bash
make run
```