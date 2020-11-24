### CS130 Backend Server

## Dependencies
1. Golang 1.9 or higher
2. Postgres
3. jwt-go

## Setup
Get the above dependencies. For jwt-go, you can use the go get command and the Github repo. 

Initialize a postgres database and update username, password, and database title variables in app-env accordingly. This can be done easily using pg Admin when downloading Postgres. 

Due to pathing issues with src, you may need to rename the root folder name to "src" and run
```export GOPATH=$(pwd)
```
in the directory containing src. 


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