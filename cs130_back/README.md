### CS130 Backend Server

## Dependencies
1. Golang 1.9 or higher
2. Postgres
3. jwt-go
4. go-sqlmock
5. httpexpect

## Setup
Get the above dependencies. For jwt-go, go-sqlmock, and httpexpect, use the "go get" command and the Github repo to retrieve the dependencies.

Initialize a postgres database and update username, password, and database title variables in app.go accordingly. This can be done easily using pg Admin when downloading Postgres. 

Due to pathing issues with src, you may need to rename the root folder name to "src" and run
```
export GOPATH=$(pwd)
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


## Testing
Testing can be done without the use of a database, as all tests are run with a mock database. The two main packages, models and handlers, have full test coverage. There are several methods of running these tests.

Option 1:
Run the following commands from the root directory to test each of the packages.
- Run `go test cs130_back/handlers`
- Run `go test cs130_back/models`

Option 2:
Change the current directory to either the handlers or the models package, and run `go test`.
