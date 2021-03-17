// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nodeListProxy

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

// NodeListABI is the input ABI used to generate the binding from.
const NodeListABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"NodeListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextEpoch\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getEpochInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"nodeList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"prevEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextEpoch\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"declaredIp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tmP2PListenAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"p2pListenAddress\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getNodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newEpoch\",\"type\":\"uint256\"}],\"name\":\"getPssStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"declaredIp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pubKx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKy\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tmP2PListenAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"p2pListenAddress\",\"type\":\"string\"}],\"name\":\"listNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"declaredIp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKy\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tmP2PListenAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"p2pListenAddress\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"nodeRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pssStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"nodeList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"prevEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextEpoch\",\"type\":\"uint256\"}],\"name\":\"updateEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"updatePssStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NodeListFuncSigs maps the 4-byte function signature to its string representation.
var NodeListFuncSigs = map[string]string{
	"3894228e": "epochInfo(uint256)",
	"135022c2": "getEpochInfo(uint256)",
	"bafb3581": "getNodeDetails(address)",
	"47de074f": "getNodes(uint256)",
	"c7aa8ff7": "getPssStatus(uint256,uint256)",
	"8f32d59b": "isOwner()",
	"7d22c35c": "isWhitelisted(uint256,address)",
	"bf2d6f81": "listNode(uint256,string,uint256,uint256,string,string)",
	"859da85f": "nodeDetails(address)",
	"86470e9e": "nodeRegistered(uint256,address)",
	"8da5cb5b": "owner()",
	"52fc47b4": "pssStatus(uint256,uint256)",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
	"ae4df20c": "updateEpoch(uint256,uint256,uint256,uint256,address[],uint256,uint256)",
	"6967ac51": "updatePssStatus(uint256,uint256,uint256)",
	"3d4602a9": "updateWhitelist(uint256,address,bool)",
	"4b25bfce": "whitelist(uint256,address)",
}

// NodeListBin is the compiled bytecode used for deploying new contracts.
var NodeListBin = "0x608060405260006100176001600160e01b0361006616565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006a565b3390565b611a1c806100796000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c8063859da85f116100a2578063ae4df20c11610071578063ae4df20c14610552578063bafb358114610610578063bf2d6f8114610781578063c7aa8ff7146108a2578063f2fde38b146108c557610116565b8063859da85f1461037b57806386470e9e146104fa5780638da5cb5b146105265780638f32d59b1461054a57610116565b80634b25bfce116100e95780634b25bfce146102a957806352fc47b4146102e95780636967ac511461031e578063715018a6146103475780637d22c35c1461034f57610116565b8063135022c21461011b5780633894228e146101b65780633d4602a91461020657806347de074f1461023c575b600080fd5b6101386004803603602081101561013157600080fd5b50356108eb565b6040518088815260200187815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019060200280838360005b8381101561019c578181015183820152602001610184565b505050509050019850505050505050505060405180910390f35b6101d3600480360360208110156101cc57600080fd5b5035610a32565b604080519687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b61023a6004803603606081101561021c57600080fd5b508035906001600160a01b0360208201351690604001351515610a69565b005b6102596004803603602081101561025257600080fd5b5035610b37565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561029557818101518382015260200161027d565b505050509050019250505060405180910390f35b6102d5600480360360408110156102bf57600080fd5b50803590602001356001600160a01b0316610bed565b604080519115158252519081900360200190f35b61030c600480360360408110156102ff57600080fd5b5080359060200135610c0d565b60408051918252519081900360200190f35b61023a6004803603606081101561033457600080fd5b5080359060208101359060400135610c2a565b61023a610d26565b6102d56004803603604081101561036557600080fd5b50803590602001356001600160a01b0316610dc2565b6103a16004803603602081101561039157600080fd5b50356001600160a01b0316610def565b6040518080602001878152602001868152602001858152602001806020018060200184810384528a818151815260200191508051906020019080838360005b838110156103f85781810151838201526020016103e0565b50505050905090810190601f1680156104255780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b83811015610458578181015183820152602001610440565b50505050905090810190601f1680156104855780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b838110156104b85781810151838201526020016104a0565b50505050905090810190601f1680156104e55780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b6102d56004803603604081101561051057600080fd5b50803590602001356001600160a01b0316610fd0565b61052e61103e565b604080516001600160a01b039092168252519081900360200190f35b6102d561104e565b61023a600480360360e081101561056857600080fd5b81359160208101359160408201359160608101359181019060a081016080820135600160201b81111561059a57600080fd5b8201836020820111156105ac57600080fd5b803590602001918460208302840111600160201b831117156105cd57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505082359350505060200135611072565b6106366004803603602081101561062657600080fd5b50356001600160a01b03166111a1565b60405180806020018581526020018060200180602001848103845288818151815260200191508051906020019080838360005b83811015610681578181015183820152602001610669565b50505050905090810190601f1680156106ae5780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b838110156106e15781810151838201526020016106c9565b50505050905090810190601f16801561070e5780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b83811015610741578181015183820152602001610729565b50505050905090810190601f16801561076e5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b61023a600480360360c081101561079757600080fd5b81359190810190604081016020820135600160201b8111156107b857600080fd5b8201836020820111156107ca57600080fd5b803590602001918460018302840111600160201b831117156107eb57600080fd5b919390928235926020810135929190606081019060400135600160201b81111561081457600080fd5b82018360208201111561082657600080fd5b803590602001918460018302840111600160201b8311171561084757600080fd5b919390929091602081019035600160201b81111561086457600080fd5b82018360208201111561087657600080fd5b803590602001918460018302840111600160201b8311171561089757600080fd5b5090925090506113d9565b61030c600480360360408110156108b857600080fd5b5080359060200135611722565b61023a600480360360208110156108db57600080fd5b50356001600160a01b031661173f565b600080808060608180878061093a576040805162461bcd60e51b815260206004820152601060248201526f045706f63682063616e277420626520360841b604482015290519081900360640190fd5b610942611857565b60008a815260026020818152604092839020835160e081018552815481526001820154818401529281015483850152600381015460608401526004810180548551818502810185019096528086529394919360808601938301828280156109d257602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116109b4575b50505050508152602001600582015481526020016006820154815250509050806000015181602001518260400151836060015184608001518560a001518660c0015182925098509850985098509850985098505050919395979092949650565b6002602081905260009182526040909120805460018201549282015460038301546005840154600690940154929493919290919086565b610a7161104e565b610abb576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b8280610b01576040805162461bcd60e51b815260206004820152601060248201526f045706f63682063616e277420626520360841b604482015290519081900360640190fd5b5060009283526001602090815260408085206001600160a01b039490941685529290529120805460ff1916911515919091179055565b60608180610b7f576040805162461bcd60e51b815260206004820152601060248201526f045706f63682063616e277420626520360841b604482015290519081900360640190fd5b60008381526002602090815260409182902060040180548351818402810184019094528084529091830182828015610be057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610bc2575b5050505050915050919050565b600160209081526000928352604080842090915290825290205460ff1681565b600460209081526000928352604080842090915290825290205481565b610c3261104e565b610c7c576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b8280610cc2576040805162461bcd60e51b815260206004820152601060248201526f045706f63682063616e277420626520360841b604482015290519081900360640190fd5b8280610d08576040805162461bcd60e51b815260206004820152601060248201526f045706f63682063616e277420626520360841b604482015290519081900360640190fd5b50506000928352600460209081526040808520938552929052912055565b610d2e61104e565b610d78576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008281526001602090815260408083206001600160a01b038516845290915290205460ff165b92915050565b60036020908152600091825260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452909291839190830182828015610e865780601f10610e5b57610100808354040283529160200191610e86565b820191906000526020600020905b815481529060010190602001808311610e6957829003601f168201915b505050505090806001015490806002015490806003015490806004018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f365780601f10610f0b57610100808354040283529160200191610f36565b820191906000526020600020905b815481529060010190602001808311610f1957829003601f168201915b5050505060058301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015610fc65780601f10610f9b57610100808354040283529160200191610fc6565b820191906000526020600020905b815481529060010190602001808311610fa957829003601f168201915b5050505050905086565b6000828152600260205260408120815b600482015481101561103357836001600160a01b031682600401828154811061100557fe5b6000918252602090912001546001600160a01b0316141561102b57600192505050610de9565b600101610fe0565b506000949350505050565b6000546001600160a01b03165b90565b600080546001600160a01b031661106361179d565b6001600160a01b031614905090565b61107a61104e565b6110c4576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b868061110a576040805162461bcd60e51b815260206004820152601060248201526f045706f63682063616e277420626520360841b604482015290519081900360640190fd5b6040805160e08101825289815260208082018a81528284018a8152606084018a8152608085018a815260a086018a905260c0860189905260008f81526002808752979020865181559351600185015591519583019590955593516003820155925180519293926111809260048501920190611894565b5060a0820151600582015560c0909101516006909101555050505050505050565b606060006060806111b06118f9565b6001600160a01b0386166000908152600360209081526040918290208251815460026001821615610100026000190190911604601f8101849004909302810160e090810190945260c081018381529093919284928491908401828280156112585780601f1061122d57610100808354040283529160200191611258565b820191906000526020600020905b81548152906001019060200180831161123b57829003601f168201915b50505050508152602001600182015481526020016002820154815260200160038201548152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113185780601f106112ed57610100808354040283529160200191611318565b820191906000526020600020905b8154815290600101906020018083116112fb57829003601f168201915b505050918352505060058201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156113ac5780601f10611381576101008083540402835291602001916113ac565b820191906000526020600020905b81548152906001019060200180831161138f57829003601f168201915b50505091909252505081516020830151608084015160a090940151919a9099509297509550909350505050565b886113e48133610dc2565b611435576040805162461bcd60e51b815260206004820181905260248201527f4e6f64652069736e27742077686974656c697374656420666f722065706f6368604482015290519081900360640190fd5b898061147b576040805162461bcd60e51b815260206004820152601060248201526f045706f63682063616e277420626520360841b604482015290519081900360640190fd5b60008b8152600260205260409020548b9081146114d7576040805162461bcd60e51b8152602060048201526015602482015274115c1bd8da08185b1c9958591e4818dc99585d1959605a1b604482015290519081900360640190fd5b6114e18c33610fd0565b15611533576040805162461bcd60e51b815260206004820152601a60248201527f4e6f646520697320616c72656164792072656769737465726564000000000000604482015290519081900360640190fd5b60008c81526002602090815260408083206004810180546001810182559085529383902090930180546001600160a01b03191633179055805160e0601f8f018490049093028101830190915260c081018d815290918291908f908f9081908501838280828437600092019190915250505090825250600483015460208083019190915260408083018e9052606083018d90528051601f8c018390048302810183019091528a8152608090920191908b908b9081908401838280828437600092019190915250505090825250604080516020601f8a018190048102820181019092528881529181019190899089908190840183828082843760009201829052509390945250503381526003602090815260409091208351805191935061165c92849291019061192f565b506020828101516001830155604083015160028301556060830151600383015560808301518051611693926004850192019061192f565b5060a082015180516116af91600584019160209091019061192f565b509050507fe2f8adb0f494dc82ccf446c031763ef3762d6396d51664611ed89aac0117339e338e836004018054905060405180846001600160a01b03166001600160a01b03168152602001838152602001828152602001935050505060405180910390a150505050505050505050505050565b600091825260046020908152604080842092845291905290205490565b61174761104e565b611791576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b61179a816117a1565b50565b3390565b6001600160a01b0381166117fc576040805162461bcd60e51b815260206004820152601860248201527f4f776e61626c653a206e6f207a65726f20616464726573730000000000000000604482015290519081900360640190fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6040518060e00160405280600081526020016000815260200160008152602001600081526020016060815260200160008152602001600081525090565b8280548282559060005260206000209081019282156118e9579160200282015b828111156118e957825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906118b4565b506118f59291506119a9565b5090565b6040518060c001604052806060815260200160008152602001600081526020016000815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061197057805160ff191683800117855561199d565b8280016001018555821561199d579182015b8281111561199d578251825591602001919060010190611982565b506118f59291506119cd565b61104b91905b808211156118f55780546001600160a01b03191681556001016119af565b61104b91905b808211156118f557600081556001016119d356fea265627a7a723158205dfb310516c9929cac3d70e82c1fd2233650a86423a9ccd75fc4b5613f3a999564736f6c63430005110032"

// DeployNodeList deploys a new Ethereum contract, binding an instance of NodeList to it.
func DeployNodeList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeList, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeListABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeList{NodeListCaller: NodeListCaller{contract: contract}, NodeListTransactor: NodeListTransactor{contract: contract}, NodeListFilterer: NodeListFilterer{contract: contract}}, nil
}

// NodeList is an auto generated Go binding around an Ethereum contract.
type NodeList struct {
	NodeListCaller     // Read-only binding to the contract
	NodeListTransactor // Write-only binding to the contract
	NodeListFilterer   // Log filterer for contract events
}

// NodeListCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeListSession struct {
	Contract     *NodeList         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeListCallerSession struct {
	Contract *NodeListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// NodeListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeListTransactorSession struct {
	Contract     *NodeListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NodeListRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeListRaw struct {
	Contract *NodeList // Generic contract binding to access the raw methods on
}

// NodeListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeListCallerRaw struct {
	Contract *NodeListCaller // Generic read-only contract binding to access the raw methods on
}

// NodeListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeListTransactorRaw struct {
	Contract *NodeListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeList creates a new instance of NodeList, bound to a specific deployed contract.
func NewNodeList(address common.Address, backend bind.ContractBackend) (*NodeList, error) {
	contract, err := bindNodeList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeList{NodeListCaller: NodeListCaller{contract: contract}, NodeListTransactor: NodeListTransactor{contract: contract}, NodeListFilterer: NodeListFilterer{contract: contract}}, nil
}

// NewNodeListCaller creates a new read-only instance of NodeList, bound to a specific deployed contract.
func NewNodeListCaller(address common.Address, caller bind.ContractCaller) (*NodeListCaller, error) {
	contract, err := bindNodeList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeListCaller{contract: contract}, nil
}

// NewNodeListTransactor creates a new write-only instance of NodeList, bound to a specific deployed contract.
func NewNodeListTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeListTransactor, error) {
	contract, err := bindNodeList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeListTransactor{contract: contract}, nil
}

// NewNodeListFilterer creates a new log filterer instance of NodeList, bound to a specific deployed contract.
func NewNodeListFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeListFilterer, error) {
	contract, err := bindNodeList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeListFilterer{contract: contract}, nil
}

// bindNodeList binds a generic wrapper to an already deployed contract.
func bindNodeList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeList *NodeListRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeList.Contract.NodeListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeList *NodeListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeList.Contract.NodeListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeList *NodeListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeList.Contract.NodeListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeList *NodeListCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeList *NodeListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeList *NodeListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeList.Contract.contract.Transact(opts, method, params...)
}

// EpochInfo is a free data retrieval call binding the contract method 0x3894228e.
//
// Solidity: function epochInfo(uint256 ) view returns(uint256 id, uint256 n, uint256 k, uint256 t, uint256 prevEpoch, uint256 nextEpoch)
func (_NodeList *NodeListCaller) EpochInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	N         *big.Int
	K         *big.Int
	T         *big.Int
	PrevEpoch *big.Int
	NextEpoch *big.Int
}, error) {
	ret := new(struct {
		Id        *big.Int
		N         *big.Int
		K         *big.Int
		T         *big.Int
		PrevEpoch *big.Int
		NextEpoch *big.Int
	})
	out := ret
	err := _NodeList.contract.Call(opts, out, "epochInfo", arg0)
	return *ret, err
}

// EpochInfo is a free data retrieval call binding the contract method 0x3894228e.
//
// Solidity: function epochInfo(uint256 ) view returns(uint256 id, uint256 n, uint256 k, uint256 t, uint256 prevEpoch, uint256 nextEpoch)
func (_NodeList *NodeListSession) EpochInfo(arg0 *big.Int) (struct {
	Id        *big.Int
	N         *big.Int
	K         *big.Int
	T         *big.Int
	PrevEpoch *big.Int
	NextEpoch *big.Int
}, error) {
	return _NodeList.Contract.EpochInfo(&_NodeList.CallOpts, arg0)
}

// EpochInfo is a free data retrieval call binding the contract method 0x3894228e.
//
// Solidity: function epochInfo(uint256 ) view returns(uint256 id, uint256 n, uint256 k, uint256 t, uint256 prevEpoch, uint256 nextEpoch)
func (_NodeList *NodeListCallerSession) EpochInfo(arg0 *big.Int) (struct {
	Id        *big.Int
	N         *big.Int
	K         *big.Int
	T         *big.Int
	PrevEpoch *big.Int
	NextEpoch *big.Int
}, error) {
	return _NodeList.Contract.EpochInfo(&_NodeList.CallOpts, arg0)
}

// GetEpochInfo is a free data retrieval call binding the contract method 0x135022c2.
//
// Solidity: function getEpochInfo(uint256 epoch) view returns(uint256 id, uint256 n, uint256 k, uint256 t, address[] nodeList, uint256 prevEpoch, uint256 nextEpoch)
func (_NodeList *NodeListCaller) GetEpochInfo(opts *bind.CallOpts, epoch *big.Int) (struct {
	Id        *big.Int
	N         *big.Int
	K         *big.Int
	T         *big.Int
	NodeList  []common.Address
	PrevEpoch *big.Int
	NextEpoch *big.Int
}, error) {
	ret := new(struct {
		Id        *big.Int
		N         *big.Int
		K         *big.Int
		T         *big.Int
		NodeList  []common.Address
		PrevEpoch *big.Int
		NextEpoch *big.Int
	})
	out := ret
	err := _NodeList.contract.Call(opts, out, "getEpochInfo", epoch)
	return *ret, err
}

// GetEpochInfo is a free data retrieval call binding the contract method 0x135022c2.
//
// Solidity: function getEpochInfo(uint256 epoch) view returns(uint256 id, uint256 n, uint256 k, uint256 t, address[] nodeList, uint256 prevEpoch, uint256 nextEpoch)
func (_NodeList *NodeListSession) GetEpochInfo(epoch *big.Int) (struct {
	Id        *big.Int
	N         *big.Int
	K         *big.Int
	T         *big.Int
	NodeList  []common.Address
	PrevEpoch *big.Int
	NextEpoch *big.Int
}, error) {
	return _NodeList.Contract.GetEpochInfo(&_NodeList.CallOpts, epoch)
}

// GetEpochInfo is a free data retrieval call binding the contract method 0x135022c2.
//
// Solidity: function getEpochInfo(uint256 epoch) view returns(uint256 id, uint256 n, uint256 k, uint256 t, address[] nodeList, uint256 prevEpoch, uint256 nextEpoch)
func (_NodeList *NodeListCallerSession) GetEpochInfo(epoch *big.Int) (struct {
	Id        *big.Int
	N         *big.Int
	K         *big.Int
	T         *big.Int
	NodeList  []common.Address
	PrevEpoch *big.Int
	NextEpoch *big.Int
}, error) {
	return _NodeList.Contract.GetEpochInfo(&_NodeList.CallOpts, epoch)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0xbafb3581.
//
// Solidity: function getNodeDetails(address nodeAddress) view returns(string declaredIp, uint256 position, string tmP2PListenAddress, string p2pListenAddress)
func (_NodeList *NodeListCaller) GetNodeDetails(opts *bind.CallOpts, nodeAddress common.Address) (struct {
	DeclaredIp         string
	Position           *big.Int
	TmP2PListenAddress string
	P2pListenAddress   string
}, error) {
	ret := new(struct {
		DeclaredIp         string
		Position           *big.Int
		TmP2PListenAddress string
		P2pListenAddress   string
	})
	out := ret
	err := _NodeList.contract.Call(opts, out, "getNodeDetails", nodeAddress)
	return *ret, err
}

// GetNodeDetails is a free data retrieval call binding the contract method 0xbafb3581.
//
// Solidity: function getNodeDetails(address nodeAddress) view returns(string declaredIp, uint256 position, string tmP2PListenAddress, string p2pListenAddress)
func (_NodeList *NodeListSession) GetNodeDetails(nodeAddress common.Address) (struct {
	DeclaredIp         string
	Position           *big.Int
	TmP2PListenAddress string
	P2pListenAddress   string
}, error) {
	return _NodeList.Contract.GetNodeDetails(&_NodeList.CallOpts, nodeAddress)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0xbafb3581.
//
// Solidity: function getNodeDetails(address nodeAddress) view returns(string declaredIp, uint256 position, string tmP2PListenAddress, string p2pListenAddress)
func (_NodeList *NodeListCallerSession) GetNodeDetails(nodeAddress common.Address) (struct {
	DeclaredIp         string
	Position           *big.Int
	TmP2PListenAddress string
	P2pListenAddress   string
}, error) {
	return _NodeList.Contract.GetNodeDetails(&_NodeList.CallOpts, nodeAddress)
}

// GetNodes is a free data retrieval call binding the contract method 0x47de074f.
//
// Solidity: function getNodes(uint256 epoch) view returns(address[])
func (_NodeList *NodeListCaller) GetNodes(opts *bind.CallOpts, epoch *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _NodeList.contract.Call(opts, out, "getNodes", epoch)
	return *ret0, err
}

// GetNodes is a free data retrieval call binding the contract method 0x47de074f.
//
// Solidity: function getNodes(uint256 epoch) view returns(address[])
func (_NodeList *NodeListSession) GetNodes(epoch *big.Int) ([]common.Address, error) {
	return _NodeList.Contract.GetNodes(&_NodeList.CallOpts, epoch)
}

// GetNodes is a free data retrieval call binding the contract method 0x47de074f.
//
// Solidity: function getNodes(uint256 epoch) view returns(address[])
func (_NodeList *NodeListCallerSession) GetNodes(epoch *big.Int) ([]common.Address, error) {
	return _NodeList.Contract.GetNodes(&_NodeList.CallOpts, epoch)
}

// GetPssStatus is a free data retrieval call binding the contract method 0xc7aa8ff7.
//
// Solidity: function getPssStatus(uint256 oldEpoch, uint256 newEpoch) view returns(uint256)
func (_NodeList *NodeListCaller) GetPssStatus(opts *bind.CallOpts, oldEpoch *big.Int, newEpoch *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeList.contract.Call(opts, out, "getPssStatus", oldEpoch, newEpoch)
	return *ret0, err
}

// GetPssStatus is a free data retrieval call binding the contract method 0xc7aa8ff7.
//
// Solidity: function getPssStatus(uint256 oldEpoch, uint256 newEpoch) view returns(uint256)
func (_NodeList *NodeListSession) GetPssStatus(oldEpoch *big.Int, newEpoch *big.Int) (*big.Int, error) {
	return _NodeList.Contract.GetPssStatus(&_NodeList.CallOpts, oldEpoch, newEpoch)
}

// GetPssStatus is a free data retrieval call binding the contract method 0xc7aa8ff7.
//
// Solidity: function getPssStatus(uint256 oldEpoch, uint256 newEpoch) view returns(uint256)
func (_NodeList *NodeListCallerSession) GetPssStatus(oldEpoch *big.Int, newEpoch *big.Int) (*big.Int, error) {
	return _NodeList.Contract.GetPssStatus(&_NodeList.CallOpts, oldEpoch, newEpoch)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_NodeList *NodeListCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NodeList.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_NodeList *NodeListSession) IsOwner() (bool, error) {
	return _NodeList.Contract.IsOwner(&_NodeList.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_NodeList *NodeListCallerSession) IsOwner() (bool, error) {
	return _NodeList.Contract.IsOwner(&_NodeList.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x7d22c35c.
//
// Solidity: function isWhitelisted(uint256 epoch, address nodeAddress) view returns(bool)
func (_NodeList *NodeListCaller) IsWhitelisted(opts *bind.CallOpts, epoch *big.Int, nodeAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NodeList.contract.Call(opts, out, "isWhitelisted", epoch, nodeAddress)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x7d22c35c.
//
// Solidity: function isWhitelisted(uint256 epoch, address nodeAddress) view returns(bool)
func (_NodeList *NodeListSession) IsWhitelisted(epoch *big.Int, nodeAddress common.Address) (bool, error) {
	return _NodeList.Contract.IsWhitelisted(&_NodeList.CallOpts, epoch, nodeAddress)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x7d22c35c.
//
// Solidity: function isWhitelisted(uint256 epoch, address nodeAddress) view returns(bool)
func (_NodeList *NodeListCallerSession) IsWhitelisted(epoch *big.Int, nodeAddress common.Address) (bool, error) {
	return _NodeList.Contract.IsWhitelisted(&_NodeList.CallOpts, epoch, nodeAddress)
}

// NodeDetails is a free data retrieval call binding the contract method 0x859da85f.
//
// Solidity: function nodeDetails(address ) view returns(string declaredIp, uint256 position, uint256 pubKx, uint256 pubKy, string tmP2PListenAddress, string p2pListenAddress)
func (_NodeList *NodeListCaller) NodeDetails(opts *bind.CallOpts, arg0 common.Address) (struct {
	DeclaredIp         string
	Position           *big.Int
	PubKx              *big.Int
	PubKy              *big.Int
	TmP2PListenAddress string
	P2pListenAddress   string
}, error) {
	ret := new(struct {
		DeclaredIp         string
		Position           *big.Int
		PubKx              *big.Int
		PubKy              *big.Int
		TmP2PListenAddress string
		P2pListenAddress   string
	})
	out := ret
	err := _NodeList.contract.Call(opts, out, "nodeDetails", arg0)
	return *ret, err
}

// NodeDetails is a free data retrieval call binding the contract method 0x859da85f.
//
// Solidity: function nodeDetails(address ) view returns(string declaredIp, uint256 position, uint256 pubKx, uint256 pubKy, string tmP2PListenAddress, string p2pListenAddress)
func (_NodeList *NodeListSession) NodeDetails(arg0 common.Address) (struct {
	DeclaredIp         string
	Position           *big.Int
	PubKx              *big.Int
	PubKy              *big.Int
	TmP2PListenAddress string
	P2pListenAddress   string
}, error) {
	return _NodeList.Contract.NodeDetails(&_NodeList.CallOpts, arg0)
}

// NodeDetails is a free data retrieval call binding the contract method 0x859da85f.
//
// Solidity: function nodeDetails(address ) view returns(string declaredIp, uint256 position, uint256 pubKx, uint256 pubKy, string tmP2PListenAddress, string p2pListenAddress)
func (_NodeList *NodeListCallerSession) NodeDetails(arg0 common.Address) (struct {
	DeclaredIp         string
	Position           *big.Int
	PubKx              *big.Int
	PubKy              *big.Int
	TmP2PListenAddress string
	P2pListenAddress   string
}, error) {
	return _NodeList.Contract.NodeDetails(&_NodeList.CallOpts, arg0)
}

// NodeRegistered is a free data retrieval call binding the contract method 0x86470e9e.
//
// Solidity: function nodeRegistered(uint256 epoch, address nodeAddress) view returns(bool)
func (_NodeList *NodeListCaller) NodeRegistered(opts *bind.CallOpts, epoch *big.Int, nodeAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NodeList.contract.Call(opts, out, "nodeRegistered", epoch, nodeAddress)
	return *ret0, err
}

// NodeRegistered is a free data retrieval call binding the contract method 0x86470e9e.
//
// Solidity: function nodeRegistered(uint256 epoch, address nodeAddress) view returns(bool)
func (_NodeList *NodeListSession) NodeRegistered(epoch *big.Int, nodeAddress common.Address) (bool, error) {
	return _NodeList.Contract.NodeRegistered(&_NodeList.CallOpts, epoch, nodeAddress)
}

// NodeRegistered is a free data retrieval call binding the contract method 0x86470e9e.
//
// Solidity: function nodeRegistered(uint256 epoch, address nodeAddress) view returns(bool)
func (_NodeList *NodeListCallerSession) NodeRegistered(epoch *big.Int, nodeAddress common.Address) (bool, error) {
	return _NodeList.Contract.NodeRegistered(&_NodeList.CallOpts, epoch, nodeAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeList *NodeListCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NodeList.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeList *NodeListSession) Owner() (common.Address, error) {
	return _NodeList.Contract.Owner(&_NodeList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeList *NodeListCallerSession) Owner() (common.Address, error) {
	return _NodeList.Contract.Owner(&_NodeList.CallOpts)
}

// PssStatus is a free data retrieval call binding the contract method 0x52fc47b4.
//
// Solidity: function pssStatus(uint256 , uint256 ) view returns(uint256)
func (_NodeList *NodeListCaller) PssStatus(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeList.contract.Call(opts, out, "pssStatus", arg0, arg1)
	return *ret0, err
}

// PssStatus is a free data retrieval call binding the contract method 0x52fc47b4.
//
// Solidity: function pssStatus(uint256 , uint256 ) view returns(uint256)
func (_NodeList *NodeListSession) PssStatus(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _NodeList.Contract.PssStatus(&_NodeList.CallOpts, arg0, arg1)
}

// PssStatus is a free data retrieval call binding the contract method 0x52fc47b4.
//
// Solidity: function pssStatus(uint256 , uint256 ) view returns(uint256)
func (_NodeList *NodeListCallerSession) PssStatus(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _NodeList.Contract.PssStatus(&_NodeList.CallOpts, arg0, arg1)
}

// Whitelist is a free data retrieval call binding the contract method 0x4b25bfce.
//
// Solidity: function whitelist(uint256 , address ) view returns(bool)
func (_NodeList *NodeListCaller) Whitelist(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NodeList.contract.Call(opts, out, "whitelist", arg0, arg1)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x4b25bfce.
//
// Solidity: function whitelist(uint256 , address ) view returns(bool)
func (_NodeList *NodeListSession) Whitelist(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _NodeList.Contract.Whitelist(&_NodeList.CallOpts, arg0, arg1)
}

// Whitelist is a free data retrieval call binding the contract method 0x4b25bfce.
//
// Solidity: function whitelist(uint256 , address ) view returns(bool)
func (_NodeList *NodeListCallerSession) Whitelist(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _NodeList.Contract.Whitelist(&_NodeList.CallOpts, arg0, arg1)
}

// ListNode is a paid mutator transaction binding the contract method 0xbf2d6f81.
//
// Solidity: function listNode(uint256 epoch, string declaredIp, uint256 pubKx, uint256 pubKy, string tmP2PListenAddress, string p2pListenAddress) returns()
func (_NodeList *NodeListTransactor) ListNode(opts *bind.TransactOpts, epoch *big.Int, declaredIp string, pubKx *big.Int, pubKy *big.Int, tmP2PListenAddress string, p2pListenAddress string) (*types.Transaction, error) {
	return _NodeList.contract.Transact(opts, "listNode", epoch, declaredIp, pubKx, pubKy, tmP2PListenAddress, p2pListenAddress)
}

// ListNode is a paid mutator transaction binding the contract method 0xbf2d6f81.
//
// Solidity: function listNode(uint256 epoch, string declaredIp, uint256 pubKx, uint256 pubKy, string tmP2PListenAddress, string p2pListenAddress) returns()
func (_NodeList *NodeListSession) ListNode(epoch *big.Int, declaredIp string, pubKx *big.Int, pubKy *big.Int, tmP2PListenAddress string, p2pListenAddress string) (*types.Transaction, error) {
	return _NodeList.Contract.ListNode(&_NodeList.TransactOpts, epoch, declaredIp, pubKx, pubKy, tmP2PListenAddress, p2pListenAddress)
}

// ListNode is a paid mutator transaction binding the contract method 0xbf2d6f81.
//
// Solidity: function listNode(uint256 epoch, string declaredIp, uint256 pubKx, uint256 pubKy, string tmP2PListenAddress, string p2pListenAddress) returns()
func (_NodeList *NodeListTransactorSession) ListNode(epoch *big.Int, declaredIp string, pubKx *big.Int, pubKy *big.Int, tmP2PListenAddress string, p2pListenAddress string) (*types.Transaction, error) {
	return _NodeList.Contract.ListNode(&_NodeList.TransactOpts, epoch, declaredIp, pubKx, pubKy, tmP2PListenAddress, p2pListenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeList *NodeListTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeList.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeList *NodeListSession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeList.Contract.RenounceOwnership(&_NodeList.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeList *NodeListTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeList.Contract.RenounceOwnership(&_NodeList.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeList *NodeListTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NodeList.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeList *NodeListSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeList.Contract.TransferOwnership(&_NodeList.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeList *NodeListTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeList.Contract.TransferOwnership(&_NodeList.TransactOpts, newOwner)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0xae4df20c.
//
// Solidity: function updateEpoch(uint256 epoch, uint256 n, uint256 k, uint256 t, address[] nodeList, uint256 prevEpoch, uint256 nextEpoch) returns()
func (_NodeList *NodeListTransactor) UpdateEpoch(opts *bind.TransactOpts, epoch *big.Int, n *big.Int, k *big.Int, t *big.Int, nodeList []common.Address, prevEpoch *big.Int, nextEpoch *big.Int) (*types.Transaction, error) {
	return _NodeList.contract.Transact(opts, "updateEpoch", epoch, n, k, t, nodeList, prevEpoch, nextEpoch)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0xae4df20c.
//
// Solidity: function updateEpoch(uint256 epoch, uint256 n, uint256 k, uint256 t, address[] nodeList, uint256 prevEpoch, uint256 nextEpoch) returns()
func (_NodeList *NodeListSession) UpdateEpoch(epoch *big.Int, n *big.Int, k *big.Int, t *big.Int, nodeList []common.Address, prevEpoch *big.Int, nextEpoch *big.Int) (*types.Transaction, error) {
	return _NodeList.Contract.UpdateEpoch(&_NodeList.TransactOpts, epoch, n, k, t, nodeList, prevEpoch, nextEpoch)
}

// UpdateEpoch is a paid mutator transaction binding the contract method 0xae4df20c.
//
// Solidity: function updateEpoch(uint256 epoch, uint256 n, uint256 k, uint256 t, address[] nodeList, uint256 prevEpoch, uint256 nextEpoch) returns()
func (_NodeList *NodeListTransactorSession) UpdateEpoch(epoch *big.Int, n *big.Int, k *big.Int, t *big.Int, nodeList []common.Address, prevEpoch *big.Int, nextEpoch *big.Int) (*types.Transaction, error) {
	return _NodeList.Contract.UpdateEpoch(&_NodeList.TransactOpts, epoch, n, k, t, nodeList, prevEpoch, nextEpoch)
}

// UpdatePssStatus is a paid mutator transaction binding the contract method 0x6967ac51.
//
// Solidity: function updatePssStatus(uint256 oldEpoch, uint256 newEpoch, uint256 status) returns()
func (_NodeList *NodeListTransactor) UpdatePssStatus(opts *bind.TransactOpts, oldEpoch *big.Int, newEpoch *big.Int, status *big.Int) (*types.Transaction, error) {
	return _NodeList.contract.Transact(opts, "updatePssStatus", oldEpoch, newEpoch, status)
}

// UpdatePssStatus is a paid mutator transaction binding the contract method 0x6967ac51.
//
// Solidity: function updatePssStatus(uint256 oldEpoch, uint256 newEpoch, uint256 status) returns()
func (_NodeList *NodeListSession) UpdatePssStatus(oldEpoch *big.Int, newEpoch *big.Int, status *big.Int) (*types.Transaction, error) {
	return _NodeList.Contract.UpdatePssStatus(&_NodeList.TransactOpts, oldEpoch, newEpoch, status)
}

// UpdatePssStatus is a paid mutator transaction binding the contract method 0x6967ac51.
//
// Solidity: function updatePssStatus(uint256 oldEpoch, uint256 newEpoch, uint256 status) returns()
func (_NodeList *NodeListTransactorSession) UpdatePssStatus(oldEpoch *big.Int, newEpoch *big.Int, status *big.Int) (*types.Transaction, error) {
	return _NodeList.Contract.UpdatePssStatus(&_NodeList.TransactOpts, oldEpoch, newEpoch, status)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x3d4602a9.
//
// Solidity: function updateWhitelist(uint256 epoch, address nodeAddress, bool allowed) returns()
func (_NodeList *NodeListTransactor) UpdateWhitelist(opts *bind.TransactOpts, epoch *big.Int, nodeAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _NodeList.contract.Transact(opts, "updateWhitelist", epoch, nodeAddress, allowed)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x3d4602a9.
//
// Solidity: function updateWhitelist(uint256 epoch, address nodeAddress, bool allowed) returns()
func (_NodeList *NodeListSession) UpdateWhitelist(epoch *big.Int, nodeAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _NodeList.Contract.UpdateWhitelist(&_NodeList.TransactOpts, epoch, nodeAddress, allowed)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0x3d4602a9.
//
// Solidity: function updateWhitelist(uint256 epoch, address nodeAddress, bool allowed) returns()
func (_NodeList *NodeListTransactorSession) UpdateWhitelist(epoch *big.Int, nodeAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _NodeList.Contract.UpdateWhitelist(&_NodeList.TransactOpts, epoch, nodeAddress, allowed)
}

// NodeListNodeListedIterator is returned from FilterNodeListed and is used to iterate over the raw logs and unpacked data for NodeListed events raised by the NodeList contract.
type NodeListNodeListedIterator struct {
	Event *NodeListNodeListed // Event containing the contract specifics and raw log

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
func (it *NodeListNodeListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeListNodeListed)
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
		it.Event = new(NodeListNodeListed)
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
func (it *NodeListNodeListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeListNodeListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeListNodeListed represents a NodeListed event raised by the NodeList contract.
type NodeListNodeListed struct {
	PublicKey common.Address
	Epoch     *big.Int
	Position  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNodeListed is a free log retrieval operation binding the contract event 0xe2f8adb0f494dc82ccf446c031763ef3762d6396d51664611ed89aac0117339e.
//
// Solidity: event NodeListed(address publicKey, uint256 epoch, uint256 position)
func (_NodeList *NodeListFilterer) FilterNodeListed(opts *bind.FilterOpts) (*NodeListNodeListedIterator, error) {

	logs, sub, err := _NodeList.contract.FilterLogs(opts, "NodeListed")
	if err != nil {
		return nil, err
	}
	return &NodeListNodeListedIterator{contract: _NodeList.contract, event: "NodeListed", logs: logs, sub: sub}, nil
}

// WatchNodeListed is a free log subscription operation binding the contract event 0xe2f8adb0f494dc82ccf446c031763ef3762d6396d51664611ed89aac0117339e.
//
// Solidity: event NodeListed(address publicKey, uint256 epoch, uint256 position)
func (_NodeList *NodeListFilterer) WatchNodeListed(opts *bind.WatchOpts, sink chan<- *NodeListNodeListed) (event.Subscription, error) {

	logs, sub, err := _NodeList.contract.WatchLogs(opts, "NodeListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeListNodeListed)
				if err := _NodeList.contract.UnpackLog(event, "NodeListed", log); err != nil {
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

// ParseNodeListed is a log parse operation binding the contract event 0xe2f8adb0f494dc82ccf446c031763ef3762d6396d51664611ed89aac0117339e.
//
// Solidity: event NodeListed(address publicKey, uint256 epoch, uint256 position)
func (_NodeList *NodeListFilterer) ParseNodeListed(log types.Log) (*NodeListNodeListed, error) {
	event := new(NodeListNodeListed)
	if err := _NodeList.contract.UnpackLog(event, "NodeListed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NodeListOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NodeList contract.
type NodeListOwnershipTransferredIterator struct {
	Event *NodeListOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NodeListOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeListOwnershipTransferred)
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
		it.Event = new(NodeListOwnershipTransferred)
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
func (it *NodeListOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeListOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeListOwnershipTransferred represents a OwnershipTransferred event raised by the NodeList contract.
type NodeListOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeList *NodeListFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodeListOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeList.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeListOwnershipTransferredIterator{contract: _NodeList.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeList *NodeListFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodeListOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeList.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeListOwnershipTransferred)
				if err := _NodeList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NodeList *NodeListFilterer) ParseOwnershipTransferred(log types.Log) (*NodeListOwnershipTransferred, error) {
	event := new(NodeListOwnershipTransferred)
	if err := _NodeList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NodeListProxyABI is the input ABI used to generate the binding from.
const NodeListProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeListContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEpoch\",\"type\":\"uint256\"}],\"name\":\"EpochChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"NodeListContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getEpochInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"nodeList\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"prevEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextEpoch\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"declaredIp\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKy\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tmP2PListenAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"p2pListenAddress\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getNodes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newEpoch\",\"type\":\"uint256\"}],\"name\":\"getPssStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newEpoch\",\"type\":\"uint256\"}],\"name\":\"setCurrentEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeListContractAddress\",\"type\":\"address\"}],\"name\":\"setNodeListContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NodeListProxyFuncSigs maps the 4-byte function signature to its string representation.
var NodeListProxyFuncSigs = map[string]string{
	"76671808": "currentEpoch()",
	"135022c2": "getEpochInfo(uint256)",
	"bafb3581": "getNodeDetails(address)",
	"47de074f": "getNodes(uint256)",
	"c7aa8ff7": "getPssStatus(uint256,uint256)",
	"8f32d59b": "isOwner()",
	"7d22c35c": "isWhitelisted(uint256,address)",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"1dd6b9b1": "setCurrentEpoch(uint256)",
	"08704c0a": "setNodeListContract(address)",
	"f2fde38b": "transferOwnership(address)",
}

// NodeListProxyBin is the compiled bytecode used for deploying new contracts.
var NodeListProxyBin = "0x608060405234801561001057600080fd5b50604051610f12380380610f128339818101604052604081101561003357600080fd5b508051602090910151600061004f6001600160e01b036100c116565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600155600280546001600160a01b0319166001600160a01b03929092169190911790556100c5565b3390565b610e3e806100d46000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80637d22c35c116100715780637d22c35c146102285780638da5cb5b146102685780638f32d59b1461028c578063bafb358114610294578063c7aa8ff714610413578063f2fde38b14610436576100b4565b806308704c0a146100b9578063135022c2146100e15780631dd6b9b11461017c57806347de074f14610199578063715018a614610206578063766718081461020e575b600080fd5b6100df600480360360208110156100cf57600080fd5b50356001600160a01b031661045c565b005b6100fe600480360360208110156100f757600080fd5b503561055e565b6040518088815260200187815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019060200280838360005b8381101561016257818101518382015260200161014a565b505050509050019850505050505050505060405180910390f35b6100df6004803603602081101561019257600080fd5b50356106cf565b6101b6600480360360208110156101af57600080fd5b5035610768565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101f25781810151838201526020016101da565b505050509050019250505060405180910390f35b6100df610882565b61021661091e565b60408051918252519081900360200190f35b6102546004803603604081101561023e57600080fd5b50803590602001356001600160a01b0316610924565b604080519115158252519081900360200190f35b6102706109af565b604080516001600160a01b039092168252519081900360200190f35b6102546109be565b6102ba600480360360208110156102aa57600080fd5b50356001600160a01b03166109e2565b6040518080602001878152602001868152602001858152602001806020018060200184810384528a818151815260200191508051906020019080838360005b838110156103115781810151838201526020016102f9565b50505050905090810190601f16801561033e5780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b83811015610371578181015183820152602001610359565b50505050905090810190601f16801561039e5780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b838110156103d15781810151838201526020016103b9565b50505050905090810190601f1680156103fe5780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b6102166004803603604081101561042957600080fd5b5080359060200135610c9e565b6100df6004803603602081101561044c57600080fd5b50356001600160a01b0316610cf1565b6104646109be565b6104ae576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b6001600160a01b0381166104fb576040805162461bcd60e51b815260206004820152600f60248201526e6e6f207a65726f206164647265737360881b604482015290519081900360640190fd5b600280546001600160a01b038381166001600160a01b0319831681179093556040805191909216808252602082019390935281517fdc1f595b72ed3d34e3bd6fac603c555426330dc53cca030f6dcf09bcf2684816929181900390910190a15050565b6000806000806060600080600260009054906101000a90046001600160a01b03166001600160a01b031663135022c2896040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b1580156105c257600080fd5b505afa1580156105d6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260e08110156105ff57600080fd5b8151602083015160408085015160608601516080870180519351959794969295919491939282019284600160201b82111561063957600080fd5b90830190602082018581111561064e57600080fd5b82518660208202830111600160201b8211171561066a57600080fd5b82525081516020918201928201910280838360005b8381101561069757818101518382015260200161067f565b505050509190910160409081526020830151920151989f50969d50949b50929950909750919550929350505050919395979092949650565b6106d76109be565b610721576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b6001805490829055604080518281526020810184905281517f528990bbb5369a7f6d5acab41233e32bddb4882673d0208805b59cbad0dc1ec8929181900390910190a15050565b600254604080516347de074f60e01b81526004810184905290516060926001600160a01b0316916347de074f916024808301926000929190829003018186803b1580156107b457600080fd5b505afa1580156107c8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156107f157600080fd5b8101908080516040519392919084600160201b82111561081057600080fd5b90830190602082018581111561082557600080fd5b82518660208202830111600160201b8211171561084157600080fd5b82525081516020918201928201910280838360005b8381101561086e578181015183820152602001610856565b505050509050016040525050509050919050565b61088a6109be565b6108d4576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60015481565b60025460408051631f48b0d760e21b8152600481018590526001600160a01b03848116602483015291516000939290921691637d22c35c91604480820192602092909190829003018186803b15801561097c57600080fd5b505afa158015610990573d6000803e3d6000fd5b505050506040513d60208110156109a657600080fd5b50519392505050565b6000546001600160a01b031690565b600080546001600160a01b03166109d3610d4f565b6001600160a01b031614905090565b6002546040805163859da85f60e01b81526001600160a01b03848116600483015291516060936000938493849387938493169163859da85f9160248083019288929190829003018186803b158015610a3957600080fd5b505afa158015610a4d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260c0811015610a7657600080fd5b8101908080516040519392919084600160201b821115610a9557600080fd5b908301906020820185811115610aaa57600080fd5b8251600160201b811182820188101715610ac357600080fd5b82525081516020918201929091019080838360005b83811015610af0578181015183820152602001610ad8565b50505050905090810190601f168015610b1d5780820380516001836020036101000a031916815260200191505b50604081815260208301519083015160608401516080909401805192969195919284600160201b821115610b5057600080fd5b908301906020820185811115610b6557600080fd5b8251600160201b811182820188101715610b7e57600080fd5b82525081516020918201929091019080838360005b83811015610bab578181015183820152602001610b93565b50505050905090810190601f168015610bd85780820380516001836020036101000a031916815260200191505b5060405260200180516040519392919084600160201b821115610bfa57600080fd5b908301906020820185811115610c0f57600080fd5b8251600160201b811182820188101715610c2857600080fd5b82525081516020918201929091019080838360005b83811015610c55578181015183820152602001610c3d565b50505050905090810190601f168015610c825780820380516001836020036101000a031916815260200191505b5060405250969e959d50939b5091995097509550909350505050565b6002546040805163c7aa8ff760e01b8152600481018590526024810184905290516000926001600160a01b03169163c7aa8ff7916044808301926020929190829003018186803b15801561097c57600080fd5b610cf96109be565b610d43576040805162461bcd60e51b815260206004820152601660248201527527bbb730b136329d103737ba103a34329037bbb732b960511b604482015290519081900360640190fd5b610d4c81610d53565b50565b3390565b6001600160a01b038116610dae576040805162461bcd60e51b815260206004820152601860248201527f4f776e61626c653a206e6f207a65726f20616464726573730000000000000000604482015290519081900360640190fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a723158205e7dee108a16bc3a495ade83042b3c9705a4c725a49ab70d6b8099b77ff5cfec64736f6c63430005110032"

// DeployNodeListProxy deploys a new Ethereum contract, binding an instance of NodeListProxy to it.
func DeployNodeListProxy(auth *bind.TransactOpts, backend bind.ContractBackend, nodeListContractAddress common.Address, epoch *big.Int) (common.Address, *types.Transaction, *NodeListProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeListProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeListProxyBin), backend, nodeListContractAddress, epoch)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeListProxy{NodeListProxyCaller: NodeListProxyCaller{contract: contract}, NodeListProxyTransactor: NodeListProxyTransactor{contract: contract}, NodeListProxyFilterer: NodeListProxyFilterer{contract: contract}}, nil
}

// NodeListProxy is an auto generated Go binding around an Ethereum contract.
type NodeListProxy struct {
	NodeListProxyCaller     // Read-only binding to the contract
	NodeListProxyTransactor // Write-only binding to the contract
	NodeListProxyFilterer   // Log filterer for contract events
}

// NodeListProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeListProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeListProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeListProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeListProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeListProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeListProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeListProxySession struct {
	Contract     *NodeListProxy    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeListProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeListProxyCallerSession struct {
	Contract *NodeListProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// NodeListProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeListProxyTransactorSession struct {
	Contract     *NodeListProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// NodeListProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeListProxyRaw struct {
	Contract *NodeListProxy // Generic contract binding to access the raw methods on
}

// NodeListProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeListProxyCallerRaw struct {
	Contract *NodeListProxyCaller // Generic read-only contract binding to access the raw methods on
}

// NodeListProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeListProxyTransactorRaw struct {
	Contract *NodeListProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeListProxy creates a new instance of NodeListProxy, bound to a specific deployed contract.
func NewNodeListProxy(address common.Address, backend bind.ContractBackend) (*NodeListProxy, error) {
	contract, err := bindNodeListProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeListProxy{NodeListProxyCaller: NodeListProxyCaller{contract: contract}, NodeListProxyTransactor: NodeListProxyTransactor{contract: contract}, NodeListProxyFilterer: NodeListProxyFilterer{contract: contract}}, nil
}

// NewNodeListProxyCaller creates a new read-only instance of NodeListProxy, bound to a specific deployed contract.
func NewNodeListProxyCaller(address common.Address, caller bind.ContractCaller) (*NodeListProxyCaller, error) {
	contract, err := bindNodeListProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeListProxyCaller{contract: contract}, nil
}

// NewNodeListProxyTransactor creates a new write-only instance of NodeListProxy, bound to a specific deployed contract.
func NewNodeListProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeListProxyTransactor, error) {
	contract, err := bindNodeListProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeListProxyTransactor{contract: contract}, nil
}

// NewNodeListProxyFilterer creates a new log filterer instance of NodeListProxy, bound to a specific deployed contract.
func NewNodeListProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeListProxyFilterer, error) {
	contract, err := bindNodeListProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeListProxyFilterer{contract: contract}, nil
}

// bindNodeListProxy binds a generic wrapper to an already deployed contract.
func bindNodeListProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeListProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeListProxy *NodeListProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeListProxy.Contract.NodeListProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeListProxy *NodeListProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeListProxy.Contract.NodeListProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeListProxy *NodeListProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeListProxy.Contract.NodeListProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeListProxy *NodeListProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeListProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeListProxy *NodeListProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeListProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeListProxy *NodeListProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeListProxy.Contract.contract.Transact(opts, method, params...)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_NodeListProxy *NodeListProxyCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeListProxy.contract.Call(opts, out, "currentEpoch")
	return *ret0, err
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_NodeListProxy *NodeListProxySession) CurrentEpoch() (*big.Int, error) {
	return _NodeListProxy.Contract.CurrentEpoch(&_NodeListProxy.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_NodeListProxy *NodeListProxyCallerSession) CurrentEpoch() (*big.Int, error) {
	return _NodeListProxy.Contract.CurrentEpoch(&_NodeListProxy.CallOpts)
}

// GetEpochInfo is a free data retrieval call binding the contract method 0x135022c2.
//
// Solidity: function getEpochInfo(uint256 epoch) view returns(uint256 id, uint256 n, uint256 k, uint256 t, address[] nodeList, uint256 prevEpoch, uint256 nextEpoch)
func (_NodeListProxy *NodeListProxyCaller) GetEpochInfo(opts *bind.CallOpts, epoch *big.Int) (struct {
	Id        *big.Int
	N         *big.Int
	K         *big.Int
	T         *big.Int
	NodeList  []common.Address
	PrevEpoch *big.Int
	NextEpoch *big.Int
}, error) {
	ret := new(struct {
		Id        *big.Int
		N         *big.Int
		K         *big.Int
		T         *big.Int
		NodeList  []common.Address
		PrevEpoch *big.Int
		NextEpoch *big.Int
	})
	out := ret
	err := _NodeListProxy.contract.Call(opts, out, "getEpochInfo", epoch)
	return *ret, err
}

// GetEpochInfo is a free data retrieval call binding the contract method 0x135022c2.
//
// Solidity: function getEpochInfo(uint256 epoch) view returns(uint256 id, uint256 n, uint256 k, uint256 t, address[] nodeList, uint256 prevEpoch, uint256 nextEpoch)
func (_NodeListProxy *NodeListProxySession) GetEpochInfo(epoch *big.Int) (struct {
	Id        *big.Int
	N         *big.Int
	K         *big.Int
	T         *big.Int
	NodeList  []common.Address
	PrevEpoch *big.Int
	NextEpoch *big.Int
}, error) {
	return _NodeListProxy.Contract.GetEpochInfo(&_NodeListProxy.CallOpts, epoch)
}

// GetEpochInfo is a free data retrieval call binding the contract method 0x135022c2.
//
// Solidity: function getEpochInfo(uint256 epoch) view returns(uint256 id, uint256 n, uint256 k, uint256 t, address[] nodeList, uint256 prevEpoch, uint256 nextEpoch)
func (_NodeListProxy *NodeListProxyCallerSession) GetEpochInfo(epoch *big.Int) (struct {
	Id        *big.Int
	N         *big.Int
	K         *big.Int
	T         *big.Int
	NodeList  []common.Address
	PrevEpoch *big.Int
	NextEpoch *big.Int
}, error) {
	return _NodeListProxy.Contract.GetEpochInfo(&_NodeListProxy.CallOpts, epoch)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0xbafb3581.
//
// Solidity: function getNodeDetails(address nodeAddress) view returns(string declaredIp, uint256 position, uint256 pubKx, uint256 pubKy, string tmP2PListenAddress, string p2pListenAddress)
func (_NodeListProxy *NodeListProxyCaller) GetNodeDetails(opts *bind.CallOpts, nodeAddress common.Address) (struct {
	DeclaredIp         string
	Position           *big.Int
	PubKx              *big.Int
	PubKy              *big.Int
	TmP2PListenAddress string
	P2pListenAddress   string
}, error) {
	ret := new(struct {
		DeclaredIp         string
		Position           *big.Int
		PubKx              *big.Int
		PubKy              *big.Int
		TmP2PListenAddress string
		P2pListenAddress   string
	})
	out := ret
	err := _NodeListProxy.contract.Call(opts, out, "getNodeDetails", nodeAddress)
	return *ret, err
}

// GetNodeDetails is a free data retrieval call binding the contract method 0xbafb3581.
//
// Solidity: function getNodeDetails(address nodeAddress) view returns(string declaredIp, uint256 position, uint256 pubKx, uint256 pubKy, string tmP2PListenAddress, string p2pListenAddress)
func (_NodeListProxy *NodeListProxySession) GetNodeDetails(nodeAddress common.Address) (struct {
	DeclaredIp         string
	Position           *big.Int
	PubKx              *big.Int
	PubKy              *big.Int
	TmP2PListenAddress string
	P2pListenAddress   string
}, error) {
	return _NodeListProxy.Contract.GetNodeDetails(&_NodeListProxy.CallOpts, nodeAddress)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0xbafb3581.
//
// Solidity: function getNodeDetails(address nodeAddress) view returns(string declaredIp, uint256 position, uint256 pubKx, uint256 pubKy, string tmP2PListenAddress, string p2pListenAddress)
func (_NodeListProxy *NodeListProxyCallerSession) GetNodeDetails(nodeAddress common.Address) (struct {
	DeclaredIp         string
	Position           *big.Int
	PubKx              *big.Int
	PubKy              *big.Int
	TmP2PListenAddress string
	P2pListenAddress   string
}, error) {
	return _NodeListProxy.Contract.GetNodeDetails(&_NodeListProxy.CallOpts, nodeAddress)
}

// GetNodes is a free data retrieval call binding the contract method 0x47de074f.
//
// Solidity: function getNodes(uint256 epoch) view returns(address[])
func (_NodeListProxy *NodeListProxyCaller) GetNodes(opts *bind.CallOpts, epoch *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _NodeListProxy.contract.Call(opts, out, "getNodes", epoch)
	return *ret0, err
}

// GetNodes is a free data retrieval call binding the contract method 0x47de074f.
//
// Solidity: function getNodes(uint256 epoch) view returns(address[])
func (_NodeListProxy *NodeListProxySession) GetNodes(epoch *big.Int) ([]common.Address, error) {
	return _NodeListProxy.Contract.GetNodes(&_NodeListProxy.CallOpts, epoch)
}

// GetNodes is a free data retrieval call binding the contract method 0x47de074f.
//
// Solidity: function getNodes(uint256 epoch) view returns(address[])
func (_NodeListProxy *NodeListProxyCallerSession) GetNodes(epoch *big.Int) ([]common.Address, error) {
	return _NodeListProxy.Contract.GetNodes(&_NodeListProxy.CallOpts, epoch)
}

// GetPssStatus is a free data retrieval call binding the contract method 0xc7aa8ff7.
//
// Solidity: function getPssStatus(uint256 oldEpoch, uint256 newEpoch) view returns(uint256)
func (_NodeListProxy *NodeListProxyCaller) GetPssStatus(opts *bind.CallOpts, oldEpoch *big.Int, newEpoch *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeListProxy.contract.Call(opts, out, "getPssStatus", oldEpoch, newEpoch)
	return *ret0, err
}

// GetPssStatus is a free data retrieval call binding the contract method 0xc7aa8ff7.
//
// Solidity: function getPssStatus(uint256 oldEpoch, uint256 newEpoch) view returns(uint256)
func (_NodeListProxy *NodeListProxySession) GetPssStatus(oldEpoch *big.Int, newEpoch *big.Int) (*big.Int, error) {
	return _NodeListProxy.Contract.GetPssStatus(&_NodeListProxy.CallOpts, oldEpoch, newEpoch)
}

// GetPssStatus is a free data retrieval call binding the contract method 0xc7aa8ff7.
//
// Solidity: function getPssStatus(uint256 oldEpoch, uint256 newEpoch) view returns(uint256)
func (_NodeListProxy *NodeListProxyCallerSession) GetPssStatus(oldEpoch *big.Int, newEpoch *big.Int) (*big.Int, error) {
	return _NodeListProxy.Contract.GetPssStatus(&_NodeListProxy.CallOpts, oldEpoch, newEpoch)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_NodeListProxy *NodeListProxyCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NodeListProxy.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_NodeListProxy *NodeListProxySession) IsOwner() (bool, error) {
	return _NodeListProxy.Contract.IsOwner(&_NodeListProxy.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_NodeListProxy *NodeListProxyCallerSession) IsOwner() (bool, error) {
	return _NodeListProxy.Contract.IsOwner(&_NodeListProxy.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x7d22c35c.
//
// Solidity: function isWhitelisted(uint256 epoch, address nodeAddress) view returns(bool)
func (_NodeListProxy *NodeListProxyCaller) IsWhitelisted(opts *bind.CallOpts, epoch *big.Int, nodeAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NodeListProxy.contract.Call(opts, out, "isWhitelisted", epoch, nodeAddress)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x7d22c35c.
//
// Solidity: function isWhitelisted(uint256 epoch, address nodeAddress) view returns(bool)
func (_NodeListProxy *NodeListProxySession) IsWhitelisted(epoch *big.Int, nodeAddress common.Address) (bool, error) {
	return _NodeListProxy.Contract.IsWhitelisted(&_NodeListProxy.CallOpts, epoch, nodeAddress)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x7d22c35c.
//
// Solidity: function isWhitelisted(uint256 epoch, address nodeAddress) view returns(bool)
func (_NodeListProxy *NodeListProxyCallerSession) IsWhitelisted(epoch *big.Int, nodeAddress common.Address) (bool, error) {
	return _NodeListProxy.Contract.IsWhitelisted(&_NodeListProxy.CallOpts, epoch, nodeAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeListProxy *NodeListProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NodeListProxy.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeListProxy *NodeListProxySession) Owner() (common.Address, error) {
	return _NodeListProxy.Contract.Owner(&_NodeListProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeListProxy *NodeListProxyCallerSession) Owner() (common.Address, error) {
	return _NodeListProxy.Contract.Owner(&_NodeListProxy.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeListProxy *NodeListProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeListProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeListProxy *NodeListProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeListProxy.Contract.RenounceOwnership(&_NodeListProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeListProxy *NodeListProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeListProxy.Contract.RenounceOwnership(&_NodeListProxy.TransactOpts)
}

// SetCurrentEpoch is a paid mutator transaction binding the contract method 0x1dd6b9b1.
//
// Solidity: function setCurrentEpoch(uint256 _newEpoch) returns()
func (_NodeListProxy *NodeListProxyTransactor) SetCurrentEpoch(opts *bind.TransactOpts, _newEpoch *big.Int) (*types.Transaction, error) {
	return _NodeListProxy.contract.Transact(opts, "setCurrentEpoch", _newEpoch)
}

// SetCurrentEpoch is a paid mutator transaction binding the contract method 0x1dd6b9b1.
//
// Solidity: function setCurrentEpoch(uint256 _newEpoch) returns()
func (_NodeListProxy *NodeListProxySession) SetCurrentEpoch(_newEpoch *big.Int) (*types.Transaction, error) {
	return _NodeListProxy.Contract.SetCurrentEpoch(&_NodeListProxy.TransactOpts, _newEpoch)
}

// SetCurrentEpoch is a paid mutator transaction binding the contract method 0x1dd6b9b1.
//
// Solidity: function setCurrentEpoch(uint256 _newEpoch) returns()
func (_NodeListProxy *NodeListProxyTransactorSession) SetCurrentEpoch(_newEpoch *big.Int) (*types.Transaction, error) {
	return _NodeListProxy.Contract.SetCurrentEpoch(&_NodeListProxy.TransactOpts, _newEpoch)
}

// SetNodeListContract is a paid mutator transaction binding the contract method 0x08704c0a.
//
// Solidity: function setNodeListContract(address nodeListContractAddress) returns()
func (_NodeListProxy *NodeListProxyTransactor) SetNodeListContract(opts *bind.TransactOpts, nodeListContractAddress common.Address) (*types.Transaction, error) {
	return _NodeListProxy.contract.Transact(opts, "setNodeListContract", nodeListContractAddress)
}

// SetNodeListContract is a paid mutator transaction binding the contract method 0x08704c0a.
//
// Solidity: function setNodeListContract(address nodeListContractAddress) returns()
func (_NodeListProxy *NodeListProxySession) SetNodeListContract(nodeListContractAddress common.Address) (*types.Transaction, error) {
	return _NodeListProxy.Contract.SetNodeListContract(&_NodeListProxy.TransactOpts, nodeListContractAddress)
}

// SetNodeListContract is a paid mutator transaction binding the contract method 0x08704c0a.
//
// Solidity: function setNodeListContract(address nodeListContractAddress) returns()
func (_NodeListProxy *NodeListProxyTransactorSession) SetNodeListContract(nodeListContractAddress common.Address) (*types.Transaction, error) {
	return _NodeListProxy.Contract.SetNodeListContract(&_NodeListProxy.TransactOpts, nodeListContractAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeListProxy *NodeListProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NodeListProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeListProxy *NodeListProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeListProxy.Contract.TransferOwnership(&_NodeListProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeListProxy *NodeListProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeListProxy.Contract.TransferOwnership(&_NodeListProxy.TransactOpts, newOwner)
}

// NodeListProxyEpochChangedIterator is returned from FilterEpochChanged and is used to iterate over the raw logs and unpacked data for EpochChanged events raised by the NodeListProxy contract.
type NodeListProxyEpochChangedIterator struct {
	Event *NodeListProxyEpochChanged // Event containing the contract specifics and raw log

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
func (it *NodeListProxyEpochChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeListProxyEpochChanged)
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
		it.Event = new(NodeListProxyEpochChanged)
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
func (it *NodeListProxyEpochChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeListProxyEpochChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeListProxyEpochChanged represents a EpochChanged event raised by the NodeListProxy contract.
type NodeListProxyEpochChanged struct {
	OldEpoch *big.Int
	NewEpoch *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEpochChanged is a free log retrieval operation binding the contract event 0x528990bbb5369a7f6d5acab41233e32bddb4882673d0208805b59cbad0dc1ec8.
//
// Solidity: event EpochChanged(uint256 oldEpoch, uint256 newEpoch)
func (_NodeListProxy *NodeListProxyFilterer) FilterEpochChanged(opts *bind.FilterOpts) (*NodeListProxyEpochChangedIterator, error) {

	logs, sub, err := _NodeListProxy.contract.FilterLogs(opts, "EpochChanged")
	if err != nil {
		return nil, err
	}
	return &NodeListProxyEpochChangedIterator{contract: _NodeListProxy.contract, event: "EpochChanged", logs: logs, sub: sub}, nil
}

// WatchEpochChanged is a free log subscription operation binding the contract event 0x528990bbb5369a7f6d5acab41233e32bddb4882673d0208805b59cbad0dc1ec8.
//
// Solidity: event EpochChanged(uint256 oldEpoch, uint256 newEpoch)
func (_NodeListProxy *NodeListProxyFilterer) WatchEpochChanged(opts *bind.WatchOpts, sink chan<- *NodeListProxyEpochChanged) (event.Subscription, error) {

	logs, sub, err := _NodeListProxy.contract.WatchLogs(opts, "EpochChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeListProxyEpochChanged)
				if err := _NodeListProxy.contract.UnpackLog(event, "EpochChanged", log); err != nil {
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

// ParseEpochChanged is a log parse operation binding the contract event 0x528990bbb5369a7f6d5acab41233e32bddb4882673d0208805b59cbad0dc1ec8.
//
// Solidity: event EpochChanged(uint256 oldEpoch, uint256 newEpoch)
func (_NodeListProxy *NodeListProxyFilterer) ParseEpochChanged(log types.Log) (*NodeListProxyEpochChanged, error) {
	event := new(NodeListProxyEpochChanged)
	if err := _NodeListProxy.contract.UnpackLog(event, "EpochChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NodeListProxyNodeListContractChangedIterator is returned from FilterNodeListContractChanged and is used to iterate over the raw logs and unpacked data for NodeListContractChanged events raised by the NodeListProxy contract.
type NodeListProxyNodeListContractChangedIterator struct {
	Event *NodeListProxyNodeListContractChanged // Event containing the contract specifics and raw log

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
func (it *NodeListProxyNodeListContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeListProxyNodeListContractChanged)
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
		it.Event = new(NodeListProxyNodeListContractChanged)
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
func (it *NodeListProxyNodeListContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeListProxyNodeListContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeListProxyNodeListContractChanged represents a NodeListContractChanged event raised by the NodeListProxy contract.
type NodeListProxyNodeListContractChanged struct {
	OldContract common.Address
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeListContractChanged is a free log retrieval operation binding the contract event 0xdc1f595b72ed3d34e3bd6fac603c555426330dc53cca030f6dcf09bcf2684816.
//
// Solidity: event NodeListContractChanged(address oldContract, address newContract)
func (_NodeListProxy *NodeListProxyFilterer) FilterNodeListContractChanged(opts *bind.FilterOpts) (*NodeListProxyNodeListContractChangedIterator, error) {

	logs, sub, err := _NodeListProxy.contract.FilterLogs(opts, "NodeListContractChanged")
	if err != nil {
		return nil, err
	}
	return &NodeListProxyNodeListContractChangedIterator{contract: _NodeListProxy.contract, event: "NodeListContractChanged", logs: logs, sub: sub}, nil
}

// WatchNodeListContractChanged is a free log subscription operation binding the contract event 0xdc1f595b72ed3d34e3bd6fac603c555426330dc53cca030f6dcf09bcf2684816.
//
// Solidity: event NodeListContractChanged(address oldContract, address newContract)
func (_NodeListProxy *NodeListProxyFilterer) WatchNodeListContractChanged(opts *bind.WatchOpts, sink chan<- *NodeListProxyNodeListContractChanged) (event.Subscription, error) {

	logs, sub, err := _NodeListProxy.contract.WatchLogs(opts, "NodeListContractChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeListProxyNodeListContractChanged)
				if err := _NodeListProxy.contract.UnpackLog(event, "NodeListContractChanged", log); err != nil {
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

// ParseNodeListContractChanged is a log parse operation binding the contract event 0xdc1f595b72ed3d34e3bd6fac603c555426330dc53cca030f6dcf09bcf2684816.
//
// Solidity: event NodeListContractChanged(address oldContract, address newContract)
func (_NodeListProxy *NodeListProxyFilterer) ParseNodeListContractChanged(log types.Log) (*NodeListProxyNodeListContractChanged, error) {
	event := new(NodeListProxyNodeListContractChanged)
	if err := _NodeListProxy.contract.UnpackLog(event, "NodeListContractChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NodeListProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NodeListProxy contract.
type NodeListProxyOwnershipTransferredIterator struct {
	Event *NodeListProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NodeListProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeListProxyOwnershipTransferred)
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
		it.Event = new(NodeListProxyOwnershipTransferred)
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
func (it *NodeListProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeListProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeListProxyOwnershipTransferred represents a OwnershipTransferred event raised by the NodeListProxy contract.
type NodeListProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeListProxy *NodeListProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodeListProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeListProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeListProxyOwnershipTransferredIterator{contract: _NodeListProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeListProxy *NodeListProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodeListProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeListProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeListProxyOwnershipTransferred)
				if err := _NodeListProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NodeListProxy *NodeListProxyFilterer) ParseOwnershipTransferred(log types.Log) (*NodeListProxyOwnershipTransferred, error) {
	event := new(NodeListProxyOwnershipTransferred)
	if err := _NodeListProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
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
