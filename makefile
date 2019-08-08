build:
	CGO_ENABLED=0 go build -o bin/app -a -installsuffix cgo ./ 
run: 
	bin/app
br: build run