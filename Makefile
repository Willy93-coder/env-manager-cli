build:
	go build -o bin/env-manager .

run:
	go run .

test:
	go test ./...

clean:
	rm -rf bin/env-manager
