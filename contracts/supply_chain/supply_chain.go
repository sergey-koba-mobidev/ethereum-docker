// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SupplyChainABI is the input ABI used to generate the binding from.
const SupplyChainABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_uid\",\"type\":\"uint256\"}],\"name\":\"getProductInfo\",\"outputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"state\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"creator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberOfProducts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"data\",\"type\":\"string\"}],\"name\":\"addProduct\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"transferProduct\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_description\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SupplyChainBin is the compiled bytecode used for deploying new contracts.
const SupplyChainBin = `0x6060604052341561000f57600080fd5b60405161086e38038061086e8339810160405280805182019190602001805190910190506000828051610046929160200190610067565b50600181805161005a929160200190610067565b5050600060045550610102565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a857805160ff19168380011785556100d5565b828001600101855582156100d5579182015b828111156100d55782518255916020019190600101906100ba565b506100e19291506100e5565b5090565b6100ff91905b808211156100e157600081556001016100eb565b90565b61075d806101116000396000f3006060604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461007c5780637284e4161461010657806379a85e6c14610119578063c66301b2146101ce578063dfe00c51146101f3578063e43059e514610258575b600080fd5b341561008757600080fd5b61008f61027a565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100cb5780820151838201526020016100b3565b50505050905090810190601f1680156100f85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561011157600080fd5b61008f610318565b341561012457600080fd5b61012f600435610383565b60405185815260208101859052600160a060020a0380841660608301528216608082015260a06040820181815290820185818151815260200191508051906020019080838360005b8381101561018f578082015183820152602001610177565b50505050905090810190601f1680156101bc5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b34156101d957600080fd5b6101e16104a5565b60405190815260200160405180910390f35b34156101fe57600080fd5b61024460046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506104ab95505050505050565b604051901515815260200160405180910390f35b341561026357600080fd5b610244600435600160a060020a03602435166105e0565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103105780601f106102e557610100808354040283529160200191610310565b820191906000526020600020905b8154815290600101906020018083116102f357829003601f168201915b505050505081565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103105780601f106102e557610100808354040283529160200191610310565b60008061038e610684565b6000848152600360205260408120600481015482919074010000000000000000000000000000000000000000900460ff1615156103ca57600080fd5b8054600180830154600484015460038501546002808701805494959094600160a060020a039485169490931692859261010092821615929092026000190116046020601f8201819004810201604051908101604052809291908181526020018280546001816001161561010002031660029004801561048a5780601f1061045f5761010080835404028352916020019161048a565b820191906000526020600020905b81548152906001019060200180831161046d57829003601f168201915b50505050509250955095509550955095505091939590929450565b60045481565b600480546001019055600060c0604051908101604090815260045480835260006020808501829052838501879052600160a060020a033316606086018190526080860152600160a08601529181526003909152208151815560208201518160010155604082015181600201908051610527929160200190610696565b50606082015160038201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055608082015160048201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560a082015160049091018054911515740100000000000000000000000000000000000000000274ff0000000000000000000000000000000000000000199092169190911790555060019050919050565b60008281526003602052604081206004015474010000000000000000000000000000000000000000900460ff16151561061857600080fd5b60008381526003602052604090206004015433600160a060020a0390811691161461064257600080fd5b5060008281526003602052604090206004018054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116179055600192915050565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106d757805160ff1916838001178555610704565b82800160010185558215610704579182015b828111156107045782518255916020019190600101906106e9565b50610710929150610714565b5090565b61072e91905b80821115610710576000815560010161071a565b905600a165627a7a72305820c920f095f804d39269708e27dcbc43e1c923d0f94e0d85993b24fcae68e618460029`

// DeploySupplyChain deploys a new Ethereum contract, binding an instance of SupplyChain to it.
func DeploySupplyChain(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _description string) (common.Address, *types.Transaction, *SupplyChain, error) {
	parsed, err := abi.JSON(strings.NewReader(SupplyChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SupplyChainBin), backend, _name, _description)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SupplyChain{SupplyChainCaller: SupplyChainCaller{contract: contract}, SupplyChainTransactor: SupplyChainTransactor{contract: contract}}, nil
}

// SupplyChain is an auto generated Go binding around an Ethereum contract.
type SupplyChain struct {
	SupplyChainCaller     // Read-only binding to the contract
	SupplyChainTransactor // Write-only binding to the contract
}

// SupplyChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type SupplyChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SupplyChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SupplyChainSession struct {
	Contract     *SupplyChain      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SupplyChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SupplyChainCallerSession struct {
	Contract *SupplyChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SupplyChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SupplyChainTransactorSession struct {
	Contract     *SupplyChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SupplyChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type SupplyChainRaw struct {
	Contract *SupplyChain // Generic contract binding to access the raw methods on
}

// SupplyChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SupplyChainCallerRaw struct {
	Contract *SupplyChainCaller // Generic read-only contract binding to access the raw methods on
}

// SupplyChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SupplyChainTransactorRaw struct {
	Contract *SupplyChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSupplyChain creates a new instance of SupplyChain, bound to a specific deployed contract.
func NewSupplyChain(address common.Address, backend bind.ContractBackend) (*SupplyChain, error) {
	contract, err := bindSupplyChain(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SupplyChain{SupplyChainCaller: SupplyChainCaller{contract: contract}, SupplyChainTransactor: SupplyChainTransactor{contract: contract}}, nil
}

// NewSupplyChainCaller creates a new read-only instance of SupplyChain, bound to a specific deployed contract.
func NewSupplyChainCaller(address common.Address, caller bind.ContractCaller) (*SupplyChainCaller, error) {
	contract, err := bindSupplyChain(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SupplyChainCaller{contract: contract}, nil
}

// NewSupplyChainTransactor creates a new write-only instance of SupplyChain, bound to a specific deployed contract.
func NewSupplyChainTransactor(address common.Address, transactor bind.ContractTransactor) (*SupplyChainTransactor, error) {
	contract, err := bindSupplyChain(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SupplyChainTransactor{contract: contract}, nil
}

// bindSupplyChain binds a generic wrapper to an already deployed contract.
func bindSupplyChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SupplyChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupplyChain *SupplyChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SupplyChain.Contract.SupplyChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupplyChain *SupplyChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupplyChain.Contract.SupplyChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupplyChain *SupplyChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupplyChain.Contract.SupplyChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupplyChain *SupplyChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SupplyChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupplyChain *SupplyChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupplyChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupplyChain *SupplyChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupplyChain.Contract.contract.Transact(opts, method, params...)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(string)
func (_SupplyChain *SupplyChainCaller) Description(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SupplyChain.contract.Call(opts, out, "description")
	return *ret0, err
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(string)
func (_SupplyChain *SupplyChainSession) Description() (string, error) {
	return _SupplyChain.Contract.Description(&_SupplyChain.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(string)
func (_SupplyChain *SupplyChainCallerSession) Description() (string, error) {
	return _SupplyChain.Contract.Description(&_SupplyChain.CallOpts)
}

// GetProductInfo is a free data retrieval call binding the contract method 0x79a85e6c.
//
// Solidity: function getProductInfo(_uid uint256) constant returns(uid uint256, state uint256, data string, owner address, creator address)
func (_SupplyChain *SupplyChainCaller) GetProductInfo(opts *bind.CallOpts, _uid *big.Int) (struct {
	Uid     *big.Int
	State   *big.Int
	Data    string
	Owner   common.Address
	Creator common.Address
}, error) {
	ret := new(struct {
		Uid     *big.Int
		State   *big.Int
		Data    string
		Owner   common.Address
		Creator common.Address
	})
	out := ret
	err := _SupplyChain.contract.Call(opts, out, "getProductInfo", _uid)
	return *ret, err
}

// GetProductInfo is a free data retrieval call binding the contract method 0x79a85e6c.
//
// Solidity: function getProductInfo(_uid uint256) constant returns(uid uint256, state uint256, data string, owner address, creator address)
func (_SupplyChain *SupplyChainSession) GetProductInfo(_uid *big.Int) (struct {
	Uid     *big.Int
	State   *big.Int
	Data    string
	Owner   common.Address
	Creator common.Address
}, error) {
	return _SupplyChain.Contract.GetProductInfo(&_SupplyChain.CallOpts, _uid)
}

// GetProductInfo is a free data retrieval call binding the contract method 0x79a85e6c.
//
// Solidity: function getProductInfo(_uid uint256) constant returns(uid uint256, state uint256, data string, owner address, creator address)
func (_SupplyChain *SupplyChainCallerSession) GetProductInfo(_uid *big.Int) (struct {
	Uid     *big.Int
	State   *big.Int
	Data    string
	Owner   common.Address
	Creator common.Address
}, error) {
	return _SupplyChain.Contract.GetProductInfo(&_SupplyChain.CallOpts, _uid)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SupplyChain *SupplyChainCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SupplyChain.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SupplyChain *SupplyChainSession) Name() (string, error) {
	return _SupplyChain.Contract.Name(&_SupplyChain.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SupplyChain *SupplyChainCallerSession) Name() (string, error) {
	return _SupplyChain.Contract.Name(&_SupplyChain.CallOpts)
}

// NumberOfProducts is a free data retrieval call binding the contract method 0xc66301b2.
//
// Solidity: function numberOfProducts() constant returns(uint256)
func (_SupplyChain *SupplyChainCaller) NumberOfProducts(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SupplyChain.contract.Call(opts, out, "numberOfProducts")
	return *ret0, err
}

// NumberOfProducts is a free data retrieval call binding the contract method 0xc66301b2.
//
// Solidity: function numberOfProducts() constant returns(uint256)
func (_SupplyChain *SupplyChainSession) NumberOfProducts() (*big.Int, error) {
	return _SupplyChain.Contract.NumberOfProducts(&_SupplyChain.CallOpts)
}

// NumberOfProducts is a free data retrieval call binding the contract method 0xc66301b2.
//
// Solidity: function numberOfProducts() constant returns(uint256)
func (_SupplyChain *SupplyChainCallerSession) NumberOfProducts() (*big.Int, error) {
	return _SupplyChain.Contract.NumberOfProducts(&_SupplyChain.CallOpts)
}

// AddProduct is a paid mutator transaction binding the contract method 0xdfe00c51.
//
// Solidity: function addProduct(data string) returns(success bool)
func (_SupplyChain *SupplyChainTransactor) AddProduct(opts *bind.TransactOpts, data string) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "addProduct", data)
}

// AddProduct is a paid mutator transaction binding the contract method 0xdfe00c51.
//
// Solidity: function addProduct(data string) returns(success bool)
func (_SupplyChain *SupplyChainSession) AddProduct(data string) (*types.Transaction, error) {
	return _SupplyChain.Contract.AddProduct(&_SupplyChain.TransactOpts, data)
}

// AddProduct is a paid mutator transaction binding the contract method 0xdfe00c51.
//
// Solidity: function addProduct(data string) returns(success bool)
func (_SupplyChain *SupplyChainTransactorSession) AddProduct(data string) (*types.Transaction, error) {
	return _SupplyChain.Contract.AddProduct(&_SupplyChain.TransactOpts, data)
}

// TransferProduct is a paid mutator transaction binding the contract method 0xe43059e5.
//
// Solidity: function transferProduct(uid uint256, addr address) returns(success bool)
func (_SupplyChain *SupplyChainTransactor) TransferProduct(opts *bind.TransactOpts, uid *big.Int, addr common.Address) (*types.Transaction, error) {
	return _SupplyChain.contract.Transact(opts, "transferProduct", uid, addr)
}

// TransferProduct is a paid mutator transaction binding the contract method 0xe43059e5.
//
// Solidity: function transferProduct(uid uint256, addr address) returns(success bool)
func (_SupplyChain *SupplyChainSession) TransferProduct(uid *big.Int, addr common.Address) (*types.Transaction, error) {
	return _SupplyChain.Contract.TransferProduct(&_SupplyChain.TransactOpts, uid, addr)
}

// TransferProduct is a paid mutator transaction binding the contract method 0xe43059e5.
//
// Solidity: function transferProduct(uid uint256, addr address) returns(success bool)
func (_SupplyChain *SupplyChainTransactorSession) TransferProduct(uid *big.Int, addr common.Address) (*types.Transaction, error) {
	return _SupplyChain.Contract.TransferProduct(&_SupplyChain.TransactOpts, uid, addr)
}
