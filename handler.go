package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// PametteHandler is a wrapper to pamette contract
type PametteHandler struct {
	pamette       *PametteCallerSession
	password      string
	blockValidity *big.Int
}

func NewPametteHandler(pc *PametteCaller, password string) *PametteHandler {
	ctx := context.TODO()
	return &PametteHandler{
		pamette: &PametteCallerSession{
			Contract: pc,
			CallOpts: bind.CallOpts{
				Context: ctx,
			},
		},
		password: password,
	}
}

func (ph *PametteHandler) GetOTP(uid int64, username string) (string, error) {

	hashToSign, bNum, err := ph.pamette.GenerateOTP(big.NewInt(uid), username, ph.password)
	if err != nil {
		return "", err
	}
	ph.blockValidity = bNum
	return common.Bytes2Hex(hashToSign[:]), nil
}

func (ph *PametteHandler) VerifyAuth(uid int64, username, signature string) (bool, error) {
	log.Println(uid, username, signature)
	//ph.pamette.IsAuthorized(username, big.NewInt(uid),
	return true, nil
}
