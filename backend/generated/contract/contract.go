// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"}],\"name\":\"AddressAssigned\",\"payable\":false,\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsDeposited\",\"payable\":false,\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsWithdrawn\",\"payable\":false,\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"}],\"name\":\"assignAddressToUsername\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"usernames\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountsInEther\",\"type\":\"uint256[]\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"}],\"name\":\"getAddressForUsername\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"temporaryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"usernameToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountInWei\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x05d01f50.
//
// Solidity: function balances(address , string ) view returns(uint256)
func (_Contract *ContractCaller) Balances(opts *bind.CallOpts, arg0 common.Address, arg1 string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x05d01f50.
//
// Solidity: function balances(address , string ) view returns(uint256)
func (_Contract *ContractSession) Balances(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0x05d01f50.
//
// Solidity: function balances(address , string ) view returns(uint256)
func (_Contract *ContractCallerSession) Balances(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _Contract.Contract.Balances(&_Contract.CallOpts, arg0, arg1)
}

// GetAddressForUsername is a free data retrieval call binding the contract method 0x095a0535.
//
// Solidity: function getAddressForUsername(string username) view returns(address)
func (_Contract *ContractCaller) GetAddressForUsername(opts *bind.CallOpts, username string) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAddressForUsername", username)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressForUsername is a free data retrieval call binding the contract method 0x095a0535.
//
// Solidity: function getAddressForUsername(string username) view returns(address)
func (_Contract *ContractSession) GetAddressForUsername(username string) (common.Address, error) {
	return _Contract.Contract.GetAddressForUsername(&_Contract.CallOpts, username)
}

// GetAddressForUsername is a free data retrieval call binding the contract method 0x095a0535.
//
// Solidity: function getAddressForUsername(string username) view returns(address)
func (_Contract *ContractCallerSession) GetAddressForUsername(username string) (common.Address, error) {
	return _Contract.Contract.GetAddressForUsername(&_Contract.CallOpts, username)
}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string username) view returns(uint256)
func (_Contract *ContractCaller) GetBalance(opts *bind.CallOpts, username string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBalance", username)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string username) view returns(uint256)
func (_Contract *ContractSession) GetBalance(username string) (*big.Int, error) {
	return _Contract.Contract.GetBalance(&_Contract.CallOpts, username)
}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string username) view returns(uint256)
func (_Contract *ContractCallerSession) GetBalance(username string) (*big.Int, error) {
	return _Contract.Contract.GetBalance(&_Contract.CallOpts, username)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// TemporaryAddress is a free data retrieval call binding the contract method 0x284c3b3b.
//
// Solidity: function temporaryAddress() view returns(address)
func (_Contract *ContractCaller) TemporaryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "temporaryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TemporaryAddress is a free data retrieval call binding the contract method 0x284c3b3b.
//
// Solidity: function temporaryAddress() view returns(address)
func (_Contract *ContractSession) TemporaryAddress() (common.Address, error) {
	return _Contract.Contract.TemporaryAddress(&_Contract.CallOpts)
}

// TemporaryAddress is a free data retrieval call binding the contract method 0x284c3b3b.
//
// Solidity: function temporaryAddress() view returns(address)
func (_Contract *ContractCallerSession) TemporaryAddress() (common.Address, error) {
	return _Contract.Contract.TemporaryAddress(&_Contract.CallOpts)
}

// UsernameToAddress is a free data retrieval call binding the contract method 0xf825f143.
//
// Solidity: function usernameToAddress(string ) view returns(address)
func (_Contract *ContractCaller) UsernameToAddress(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "usernameToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsernameToAddress is a free data retrieval call binding the contract method 0xf825f143.
//
// Solidity: function usernameToAddress(string ) view returns(address)
func (_Contract *ContractSession) UsernameToAddress(arg0 string) (common.Address, error) {
	return _Contract.Contract.UsernameToAddress(&_Contract.CallOpts, arg0)
}

// UsernameToAddress is a free data retrieval call binding the contract method 0xf825f143.
//
// Solidity: function usernameToAddress(string ) view returns(address)
func (_Contract *ContractCallerSession) UsernameToAddress(arg0 string) (common.Address, error) {
	return _Contract.Contract.UsernameToAddress(&_Contract.CallOpts, arg0)
}

// AssignAddressToUsername is a paid mutator transaction binding the contract method 0xbacd359a.
//
// Solidity: function assignAddressToUsername(string username, address walletAddress) returns()
func (_Contract *ContractTransactor) AssignAddressToUsername(opts *bind.TransactOpts, username string, walletAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "assignAddressToUsername", username, walletAddress)
}

// AssignAddressToUsername is a paid mutator transaction binding the contract method 0xbacd359a.
//
// Solidity: function assignAddressToUsername(string username, address walletAddress) returns()
func (_Contract *ContractSession) AssignAddressToUsername(username string, walletAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AssignAddressToUsername(&_Contract.TransactOpts, username, walletAddress)
}

// AssignAddressToUsername is a paid mutator transaction binding the contract method 0xbacd359a.
//
// Solidity: function assignAddressToUsername(string username, address walletAddress) returns()
func (_Contract *ContractTransactorSession) AssignAddressToUsername(username string, walletAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AssignAddressToUsername(&_Contract.TransactOpts, username, walletAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0x877e531e.
//
// Solidity: function deposit(string[] usernames, uint256[] amountsInEther) payable returns()
func (_Contract *ContractTransactor) Deposit(opts *bind.TransactOpts, usernames []string, amountsInEther []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deposit", usernames, amountsInEther)
}

// Deposit is a paid mutator transaction binding the contract method 0x877e531e.
//
// Solidity: function deposit(string[] usernames, uint256[] amountsInEther) payable returns()
func (_Contract *ContractSession) Deposit(usernames []string, amountsInEther []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, usernames, amountsInEther)
}

// Deposit is a paid mutator transaction binding the contract method 0x877e531e.
//
// Solidity: function deposit(string[] usernames, uint256[] amountsInEther) payable returns()
func (_Contract *ContractTransactorSession) Deposit(usernames []string, amountsInEther []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, usernames, amountsInEther)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string username, uint256 amountInWei) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, username string, amountInWei *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", username, amountInWei)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string username, uint256 amountInWei) returns()
func (_Contract *ContractSession) Withdraw(username string, amountInWei *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, username, amountInWei)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string username, uint256 amountInWei) returns()
func (_Contract *ContractTransactorSession) Withdraw(username string, amountInWei *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, username, amountInWei)
}

// ContractAddressAssignedIterator is returned from FilterAddressAssigned and is used to iterate over the raw logs and unpacked data for AddressAssigned events raised by the Contract contract.
type ContractAddressAssignedIterator struct {
	Event *ContractAddressAssigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAddressAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAddressAssigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAddressAssigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAddressAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAddressAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAddressAssigned represents a AddressAssigned event raised by the Contract contract.
type ContractAddressAssigned struct {
	Username      common.Hash
	WalletAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddressAssigned is a free log retrieval operation binding the contract event 0x5face5b6b37ef00f3594696aab234cb398f4071e830982482a2014002278559f.
//
// Solidity: event AddressAssigned(string indexed username, address walletAddress)
func (_Contract *ContractFilterer) FilterAddressAssigned(opts *bind.FilterOpts, username []string) (*ContractAddressAssignedIterator, error) {

	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AddressAssigned", usernameRule)
	if err != nil {
		return nil, err
	}
	return &ContractAddressAssignedIterator{contract: _Contract.contract, event: "AddressAssigned", logs: logs, sub: sub}, nil
}

// WatchAddressAssigned is a free log subscription operation binding the contract event 0x5face5b6b37ef00f3594696aab234cb398f4071e830982482a2014002278559f.
//
// Solidity: event AddressAssigned(string indexed username, address walletAddress)
func (_Contract *ContractFilterer) WatchAddressAssigned(opts *bind.WatchOpts, sink chan<- *ContractAddressAssigned, username []string) (event.Subscription, error) {

	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AddressAssigned", usernameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAddressAssigned)
				if err := _Contract.contract.UnpackLog(event, "AddressAssigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddressAssigned is a log parse operation binding the contract event 0x5face5b6b37ef00f3594696aab234cb398f4071e830982482a2014002278559f.
//
// Solidity: event AddressAssigned(string indexed username, address walletAddress)
func (_Contract *ContractFilterer) ParseAddressAssigned(log types.Log) (*ContractAddressAssigned, error) {
	event := new(ContractAddressAssigned)
	if err := _Contract.contract.UnpackLog(event, "AddressAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFundsDepositedIterator is returned from FilterFundsDeposited and is used to iterate over the raw logs and unpacked data for FundsDeposited events raised by the Contract contract.
type ContractFundsDepositedIterator struct {
	Event *ContractFundsDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFundsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFundsDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFundsDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFundsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFundsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFundsDeposited represents a FundsDeposited event raised by the Contract contract.
type ContractFundsDeposited struct {
	Receiver common.Address
	Username common.Hash
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFundsDeposited is a free log retrieval operation binding the contract event 0xf05aa94f1ac258aebd610cdbdafea4e9d59a4fd25eda90cd991abbcdf3dc7f14.
//
// Solidity: event FundsDeposited(address indexed receiver, string indexed username, uint256 amount)
func (_Contract *ContractFilterer) FilterFundsDeposited(opts *bind.FilterOpts, receiver []common.Address, username []string) (*ContractFundsDepositedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FundsDeposited", receiverRule, usernameRule)
	if err != nil {
		return nil, err
	}
	return &ContractFundsDepositedIterator{contract: _Contract.contract, event: "FundsDeposited", logs: logs, sub: sub}, nil
}

// WatchFundsDeposited is a free log subscription operation binding the contract event 0xf05aa94f1ac258aebd610cdbdafea4e9d59a4fd25eda90cd991abbcdf3dc7f14.
//
// Solidity: event FundsDeposited(address indexed receiver, string indexed username, uint256 amount)
func (_Contract *ContractFilterer) WatchFundsDeposited(opts *bind.WatchOpts, sink chan<- *ContractFundsDeposited, receiver []common.Address, username []string) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FundsDeposited", receiverRule, usernameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFundsDeposited)
				if err := _Contract.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsDeposited is a log parse operation binding the contract event 0xf05aa94f1ac258aebd610cdbdafea4e9d59a4fd25eda90cd991abbcdf3dc7f14.
//
// Solidity: event FundsDeposited(address indexed receiver, string indexed username, uint256 amount)
func (_Contract *ContractFilterer) ParseFundsDeposited(log types.Log) (*ContractFundsDeposited, error) {
	event := new(ContractFundsDeposited)
	if err := _Contract.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the Contract contract.
type ContractFundsWithdrawnIterator struct {
	Event *ContractFundsWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFundsWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractFundsWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFundsWithdrawn represents a FundsWithdrawn event raised by the Contract contract.
type ContractFundsWithdrawn struct {
	Receiver common.Address
	Username common.Hash
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawn is a free log retrieval operation binding the contract event 0xe72e45c937393adb200d2f0b1e3ca5ab4d3c67ac00d0c8ffdfdcc9bc24d15ba1.
//
// Solidity: event FundsWithdrawn(address indexed receiver, string indexed username, uint256 amount)
func (_Contract *ContractFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, receiver []common.Address, username []string) (*ContractFundsWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FundsWithdrawn", receiverRule, usernameRule)
	if err != nil {
		return nil, err
	}
	return &ContractFundsWithdrawnIterator{contract: _Contract.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0xe72e45c937393adb200d2f0b1e3ca5ab4d3c67ac00d0c8ffdfdcc9bc24d15ba1.
//
// Solidity: event FundsWithdrawn(address indexed receiver, string indexed username, uint256 amount)
func (_Contract *ContractFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractFundsWithdrawn, receiver []common.Address, username []string) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var usernameRule []interface{}
	for _, usernameItem := range username {
		usernameRule = append(usernameRule, usernameItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FundsWithdrawn", receiverRule, usernameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFundsWithdrawn)
				if err := _Contract.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsWithdrawn is a log parse operation binding the contract event 0xe72e45c937393adb200d2f0b1e3ca5ab4d3c67ac00d0c8ffdfdcc9bc24d15ba1.
//
// Solidity: event FundsWithdrawn(address indexed receiver, string indexed username, uint256 amount)
func (_Contract *ContractFilterer) ParseFundsWithdrawn(log types.Log) (*ContractFundsWithdrawn, error) {
	event := new(ContractFundsWithdrawn)
	if err := _Contract.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
