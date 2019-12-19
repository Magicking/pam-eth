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
	"log"
	"log/syslog"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens"
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
	log.Printf(format, args...)
}

func initEth(argv []string) (*PametteCaller, error) {
	var contractAddress, rpcEndpointOpt, password string

	for _, arg := range argv {
		opt := strings.Split(arg, "=")
		switch opt[0] {
		case "contract-address":
			contractAddress = opt[1]
		case "rpc-endpoint":
			rpcEndpointOpt = opt[1]
			pamLog("RPC endpoint set is set to %s", rpcEndpointOpt)
		case "password":
			password = opt[1]
		default:
			pamLog("unkown option: %s\n", opt[0])
			return nil, fmt.Errorf("unknown option passed")
		}
	}

	// Connect to RPC
	c, err := ethclient.Dial(rpcEndpointOpt)
	if err != nil {
		return nil, fmt.Errorf("Could not initialize client: %v", err)
	}
	// Resolve ens
	address, err := ens.Resolve(c, contractAddress)
	if err != nil {
		return nil, err
	}
	pamLog("%s resolved to %s", contractAddress, address.Hex())
	ctct, err := NewPametteCaller(address, c)
	if err != nil {
		return nil, fmt.Errorf("Could not instanciate Pamette: %v", err)
	}

	pamLog("Check with password %v", password)
	return ctct, nil
}

func pamAuthenticate(w io.Writer, uid int, username string, pamette *PametteCaller) AuthResult {

	str, err := pamette.GenerateOTP(nil)
	if err != nil {
		log.Printf("Could not generate OTP: %v", err)
		return AuthError
	}
	log.Println(str)
	/*
		ppf, err := examples.NewPingPongFilterer(addr, c)
		if err != nil {
			log.Fatalf("Could not watch for events: %v", err)
		}*/
	// TODO check signature here
	return AuthSuccess
}

func main() {}
