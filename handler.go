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

func (ph *PametteHandler) GetOTP(uid int64, username string) (string, *big.Int, error) {

	hashToSign, bNum, err := ph.pamette.GenerateOTP(big.NewInt(uid), username, ph.password)
	if err != nil {
		return "", nil, err
	}
	ph.blockValidity = bNum
	return common.Bytes2Hex(hashToSign[:]), bNum, nil
}

func (ph *PametteHandler) VerifyAuth(uid int64, username, signature string) (bool, error) {
	log.Println(uid, username, signature)
	// todo check sig length
	v, r, s := extractSignature(signature)
	ret, err := ph.pamette.IsAuthorized(big.NewInt(uid), username, ph.password, ph.blockValidity, v, r, s)
	log.Println("ret", ret)
	if err != nil {
		return false, err
	}
	if big.NewInt(0).Cmp(ret) != 0 {
		pamLog(ret.String())
		return false, err
	}
	return true, nil
}

func extractSignature(signature string) (v uint8, r, s [32]byte) {
	sig := common.FromHex(signature)
	v = uint8(sig[64])
	copy(r[:], sig[0:32])
	copy(s[:], sig[32:64])

	return v, r, s
}
