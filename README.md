### CS130 Full Application Directions

## Setup
Go into cs130_back and cs130_front and follow the instructions in their respective READMEs for both to get dependencies and set-up. 


## Instructions to run client and server
1. In a terminal, run "make run" inside of the cs130_back directory. This will start the server.
2. In a separate terminal, run "yarn start" in the cs130_front directory. This will start the client. 
3. Go to localhost:3000/ on a browser to use the application.


## Testing Instructions
The testing instructions for the server and client are in their respective READMEs.


## Project directory 
The Studdie is comprised of a Golang server located under cs130_back and a React client under cs130_front. The application code as well as testing suites are located under these folders.

The most important directories and files are highlighted below.

- cs130_back
    - handlers
        - (multiple handler files)
        - (multiple test files to test handler functions)
    - models
        - (multiple model files)
        - (multiple test files to test model functions)
    - app.go
- cs130_front
    - src
        - components
            - views
                - (React components that make up the pages in the app)
            - (various common React components)
        - App.js
        - Component.test.js
        - View.test.js


- Backend
    - Handlers are the functions that are executed when specific endpoints are triggered. 
    - Models describe the database schema of the different classes and contain functions that handle database operations.
    - app.go initializes and configures the server and defines the paths of the different endpoints on the server. 
    - The test files under handlers and models end in "_test.go"

- Frontend
    - Components make up the entirety of the React app. views contain components that are comglomerates of smaller components in order to create a page on the app. components also contains various other smaller components that make up parts of a page. 
    - App.js defines the different routes/paths available on the client. 
    - Component.test.js and View.test.js contain the tests for the React client.
