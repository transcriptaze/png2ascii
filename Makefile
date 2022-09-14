CMD = ./bin/png2ascii
.PHONY: format

format:
	go fmt ./...

build: format
	mkdir -p bin
	go build -o bin ./...

build-all: build
	mkdir -p dist/$(DIST)/windows
	mkdir -p dist/$(DIST)/darwin
	mkdir -p dist/$(DIST)/linux
	env GOOS=linux   GOARCH=amd64 GOWORK=off go build -trimpath -o dist/$(DIST)/linux/   ./...
	env GOOS=darwin  GOARCH=amd64 GOWORK=off go build -trimpath -o dist/$(DIST)/darwin/  ./...
	env GOOS=windows GOARCH=amd64 GOWORK=off go build -trimpath -o dist/$(DIST)/windows/ ./...

text: build
	$(CMD) --debug --format text -out ./runtime/mp42asc.txt ./images/reference.png
	cat ./runtime/mp42asc.txt

png: build
	$(CMD) --debug --profile .default.json --format png -out ./runtime/reference.png ./images/reference.png
	open ./runtime/reference.png

nosquoosh: build
	$(CMD) --debug --no-squoosh --format png -out ./runtime/squooshed.png ./images/squooshed.png
	open ./runtime/squooshed.png

mp4: build
	rm -rf ./runtime/mp4
	$(CMD)  --debug --format png -out ./runtime/mp4 ./runtime/movie

gradient: build
	$(CMD)  --debug --format png -out ./runtime/gradient.png  ./images/gradient.png
	open ./runtime/gradient.png

ancient: build
	$(CMD)  --debug --format png -out ./runtime/mp42asc.png ./runtime/ancient-dust.png
	open ./runtime/mp42asc.png

floating: build
	$(CMD)  --debug --format png -out ./runtime/mp42asc.png ./runtime/floating.png
	open ./runtime/mp42asc.png
