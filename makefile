build:
	CGO_ENABLED=0 go build -o bin/app
run:
	bin/app
br: build run

pre_test:
	rm -R test/
	mkdir test/

t: pre_test
	go test .