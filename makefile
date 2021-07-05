gorun: gocompile
	./build/processing-file --file-path ./exmpl.txt
gocompile:
	CGO_ENABLED=0 GOOS=linux GOARH=amd64 go build -o build/processing-file ./cmd/main.go
gotest:
	go test -v ./internal