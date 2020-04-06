// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifierlist

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

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VerifierListABI is the input ABI used to generate the binding from.
const VerifierListABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"verifier\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeOfVerifier\",\"type\":\"string\"}],\"name\":\"VerifierAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"verifier\",\"type\":\"string\"}],\"name\":\"VerifierUpdated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_verifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_typeOfVerifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_verifierParams\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"addVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_verifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_verifierParams\",\"type\":\"string\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"verifierList\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"verifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"typeOfVerifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verifierParams\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isCreated\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VerifierListFuncSigs maps the 4-byte function signature to its string representation.
var VerifierListFuncSigs = map[string]string{
	"43e9c5b6": "addVerifier(string,string,string,address)",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
	"59fb00f9": "updateVerifier(string,string)",
	"57762ddd": "verifierList(uint256)",
	"ac205deb": "verifiers(string)",
}

// VerifierListBin is the compiled bytecode used for deploying new contracts.
var VerifierListBin = "0x608060405260006100176001600160e01b0361006616565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006a565b3390565b610fb0806100796000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638da5cb5b1161005b5780638da5cb5b146102fe5780638f32d59b14610322578063ac205deb1461033e578063f2fde38b146104e457610088565b806343e9c5b61461008d57806357762ddd146101a657806359fb00f914610238578063715018a6146102f6575b600080fd5b6101a4600480360360808110156100a357600080fd5b810190602081018135600160201b8111156100bd57600080fd5b8201836020820111156100cf57600080fd5b803590602001918460018302840111600160201b831117156100f057600080fd5b919390929091602081019035600160201b81111561010d57600080fd5b82018360208201111561011f57600080fd5b803590602001918460018302840111600160201b8311171561014057600080fd5b919390929091602081019035600160201b81111561015d57600080fd5b82018360208201111561016f57600080fd5b803590602001918460018302840111600160201b8311171561019057600080fd5b9193509150356001600160a01b031661050a565b005b6101c3600480360360208110156101bc57600080fd5b503561084e565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101fd5781810151838201526020016101e5565b50505050905090810190601f16801561022a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101a46004803603604081101561024e57600080fd5b810190602081018135600160201b81111561026857600080fd5b82018360208201111561027a57600080fd5b803590602001918460018302840111600160201b8311171561029b57600080fd5b919390929091602081019035600160201b8111156102b857600080fd5b8201836020820111156102ca57600080fd5b803590602001918460018302840111600160201b831117156102eb57600080fd5b5090925090506108f4565b6101a4610b32565b610306610bce565b604080516001600160a01b039092168252519081900360200190f35b61032a610bde565b604080519115158252519081900360200190f35b6103e26004803603602081101561035457600080fd5b810190602081018135600160201b81111561036e57600080fd5b82018360208201111561038057600080fd5b803590602001918460018302840111600160201b831117156103a157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c02945050505050565b60405180856001600160a01b03166001600160a01b03168152602001806020018060200184151515158152602001838103835286818151815260200191508051906020019080838360005b8381101561044557818101518382015260200161042d565b50505050905090810190601f1680156104725780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156104a557818101518382015260200161048d565b50505050905090810190601f1680156104d25780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b6101a4600480360360208110156104fa57600080fd5b50356001600160a01b0316610d5d565b610512610bde565b61055c576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b86868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604051835160019350849250819060208401908083835b602083106105c65780518252601f1990920191602091820191016105a7565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206003015460ff1615915061064e9050576040805162461bcd60e51b815260206004820152601760248201527f766572696669657220616c726561647920657869737473000000000000000000604482015290519081900360640190fd5b6040518060800160405280836001600160a01b0316815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f880181900481028201810190925286815291810191908790879081908401838280828437600092019190915250505090825250600160209091018190526040518a908a908083838082843791909101948552505060405160209381900384019020845181546001600160a01b0319166001600160a01b03909116178155848401518051919461074094506001860193500190610e75565b506040820151805161075c916002840191602090910190610e75565b50606091909101516003909101805460ff191691151591909117905560028054600181018083556000929092526107b6907f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace018a8a610ef3565b50507fda2f0cdcc06a08d65b7aab3dc839c3414686e2d8cd4aa2d4b3e14313f8df34e8888888886040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a15050505050505050565b6002818154811061085b57fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152935090918301828280156108ec5780601f106108c1576101008083540402835291602001916108ec565b820191906000526020600020905b8154815290600101906020018083116108cf57829003601f168201915b505050505081565b83838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604051835133935060019250849190819060208401908083835b602083106109615780518252601f199092019160209182019101610942565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b03169290921491506109a7905057600080fd5b84848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604051835160019350849250819060208401908083835b60208310610a115780518252601f1990920191602091820191016109f2565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206003015460ff169150610a909050576040805162461bcd60e51b81526020600482015260156024820152741d995c9a599a595c88191bd95cdb9d08195e1a5cdd605a1b604482015290519081900360640190fd5b83836001888860405180838380828437808301925050509250505090815260200160405180910390206002019190610ac9929190610ef3565b507f24d1bc6b4f3a4d03649d6c40376d48f89866293956598a1e6c5f974fbd0021e4868660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a1505050505050565b610b3a610bde565b610b84576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b03165b90565b600080546001600160a01b0316610bf3610dbb565b6001600160a01b031614905090565b8051808201602090810180516001808352938301948301949094209390528254838301805460408051601f60026000199885161561010002989098019093169690960491820185900485028601850190528085526001600160a01b03909216949392909190830182828015610cb85780601f10610c8d57610100808354040283529160200191610cb8565b820191906000526020600020905b815481529060010190602001808311610c9b57829003601f168201915b50505060028085018054604080516020601f6000196101006001871615020190941695909504928301859004850281018501909152818152959695945090925090830182828015610d4a5780601f10610d1f57610100808354040283529160200191610d4a565b820191906000526020600020905b815481529060010190602001808311610d2d57829003601f168201915b5050506003909301549192505060ff1684565b610d65610bde565b610daf576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b610db881610dbf565b50565b3390565b6001600160a01b038116610e1a576040805162461bcd60e51b815260206004820152601860248201527f4f776e61626c653a206e6f207a65726f20616464726573730000000000000000604482015290519081900360640190fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610eb657805160ff1916838001178555610ee3565b82800160010185558215610ee3579182015b82811115610ee3578251825591602001919060010190610ec8565b50610eef929150610f61565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f345782800160ff19823516178555610ee3565b82800160010185558215610ee3579182015b82811115610ee3578235825591602001919060010190610f46565b610bdb91905b80821115610eef5760008155600101610f6756fea265627a7a72315820037c931da0bcf5679af5da9451bf07eaee331b6f9bb7f4ce1c5f43bde4bbf01164736f6c634300050f0032"

// DeployVerifierList deploys a new Ethereum contract, binding an instance of VerifierList to it.
func DeployVerifierList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VerifierList, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierListABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VerifierListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierList{VerifierListCaller: VerifierListCaller{contract: contract}, VerifierListTransactor: VerifierListTransactor{contract: contract}, VerifierListFilterer: VerifierListFilterer{contract: contract}}, nil
}

// VerifierList is an auto generated Go binding around an Ethereum contract.
type VerifierList struct {
	VerifierListCaller     // Read-only binding to the contract
	VerifierListTransactor // Write-only binding to the contract
	VerifierListFilterer   // Log filterer for contract events
}

// VerifierListCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierListSession struct {
	Contract     *VerifierList     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierListCallerSession struct {
	Contract *VerifierListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VerifierListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierListTransactorSession struct {
	Contract     *VerifierListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VerifierListRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierListRaw struct {
	Contract *VerifierList // Generic contract binding to access the raw methods on
}

// VerifierListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierListCallerRaw struct {
	Contract *VerifierListCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierListTransactorRaw struct {
	Contract *VerifierListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifierList creates a new instance of VerifierList, bound to a specific deployed contract.
func NewVerifierList(address common.Address, backend bind.ContractBackend) (*VerifierList, error) {
	contract, err := bindVerifierList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierList{VerifierListCaller: VerifierListCaller{contract: contract}, VerifierListTransactor: VerifierListTransactor{contract: contract}, VerifierListFilterer: VerifierListFilterer{contract: contract}}, nil
}

// NewVerifierListCaller creates a new read-only instance of VerifierList, bound to a specific deployed contract.
func NewVerifierListCaller(address common.Address, caller bind.ContractCaller) (*VerifierListCaller, error) {
	contract, err := bindVerifierList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierListCaller{contract: contract}, nil
}

// NewVerifierListTransactor creates a new write-only instance of VerifierList, bound to a specific deployed contract.
func NewVerifierListTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierListTransactor, error) {
	contract, err := bindVerifierList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierListTransactor{contract: contract}, nil
}

// NewVerifierListFilterer creates a new log filterer instance of VerifierList, bound to a specific deployed contract.
func NewVerifierListFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierListFilterer, error) {
	contract, err := bindVerifierList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierListFilterer{contract: contract}, nil
}

// bindVerifierList binds a generic wrapper to an already deployed contract.
func bindVerifierList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierList *VerifierListRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VerifierList.Contract.VerifierListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierList *VerifierListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierList.Contract.VerifierListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierList *VerifierListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierList.Contract.VerifierListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierList *VerifierListCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VerifierList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierList *VerifierListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierList *VerifierListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierList.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_VerifierList *VerifierListCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VerifierList.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_VerifierList *VerifierListSession) IsOwner() (bool, error) {
	return _VerifierList.Contract.IsOwner(&_VerifierList.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_VerifierList *VerifierListCallerSession) IsOwner() (bool, error) {
	return _VerifierList.Contract.IsOwner(&_VerifierList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VerifierList *VerifierListCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VerifierList.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VerifierList *VerifierListSession) Owner() (common.Address, error) {
	return _VerifierList.Contract.Owner(&_VerifierList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VerifierList *VerifierListCallerSession) Owner() (common.Address, error) {
	return _VerifierList.Contract.Owner(&_VerifierList.CallOpts)
}

// VerifierList is a free data retrieval call binding the contract method 0x57762ddd.
//
// Solidity: function verifierList(uint256 ) constant returns(string)
func (_VerifierList *VerifierListCaller) VerifierList(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _VerifierList.contract.Call(opts, out, "verifierList", arg0)
	return *ret0, err
}

// VerifierList is a free data retrieval call binding the contract method 0x57762ddd.
//
// Solidity: function verifierList(uint256 ) constant returns(string)
func (_VerifierList *VerifierListSession) VerifierList(arg0 *big.Int) (string, error) {
	return _VerifierList.Contract.VerifierList(&_VerifierList.CallOpts, arg0)
}

// VerifierList is a free data retrieval call binding the contract method 0x57762ddd.
//
// Solidity: function verifierList(uint256 ) constant returns(string)
func (_VerifierList *VerifierListCallerSession) VerifierList(arg0 *big.Int) (string, error) {
	return _VerifierList.Contract.VerifierList(&_VerifierList.CallOpts, arg0)
}

// Verifiers is a free data retrieval call binding the contract method 0xac205deb.
//
// Solidity: function verifiers(string ) constant returns(address owner, string typeOfVerifier, string verifierParams, bool isCreated)
func (_VerifierList *VerifierListCaller) Verifiers(opts *bind.CallOpts, arg0 string) (struct {
	Owner          common.Address
	TypeOfVerifier string
	VerifierParams string
	IsCreated      bool
}, error) {
	ret := new(struct {
		Owner          common.Address
		TypeOfVerifier string
		VerifierParams string
		IsCreated      bool
	})
	out := ret
	err := _VerifierList.contract.Call(opts, out, "verifiers", arg0)
	return *ret, err
}

// Verifiers is a free data retrieval call binding the contract method 0xac205deb.
//
// Solidity: function verifiers(string ) constant returns(address owner, string typeOfVerifier, string verifierParams, bool isCreated)
func (_VerifierList *VerifierListSession) Verifiers(arg0 string) (struct {
	Owner          common.Address
	TypeOfVerifier string
	VerifierParams string
	IsCreated      bool
}, error) {
	return _VerifierList.Contract.Verifiers(&_VerifierList.CallOpts, arg0)
}

// Verifiers is a free data retrieval call binding the contract method 0xac205deb.
//
// Solidity: function verifiers(string ) constant returns(address owner, string typeOfVerifier, string verifierParams, bool isCreated)
func (_VerifierList *VerifierListCallerSession) Verifiers(arg0 string) (struct {
	Owner          common.Address
	TypeOfVerifier string
	VerifierParams string
	IsCreated      bool
}, error) {
	return _VerifierList.Contract.Verifiers(&_VerifierList.CallOpts, arg0)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x43e9c5b6.
//
// Solidity: function addVerifier(string _verifier, string _typeOfVerifier, string _verifierParams, address _owner) returns()
func (_VerifierList *VerifierListTransactor) AddVerifier(opts *bind.TransactOpts, _verifier string, _typeOfVerifier string, _verifierParams string, _owner common.Address) (*types.Transaction, error) {
	return _VerifierList.contract.Transact(opts, "addVerifier", _verifier, _typeOfVerifier, _verifierParams, _owner)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x43e9c5b6.
//
// Solidity: function addVerifier(string _verifier, string _typeOfVerifier, string _verifierParams, address _owner) returns()
func (_VerifierList *VerifierListSession) AddVerifier(_verifier string, _typeOfVerifier string, _verifierParams string, _owner common.Address) (*types.Transaction, error) {
	return _VerifierList.Contract.AddVerifier(&_VerifierList.TransactOpts, _verifier, _typeOfVerifier, _verifierParams, _owner)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x43e9c5b6.
//
// Solidity: function addVerifier(string _verifier, string _typeOfVerifier, string _verifierParams, address _owner) returns()
func (_VerifierList *VerifierListTransactorSession) AddVerifier(_verifier string, _typeOfVerifier string, _verifierParams string, _owner common.Address) (*types.Transaction, error) {
	return _VerifierList.Contract.AddVerifier(&_VerifierList.TransactOpts, _verifier, _typeOfVerifier, _verifierParams, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifierList *VerifierListTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierList.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifierList *VerifierListSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifierList.Contract.RenounceOwnership(&_VerifierList.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifierList *VerifierListTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifierList.Contract.RenounceOwnership(&_VerifierList.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifierList *VerifierListTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VerifierList.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifierList *VerifierListSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifierList.Contract.TransferOwnership(&_VerifierList.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifierList *VerifierListTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifierList.Contract.TransferOwnership(&_VerifierList.TransactOpts, newOwner)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x59fb00f9.
//
// Solidity: function updateVerifier(string _verifier, string _verifierParams) returns()
func (_VerifierList *VerifierListTransactor) UpdateVerifier(opts *bind.TransactOpts, _verifier string, _verifierParams string) (*types.Transaction, error) {
	return _VerifierList.contract.Transact(opts, "updateVerifier", _verifier, _verifierParams)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x59fb00f9.
//
// Solidity: function updateVerifier(string _verifier, string _verifierParams) returns()
func (_VerifierList *VerifierListSession) UpdateVerifier(_verifier string, _verifierParams string) (*types.Transaction, error) {
	return _VerifierList.Contract.UpdateVerifier(&_VerifierList.TransactOpts, _verifier, _verifierParams)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x59fb00f9.
//
// Solidity: function updateVerifier(string _verifier, string _verifierParams) returns()
func (_VerifierList *VerifierListTransactorSession) UpdateVerifier(_verifier string, _verifierParams string) (*types.Transaction, error) {
	return _VerifierList.Contract.UpdateVerifier(&_VerifierList.TransactOpts, _verifier, _verifierParams)
}

// VerifierListOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VerifierList contract.
type VerifierListOwnershipTransferredIterator struct {
	Event *VerifierListOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VerifierListOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierListOwnershipTransferred)
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
		it.Event = new(VerifierListOwnershipTransferred)
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
func (it *VerifierListOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifierListOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifierListOwnershipTransferred represents a OwnershipTransferred event raised by the VerifierList contract.
type VerifierListOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifierList *VerifierListFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VerifierListOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifierList.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VerifierListOwnershipTransferredIterator{contract: _VerifierList.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifierList *VerifierListFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierListOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifierList.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifierListOwnershipTransferred)
				if err := _VerifierList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifierList *VerifierListFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierListOwnershipTransferred, error) {
	event := new(VerifierListOwnershipTransferred)
	if err := _VerifierList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VerifierListVerifierAddedIterator is returned from FilterVerifierAdded and is used to iterate over the raw logs and unpacked data for VerifierAdded events raised by the VerifierList contract.
type VerifierListVerifierAddedIterator struct {
	Event *VerifierListVerifierAdded // Event containing the contract specifics and raw log

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
func (it *VerifierListVerifierAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierListVerifierAdded)
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
		it.Event = new(VerifierListVerifierAdded)
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
func (it *VerifierListVerifierAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifierListVerifierAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifierListVerifierAdded represents a VerifierAdded event raised by the VerifierList contract.
type VerifierListVerifierAdded struct {
	Verifier       string
	TypeOfVerifier string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVerifierAdded is a free log retrieval operation binding the contract event 0xda2f0cdcc06a08d65b7aab3dc839c3414686e2d8cd4aa2d4b3e14313f8df34e8.
//
// Solidity: event VerifierAdded(string verifier, string typeOfVerifier)
func (_VerifierList *VerifierListFilterer) FilterVerifierAdded(opts *bind.FilterOpts) (*VerifierListVerifierAddedIterator, error) {

	logs, sub, err := _VerifierList.contract.FilterLogs(opts, "VerifierAdded")
	if err != nil {
		return nil, err
	}
	return &VerifierListVerifierAddedIterator{contract: _VerifierList.contract, event: "VerifierAdded", logs: logs, sub: sub}, nil
}

// WatchVerifierAdded is a free log subscription operation binding the contract event 0xda2f0cdcc06a08d65b7aab3dc839c3414686e2d8cd4aa2d4b3e14313f8df34e8.
//
// Solidity: event VerifierAdded(string verifier, string typeOfVerifier)
func (_VerifierList *VerifierListFilterer) WatchVerifierAdded(opts *bind.WatchOpts, sink chan<- *VerifierListVerifierAdded) (event.Subscription, error) {

	logs, sub, err := _VerifierList.contract.WatchLogs(opts, "VerifierAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifierListVerifierAdded)
				if err := _VerifierList.contract.UnpackLog(event, "VerifierAdded", log); err != nil {
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

// ParseVerifierAdded is a log parse operation binding the contract event 0xda2f0cdcc06a08d65b7aab3dc839c3414686e2d8cd4aa2d4b3e14313f8df34e8.
//
// Solidity: event VerifierAdded(string verifier, string typeOfVerifier)
func (_VerifierList *VerifierListFilterer) ParseVerifierAdded(log types.Log) (*VerifierListVerifierAdded, error) {
	event := new(VerifierListVerifierAdded)
	if err := _VerifierList.contract.UnpackLog(event, "VerifierAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VerifierListVerifierUpdatedIterator is returned from FilterVerifierUpdated and is used to iterate over the raw logs and unpacked data for VerifierUpdated events raised by the VerifierList contract.
type VerifierListVerifierUpdatedIterator struct {
	Event *VerifierListVerifierUpdated // Event containing the contract specifics and raw log

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
func (it *VerifierListVerifierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierListVerifierUpdated)
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
		it.Event = new(VerifierListVerifierUpdated)
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
func (it *VerifierListVerifierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifierListVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifierListVerifierUpdated represents a VerifierUpdated event raised by the VerifierList contract.
type VerifierListVerifierUpdated struct {
	Verifier string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierUpdated is a free log retrieval operation binding the contract event 0x24d1bc6b4f3a4d03649d6c40376d48f89866293956598a1e6c5f974fbd0021e4.
//
// Solidity: event VerifierUpdated(string verifier)
func (_VerifierList *VerifierListFilterer) FilterVerifierUpdated(opts *bind.FilterOpts) (*VerifierListVerifierUpdatedIterator, error) {

	logs, sub, err := _VerifierList.contract.FilterLogs(opts, "VerifierUpdated")
	if err != nil {
		return nil, err
	}
	return &VerifierListVerifierUpdatedIterator{contract: _VerifierList.contract, event: "VerifierUpdated", logs: logs, sub: sub}, nil
}

// WatchVerifierUpdated is a free log subscription operation binding the contract event 0x24d1bc6b4f3a4d03649d6c40376d48f89866293956598a1e6c5f974fbd0021e4.
//
// Solidity: event VerifierUpdated(string verifier)
func (_VerifierList *VerifierListFilterer) WatchVerifierUpdated(opts *bind.WatchOpts, sink chan<- *VerifierListVerifierUpdated) (event.Subscription, error) {

	logs, sub, err := _VerifierList.contract.WatchLogs(opts, "VerifierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifierListVerifierUpdated)
				if err := _VerifierList.contract.UnpackLog(event, "VerifierUpdated", log); err != nil {
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

// ParseVerifierUpdated is a log parse operation binding the contract event 0x24d1bc6b4f3a4d03649d6c40376d48f89866293956598a1e6c5f974fbd0021e4.
//
// Solidity: event VerifierUpdated(string verifier)
func (_VerifierList *VerifierListFilterer) ParseVerifierUpdated(log types.Log) (*VerifierListVerifierUpdated, error) {
	event := new(VerifierListVerifierUpdated)
	if err := _VerifierList.contract.UnpackLog(event, "VerifierUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
