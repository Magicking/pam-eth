// +build darwin linux

/*
Copyright (c) 2017 Uber Technologies, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"fmt"
	"io"
	"log/syslog"
	"strings"
)

// AuthResult is the result of the authentcate function.
type AuthResult int

const (
	// AuthError is a failure.
	AuthError AuthResult = iota
	// AuthSuccess is a success.
	AuthSuccess
)

func pamLog(format string, args ...interface{}) {
	l, err := syslog.New(syslog.LOG_AUTH|syslog.LOG_WARNING, "pam-eth")
	if err != nil {
		return
	}
	l.Warning(fmt.Sprintf(format, args...))
}

func pamAuthenticate(w io.Writer, uid int, username string, argv []string) AuthResult {

	var contractAddressOpt, rpcEndpointOpt string

	for _, arg := range argv {
		opt := strings.Split(arg, "=")
		switch opt[0] {
		case "contract-address":
			contractAddressOpt = opt[1]
			pamLog("contract address is set to %s", contractAddressOpt)
		case "rpc-endpoint":
			rpcEndpointOpt = opt[1]
			pamLog("RPC endpoint set is set to %s", rpcEndpointOpt)
		default:
			pamLog("unkown option: %s\n", opt[0])
		}
	}

	// TODO check signature here
	return AuthSuccess
}

func main() {}
