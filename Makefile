assets = assets
appFile = cmd/app/app.go
appName = app

.PHONY: build
build:
	go build -o $(appName) -v $(appFile)

.PHONY: run
run: build
	./$(appName) $(assets)/configs.txt $(assets)/content.txt $(assets)/output.txt

.PHONY: clean
clean:
	rm $(appName)