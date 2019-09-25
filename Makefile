MODULE := pam_eth

module: test
	go build -buildmode=c-shared -o ${MODULE}.so
	-chmod +x ${MODULE}.so

test: *.go
	go test -cover

clean:
	go clean
	-rm -f ${MODULE}.so ${MODULE}.h

.PHONY: test module download_deps clean
