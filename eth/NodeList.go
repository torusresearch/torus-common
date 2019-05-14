// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// NodeListABI is the input ABI used to generate the binding from.
const NodeListABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getEpochInfo\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"n\",\"type\":\"uint256\"},{\"name\":\"k\",\"type\":\"uint256\"},{\"name\":\"t\",\"type\":\"uint256\"},{\"name\":\"nodeList\",\"type\":\"address[]\"},{\"name\":\"prevEpoch\",\"type\":\"uint256\"},{\"name\":\"nextEpoch\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochInfo\",\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"n\",\"type\":\"uint256\"},{\"name\":\"k\",\"type\":\"uint256\"},{\"name\":\"t\",\"type\":\"uint256\"},{\"name\":\"prevEpoch\",\"type\":\"uint256\"},{\"name\":\"nextEpoch\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\"},{\"name\":\"nodeAddress\",\"type\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\"},{\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"IsWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeDetails\",\"outputs\":[{\"name\":\"declaredIp\",\"type\":\"string\"},{\"name\":\"position\",\"type\":\"uint256\"},{\"name\":\"pubKx\",\"type\":\"uint256\"},{\"name\":\"pubKy\",\"type\":\"uint256\"},{\"name\":\"tmP2PListenAddress\",\"type\":\"string\"},{\"name\":\"p2pListenAddress\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\"},{\"name\":\"n\",\"type\":\"uint256\"},{\"name\":\"k\",\"type\":\"uint256\"},{\"name\":\"t\",\"type\":\"uint256\"},{\"name\":\"nodeList\",\"type\":\"address[]\"},{\"name\":\"prevEpoch\",\"type\":\"uint256\"},{\"name\":\"nextEpoch\",\"type\":\"uint256\"}],\"name\":\"updateEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeDetails\",\"outputs\":[{\"name\":\"declaredIp\",\"type\":\"string\"},{\"name\":\"position\",\"type\":\"uint256\"},{\"name\":\"tmP2PListenAddress\",\"type\":\"string\"},{\"name\":\"p2pListenAddress\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\"},{\"name\":\"declaredIp\",\"type\":\"string\"},{\"name\":\"pubKx\",\"type\":\"uint256\"},{\"name\":\"pubKy\",\"type\":\"uint256\"},{\"name\":\"tmP2PListenAddress\",\"type\":\"string\"},{\"name\":\"p2pListenAddress\",\"type\":\"string\"}],\"name\":\"listNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"publicKey\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"NodeListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// NodeListBin is the compiled bytecode used for deploying new contracts.
const NodeListBin = `0x608060405234801561001057600080fd5b50600080546001600160a01b03191633178082556040516001600160a01b039190911691907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3611426806100696000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063859da85f1161008c578063ae4df20c11610066578063ae4df20c1461049c578063bafb35811461055a578063bf2d6f81146106cb578063f2fde38b146107ec576100ea565b8063859da85f146102f15780638da5cb5b146104705780638f32d59b14610494576100ea565b806347de074f116100c857806347de074f146102105780634b25bfce1461027d578063715018a6146102bd5780637b08d896146102c5576100ea565b8063135022c2146100ef5780633894228e1461018a5780633d4602a9146101da575b600080fd5b61010c6004803603602081101561010557600080fd5b5035610812565b6040518088815260200187815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019060200280838360005b83811015610170578181015183820152602001610158565b505050509050019850505050505050505060405180910390f35b6101a7600480360360208110156101a057600080fd5b503561091e565b604080519687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b61020e600480360360608110156101f057600080fd5b508035906001600160a01b0360208201351690604001351515610955565b005b61022d6004803603602081101561022657600080fd5b50356109a7565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610269578181015183820152602001610251565b505050509050019250505060405180910390f35b6102a96004803603604081101561029357600080fd5b50803590602001356001600160a01b0316610a22565b604080519115158252519081900360200190f35b61020e610a42565b6102a9600480360360408110156102db57600080fd5b50803590602001356001600160a01b0316610a9d565b6103176004803603602081101561030757600080fd5b50356001600160a01b0316610ac8565b6040518080602001878152602001868152602001858152602001806020018060200184810384528a818151815260200191508051906020019080838360005b8381101561036e578181015183820152602001610356565b50505050905090810190601f16801561039b5780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b838110156103ce5781810151838201526020016103b6565b50505050905090810190601f1680156103fb5780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b8381101561042e578181015183820152602001610416565b50505050905090810190601f16801561045b5780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b610478610ca9565b604080516001600160a01b039092168252519081900360200190f35b6102a9610cb9565b61020e600480360360e08110156104b257600080fd5b81359160208101359160408201359160608101359181019060a081016080820135600160201b8111156104e457600080fd5b8201836020820111156104f657600080fd5b803590602001918460208302840111600160201b8311171561051757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505082359350505060200135610cca565b6105806004803603602081101561057057600080fd5b50356001600160a01b0316610d7d565b60405180806020018581526020018060200180602001848103845288818151815260200191508051906020019080838360005b838110156105cb5781810151838201526020016105b3565b50505050905090810190601f1680156105f85780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b8381101561062b578181015183820152602001610613565b50505050905090810190601f1680156106585780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b8381101561068b578181015183820152602001610673565b50505050905090810190601f1680156106b85780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b61020e600480360360c08110156106e157600080fd5b81359190810190604081016020820135600160201b81111561070257600080fd5b82018360208201111561071457600080fd5b803590602001918460018302840111600160201b8311171561073557600080fd5b919390928235926020810135929190606081019060400135600160201b81111561075e57600080fd5b82018360208201111561077057600080fd5b803590602001918460018302840111600160201b8311171561079157600080fd5b919390929091602081019035600160201b8111156107ae57600080fd5b8201836020820111156107c057600080fd5b803590602001918460018302840111600160201b831117156107e157600080fd5b509092509050610fb5565b61020e6004803603602081101561080257600080fd5b50356001600160a01b03166111df565b600080808060608180878061082657600080fd5b61082e61126a565b60008a815260026020818152604092839020835160e081018552815481526001820154818401529281015483850152600381015460608401526004810180548551818502810185019096528086529394919360808601938301828280156108be57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116108a0575b50505050508152602001600582015481526020016006820154815250509050806000015181602001518260400151836060015184608001518560a001518660c0015182925098509850985098509850985098505050919395979092949650565b6002602081905260009182526040909120805460018201549282015460038301546005840154600690940154929493919290919086565b61095d610cb9565b61096657600080fd5b828061097157600080fd5b5060009283526001602090815260408085206001600160a01b039490941685529290529120805460ff1916911515919091179055565b606081806109b457600080fd5b60008381526002602090815260409182902060040180548351818402810184019094528084529091830182828015610a1557602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116109f7575b5050505050915050919050565b600160209081526000928352604080842090915290825290205460ff1681565b610a4a610cb9565b610a5357600080fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60009182526001602090815260408084206001600160a01b0393909316845291905290205460ff1690565b60036020908152600091825260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452909291839190830182828015610b5f5780601f10610b3457610100808354040283529160200191610b5f565b820191906000526020600020905b815481529060010190602001808311610b4257829003601f168201915b505050505090806001015490806002015490806003015490806004018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c0f5780601f10610be457610100808354040283529160200191610c0f565b820191906000526020600020905b815481529060010190602001808311610bf257829003601f168201915b5050505060058301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015610c9f5780601f10610c7457610100808354040283529160200191610c9f565b820191906000526020600020905b815481529060010190602001808311610c8257829003601f168201915b5050505050905086565b6000546001600160a01b03165b90565b6000546001600160a01b0316331490565b610cd2610cb9565b610cdb57600080fd5b8680610ce657600080fd5b6040805160e08101825289815260208082018a81528284018a8152606084018a8152608085018a815260a086018a905260c0860189905260008f8152600280875297902086518155935160018501559151958301959095559351600382015592518051929392610d5c92600485019201906112a7565b5060a0820151600582015560c0909101516006909101555050505050505050565b60606000606080610d8c61130c565b6001600160a01b0386166000908152600360209081526040918290208251815460026001821615610100026000190190911604601f8101849004909302810160e090810190945260c08101838152909391928492849190840182828015610e345780601f10610e0957610100808354040283529160200191610e34565b820191906000526020600020905b815481529060010190602001808311610e1757829003601f168201915b50505050508152602001600182015481526020016002820154815260200160038201548152602001600482018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ef45780601f10610ec957610100808354040283529160200191610ef4565b820191906000526020600020905b815481529060010190602001808311610ed757829003601f168201915b505050918352505060058201805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152938201939291830182828015610f885780601f10610f5d57610100808354040283529160200191610f88565b820191906000526020600020905b815481529060010190602001808311610f6b57829003601f168201915b50505091909252505081516020830151608084015160a090940151919a9099509297509550909350505050565b88610fc08133610a9d565b610fc957600080fd5b8980610fd457600080fd5b60008b8152600260205260409020548b908114610ff057600080fd5b60008c81526002602090815260408083206004810180546001810182559085529383902090930180546001600160a01b03191633179055805160e0601f8f018490049093028101830190915260c081018d815290918291908f908f9081908501838280828437600092019190915250505090825250600483015460208083019190915260408083018e9052606083018d90528051601f8c018390048302810183019091528a8152608090920191908b908b9081908401838280828437600092019190915250505090825250604080516020601f8a0181900481028201810190925288815291810191908990899081908401838280828437600092018290525093909452505033815260036020908152604090912083518051919350611119928492910190611342565b5060208281015160018301556040830151600283015560608301516003830155608083015180516111509260048501920190611342565b5060a0820151805161116c916005840191602090910190611342565b509050507fe2f8adb0f494dc82ccf446c031763ef3762d6396d51664611ed89aac0117339e338e836004018054905060405180846001600160a01b03166001600160a01b03168152602001838152602001828152602001935050505060405180910390a150505050505050505050505050565b6111e7610cb9565b6111f057600080fd5b6111f9816111fc565b50565b6001600160a01b03811661120f57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6040518060e00160405280600081526020016000815260200160008152602001600081526020016060815260200160008152602001600081525090565b8280548282559060005260206000209081019282156112fc579160200282015b828111156112fc57825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906112c7565b506113089291506113bc565b5090565b6040518060c001604052806060815260200160008152602001600081526020016000815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061138357805160ff19168380011785556113b0565b828001600101855582156113b0579182015b828111156113b0578251825591602001919060010190611395565b506113089291506113e0565b610cb691905b808211156113085780546001600160a01b03191681556001016113c2565b610cb691905b8082111561130857600081556001016113e656fea165627a7a7230582043f5894585c291929bd1a26834120526009bda0d533547efd65a237d167dcfc90029`

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

// IsWhitelisted is a free data retrieval call binding the contract method 0x7b08d896.
//
// Solidity: function IsWhitelisted(uint256 epoch, address nodeAddress) constant returns(bool)
func (_NodeList *NodeListCaller) IsWhitelisted(opts *bind.CallOpts, epoch *big.Int, nodeAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NodeList.contract.Call(opts, out, "IsWhitelisted", epoch, nodeAddress)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x7b08d896.
//
// Solidity: function IsWhitelisted(uint256 epoch, address nodeAddress) constant returns(bool)
func (_NodeList *NodeListSession) IsWhitelisted(epoch *big.Int, nodeAddress common.Address) (bool, error) {
	return _NodeList.Contract.IsWhitelisted(&_NodeList.CallOpts, epoch, nodeAddress)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x7b08d896.
//
// Solidity: function IsWhitelisted(uint256 epoch, address nodeAddress) constant returns(bool)
func (_NodeList *NodeListCallerSession) IsWhitelisted(epoch *big.Int, nodeAddress common.Address) (bool, error) {
	return _NodeList.Contract.IsWhitelisted(&_NodeList.CallOpts, epoch, nodeAddress)
}

// EpochInfo is a free data retrieval call binding the contract method 0x3894228e.
//
// Solidity: function epochInfo(uint256 ) constant returns(uint256 id, uint256 n, uint256 k, uint256 t, uint256 prevEpoch, uint256 nextEpoch)
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
// Solidity: function epochInfo(uint256 ) constant returns(uint256 id, uint256 n, uint256 k, uint256 t, uint256 prevEpoch, uint256 nextEpoch)
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
// Solidity: function epochInfo(uint256 ) constant returns(uint256 id, uint256 n, uint256 k, uint256 t, uint256 prevEpoch, uint256 nextEpoch)
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
// Solidity: function getEpochInfo(uint256 epoch) constant returns(uint256 id, uint256 n, uint256 k, uint256 t, address[] nodeList, uint256 prevEpoch, uint256 nextEpoch)
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
// Solidity: function getEpochInfo(uint256 epoch) constant returns(uint256 id, uint256 n, uint256 k, uint256 t, address[] nodeList, uint256 prevEpoch, uint256 nextEpoch)
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
// Solidity: function getEpochInfo(uint256 epoch) constant returns(uint256 id, uint256 n, uint256 k, uint256 t, address[] nodeList, uint256 prevEpoch, uint256 nextEpoch)
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
// Solidity: function getNodeDetails(address nodeAddress) constant returns(string declaredIp, uint256 position, string tmP2PListenAddress, string p2pListenAddress)
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
// Solidity: function getNodeDetails(address nodeAddress) constant returns(string declaredIp, uint256 position, string tmP2PListenAddress, string p2pListenAddress)
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
// Solidity: function getNodeDetails(address nodeAddress) constant returns(string declaredIp, uint256 position, string tmP2PListenAddress, string p2pListenAddress)
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
// Solidity: function getNodes(uint256 epoch) constant returns(address[])
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
// Solidity: function getNodes(uint256 epoch) constant returns(address[])
func (_NodeList *NodeListSession) GetNodes(epoch *big.Int) ([]common.Address, error) {
	return _NodeList.Contract.GetNodes(&_NodeList.CallOpts, epoch)
}

// GetNodes is a free data retrieval call binding the contract method 0x47de074f.
//
// Solidity: function getNodes(uint256 epoch) constant returns(address[])
func (_NodeList *NodeListCallerSession) GetNodes(epoch *big.Int) ([]common.Address, error) {
	return _NodeList.Contract.GetNodes(&_NodeList.CallOpts, epoch)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
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
// Solidity: function isOwner() constant returns(bool)
func (_NodeList *NodeListSession) IsOwner() (bool, error) {
	return _NodeList.Contract.IsOwner(&_NodeList.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_NodeList *NodeListCallerSession) IsOwner() (bool, error) {
	return _NodeList.Contract.IsOwner(&_NodeList.CallOpts)
}

// NodeDetails is a free data retrieval call binding the contract method 0x859da85f.
//
// Solidity: function nodeDetails(address ) constant returns(string declaredIp, uint256 position, uint256 pubKx, uint256 pubKy, string tmP2PListenAddress, string p2pListenAddress)
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
// Solidity: function nodeDetails(address ) constant returns(string declaredIp, uint256 position, uint256 pubKx, uint256 pubKy, string tmP2PListenAddress, string p2pListenAddress)
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
// Solidity: function nodeDetails(address ) constant returns(string declaredIp, uint256 position, uint256 pubKx, uint256 pubKy, string tmP2PListenAddress, string p2pListenAddress)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
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
// Solidity: function owner() constant returns(address)
func (_NodeList *NodeListSession) Owner() (common.Address, error) {
	return _NodeList.Contract.Owner(&_NodeList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NodeList *NodeListCallerSession) Owner() (common.Address, error) {
	return _NodeList.Contract.Owner(&_NodeList.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x4b25bfce.
//
// Solidity: function whitelist(uint256 , address ) constant returns(bool)
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
// Solidity: function whitelist(uint256 , address ) constant returns(bool)
func (_NodeList *NodeListSession) Whitelist(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _NodeList.Contract.Whitelist(&_NodeList.CallOpts, arg0, arg1)
}

// Whitelist is a free data retrieval call binding the contract method 0x4b25bfce.
//
// Solidity: function whitelist(uint256 , address ) constant returns(bool)
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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
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
