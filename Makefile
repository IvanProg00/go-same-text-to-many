appFile = cmd/app/app.go
appName = app

.PHONY: build
build:
	go build -o $(appName) -v $(appFile)

.PHONY: run
run: build
	./${appName} configs.txt content.txt output.txt

.PHONY: clean
clean:
	rm $(appName)