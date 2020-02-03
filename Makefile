MODULE := pam_eth

module: test
	go build -v -buildmode=c-shared -a -o ${MODULE}.so
	-chmod +x ${MODULE}.so
	-ls -hl ${MODULE}.so

test: *.go
	go test -cover

clean:
	go clean
	-rm -f ${MODULE}.so ${MODULE}.h

.PHONY: test module download_deps clean
