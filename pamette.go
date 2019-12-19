// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PametteABI is the input ABI used to generate the binding from.
const PametteABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"generateOTP\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PametteFuncSigs maps the 4-byte function signature to its string representation.
var PametteFuncSigs = map[string]string{
	"08e465d0": "generateOTP()",
	"f54aca07": "isAuthorized(bytes,bytes)",
}

// PametteBin is the compiled bytecode used for deploying new contracts.
var PametteBin = "0x608060405234801561001057600080fd5b5061027e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806308e465d01461003b578063f54aca07146100b8575b600080fd5b6100436101f9565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561007d578181015183820152602001610065565b50505050905090810190601f1680156100aa5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101e5600480360360408110156100ce57600080fd5b8101906020810181356401000000008111156100e957600080fd5b8201836020820111156100fb57600080fd5b8035906020019184600183028401116401000000008311171561011d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561017057600080fd5b82018360208201111561018257600080fd5b803590602001918460018302840111640100000000831117156101a457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061021b945050505050565b604080519115158252519081900360200190f35b6040805180820190915260088152671c185cdcdddbdc9960c21b602082015290565b60019291505056fea265627a7a723158208f08c8210f9dc774d17287dcbfe884a7763423416109a1ec27239976411d2ce364736f6c637828302e352e31342d646576656c6f702e323031392e31322e372b636f6d6d69742e31666531343539620058"

// DeployPamette deploys a new Ethereum contract, binding an instance of Pamette to it.
func DeployPamette(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pamette, error) {
	parsed, err := abi.JSON(strings.NewReader(PametteABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PametteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pamette{PametteCaller: PametteCaller{contract: contract}, PametteTransactor: PametteTransactor{contract: contract}, PametteFilterer: PametteFilterer{contract: contract}}, nil
}

// Pamette is an auto generated Go binding around an Ethereum contract.
type Pamette struct {
	PametteCaller     // Read-only binding to the contract
	PametteTransactor // Write-only binding to the contract
	PametteFilterer   // Log filterer for contract events
}

// PametteCaller is an auto generated read-only Go binding around an Ethereum contract.
type PametteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PametteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PametteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PametteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PametteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PametteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PametteSession struct {
	Contract     *Pamette          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PametteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PametteCallerSession struct {
	Contract *PametteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PametteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PametteTransactorSession struct {
	Contract     *PametteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PametteRaw is an auto generated low-level Go binding around an Ethereum contract.
type PametteRaw struct {
	Contract *Pamette // Generic contract binding to access the raw methods on
}

// PametteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PametteCallerRaw struct {
	Contract *PametteCaller // Generic read-only contract binding to access the raw methods on
}

// PametteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PametteTransactorRaw struct {
	Contract *PametteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPamette creates a new instance of Pamette, bound to a specific deployed contract.
func NewPamette(address common.Address, backend bind.ContractBackend) (*Pamette, error) {
	contract, err := bindPamette(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pamette{PametteCaller: PametteCaller{contract: contract}, PametteTransactor: PametteTransactor{contract: contract}, PametteFilterer: PametteFilterer{contract: contract}}, nil
}

// NewPametteCaller creates a new read-only instance of Pamette, bound to a specific deployed contract.
func NewPametteCaller(address common.Address, caller bind.ContractCaller) (*PametteCaller, error) {
	contract, err := bindPamette(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PametteCaller{contract: contract}, nil
}

// NewPametteTransactor creates a new write-only instance of Pamette, bound to a specific deployed contract.
func NewPametteTransactor(address common.Address, transactor bind.ContractTransactor) (*PametteTransactor, error) {
	contract, err := bindPamette(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PametteTransactor{contract: contract}, nil
}

// NewPametteFilterer creates a new log filterer instance of Pamette, bound to a specific deployed contract.
func NewPametteFilterer(address common.Address, filterer bind.ContractFilterer) (*PametteFilterer, error) {
	contract, err := bindPamette(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PametteFilterer{contract: contract}, nil
}

// bindPamette binds a generic wrapper to an already deployed contract.
func bindPamette(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PametteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pamette *PametteRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pamette.Contract.PametteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pamette *PametteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pamette.Contract.PametteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pamette *PametteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pamette.Contract.PametteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pamette *PametteCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pamette.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pamette *PametteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pamette.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pamette *PametteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pamette.Contract.contract.Transact(opts, method, params...)
}

// GenerateOTP is a free data retrieval call binding the contract method 0x08e465d0.
//
// Solidity: function generateOTP() constant returns(string)
func (_Pamette *PametteCaller) GenerateOTP(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pamette.contract.Call(opts, out, "generateOTP")
	return *ret0, err
}

// GenerateOTP is a free data retrieval call binding the contract method 0x08e465d0.
//
// Solidity: function generateOTP() constant returns(string)
func (_Pamette *PametteSession) GenerateOTP() (string, error) {
	return _Pamette.Contract.GenerateOTP(&_Pamette.CallOpts)
}

// GenerateOTP is a free data retrieval call binding the contract method 0x08e465d0.
//
// Solidity: function generateOTP() constant returns(string)
func (_Pamette *PametteCallerSession) GenerateOTP() (string, error) {
	return _Pamette.Contract.GenerateOTP(&_Pamette.CallOpts)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xf54aca07.
//
// Solidity: function isAuthorized(bytes data, bytes signature) constant returns(bool)
func (_Pamette *PametteCaller) IsAuthorized(opts *bind.CallOpts, data []byte, signature []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pamette.contract.Call(opts, out, "isAuthorized", data, signature)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0xf54aca07.
//
// Solidity: function isAuthorized(bytes data, bytes signature) constant returns(bool)
func (_Pamette *PametteSession) IsAuthorized(data []byte, signature []byte) (bool, error) {
	return _Pamette.Contract.IsAuthorized(&_Pamette.CallOpts, data, signature)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xf54aca07.
//
// Solidity: function isAuthorized(bytes data, bytes signature) constant returns(bool)
func (_Pamette *PametteCallerSession) IsAuthorized(data []byte, signature []byte) (bool, error) {
	return _Pamette.Contract.IsAuthorized(&_Pamette.CallOpts, data, signature)
}
