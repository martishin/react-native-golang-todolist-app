# Todo React Native App
React Native todo application with backend written in Go

<div>
  <img src="https://github.com/tty-monkey/react-native-golang-todolist-app/blob/main/ios.png" width="200" style="display: inline-block; margin-right: 10px;"/>
  <img src="https://github.com/tty-monkey/react-native-golang-todolist-app/blob/main/android.png" width="195" style="display: inline-block;"/>
</div>

## Running Locally
### Server
* Navigate to server folder: `cd server`
* Run PostgreSQL: `make run-postgres`
* Start the server: `cd cmd && go run .`
* API will be available at http://localhost:3000/
### Client
* Navigate to client folder: `cd app`
* Install dependencies: `npm install`
* Start expo: `npm run start`
