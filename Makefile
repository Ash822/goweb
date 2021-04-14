test:
	go test ./...

mock-gen:
	go generate ./...

swagger-gen:
	swagger generate spec -o ./swagger.yaml --scan-models

swagger-run: swagger-gen
	swagger serve -F=swagger swagger.yaml

dev-mode:
	watcher

run:
	go run .