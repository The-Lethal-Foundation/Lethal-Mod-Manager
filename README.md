# lorca-ts-react-starter

A starter project for building modern cross-platform desktop apps in Go, HTML, Typescript and React.

This is really nothing more than a [Typescript][3]-flavoured ["Create React App"][1]-bootstrapped app combined with a modified example from the fine ["Lorca"][2] project.

Check out those projects for how to use each component.

[1]: https://facebook.github.io/create-react-app/docs/adding-typescript "Create React App"
[2]: https://github.com/zserge/lorca "Lorca"
[3]: https://www.typescriptlang.org "Typescript"

## Getting started

Either just clone this repo and start coding, or repeat the simple steps taken to create it yourself to get that warm fuzzy feeling of having done it all "by hand".

### Prerequisites

Since this project builds on [Lorca][2], you need Chrome installed to develop and run your app. You also need recent enough versions of Go (1.12), Node.js (v8.X) and npm (v5.2) installed.

### Cloning

1. Clone this repo: `git clone --depth=1 https://github.com/erkkah/lorca-ts-react-starter/lorca-ts-react-starter.git <NEWPROJECTNAME>`
1. Go to the newly created directory: `cd <NEWPROJECTNAME>`
1. Install dependencies: `npm install`
1. Launch the app in development mode: `npm start`
1. Read the code and start building your app!

### By hand

1. Create a React app with Typescript: `npx create-react-app <NEWPROJECTNAME> --typescript`
1. Go to the newly created directory: `cd <NEWPROJECTNAME>`
1. Initialize the Go part of the project: `go mod init <APPNAME>`
1. Copy the Go stub files from this project to your project
1. Copy the `scripts` part from `package.json` to your project
    1. Make sure to update the `BROWSER=./app` to `<APPNAME>` from above
1. Install `npm-run-all` since the scripts need it: `npm install --save-dev npm-run-all`
1. Launch the app in development mode: `npm start`
1. Read the code and start building your app!

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode. The "backend" Go part will launch Chrome to display the HTML content.

The page will reload if you make edits to the React parts.
You will also see any lint errors in the console.

### `npm run build`

Builds the app for production, creating a single executable will all assets bundled. It correctly bundles React in production mode and optimizes the build for the best performance.

### `npm test`

Runs the test suites for both Go and React parts.
To launch React tests in watch mode, run `npm run react:test -- --watchAll=true`
