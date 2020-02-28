build:
	CGO_ENABLED=0 go build -o bin/app
run:
	bin/app
rund:
	LOG_LEVEL=debug	bin/app
br: build run
brd: build rund

pre_test:
	if [ -d "test/" ]; then \
		rm -R test/; \
	fi;	
	mkdir test/

t: pre_test
	go test .


vendor:
	go mod vendor
	go mod download


build_tst:
	CGO_ENABLED=0 go build -o bin/app_tst testapp/*.go
run_tst:
	bin/app_tst
brt: build_tst run_tst
