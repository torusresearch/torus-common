// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifierListAggregator

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
// Solidity: function isOwner() view returns(bool)
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
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
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
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
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
const VerifierListABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"verifier\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeOfVerifier\",\"type\":\"string\"}],\"name\":\"VerifierAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"verifier\",\"type\":\"string\"}],\"name\":\"VerifierUpdated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_verifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_typeOfVerifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_verifierParams\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"addVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVerifierListCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"verifierListCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_verifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_verifierParams\",\"type\":\"string\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"verifierList\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"verifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"typeOfVerifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verifierParams\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isCreated\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VerifierListFuncSigs maps the 4-byte function signature to its string representation.
var VerifierListFuncSigs = map[string]string{
	"43e9c5b6": "addVerifier(string,string,string,address)",
	"16feefe8": "getVerifierListCount()",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
	"59fb00f9": "updateVerifier(string,string)",
	"57762ddd": "verifierList(uint256)",
	"ac205deb": "verifiers(string)",
}

// VerifierListBin is the compiled bytecode used for deploying new contracts.
var VerifierListBin = "0x608060405260006100176001600160e01b0361006616565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006a565b3390565b61101b806100796000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063715018a611610066578063715018a61461031b5780638da5cb5b146103235780638f32d59b14610347578063ac205deb14610363578063f2fde38b1461050957610093565b806316feefe81461009857806343e9c5b6146100b257806357762ddd146101cb57806359fb00f91461025d575b600080fd5b6100a061052f565b60408051918252519081900360200190f35b6101c9600480360360808110156100c857600080fd5b810190602081018135600160201b8111156100e257600080fd5b8201836020820111156100f457600080fd5b803590602001918460018302840111600160201b8311171561011557600080fd5b919390929091602081019035600160201b81111561013257600080fd5b82018360208201111561014457600080fd5b803590602001918460018302840111600160201b8311171561016557600080fd5b919390929091602081019035600160201b81111561018257600080fd5b82018360208201111561019457600080fd5b803590602001918460018302840111600160201b831117156101b557600080fd5b9193509150356001600160a01b0316610536565b005b6101e8600480360360208110156101e157600080fd5b503561087a565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561022257818101518382015260200161020a565b50505050905090810190601f16801561024f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101c96004803603604081101561027357600080fd5b810190602081018135600160201b81111561028d57600080fd5b82018360208201111561029f57600080fd5b803590602001918460018302840111600160201b831117156102c057600080fd5b919390929091602081019035600160201b8111156102dd57600080fd5b8201836020820111156102ef57600080fd5b803590602001918460018302840111600160201b8311171561031057600080fd5b509092509050610920565b6101c9610b9e565b61032b610c3a565b604080516001600160a01b039092168252519081900360200190f35b61034f610c49565b604080519115158252519081900360200190f35b6104076004803603602081101561037957600080fd5b810190602081018135600160201b81111561039357600080fd5b8201836020820111156103a557600080fd5b803590602001918460018302840111600160201b831117156103c657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c6d945050505050565b60405180856001600160a01b03166001600160a01b03168152602001806020018060200184151515158152602001838103835286818151815260200191508051906020019080838360005b8381101561046a578181015183820152602001610452565b50505050905090810190601f1680156104975780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156104ca5781810151838201526020016104b2565b50505050905090810190601f1680156104f75780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b6101c96004803603602081101561051f57600080fd5b50356001600160a01b0316610dc8565b6002545b90565b61053e610c49565b610588576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b86868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604051835160019350849250819060208401908083835b602083106105f25780518252601f1990920191602091820191016105d3565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206003015460ff1615915061067a9050576040805162461bcd60e51b815260206004820152601760248201527f766572696669657220616c726561647920657869737473000000000000000000604482015290519081900360640190fd5b6040518060800160405280836001600160a01b0316815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f880181900481028201810190925286815291810191908790879081908401838280828437600092019190915250505090825250600160209091018190526040518a908a908083838082843791909101948552505060405160209381900384019020845181546001600160a01b0319166001600160a01b03909116178155848401518051919461076c94506001860193500190610ee0565b5060408201518051610788916002840191602090910190610ee0565b50606091909101516003909101805460ff191691151591909117905560028054600181018083556000929092526107e2907f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace018a8a610f5e565b50507fda2f0cdcc06a08d65b7aab3dc839c3414686e2d8cd4aa2d4b3e14313f8df34e8888888886040518080602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600083820152604051601f909101601f19169092018290039850909650505050505050a15050505050505050565b6002818154811061088757fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152935090918301828280156109185780601f106108ed57610100808354040283529160200191610918565b820191906000526020600020905b8154815290600101906020018083116108fb57829003601f168201915b505050505081565b83838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604051835133935060019250849190819060208401908083835b6020831061098d5780518252601f19909201916020918201910161096e565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b0316929092149150610a139050576040805162461bcd60e51b81526020600482015260156024820152743737ba1037bbb732b91037b3103b32b934b334b2b960591b604482015290519081900360640190fd5b84848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604051835160019350849250819060208401908083835b60208310610a7d5780518252601f199092019160209182019101610a5e565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206003015460ff169150610afc9050576040805162461bcd60e51b81526020600482015260156024820152741d995c9a599a595c88191bd95cdb9d08195e1a5cdd605a1b604482015290519081900360640190fd5b83836001888860405180838380828437808301925050509250505090815260200160405180910390206002019190610b35929190610f5e565b507f24d1bc6b4f3a4d03649d6c40376d48f89866293956598a1e6c5f974fbd0021e4868660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a1505050505050565b610ba6610c49565b610bf0576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b600080546001600160a01b0316610c5e610e26565b6001600160a01b031614905090565b8051808201602090810180516001808352938301948301949094209390528254838301805460408051601f60026000199885161561010002989098019093169690960491820185900485028601850190528085526001600160a01b03909216949392909190830182828015610d235780601f10610cf857610100808354040283529160200191610d23565b820191906000526020600020905b815481529060010190602001808311610d0657829003601f168201915b50505060028085018054604080516020601f6000196101006001871615020190941695909504928301859004850281018501909152818152959695945090925090830182828015610db55780601f10610d8a57610100808354040283529160200191610db5565b820191906000526020600020905b815481529060010190602001808311610d9857829003601f168201915b5050506003909301549192505060ff1684565b610dd0610c49565b610e1a576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b610e2381610e2a565b50565b3390565b6001600160a01b038116610e85576040805162461bcd60e51b815260206004820152601860248201527f4f776e61626c653a206e6f207a65726f20616464726573730000000000000000604482015290519081900360640190fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f2157805160ff1916838001178555610f4e565b82800160010185558215610f4e579182015b82811115610f4e578251825591602001919060010190610f33565b50610f5a929150610fcc565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f9f5782800160ff19823516178555610f4e565b82800160010185558215610f4e579182015b82811115610f4e578235825591602001919060010190610fb1565b61053391905b80821115610f5a5760008155600101610fd256fea265627a7a72315820c815dfe0c4d8cb412442c1e139eca079ae209b081bfe0551e98a784b70a57fab64736f6c63430005110032"

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

// GetVerifierListCount is a free data retrieval call binding the contract method 0x16feefe8.
//
// Solidity: function getVerifierListCount() view returns(uint256 verifierListCount)
func (_VerifierList *VerifierListCaller) GetVerifierListCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VerifierList.contract.Call(opts, out, "getVerifierListCount")
	return *ret0, err
}

// GetVerifierListCount is a free data retrieval call binding the contract method 0x16feefe8.
//
// Solidity: function getVerifierListCount() view returns(uint256 verifierListCount)
func (_VerifierList *VerifierListSession) GetVerifierListCount() (*big.Int, error) {
	return _VerifierList.Contract.GetVerifierListCount(&_VerifierList.CallOpts)
}

// GetVerifierListCount is a free data retrieval call binding the contract method 0x16feefe8.
//
// Solidity: function getVerifierListCount() view returns(uint256 verifierListCount)
func (_VerifierList *VerifierListCallerSession) GetVerifierListCount() (*big.Int, error) {
	return _VerifierList.Contract.GetVerifierListCount(&_VerifierList.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
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
// Solidity: function isOwner() view returns(bool)
func (_VerifierList *VerifierListSession) IsOwner() (bool, error) {
	return _VerifierList.Contract.IsOwner(&_VerifierList.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_VerifierList *VerifierListCallerSession) IsOwner() (bool, error) {
	return _VerifierList.Contract.IsOwner(&_VerifierList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
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
// Solidity: function owner() view returns(address)
func (_VerifierList *VerifierListSession) Owner() (common.Address, error) {
	return _VerifierList.Contract.Owner(&_VerifierList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifierList *VerifierListCallerSession) Owner() (common.Address, error) {
	return _VerifierList.Contract.Owner(&_VerifierList.CallOpts)
}

// VerifierList is a free data retrieval call binding the contract method 0x57762ddd.
//
// Solidity: function verifierList(uint256 ) view returns(string)
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
// Solidity: function verifierList(uint256 ) view returns(string)
func (_VerifierList *VerifierListSession) VerifierList(arg0 *big.Int) (string, error) {
	return _VerifierList.Contract.VerifierList(&_VerifierList.CallOpts, arg0)
}

// VerifierList is a free data retrieval call binding the contract method 0x57762ddd.
//
// Solidity: function verifierList(uint256 ) view returns(string)
func (_VerifierList *VerifierListCallerSession) VerifierList(arg0 *big.Int) (string, error) {
	return _VerifierList.Contract.VerifierList(&_VerifierList.CallOpts, arg0)
}

// Verifiers is a free data retrieval call binding the contract method 0xac205deb.
//
// Solidity: function verifiers(string ) view returns(address owner, string typeOfVerifier, string verifierParams, bool isCreated)
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
// Solidity: function verifiers(string ) view returns(address owner, string typeOfVerifier, string verifierParams, bool isCreated)
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
// Solidity: function verifiers(string ) view returns(address owner, string typeOfVerifier, string verifierParams, bool isCreated)
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

// VerifierListAggregatorABI is the input ABI used to generate the binding from.
const VerifierListAggregatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifierListContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifierOwnerManagerContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllVerifiers\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setVerifierOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VerifierListAggregatorFuncSigs maps the 4-byte function signature to its string representation.
var VerifierListAggregatorFuncSigs = map[string]string{
	"af8d1555": "getAllVerifiers()",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"5437988d": "setVerifier(address)",
	"2c713f9f": "setVerifierOwner(address)",
	"f2fde38b": "transferOwnership(address)",
}

// VerifierListAggregatorBin is the compiled bytecode used for deploying new contracts.
var VerifierListAggregatorBin = "0x608060405234801561001057600080fd5b50604051610d99380380610d9983398101604081905261002f916100d2565b60006100426001600160e01b036100bd16565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600180546001600160a01b039384166001600160a01b03199182161790915560028054929093169116179055610134565b3390565b80516100cc8161011d565b92915050565b600080604083850312156100e557600080fd5b60006100f185856100c1565b9250506020610102858286016100c1565b9150509250929050565b60006001600160a01b0382166100cc565b6101268161010c565b811461013157600080fd5b50565b610c56806101436000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100b25780638f32d59b146100d0578063af8d1555146100e5578063f2fde38b146100fd5761007d565b80632c713f9f146100825780635437988d14610097578063715018a6146100aa575b600080fd5b6100956100903660046107a6565b610110565b005b6100956100a53660046107a6565b61015f565b6100956101a5565b6100ba610213565b6040516100c79190610a80565b60405180910390f35b6100d8610222565b6040516100c79190610ae5565b6100ed610246565b6040516100c79493929190610a8e565b61009561010b3660046107a6565b61066b565b610118610222565b61013d5760405162461bcd60e51b815260040161013490610b04565b60405180910390fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b610167610222565b6101835760405162461bcd60e51b815260040161013490610b04565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6101ad610222565b6101c95760405162461bcd60e51b815260040161013490610b04565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b600080546001600160a01b031661023761069b565b6001600160a01b031614905090565b6060806060806000600160009054906101000a90046001600160a01b03166001600160a01b03166316feefe86040518163ffffffff1660e01b815260040160206040518083038186803b15801561029c57600080fd5b505afa1580156102b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506102d491908101906108b0565b905060608160405190808252806020026020018201604052801561030c57816020015b60608152602001906001900390816102f75790505b50905060608260405190808252806020026020018201604052801561034557816020015b60608152602001906001900390816103305790505b50905060608360405190808252806020026020018201604052801561037e57816020015b60608152602001906001900390816103695790505b5090506060846040519080825280602002602001820160405280156103ad578160200160208202803883390190505b50905060005b85811015610463576001546040516357762ddd60e01b81526060916001600160a01b0316906357762ddd906103ec908590600401610b24565b60006040518083038186803b15801561040457600080fd5b505afa158015610418573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610440919081019061087b565b90508086838151811061044f57fe5b6020908102919091010152506001016103b3565b5060005b8581101561065b57600154855160009182916001600160a01b039091169063ac205deb9089908690811061049757fe5b60200260200101516040518263ffffffff1660e01b81526004016104bb9190610af3565b60006040518083038186803b1580156104d357600080fd5b505afa1580156104e7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261050f91908101906107ea565b89878151811061051b57fe5b6020026020010189888151811061052e57fe5b60209081029190910101929092529190526002549093509091506001600160a01b03808316911614156106235760025487516001600160a01b039091169063dc578e0c9089908690811061057e57fe5b60200260200101516040518263ffffffff1660e01b81526004016105a29190610af3565b60206040518083038186803b1580156105ba57600080fd5b505afa1580156105ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506105f291908101906107cc565b8484815181106105fe57fe5b60200260200101906001600160a01b031690816001600160a01b031681525050610651565b8084848151811061063057fe5b60200260200101906001600160a01b031690816001600160a01b0316815250505b5050600101610467565b5092989197509550909350915050565b610673610222565b61068f5760405162461bcd60e51b815260040161013490610b04565b6106988161069f565b50565b3390565b6001600160a01b0381166106c55760405162461bcd60e51b815260040161013490610b14565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b803561072b81610bed565b92915050565b805161072b81610bed565b805161072b81610c01565b600082601f83011261075857600080fd5b815161076b61076682610b59565b610b32565b9150808252602083016020830185838301111561078757600080fd5b610792838284610bb3565b50505092915050565b805161072b81610c0a565b6000602082840312156107b857600080fd5b60006107c48484610720565b949350505050565b6000602082840312156107de57600080fd5b60006107c48484610731565b6000806000806080858703121561080057600080fd5b600061080c8787610731565b945050602085015167ffffffffffffffff81111561082957600080fd5b61083587828801610747565b935050604085015167ffffffffffffffff81111561085257600080fd5b61085e87828801610747565b925050606061086f8782880161073c565b91505092959194509250565b60006020828403121561088d57600080fd5b815167ffffffffffffffff8111156108a457600080fd5b6107c484828501610747565b6000602082840312156108c257600080fd5b60006107c4848461079b565b60006108da83836108f5565b505060200190565b60006108ee83836109d4565b9392505050565b6108fe81610b94565b82525050565b600061090f82610b87565b6109198185610b8b565b935061092483610b81565b8060005b8381101561095257815161093c88826108ce565b975061094783610b81565b925050600101610928565b509495945050505050565b600061096882610b87565b6109728185610b8b565b93508360208202850161098485610b81565b8060005b858110156109be57848403895281516109a185826108e2565b94506109ac83610b81565b60209a909a0199925050600101610988565b5091979650505050505050565b6108fe81610b9f565b60006109df82610b87565b6109e98185610b8b565b93506109f9818560208601610bb3565b610a0281610be3565b9093019392505050565b6000610a19601683610b8b565b7527bbb730b136329d103737ba103a34329037bbb732b960511b815260200192915050565b6000610a4b601883610b8b565b7f4f776e61626c653a206e6f207a65726f20616464726573730000000000000000815260200192915050565b6108fe81610bb0565b6020810161072b82846108f5565b60808082528101610a9f818761095d565b90508181036020830152610ab3818661095d565b90508181036040830152610ac7818561095d565b90508181036060830152610adb8184610904565b9695505050505050565b6020810161072b82846109cb565b602080825281016108ee81846109d4565b6020808252810161072b81610a0c565b6020808252810161072b81610a3e565b6020810161072b8284610a77565b60405181810167ffffffffffffffff81118282101715610b5157600080fd5b604052919050565b600067ffffffffffffffff821115610b7057600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b600061072b82610ba4565b151590565b6001600160a01b031690565b90565b60005b83811015610bce578181015183820152602001610bb6565b83811115610bdd576000848401525b50505050565b601f01601f191690565b610bf681610b94565b811461069857600080fd5b610bf681610b9f565b610bf681610bb056fea365627a7a723158209af0ed2a9c0c220cc71cb8d88c9c26938d2bf1decc9472d54f01fb82a5c3a5456c6578706572696d656e74616cf564736f6c63430005110040"

// DeployVerifierListAggregator deploys a new Ethereum contract, binding an instance of VerifierListAggregator to it.
func DeployVerifierListAggregator(auth *bind.TransactOpts, backend bind.ContractBackend, _verifierListContract common.Address, _verifierOwnerManagerContract common.Address) (common.Address, *types.Transaction, *VerifierListAggregator, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierListAggregatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VerifierListAggregatorBin), backend, _verifierListContract, _verifierOwnerManagerContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierListAggregator{VerifierListAggregatorCaller: VerifierListAggregatorCaller{contract: contract}, VerifierListAggregatorTransactor: VerifierListAggregatorTransactor{contract: contract}, VerifierListAggregatorFilterer: VerifierListAggregatorFilterer{contract: contract}}, nil
}

// VerifierListAggregator is an auto generated Go binding around an Ethereum contract.
type VerifierListAggregator struct {
	VerifierListAggregatorCaller     // Read-only binding to the contract
	VerifierListAggregatorTransactor // Write-only binding to the contract
	VerifierListAggregatorFilterer   // Log filterer for contract events
}

// VerifierListAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierListAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierListAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierListAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierListAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierListAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierListAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierListAggregatorSession struct {
	Contract     *VerifierListAggregator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VerifierListAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierListAggregatorCallerSession struct {
	Contract *VerifierListAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// VerifierListAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierListAggregatorTransactorSession struct {
	Contract     *VerifierListAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// VerifierListAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierListAggregatorRaw struct {
	Contract *VerifierListAggregator // Generic contract binding to access the raw methods on
}

// VerifierListAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierListAggregatorCallerRaw struct {
	Contract *VerifierListAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierListAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierListAggregatorTransactorRaw struct {
	Contract *VerifierListAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifierListAggregator creates a new instance of VerifierListAggregator, bound to a specific deployed contract.
func NewVerifierListAggregator(address common.Address, backend bind.ContractBackend) (*VerifierListAggregator, error) {
	contract, err := bindVerifierListAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierListAggregator{VerifierListAggregatorCaller: VerifierListAggregatorCaller{contract: contract}, VerifierListAggregatorTransactor: VerifierListAggregatorTransactor{contract: contract}, VerifierListAggregatorFilterer: VerifierListAggregatorFilterer{contract: contract}}, nil
}

// NewVerifierListAggregatorCaller creates a new read-only instance of VerifierListAggregator, bound to a specific deployed contract.
func NewVerifierListAggregatorCaller(address common.Address, caller bind.ContractCaller) (*VerifierListAggregatorCaller, error) {
	contract, err := bindVerifierListAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierListAggregatorCaller{contract: contract}, nil
}

// NewVerifierListAggregatorTransactor creates a new write-only instance of VerifierListAggregator, bound to a specific deployed contract.
func NewVerifierListAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierListAggregatorTransactor, error) {
	contract, err := bindVerifierListAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierListAggregatorTransactor{contract: contract}, nil
}

// NewVerifierListAggregatorFilterer creates a new log filterer instance of VerifierListAggregator, bound to a specific deployed contract.
func NewVerifierListAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierListAggregatorFilterer, error) {
	contract, err := bindVerifierListAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierListAggregatorFilterer{contract: contract}, nil
}

// bindVerifierListAggregator binds a generic wrapper to an already deployed contract.
func bindVerifierListAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierListAggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierListAggregator *VerifierListAggregatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VerifierListAggregator.Contract.VerifierListAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierListAggregator *VerifierListAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierListAggregator.Contract.VerifierListAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierListAggregator *VerifierListAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierListAggregator.Contract.VerifierListAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierListAggregator *VerifierListAggregatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VerifierListAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierListAggregator *VerifierListAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierListAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierListAggregator *VerifierListAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierListAggregator.Contract.contract.Transact(opts, method, params...)
}

// GetAllVerifiers is a free data retrieval call binding the contract method 0xaf8d1555.
//
// Solidity: function getAllVerifiers() view returns(string[], string[], string[], address[])
func (_VerifierListAggregator *VerifierListAggregatorCaller) GetAllVerifiers(opts *bind.CallOpts) ([]string, []string, []string, []common.Address, error) {
	var (
		ret0 = new([]string)
		ret1 = new([]string)
		ret2 = new([]string)
		ret3 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _VerifierListAggregator.contract.Call(opts, out, "getAllVerifiers")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetAllVerifiers is a free data retrieval call binding the contract method 0xaf8d1555.
//
// Solidity: function getAllVerifiers() view returns(string[], string[], string[], address[])
func (_VerifierListAggregator *VerifierListAggregatorSession) GetAllVerifiers() ([]string, []string, []string, []common.Address, error) {
	return _VerifierListAggregator.Contract.GetAllVerifiers(&_VerifierListAggregator.CallOpts)
}

// GetAllVerifiers is a free data retrieval call binding the contract method 0xaf8d1555.
//
// Solidity: function getAllVerifiers() view returns(string[], string[], string[], address[])
func (_VerifierListAggregator *VerifierListAggregatorCallerSession) GetAllVerifiers() ([]string, []string, []string, []common.Address, error) {
	return _VerifierListAggregator.Contract.GetAllVerifiers(&_VerifierListAggregator.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_VerifierListAggregator *VerifierListAggregatorCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VerifierListAggregator.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_VerifierListAggregator *VerifierListAggregatorSession) IsOwner() (bool, error) {
	return _VerifierListAggregator.Contract.IsOwner(&_VerifierListAggregator.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_VerifierListAggregator *VerifierListAggregatorCallerSession) IsOwner() (bool, error) {
	return _VerifierListAggregator.Contract.IsOwner(&_VerifierListAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifierListAggregator *VerifierListAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VerifierListAggregator.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifierListAggregator *VerifierListAggregatorSession) Owner() (common.Address, error) {
	return _VerifierListAggregator.Contract.Owner(&_VerifierListAggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifierListAggregator *VerifierListAggregatorCallerSession) Owner() (common.Address, error) {
	return _VerifierListAggregator.Contract.Owner(&_VerifierListAggregator.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifierListAggregator *VerifierListAggregatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierListAggregator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifierListAggregator *VerifierListAggregatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifierListAggregator.Contract.RenounceOwnership(&_VerifierListAggregator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifierListAggregator *VerifierListAggregatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifierListAggregator.Contract.RenounceOwnership(&_VerifierListAggregator.TransactOpts)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address addr) returns()
func (_VerifierListAggregator *VerifierListAggregatorTransactor) SetVerifier(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _VerifierListAggregator.contract.Transact(opts, "setVerifier", addr)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address addr) returns()
func (_VerifierListAggregator *VerifierListAggregatorSession) SetVerifier(addr common.Address) (*types.Transaction, error) {
	return _VerifierListAggregator.Contract.SetVerifier(&_VerifierListAggregator.TransactOpts, addr)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address addr) returns()
func (_VerifierListAggregator *VerifierListAggregatorTransactorSession) SetVerifier(addr common.Address) (*types.Transaction, error) {
	return _VerifierListAggregator.Contract.SetVerifier(&_VerifierListAggregator.TransactOpts, addr)
}

// SetVerifierOwner is a paid mutator transaction binding the contract method 0x2c713f9f.
//
// Solidity: function setVerifierOwner(address addr) returns()
func (_VerifierListAggregator *VerifierListAggregatorTransactor) SetVerifierOwner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _VerifierListAggregator.contract.Transact(opts, "setVerifierOwner", addr)
}

// SetVerifierOwner is a paid mutator transaction binding the contract method 0x2c713f9f.
//
// Solidity: function setVerifierOwner(address addr) returns()
func (_VerifierListAggregator *VerifierListAggregatorSession) SetVerifierOwner(addr common.Address) (*types.Transaction, error) {
	return _VerifierListAggregator.Contract.SetVerifierOwner(&_VerifierListAggregator.TransactOpts, addr)
}

// SetVerifierOwner is a paid mutator transaction binding the contract method 0x2c713f9f.
//
// Solidity: function setVerifierOwner(address addr) returns()
func (_VerifierListAggregator *VerifierListAggregatorTransactorSession) SetVerifierOwner(addr common.Address) (*types.Transaction, error) {
	return _VerifierListAggregator.Contract.SetVerifierOwner(&_VerifierListAggregator.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifierListAggregator *VerifierListAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VerifierListAggregator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifierListAggregator *VerifierListAggregatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifierListAggregator.Contract.TransferOwnership(&_VerifierListAggregator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifierListAggregator *VerifierListAggregatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifierListAggregator.Contract.TransferOwnership(&_VerifierListAggregator.TransactOpts, newOwner)
}

// VerifierListAggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VerifierListAggregator contract.
type VerifierListAggregatorOwnershipTransferredIterator struct {
	Event *VerifierListAggregatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VerifierListAggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierListAggregatorOwnershipTransferred)
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
		it.Event = new(VerifierListAggregatorOwnershipTransferred)
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
func (it *VerifierListAggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifierListAggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifierListAggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the VerifierListAggregator contract.
type VerifierListAggregatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifierListAggregator *VerifierListAggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VerifierListAggregatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifierListAggregator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VerifierListAggregatorOwnershipTransferredIterator{contract: _VerifierListAggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifierListAggregator *VerifierListAggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierListAggregatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifierListAggregator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifierListAggregatorOwnershipTransferred)
				if err := _VerifierListAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VerifierListAggregator *VerifierListAggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierListAggregatorOwnershipTransferred, error) {
	event := new(VerifierListAggregatorOwnershipTransferred)
	if err := _VerifierListAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VerifierOwnerManagerABI is the input ABI used to generate the binding from.
const VerifierOwnerManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifierListContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"verifier\",\"type\":\"string\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"verifier\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifierListContract\",\"type\":\"address\"}],\"name\":\"setVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"verifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"verifierParams\",\"type\":\"string\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"verifierOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VerifierOwnerManagerFuncSigs maps the 4-byte function signature to its string representation.
var VerifierOwnerManagerFuncSigs = map[string]string{
	"4aaf4a12": "getOwner(string)",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"092b25e9": "setOwner(string,address)",
	"5437988d": "setVerifier(address)",
	"f2fde38b": "transferOwnership(address)",
	"59fb00f9": "updateVerifier(string,string)",
	"dc578e0c": "verifierOwner(string)",
}

// VerifierOwnerManagerBin is the compiled bytecode used for deploying new contracts.
var VerifierOwnerManagerBin = "0x608060405234801561001057600080fd5b50604051610bc6380380610bc68339818101604052602081101561003357600080fd5b505160006100486001600160e01b036100b716565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600180546001600160a01b0319166001600160a01b03929092169190911790556100bb565b3390565b610afc806100ca6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063715018a611610066578063715018a61461027f5780638da5cb5b146102875780638f32d59b1461028f578063dc578e0c146102ab578063f2fde38b1461034f57610093565b8063092b25e9146100985780634aaf4a12146101115780635437988d1461019b57806359fb00f9146101c1575b600080fd5b61010f600480360360408110156100ae57600080fd5b810190602081018135600160201b8111156100c857600080fd5b8201836020820111156100da57600080fd5b803590602001918460018302840111600160201b831117156100fb57600080fd5b9193509150356001600160a01b0316610375565b005b61017f6004803603602081101561012757600080fd5b810190602081018135600160201b81111561014157600080fd5b82018360208201111561015357600080fd5b803590602001918460018302840111600160201b8311171561017457600080fd5b509092509050610585565b604080516001600160a01b039092168252519081900360200190f35b61010f600480360360208110156101b157600080fd5b50356001600160a01b03166105bf565b61010f600480360360408110156101d757600080fd5b810190602081018135600160201b8111156101f157600080fd5b82018360208201111561020357600080fd5b803590602001918460018302840111600160201b8311171561022457600080fd5b919390929091602081019035600160201b81111561024157600080fd5b82018360208201111561025357600080fd5b803590602001918460018302840111600160201b8311171561027457600080fd5b509092509050610633565b61010f6108ba565b61017f610956565b610297610965565b604080519115158252519081900360200190f35b61017f600480360360208110156102c157600080fd5b810190602081018135600160201b8111156102db57600080fd5b8201836020820111156102ed57600080fd5b803590602001918460018302840111600160201b8311171561030e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610989945050505050565b61010f6004803603602081101561036557600080fd5b50356001600160a01b03166109af565b82828080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060405185519194506002935085925090819060208401908083835b602083106103e15780518252601f1990920191602091820191016103c2565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b03169290921415915061047590505761042b610965565b610470576040805162461bcd60e51b81526020600482015260116024820152703737ba103232b330bab63a1037bbb732b960791b604482015290519081900360640190fd5b610537565b336001600160a01b03166002826040518082805190602001908083835b602083106104b15780518252601f199092019160209182019101610492565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b03169290921491506105379050576040805162461bcd60e51b81526020600482015260156024820152743737ba1037bbb732b91037b3103b32b934b334b2b960591b604482015290519081900360640190fd5b8160028585604051808383808284379190910194855250506040519283900360200190922080546001600160a01b03949094166001600160a01b031990941693909317909255505050505050565b6000600283836040518083838082843791909101948552505060405192839003602001909220546001600160a01b03169250505092915050565b6105c7610965565b610611576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b83838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525060405185519194506002935085925090819060208401908083835b6020831061069f5780518252601f199092019160209182019101610680565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b0316929092141591506107339050576106e9610965565b61072e576040805162461bcd60e51b81526020600482015260116024820152703737ba103232b330bab63a1037bbb732b960791b604482015290519081900360640190fd5b6107f5565b336001600160a01b03166002826040518082805190602001908083835b6020831061076f5780518252601f199092019160209182019101610750565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b03169290921491506107f59050576040805162461bcd60e51b81526020600482015260156024820152743737ba1037bbb732b91037b3103b32b934b334b2b960591b604482015290519081900360640190fd5b600154604080516359fb00f960e01b815260048101918252604481018790526001600160a01b03909216916359fb00f99188918891889188919081906024810190606401878780828437600083820152601f01601f191690910184810383528581526020019050858580828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b15801561089b57600080fd5b505af11580156108af573d6000803e3d6000fd5b505050505050505050565b6108c2610965565b61090c576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b600080546001600160a01b031661097a610a0d565b6001600160a01b031614905090565b80516020818301810180516002825292820191909301209152546001600160a01b031681565b6109b7610965565b610a01576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b610a0a81610a11565b50565b3390565b6001600160a01b038116610a6c576040805162461bcd60e51b815260206004820152601860248201527f4f776e61626c653a206e6f207a65726f20616464726573730000000000000000604482015290519081900360640190fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a723158205df2c21880597bffeb2e1f8eb49f88418624c0816973f3da8516ef0cb96fce0c64736f6c63430005110032"

// DeployVerifierOwnerManager deploys a new Ethereum contract, binding an instance of VerifierOwnerManager to it.
func DeployVerifierOwnerManager(auth *bind.TransactOpts, backend bind.ContractBackend, _verifierListContract common.Address) (common.Address, *types.Transaction, *VerifierOwnerManager, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierOwnerManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VerifierOwnerManagerBin), backend, _verifierListContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierOwnerManager{VerifierOwnerManagerCaller: VerifierOwnerManagerCaller{contract: contract}, VerifierOwnerManagerTransactor: VerifierOwnerManagerTransactor{contract: contract}, VerifierOwnerManagerFilterer: VerifierOwnerManagerFilterer{contract: contract}}, nil
}

// VerifierOwnerManager is an auto generated Go binding around an Ethereum contract.
type VerifierOwnerManager struct {
	VerifierOwnerManagerCaller     // Read-only binding to the contract
	VerifierOwnerManagerTransactor // Write-only binding to the contract
	VerifierOwnerManagerFilterer   // Log filterer for contract events
}

// VerifierOwnerManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierOwnerManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierOwnerManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierOwnerManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierOwnerManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierOwnerManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierOwnerManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierOwnerManagerSession struct {
	Contract     *VerifierOwnerManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VerifierOwnerManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierOwnerManagerCallerSession struct {
	Contract *VerifierOwnerManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// VerifierOwnerManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierOwnerManagerTransactorSession struct {
	Contract     *VerifierOwnerManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// VerifierOwnerManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierOwnerManagerRaw struct {
	Contract *VerifierOwnerManager // Generic contract binding to access the raw methods on
}

// VerifierOwnerManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierOwnerManagerCallerRaw struct {
	Contract *VerifierOwnerManagerCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierOwnerManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierOwnerManagerTransactorRaw struct {
	Contract *VerifierOwnerManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifierOwnerManager creates a new instance of VerifierOwnerManager, bound to a specific deployed contract.
func NewVerifierOwnerManager(address common.Address, backend bind.ContractBackend) (*VerifierOwnerManager, error) {
	contract, err := bindVerifierOwnerManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierOwnerManager{VerifierOwnerManagerCaller: VerifierOwnerManagerCaller{contract: contract}, VerifierOwnerManagerTransactor: VerifierOwnerManagerTransactor{contract: contract}, VerifierOwnerManagerFilterer: VerifierOwnerManagerFilterer{contract: contract}}, nil
}

// NewVerifierOwnerManagerCaller creates a new read-only instance of VerifierOwnerManager, bound to a specific deployed contract.
func NewVerifierOwnerManagerCaller(address common.Address, caller bind.ContractCaller) (*VerifierOwnerManagerCaller, error) {
	contract, err := bindVerifierOwnerManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierOwnerManagerCaller{contract: contract}, nil
}

// NewVerifierOwnerManagerTransactor creates a new write-only instance of VerifierOwnerManager, bound to a specific deployed contract.
func NewVerifierOwnerManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierOwnerManagerTransactor, error) {
	contract, err := bindVerifierOwnerManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierOwnerManagerTransactor{contract: contract}, nil
}

// NewVerifierOwnerManagerFilterer creates a new log filterer instance of VerifierOwnerManager, bound to a specific deployed contract.
func NewVerifierOwnerManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierOwnerManagerFilterer, error) {
	contract, err := bindVerifierOwnerManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierOwnerManagerFilterer{contract: contract}, nil
}

// bindVerifierOwnerManager binds a generic wrapper to an already deployed contract.
func bindVerifierOwnerManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerifierOwnerManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierOwnerManager *VerifierOwnerManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VerifierOwnerManager.Contract.VerifierOwnerManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierOwnerManager *VerifierOwnerManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.VerifierOwnerManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierOwnerManager *VerifierOwnerManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.VerifierOwnerManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifierOwnerManager *VerifierOwnerManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VerifierOwnerManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifierOwnerManager *VerifierOwnerManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifierOwnerManager *VerifierOwnerManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x4aaf4a12.
//
// Solidity: function getOwner(string verifier) view returns(address)
func (_VerifierOwnerManager *VerifierOwnerManagerCaller) GetOwner(opts *bind.CallOpts, verifier string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VerifierOwnerManager.contract.Call(opts, out, "getOwner", verifier)
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x4aaf4a12.
//
// Solidity: function getOwner(string verifier) view returns(address)
func (_VerifierOwnerManager *VerifierOwnerManagerSession) GetOwner(verifier string) (common.Address, error) {
	return _VerifierOwnerManager.Contract.GetOwner(&_VerifierOwnerManager.CallOpts, verifier)
}

// GetOwner is a free data retrieval call binding the contract method 0x4aaf4a12.
//
// Solidity: function getOwner(string verifier) view returns(address)
func (_VerifierOwnerManager *VerifierOwnerManagerCallerSession) GetOwner(verifier string) (common.Address, error) {
	return _VerifierOwnerManager.Contract.GetOwner(&_VerifierOwnerManager.CallOpts, verifier)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_VerifierOwnerManager *VerifierOwnerManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VerifierOwnerManager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_VerifierOwnerManager *VerifierOwnerManagerSession) IsOwner() (bool, error) {
	return _VerifierOwnerManager.Contract.IsOwner(&_VerifierOwnerManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_VerifierOwnerManager *VerifierOwnerManagerCallerSession) IsOwner() (bool, error) {
	return _VerifierOwnerManager.Contract.IsOwner(&_VerifierOwnerManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifierOwnerManager *VerifierOwnerManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VerifierOwnerManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifierOwnerManager *VerifierOwnerManagerSession) Owner() (common.Address, error) {
	return _VerifierOwnerManager.Contract.Owner(&_VerifierOwnerManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifierOwnerManager *VerifierOwnerManagerCallerSession) Owner() (common.Address, error) {
	return _VerifierOwnerManager.Contract.Owner(&_VerifierOwnerManager.CallOpts)
}

// VerifierOwner is a free data retrieval call binding the contract method 0xdc578e0c.
//
// Solidity: function verifierOwner(string ) view returns(address)
func (_VerifierOwnerManager *VerifierOwnerManagerCaller) VerifierOwner(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VerifierOwnerManager.contract.Call(opts, out, "verifierOwner", arg0)
	return *ret0, err
}

// VerifierOwner is a free data retrieval call binding the contract method 0xdc578e0c.
//
// Solidity: function verifierOwner(string ) view returns(address)
func (_VerifierOwnerManager *VerifierOwnerManagerSession) VerifierOwner(arg0 string) (common.Address, error) {
	return _VerifierOwnerManager.Contract.VerifierOwner(&_VerifierOwnerManager.CallOpts, arg0)
}

// VerifierOwner is a free data retrieval call binding the contract method 0xdc578e0c.
//
// Solidity: function verifierOwner(string ) view returns(address)
func (_VerifierOwnerManager *VerifierOwnerManagerCallerSession) VerifierOwner(arg0 string) (common.Address, error) {
	return _VerifierOwnerManager.Contract.VerifierOwner(&_VerifierOwnerManager.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifierOwnerManager *VerifierOwnerManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierOwnerManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifierOwnerManager *VerifierOwnerManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.RenounceOwnership(&_VerifierOwnerManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifierOwnerManager *VerifierOwnerManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.RenounceOwnership(&_VerifierOwnerManager.TransactOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x092b25e9.
//
// Solidity: function setOwner(string verifier, address newOwner) returns()
func (_VerifierOwnerManager *VerifierOwnerManagerTransactor) SetOwner(opts *bind.TransactOpts, verifier string, newOwner common.Address) (*types.Transaction, error) {
	return _VerifierOwnerManager.contract.Transact(opts, "setOwner", verifier, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x092b25e9.
//
// Solidity: function setOwner(string verifier, address newOwner) returns()
func (_VerifierOwnerManager *VerifierOwnerManagerSession) SetOwner(verifier string, newOwner common.Address) (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.SetOwner(&_VerifierOwnerManager.TransactOpts, verifier, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x092b25e9.
//
// Solidity: function setOwner(string verifier, address newOwner) returns()
func (_VerifierOwnerManager *VerifierOwnerManagerTransactorSession) SetOwner(verifier string, newOwner common.Address) (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.SetOwner(&_VerifierOwnerManager.TransactOpts, verifier, newOwner)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address _verifierListContract) returns()
func (_VerifierOwnerManager *VerifierOwnerManagerTransactor) SetVerifier(opts *bind.TransactOpts, _verifierListContract common.Address) (*types.Transaction, error) {
	return _VerifierOwnerManager.contract.Transact(opts, "setVerifier", _verifierListContract)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address _verifierListContract) returns()
func (_VerifierOwnerManager *VerifierOwnerManagerSession) SetVerifier(_verifierListContract common.Address) (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.SetVerifier(&_VerifierOwnerManager.TransactOpts, _verifierListContract)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address _verifierListContract) returns()
func (_VerifierOwnerManager *VerifierOwnerManagerTransactorSession) SetVerifier(_verifierListContract common.Address) (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.SetVerifier(&_VerifierOwnerManager.TransactOpts, _verifierListContract)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifierOwnerManager *VerifierOwnerManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VerifierOwnerManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifierOwnerManager *VerifierOwnerManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.TransferOwnership(&_VerifierOwnerManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifierOwnerManager *VerifierOwnerManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.TransferOwnership(&_VerifierOwnerManager.TransactOpts, newOwner)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x59fb00f9.
//
// Solidity: function updateVerifier(string verifier, string verifierParams) returns()
func (_VerifierOwnerManager *VerifierOwnerManagerTransactor) UpdateVerifier(opts *bind.TransactOpts, verifier string, verifierParams string) (*types.Transaction, error) {
	return _VerifierOwnerManager.contract.Transact(opts, "updateVerifier", verifier, verifierParams)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x59fb00f9.
//
// Solidity: function updateVerifier(string verifier, string verifierParams) returns()
func (_VerifierOwnerManager *VerifierOwnerManagerSession) UpdateVerifier(verifier string, verifierParams string) (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.UpdateVerifier(&_VerifierOwnerManager.TransactOpts, verifier, verifierParams)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x59fb00f9.
//
// Solidity: function updateVerifier(string verifier, string verifierParams) returns()
func (_VerifierOwnerManager *VerifierOwnerManagerTransactorSession) UpdateVerifier(verifier string, verifierParams string) (*types.Transaction, error) {
	return _VerifierOwnerManager.Contract.UpdateVerifier(&_VerifierOwnerManager.TransactOpts, verifier, verifierParams)
}

// VerifierOwnerManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VerifierOwnerManager contract.
type VerifierOwnerManagerOwnershipTransferredIterator struct {
	Event *VerifierOwnerManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VerifierOwnerManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierOwnerManagerOwnershipTransferred)
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
		it.Event = new(VerifierOwnerManagerOwnershipTransferred)
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
func (it *VerifierOwnerManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifierOwnerManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifierOwnerManagerOwnershipTransferred represents a OwnershipTransferred event raised by the VerifierOwnerManager contract.
type VerifierOwnerManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifierOwnerManager *VerifierOwnerManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VerifierOwnerManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifierOwnerManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VerifierOwnerManagerOwnershipTransferredIterator{contract: _VerifierOwnerManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifierOwnerManager *VerifierOwnerManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierOwnerManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifierOwnerManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifierOwnerManagerOwnershipTransferred)
				if err := _VerifierOwnerManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VerifierOwnerManager *VerifierOwnerManagerFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierOwnerManagerOwnershipTransferred, error) {
	event := new(VerifierOwnerManagerOwnershipTransferred)
	if err := _VerifierOwnerManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
