
## Dependencies and setup
### `yarn`

*Very important to do this*

Downloads dependencies in yarn.lock. Do this before yarn start.

## Instructions to run client
### `yarn start`

Runs the app in the development mode.<br />
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br />
You will also see any lint errors in the console.

## Testing instructions
### `yarn test`

There are two test files on the front end: Component.test.js and View.test.js.
The first file focuses on testing the components that show up across many pages of the app.
The second file focuses on testing the pages/components that are page specific. 
Running yarn test will automatically run the tests and output their result. 

If yarn test gives a MutationObserver error, use `yarn test --env jest-environment-jsdom-fourteen`

See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.
