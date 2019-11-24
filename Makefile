all: build image compose
	rm crdb-test

# Build the test binary to containerise.
build:
	GOOS=linux go test ./... -v -c -o crdb-test

# Containerise the test binary.
image:
	docker build -t crdb-test .

# Run integration tests.
compose:
	docker-compose --no-ansi -f crdb-test.yaml up --abort-on-container-exit --force-recreate