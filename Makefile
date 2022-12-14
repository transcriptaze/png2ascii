CMD = ./bin/png2ascii
PROFILE = .profiles/default.json
DIST = development

.PHONY: format

format:
	go fmt ./...

build: format
	mkdir -p bin
	go build -o bin ./...

test: build
	go test ./...

build-all: build
	mkdir -p dist/$(DIST)/windows
	mkdir -p dist/$(DIST)/darwin
	mkdir -p dist/$(DIST)/linux
	env GOOS=linux   GOARCH=amd64 GOWORK=off go build -trimpath -o dist/$(DIST)/linux/   ./...
	env GOOS=darwin  GOARCH=amd64 GOWORK=off go build -trimpath -o dist/$(DIST)/darwin/  ./...
	env GOOS=windows GOARCH=amd64 GOWORK=off go build -trimpath -o dist/$(DIST)/windows/ ./...

release: build-all
	cd dist && tar cvzf $(DIST).tar.gz $(DIST)/*

debug: build
	$(CMD) --debug --profile $(PROFILE) --format png --bgcolor '#ff00ff'   --fgcolor '#00ff00'   --out ./runtime/reference.png ./images/reference.png
	$(CMD) --debug --profile $(PROFILE) --format png --bgcolor '#44444444' --fgcolor '#ff000080' --out ./runtime/reference.png ./images/reference.png
# 	open ./runtime/reference.png

text: build
	$(CMD) --debug --format text --out ./runtime/mp42asc.txt ./images/reference.png
	cat ./runtime/mp42asc.txt

png: build
	$(CMD) --debug --profile $(PROFILE) --format png --out ./runtime/reference.png ./images/reference.png
	open ./runtime/reference.png

squoosh: build
# 	$(CMD) --debug --profile $(PROFILE) --format png --squoosh width:200 --out ./runtime/reference.png ./images/reference.png
# 	open ./runtime/reference.png
# 	$(CMD) --debug --format text --squoosh width:132 --out ./runtime/mp42asc.txt ./images/reference.png
# 	cat ./runtime/mp42asc.txt
	$(CMD) --debug --profile $(PROFILE) --format png --squoosh width:400 --out ./runtime/reference.png ./runtime/floating.png
	open ./runtime/reference.png

nosquoosh: build
	$(CMD) --debug --squoosh no --format png --out ./runtime/squooshed.png ./images/squooshed.png
	open ./runtime/squooshed.png

font: build
# 	$(CMD) --debug --profile $(PROFILE) --font gomonoitalic:16:100 --format png --out ./runtime/reference.png ./images/reference.png
	$(CMD) --debug --profile $(PROFILE) --font .fonts/repetition.ttf:48:72 --squoosh width:100 --format png --out ./runtime/reference.png ./images/reference.png
	open ./runtime/reference.png

mp4: build
	rm -rf ./runtime/mp4
	$(CMD)  --debug --format png --out ./runtime/mp4 ./runtime/movie

gradient: build
	$(CMD)  --debug --format png --out ./runtime/gradient.png  ./images/gradient.png
	open ./runtime/gradient.png

