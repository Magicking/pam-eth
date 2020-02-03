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
const PametteABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"AllowedUser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"userName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sharedPassword\",\"type\":\"string\"}],\"name\":\"generateOTP\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"userName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sharedPassword\",\"type\":\"string\"}],\"name\":\"generateUserHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"userId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"userName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sharedPassword\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"setUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PametteFuncSigs maps the 4-byte function signature to its string representation.
var PametteFuncSigs = map[string]string{
	"a46dfcd0": "AllowedUser(bytes32)",
	"4d0cb814": "generateOTP(uint256,string,string)",
	"098d512c": "generateUserHash(uint256,string,string)",
	"b6c50074": "isAuthorized(uint256,string,string,uint256,uint8,bytes32,bytes32)",
	"8da5cb5b": "owner()",
	"9476b3e2": "setUser(bytes32,address)",
}

// PametteBin is the compiled bytecode used for deploying new contracts.
var PametteBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b0319163317905560408051808201825260098152686d616769636b696e6760b81b60208083019190915282518084019093526004835263746f746f60e01b9083015261008791610078916103e8916001600160e01b0361008c16565b336001600160e01b0361015516565b61021e565b60008284836040516020018084805190602001908083835b602083106100c35780518252601f1990920191602091820191016100a4565b51815160209384036101000a60001901801990921691161790529201858152845190830192850191508083835b6020831061010f5780518252601f1990920191602091820191016100f0565b6001836020036101000a03801982511681845116808217855250505050505090500193505050506040516020818303038152906040528051906020012090509392505050565b6000546001600160a01b031633146101b4576040805162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722063616e2075736520746869732066756e6374696f6e604482015290519081900360640190fd5b6000828152600160205260409020546001600160a01b0316156101f257600082815260016020526040902080546001600160a01b031916905561021a565b600082815260016020526040902080546001600160a01b0319166001600160a01b0383161790555b5050565b6108c28061022d6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063098d512c146100675780634d0cb814146101a95780638da5cb5b146102f25780639476b3e214610316578063a46dfcd014610344578063b6c5007414610361575b600080fd5b6101976004803603606081101561007d57600080fd5b81359190810190604081016020820135600160201b81111561009e57600080fd5b8201836020820111156100b057600080fd5b803590602001918460018302840111600160201b831117156100d157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561012357600080fd5b82018360208201111561013557600080fd5b803590602001918460018302840111600160201b8311171561015657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104a5945050505050565b60408051918252519081900360200190f35b6102d9600480360360608110156101bf57600080fd5b81359190810190604081016020820135600160201b8111156101e057600080fd5b8201836020820111156101f257600080fd5b803590602001918460018302840111600160201b8311171561021357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561026557600080fd5b82018360208201111561027757600080fd5b803590602001918460018302840111600160201b8311171561029857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061056e945050505050565b6040805192835260208301919091528051918290030190f35b6102fa6105b7565b604080516001600160a01b039092168252519081900360200190f35b6103426004803603604081101561032c57600080fd5b50803590602001356001600160a01b03166105c6565b005b6102fa6004803603602081101561035a57600080fd5b503561068f565b610197600480360360e081101561037757600080fd5b81359190810190604081016020820135600160201b81111561039857600080fd5b8201836020820111156103aa57600080fd5b803590602001918460018302840111600160201b831117156103cb57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561041d57600080fd5b82018360208201111561042f57600080fd5b803590602001918460018302840111600160201b8311171561045057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505060ff6020830135169160408101359150606001356106aa565b60008284836040516020018084805190602001908083835b602083106104dc5780518252601f1990920191602091820191016104bd565b51815160209384036101000a60001901801990921691161790529201858152845190830192850191508083835b602083106105285780518252601f199092019160209182019101610509565b6001836020036101000a03801982511681845116808217855250505050505090500193505050506040516020818303038152906040528051906020012090509392505050565b600080600061057e8686866104a5565b60408051602080820193909352438183018190528251808303840181526060909201909252805192019190912097909650945050505050565b6000546001600160a01b031681565b6000546001600160a01b03163314610625576040805162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f776e65722063616e2075736520746869732066756e6374696f6e604482015290519081900360640190fd5b6000828152600160205260409020546001600160a01b03161561066357600082815260016020526040902080546001600160a01b031916905561068b565b600082815260016020526040902080546001600160a01b0319166001600160a01b0383161790555b5050565b6001602052600090815260409020546001600160a01b031681565b6000600585430311156106bf5750600161085c565b60006106cc8989896104a5565b6000818152600160205260409020549091506001600160a01b03166106f557600291505061085c565b6040805160208082018490528183018990528251808303840181526060830180855281519183019190912060a084018552601c8083527f19457468657265756d205369676e6564204d6573736167653a0a3332000000006080909501948552945190949193600093859387939091019182918083835b6020831061078a5780518252601f19909201916020918201910161076b565b51815160209384036101000a6000190180199092169116179052920193845250604080518085038152848301808352815191840191909120600090915281850180835281905260ff8e166060860152608085018d905260a085018c905290519095506001945060c080850194929350601f198201928290030190855afa158015610818573d6000803e3d6000fd5b505060408051601f1901516000878152600160205291909120546001600160a01b03908116911614905061085357600394505050505061085c565b60009450505050505b97965050505050505056fea265627a7a72315820efeaff7d12d8b3ee6de0af670820b95b84d937add3e00e0083702c313f43426064736f6c637828302e352e31342d646576656c6f702e323031392e31322e372b636f6d6d69742e31666531343539620058"

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

// GenerateOTP is a free data retrieval call binding the contract method 0x4d0cb814.
//
// Solidity: function generateOTP(uint256 userId, string userName, string sharedPassword) constant returns(bytes32, uint256)
func (_Pamette *PametteCaller) GenerateOTP(opts *bind.CallOpts, userId *big.Int, userName string, sharedPassword string) ([32]byte, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Pamette.contract.Call(opts, out, "generateOTP", userId, userName, sharedPassword)
	return *ret0, *ret1, err
}

// GenerateOTP is a free data retrieval call binding the contract method 0x4d0cb814.
//
// Solidity: function generateOTP(uint256 userId, string userName, string sharedPassword) constant returns(bytes32, uint256)
func (_Pamette *PametteSession) GenerateOTP(userId *big.Int, userName string, sharedPassword string) ([32]byte, *big.Int, error) {
	return _Pamette.Contract.GenerateOTP(&_Pamette.CallOpts, userId, userName, sharedPassword)
}

// GenerateOTP is a free data retrieval call binding the contract method 0x4d0cb814.
//
// Solidity: function generateOTP(uint256 userId, string userName, string sharedPassword) constant returns(bytes32, uint256)
func (_Pamette *PametteCallerSession) GenerateOTP(userId *big.Int, userName string, sharedPassword string) ([32]byte, *big.Int, error) {
	return _Pamette.Contract.GenerateOTP(&_Pamette.CallOpts, userId, userName, sharedPassword)
}

// GenerateUserHash is a free data retrieval call binding the contract method 0x098d512c.
//
// Solidity: function generateUserHash(uint256 userId, string userName, string sharedPassword) constant returns(bytes32)
func (_Pamette *PametteCaller) GenerateUserHash(opts *bind.CallOpts, userId *big.Int, userName string, sharedPassword string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Pamette.contract.Call(opts, out, "generateUserHash", userId, userName, sharedPassword)
	return *ret0, err
}

// GenerateUserHash is a free data retrieval call binding the contract method 0x098d512c.
//
// Solidity: function generateUserHash(uint256 userId, string userName, string sharedPassword) constant returns(bytes32)
func (_Pamette *PametteSession) GenerateUserHash(userId *big.Int, userName string, sharedPassword string) ([32]byte, error) {
	return _Pamette.Contract.GenerateUserHash(&_Pamette.CallOpts, userId, userName, sharedPassword)
}

// GenerateUserHash is a free data retrieval call binding the contract method 0x098d512c.
//
// Solidity: function generateUserHash(uint256 userId, string userName, string sharedPassword) constant returns(bytes32)
func (_Pamette *PametteCallerSession) GenerateUserHash(userId *big.Int, userName string, sharedPassword string) ([32]byte, error) {
	return _Pamette.Contract.GenerateUserHash(&_Pamette.CallOpts, userId, userName, sharedPassword)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xb6c50074.
//
// Solidity: function isAuthorized(uint256 userId, string userName, string sharedPassword, uint256 blockNumber, uint8 v, bytes32 r, bytes32 s) constant returns(uint256)
func (_Pamette *PametteCaller) IsAuthorized(opts *bind.CallOpts, userId *big.Int, userName string, sharedPassword string, blockNumber *big.Int, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pamette.contract.Call(opts, out, "isAuthorized", userId, userName, sharedPassword, blockNumber, v, r, s)
	return *ret0, err
}

// IsAuthorized is a free data retrieval call binding the contract method 0xb6c50074.
//
// Solidity: function isAuthorized(uint256 userId, string userName, string sharedPassword, uint256 blockNumber, uint8 v, bytes32 r, bytes32 s) constant returns(uint256)
func (_Pamette *PametteSession) IsAuthorized(userId *big.Int, userName string, sharedPassword string, blockNumber *big.Int, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _Pamette.Contract.IsAuthorized(&_Pamette.CallOpts, userId, userName, sharedPassword, blockNumber, v, r, s)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xb6c50074.
//
// Solidity: function isAuthorized(uint256 userId, string userName, string sharedPassword, uint256 blockNumber, uint8 v, bytes32 r, bytes32 s) constant returns(uint256)
func (_Pamette *PametteCallerSession) IsAuthorized(userId *big.Int, userName string, sharedPassword string, blockNumber *big.Int, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _Pamette.Contract.IsAuthorized(&_Pamette.CallOpts, userId, userName, sharedPassword, blockNumber, v, r, s)
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
