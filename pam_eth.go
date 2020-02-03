// +build darwin linux

package main

import (
	"context"
	"fmt"
	"log/syslog"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens"
)

func pamLog(format string, args ...interface{}) {
	l, err := syslog.New(syslog.LOG_AUTH|syslog.LOG_WARNING, "pam-eth")
	if err != nil {
		return
	}
	l.Warning(fmt.Sprintf(format, args...))
	//	log.Printf(format, args...)
}

func initEth(argv []string) (*PametteHandler, error) {
	var contractAddress, rpcEndpointOpt, password string

	for _, arg := range argv {
		opt := strings.Split(arg, "=")
		switch opt[0] {
		case "contract-address":
			contractAddress = opt[1]
		case "rpc-endpoint":
			rpcEndpointOpt = opt[1]
		case "password":
			password = opt[1]
		default:
			pamLog("unkown option: %s\n", opt[0])
			return nil, fmt.Errorf("unknown option passed")
		}
	}

	// Connect to RPC
	pamLog("Connecting to %s", rpcEndpointOpt)
	c, err := ethclient.Dial(rpcEndpointOpt)
	if err != nil {
		return nil, fmt.Errorf("Could not initialize client: %v", err)
	}
	pamLog("Resolving ENS %s", contractAddress)
	netId, err := c.NetworkID(context.TODO())
	if err != nil {
		pamLog("NetworkID err %v", err)
	}
	pamLog("NetworkID %v", netId)
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

	ph := NewPametteHandler(ctct, password)
	return ph, nil
}

func main() {}
