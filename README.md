# golang-ecolabel-backend

Golang implementation for EcoLabel project

## Run api

Start your API by running `go run main.go` in the terminal.
Your API should start listening on port `3000`.

In another terminal window, first, send a login request to obtain a JWT token:

```sh
curl -X POST -H "Content-Type: application/json"
-d '{"username": "test", "password": "test"}' 
http://localhost:3000/login
```

### This command sends a POST request with the JSON payload

`{"username": "test", "password": "test"}` to the `/login` route.
If the credentials are correct, the API will return a JWT token as a response.
The response should look like

```json
{"token":"<your_jwt_token>"}
```

Copy the value of the token field (without quotes) for the next step.

Now, send a request to the dashboard route using the JWT token:

```sh
curl -X GET -H "Authorization: Bearer <your_jwt_token>" http://localhost:3000/dashboard
```

Replace <your_jwt_token> with the token value you obtained in the previous step.
If the token is valid, you should see a response similar to:

```json
{"message":"Welcome to the dashboard"}
```

If the token is invalid or not provided, the API will return a `401 Unauthorized` status with an error message.
