# cockroachdb-testing
An example of integration testing CockroachDB database connections.

## Installation

```
$ go get -u -d github.com/codingconcepts/cockroachdb-testing
```

You'll also need to install Docker and docker-compose.

## Run

Build a test executable to containerise:

```
$ make build
```

Containerise the test executable:

```
$ make image
```

Run the integration tests:

```
$ make compose
```

Alternatively, you can build, containerise, and run the integration tests by running:

```
$ make all
```