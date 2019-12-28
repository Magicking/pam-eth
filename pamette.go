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
const PametteABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"AllowedUser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"userId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machineId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sharedPassword\",\"type\":\"bytes32\"}],\"name\":\"generateOTP\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"userId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machineId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sharedPassword\",\"type\":\"bytes32\"}],\"name\":\"generateUserHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"userId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machineId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sharedPassword\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"setUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PametteFuncSigs maps the 4-byte function signature to its string representation.
var PametteFuncSigs = map[string]string{
	"a46dfcd0": "AllowedUser(bytes32)",
	"a247e33b": "generateOTP(string,uint256,uint256,bytes32)",
	"2bd5c8d9": "generateUserHash(string,uint256,uint256,bytes32)",
	"210d879e": "isAuthorized(string,uint256,uint256,bytes32,uint256,uint8,bytes32,bytes32)",
	"8da5cb5b": "owner()",
	"9476b3e2": "setUser(bytes32,address)",
}

// PametteBin is the compiled bytecode used for deploying new contracts.
var PametteBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556107c2806100326000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063210d879e146100675780632bd5c8d9146101485780638da5cb5b1461020b5780639476b3e21461022f578063a247e33b1461025d578063a46dfcd014610327575b600080fd5b610134600480360361010081101561007e57600080fd5b81019060208101813564010000000081111561009957600080fd5b8201836020820111156100ab57600080fd5b803590602001918460018302840111640100000000831117156100cd57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060208101359060408101359060608101359060ff6080820135169060a08101359060c00135610344565b604080519115158252519081900360200190f35b6101f96004803603608081101561015e57600080fd5b81019060208101813564010000000081111561017957600080fd5b82018360208201111561018b57600080fd5b803590602001918460018302840111640100000000831117156101ad57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602081013590604001356105a1565b60408051918252519081900360200190f35b61021361064e565b604080516001600160a01b039092168252519081900360200190f35b61025b6004803603604081101561024557600080fd5b50803590602001356001600160a01b031661065d565b005b61030e6004803603608081101561027357600080fd5b81019060208101813564010000000081111561028e57600080fd5b8201836020820111156102a057600080fd5b803590602001918460018302840111640100000000831117156102c257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060208101359060400135610726565b6040805192835260208301919091528051918290030190f35b6102136004803603602081101561033d57600080fd5b5035610771565b600060058543031115610392576040805162461bcd60e51b815260206004820152601160248201527012185cda081d1a5b59481a5b9d985b1a59607a1b604482015290519081900360640190fd5b60006103a08a8a8a8a6105a1565b6000818152600160205260409020549091506001600160a01b03166103fe576040805162461bcd60e51b815260206004820152600f60248201526e155cd95c881a5b995e1a5cdd185b9d608a1b604482015290519081900360640190fd5b60408051602080820184905243828401528251808303840181526060830180855281519183019190912060a084018552601c8083527f19457468657265756d205369676e6564204d6573736167653a0a3133000000006080909501948552945190949193600093859387939091019182918083835b602083106104925780518252601f199092019160209182019101610473565b51815160209384036101000a6000190180199092169116179052920193845250604080518085038152848301808352815191840191909120600090915281850180835281905260ff8e166060860152608085018d905260a085018c905290519095506001945060c080850194929350601f198201928290030190855afa158015610520573d6000803e3d6000fd5b505060408051601f1901516000878152600160205291909120546001600160a01b03908116911614905061058e576040805162461bcd60e51b815260206004820152601060248201526f155cd95c881d5b9c195c9b5a5d1d195960821b604482015290519081900360640190fd5b5060019c9b505050505050505050505050565b6000848484846040516020018080602001858152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b838110156105f95781810151838201526020016105e1565b50505050905090810190601f1680156106265780820380516001836020036101000a031916815260200191505b5095505050505050604051602081830303815290604052805190602001209050949350505050565b6000546001600160a01b031681565b6000546001600160a01b031633146106bc576040805162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722063616e2075736520746869732066756e6374696f6e604482015290519081900360640190fd5b6000828152600160205260409020546001600160a01b0316156106fa57600082815260016020526040902080546001600160a01b0319169055610722565b600082815260016020526040902080546001600160a01b0319166001600160a01b0383161790555b5050565b6000806000610737878787876105a1565b6040805160208082019390935243818301819052825180830384018152606090920190925280519201919091209890975095505050505050565b6001602052600090815260409020546001600160a01b03168156fea2646970667358221220fbe1ed87040886b93d60e3a0bdfd77520566661b61ad065e060f634141d8bcc664736f6c63430006000033"

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

// AllowedUser is a free data retrieval call binding the contract method 0xa46dfcd0.
//
// Solidity: function AllowedUser(bytes32 ) constant returns(address)
func (_Pamette *PametteCaller) AllowedUser(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pamette.contract.Call(opts, out, "AllowedUser", arg0)
	return *ret0, err
}

// AllowedUser is a free data retrieval call binding the contract method 0xa46dfcd0.
//
// Solidity: function AllowedUser(bytes32 ) constant returns(address)
func (_Pamette *PametteSession) AllowedUser(arg0 [32]byte) (common.Address, error) {
	return _Pamette.Contract.AllowedUser(&_Pamette.CallOpts, arg0)
}

// AllowedUser is a free data retrieval call binding the contract method 0xa46dfcd0.
//
// Solidity: function AllowedUser(bytes32 ) constant returns(address)
func (_Pamette *PametteCallerSession) AllowedUser(arg0 [32]byte) (common.Address, error) {
	return _Pamette.Contract.AllowedUser(&_Pamette.CallOpts, arg0)
}

// GenerateOTP is a free data retrieval call binding the contract method 0xa247e33b.
//
// Solidity: function generateOTP(string userName, uint256 userId, uint256 machineId, bytes32 sharedPassword) constant returns(bytes32, uint256)
func (_Pamette *PametteCaller) GenerateOTP(opts *bind.CallOpts, userName string, userId *big.Int, machineId *big.Int, sharedPassword [32]byte) ([32]byte, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Pamette.contract.Call(opts, out, "generateOTP", userName, userId, machineId, sharedPassword)
	return *ret0, *ret1, err
}

// GenerateOTP is a free data retrieval call binding the contract method 0xa247e33b.
//
// Solidity: function generateOTP(string userName, uint256 userId, uint256 machineId, bytes32 sharedPassword) constant returns(bytes32, uint256)
func (_Pamette *PametteSession) GenerateOTP(userName string, userId *big.Int, machineId *big.Int, sharedPassword [32]byte) ([32]byte, *big.Int, error) {
	return _Pamette.Contract.GenerateOTP(&_Pamette.CallOpts, userName, userId, machineId, sharedPassword)
}

// GenerateOTP is a free data retrieval call binding the contract method 0xa247e33b.
//
// Solidity: function generateOTP(string userName, uint256 userId, uint256 machineId, bytes32 sharedPassword) constant returns(bytes32, uint256)
func (_Pamette *PametteCallerSession) GenerateOTP(userName string, userId *big.Int, machineId *big.Int, sharedPassword [32]byte) ([32]byte, *big.Int, error) {
	return _Pamette.Contract.GenerateOTP(&_Pamette.CallOpts, userName, userId, machineId, sharedPassword)
}

// GenerateUserHash is a free data retrieval call binding the contract method 0x2bd5c8d9.
//
// Solidity: function generateUserHash(string userName, uint256 userId, uint256 machineId, bytes32 sharedPassword) constant returns(bytes32)
func (_Pamette *PametteCaller) GenerateUserHash(opts *bind.CallOpts, userName string, userId *big.Int, machineId *big.Int, sharedPassword [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Pamette.contract.Call(opts, out, "generateUserHash", userName, userId, machineId, sharedPassword)
	return *ret0, err
}

// GenerateUserHash is a free data retrieval call binding the contract method 0x2bd5c8d9.
//
// Solidity: function generateUserHash(string userName, uint256 userId, uint256 machineId, bytes32 sharedPassword) constant returns(bytes32)
func (_Pamette *PametteSession) GenerateUserHash(userName string, userId *big.Int, machineId *big.Int, sharedPassword [32]byte) ([32]byte, error) {
	return _Pamette.Contract.GenerateUserHash(&_Pamette.CallOpts, userName, userId, machineId, sharedPassword)
}

// GenerateUserHash is a free data retrieval call binding the contract method 0x2bd5c8d9.
//
// Solidity: function generateUserHash(string userName, uint256 userId, uint256 machineId, bytes32 sharedPassword) constant returns(bytes32)
func (_Pamette *PametteCallerSession) GenerateUserHash(userName string, userId *big.Int, machineId *big.Int, sharedPassword [32]byte) ([32]byte, error) {
	return _Pamette.Contract.GenerateUserHash(&_Pamette.CallOpts, userName, userId, machineId, sharedPassword)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x210d879e.
//
// Solidity: function isAuthorized(string userName, uint256 userId, uint256 machineId, bytes32 sharedPassword, uint256 blockNumber, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_Pamette *PametteCaller) IsAuthorized(opts *bind.CallOpts, userName string, userId *big.Int, machineId *big.Int, sharedPassword [32]byte, blockNumber *big.Int, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pamette.contract.Call(opts, out, "isAuthorized", userName, userId, machineId, sharedPassword, blockNumber, v, r, s)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0x210d879e.
//
// Solidity: function isAuthorized(string userName, uint256 userId, uint256 machineId, bytes32 sharedPassword, uint256 blockNumber, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_Pamette *PametteSession) IsAuthorized(userName string, userId *big.Int, machineId *big.Int, sharedPassword [32]byte, blockNumber *big.Int, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Pamette.Contract.IsAuthorized(&_Pamette.CallOpts, userName, userId, machineId, sharedPassword, blockNumber, v, r, s)
}

// IsAuthorized is a free data retrieval call binding the contract method 0x210d879e.
//
// Solidity: function isAuthorized(string userName, uint256 userId, uint256 machineId, bytes32 sharedPassword, uint256 blockNumber, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_Pamette *PametteCallerSession) IsAuthorized(userName string, userId *big.Int, machineId *big.Int, sharedPassword [32]byte, blockNumber *big.Int, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Pamette.Contract.IsAuthorized(&_Pamette.CallOpts, userName, userId, machineId, sharedPassword, blockNumber, v, r, s)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pamette *PametteCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pamette.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pamette *PametteSession) Owner() (common.Address, error) {
	return _Pamette.Contract.Owner(&_Pamette.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pamette *PametteCallerSession) Owner() (common.Address, error) {
	return _Pamette.Contract.Owner(&_Pamette.CallOpts)
}

// SetUser is a paid mutator transaction binding the contract method 0x9476b3e2.
//
// Solidity: function setUser(bytes32 hash, address user) returns()
func (_Pamette *PametteTransactor) SetUser(opts *bind.TransactOpts, hash [32]byte, user common.Address) (*types.Transaction, error) {
	return _Pamette.contract.Transact(opts, "setUser", hash, user)
}

// SetUser is a paid mutator transaction binding the contract method 0x9476b3e2.
//
// Solidity: function setUser(bytes32 hash, address user) returns()
func (_Pamette *PametteSession) SetUser(hash [32]byte, user common.Address) (*types.Transaction, error) {
	return _Pamette.Contract.SetUser(&_Pamette.TransactOpts, hash, user)
}

// SetUser is a paid mutator transaction binding the contract method 0x9476b3e2.
//
// Solidity: function setUser(bytes32 hash, address user) returns()
func (_Pamette *PametteTransactorSession) SetUser(hash [32]byte, user common.Address) (*types.Transaction, error) {
	return _Pamette.Contract.SetUser(&_Pamette.TransactOpts, hash, user)
}

