# golang-ecolabel-backend

🔷 Golang implementation for 📗 EcoLabel project

## Run api with Go CLI

Start your API by running `go run main.go` in the terminal.
Your API should start listening on port `3000`.

In another terminal window, first, send a login request to obtain a JWT token:

```sh
curl -X POST -H "Content-Type: application/json"
-d '{"username": "test", "password": "test"}' 
http://localhost:3000/login
```

This command sends a POST request with the JSON payload

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

### Testing

```sh
go test -v ./...
```

### Benchmark

```sh
go test -bench=. -benchmem ./internal/...
```

## Run api with Makefile 💫

### Building the project

To build the project, run the following command:

```sh
make build
```

This will compile the project and create an executable in the bin/ directory called ecolabel.

### Running the project

To run the compiled executable, execute the following command:

```sh
make run
```

This command will first build the project and then run the compiled ecolabel executable.

### Running tests

To run all the tests in the project, use the following command:

```sh
make test
```

This will execute all the tests in the project and display the results.

### Running benchmarks

To run benchmark tests for the project, use the following command:

```sh
make bench
```

This command will run the benchmark tests in the internal/ directory and display the results, including the number of operations, the time taken per operation, and the memory usage.

```sh
make cover
```

This target runs the test suite for the entire project while collecting code coverage information. It generates a coverage report in the coverage.out file. After running the tests, it uses the go tool cover command with the -func flag to display the coverage percentages for each function in your code.

```sh
make cover-html
```

Similar to the cover target, this target also runs the test suite for the entire project while collecting code coverage information and generates a coverage report in the coverage.out file. However, instead of displaying the coverage percentages for each function, it generates an HTML coverage report using the go tool cover command with the -html flag. The HTML report is saved as coverage.html. The last line attempts to open the generated coverage.html file in the default web browser. The `xdg-open` command is used for Linux systems, and the `open` command is used for macOS systems, and the `start` is for Windows users.

Here are curl examples for setting and getting session values:

Set a session value:

```sh
curl -X POST http://localhost:3000/set-session-value -c cookies.txt
```

This will send a POST request to the /set-session-value endpoint and save the session cookie in a file called cookies.txt.

Get the session value:

```sh
curl -X GET http://localhost:3000/get-session-value -b cookies.txt
```

This will send a GET request to the /get-session-value endpoint and use the session cookie stored in cookies.txt. The response should contain the session value you set earlier.
