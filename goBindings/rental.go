// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
)

// RentalContractMetaData contains all meta data concerning the RentalContract contract.
var RentalContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"}],\"name\":\"PropertyListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"checkin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"checkout\",\"type\":\"uint256\"}],\"name\":\"PropertyRented\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Time_call\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"available\",\"type\":\"bool\"}],\"name\":\"UpdatePropertyAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"dispute\",\"type\":\"bool\"}],\"name\":\"fileDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"perNight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"payInFull\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"securityDeposit\",\"type\":\"uint256\"}],\"name\":\"listProperty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"properties\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"perNight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"available\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"payInFull\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"securityDeposit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"renter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"releaseDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"t1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"t2\",\"type\":\"uint256\"}],\"name\":\"rentProperty\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rentals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"t1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"t2\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"securityDeposits\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"dispute\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"updateRentalPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"03261030": "Time_call()",
		"f913a4bc": "UpdatePropertyAvailability(uint256,bool)",
		"f40d0602": "fileDispute(address,uint256,bool)",
		"a8d659bb": "listProperty(uint256,bool,uint256)",
		"14bfcdd7": "properties(address,uint256)",
		"524dd059": "releaseDeposit(address,uint256)",
		"3b88fa21": "rentProperty(address,uint256,uint256,uint256)",
		"090fc58b": "rentals(address,uint256)",
		"f12fdc4e": "securityDeposits(address,uint256)",
		"90abcd98": "updateRentalPrice(uint256,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610fc8806100206000396000f3fe6080604052600436106100915760003560e01c806390abcd981161005957806390abcd9814610228578063a8d659bb14610248578063f12fdc4e14610268578063f40d06021461030e578063f913a4bc1461032e57600080fd5b80630326103014610096578063090fc58b146100b657806314bfcdd71461014f5780633b88fa21146101f3578063524dd05914610208575b600080fd5b3480156100a257600080fd5b506040514281526020015b60405180910390f35b3480156100c257600080fd5b506101166100d1366004610d9e565b60016020818152600093845260408085209091529183529120805491810154600282015460038301546004909301546001600160a01b03948516949293919092169185565b604080516001600160a01b039687168152602081019590955292909416918301919091526060820152608081019190915260a0016100ad565b34801561015b57600080fd5b506101b861016a366004610d9e565b6000602081815292815260408082209093529081522080546001820154600283015460038401546004909401546001600160a01b03909316939192909160ff80821692610100909204169086565b604080516001600160a01b039097168752602087019590955293850192909252151560608401521515608083015260a082015260c0016100ad565b610206610201366004610dca565b610377565b005b34801561021457600080fd5b50610206610223366004610d9e565b610a33565b34801561023457600080fd5b50610206610243366004610e05565b610be3565b34801561025457600080fd5b50610206610263366004610e3c565b610c42565b34801561027457600080fd5b506102d0610283366004610d9e565b600460208181526000938452604080852090915291835291208054600182015460028301546003840154948401546005909401546001600160a01b03938416959294919093169260ff1686565b604080516001600160a01b039788168152602081019690965293909516928401929092526060830152608082015290151560a082015260c0016100ad565b34801561031a57600080fd5b50610206610329366004610e71565b610d24565b34801561033a57600080fd5b50610206610349366004610eaf565b3360009081526020818152604080832094835293905291909120600301805460ff1916911515919091179055565b6001600160a01b038416600090815260026020526040812054819081908190819081906103a690600190610ef1565b8911156103e95760405162461bcd60e51b815260206004820152600c60248201526b1251081b9bdd08199bdd5b9960a21b60448201526064015b60405180910390fd5b6103f38888610ef1565b93506201518084111561040557600080fd5b60005b6001600160a01b038b166000908152602081815260408083208d84529091529020600501548110156104e9576001600160a01b038b166000908152602081815260408083208d8452909152902060050180548290811061046a5761046a610f08565b60009182526020808320909101546001600160a01b038e16835282825260408084208e85529092529120600601805491985090829081106104ad576104ad610f08565b9060005260206000200154955088861115806104c95750878710155b15156001146104d757600080fd5b806104e181610f1e565b915050610408565b50506001600160a01b0389166000908152602081815260408083208b845290915290206004810154600290910154819062015180906105289087610f39565b6105329190610f58565b61053c9190610f7a565b6001600160a01b038b166000908152602081815260408083208d845290915290206002015490935062015180906105739086610f39565b61057d9190610f58565b91508234101561058c57600080fd5b60036000336001600160a01b03166001600160a01b03168152602001908152602001600020805490506005819055508960016000336001600160a01b03166001600160a01b031681526020019081526020016000206000600554815260200190815260200160002060000160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508860016000336001600160a01b03166001600160a01b0316815260200190815260200160002060006005548152602001908152602001600020600101819055503360016000336001600160a01b03166001600160a01b031681526020019081526020016000206000600554815260200190815260200160002060020160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508760016000336001600160a01b03166001600160a01b0316815260200190815260200160002060006005548152602001908152602001600020600301819055508660016000336001600160a01b03166001600160a01b0316815260200190815260200160002060006005548152602001908152602001600020600401819055506000808b6001600160a01b03166001600160a01b0316815260200190815260200160002060008a81526020019081526020016000206005018890806001815401808255809150506001900390600052602060002001600090919091909150556000808b6001600160a01b03166001600160a01b0316815260200190815260200160002060008a81526020019081526020016000206006018790806001815401808255809150506001900390600052602060002001600090919091909150558960046000336001600160a01b03166001600160a01b031681526020019081526020016000206000600554815260200190815260200160002060000160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508860046000336001600160a01b03166001600160a01b0316815260200190815260200160002060006005548152602001908152602001600020600101819055503360046000336001600160a01b03166001600160a01b031681526020019081526020016000206000600554815260200190815260200160002060020160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508060046000336001600160a01b03166001600160a01b031681526020019081526020016000206000600554815260200190815260200160002060030160008282546109449190610f7a565b90915550503360008181526004602081815260408084206005805486529083528185208101805460ff19169055805485528185204294019390935593835260038152838320915482546001810184559284529083209091015590516001600160a01b038c16916108fc851502918591818181858888f193505050501580156109d0573d6000803e3d6000fd5b50604080516001600160a01b038c168152602081018b90523381830152606081018a90526080810189905290517fda21554a181142009a8ac60c8d97feeae9dcd4433f8d31bbe4ee4080692a4f3b9181900360a00190a150505050505050505050565b6001600160a01b038281166000908152600460209081526040808320858452909152812054909116331480610a8f57506001600160a01b0383811660009081526004602090815260408083208684529091529020600201541633145b610a9857600080fd5b6001600160a01b038316600090815260046020818152604080842086855290915290912001544290610acd9062093a80610f7a565b10610b125760405162461bcd60e51b815260206004820152601560248201527411195c1bdcda5d081cdd1a5b1b081bdb881a1bdb19605a1b60448201526064016103e0565b6001600160a01b038316600090815260046020908152604080832085845290915290206005015460ff1615610b895760405162461bcd60e51b815260206004820152601960248201527f4f776e6572206861732066696c6564206120646973707574650000000000000060448201526064016103e0565b506001600160a01b038216600081815260046020908152604080832085845290915280822060030154905190929183156108fc02918491818181858888f19350505050158015610bdd573d6000803e3d6000fd5b50505050565b336000818152602081815260408083208584528252918290206002018590558151928352820184905281018290527fb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c09060600160405180910390a15050565b3360008181526002602081815260408084208054600581815586855283872091875290845282862080546001600160a01b03191688179055805480875283872060018181019290925586018b9055815487528387206003908101805460ff1916831790558254885284882001805461ff0019166101008c15150217905581548752838720600401899055948452805482549586018355918652948390209093019290925591548151938452918301919091527fd45ae8d02a95e46a15b86aa93911a32c7e0bf35515df8de4945f3afb1ba11039910160405180910390a1505050565b6001600160a01b038381166000908152600460209081526040808320868452909152902054163314610d5557600080fd5b33600090815260046020908152604080832094835293905291909120600501805460ff191691151591909117905550565b6001600160a01b0381168114610d9b57600080fd5b50565b60008060408385031215610db157600080fd5b8235610dbc81610d86565b946020939093013593505050565b60008060008060808587031215610de057600080fd5b8435610deb81610d86565b966020860135965060408601359560600135945092505050565b60008060408385031215610e1857600080fd5b50508035926020909101359150565b80358015158114610e3757600080fd5b919050565b600080600060608486031215610e5157600080fd5b83359250610e6160208501610e27565b9150604084013590509250925092565b600080600060608486031215610e8657600080fd5b8335610e9181610d86565b925060208401359150610ea660408501610e27565b90509250925092565b60008060408385031215610ec257600080fd5b82359150610ed260208401610e27565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b600082821015610f0357610f03610edb565b500390565b634e487b7160e01b600052603260045260246000fd5b6000600019821415610f3257610f32610edb565b5060010190565b6000816000190483118215151615610f5357610f53610edb565b500290565b600082610f7557634e487b7160e01b600052601260045260246000fd5b500490565b60008219821115610f8d57610f8d610edb565b50019056fea26469706673582212204d47fc8a407712146dd4e152cdc7c6b6484b0dad0a501cf0f36370e3ff359b7364736f6c634300080a0033",
}

// RentalContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RentalContractMetaData.ABI instead.
var RentalContractABI = RentalContractMetaData.ABI

// Deprecated: Use RentalContractMetaData.Sigs instead.
// RentalContractFuncSigs maps the 4-byte function signature to its string representation.
var RentalContractFuncSigs = RentalContractMetaData.Sigs

// RentalContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RentalContractMetaData.Bin instead.
var RentalContractBin = RentalContractMetaData.Bin

// DeployRentalContract deploys a new Ethereum contract, binding an instance of RentalContract to it.
func DeployRentalContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RentalContract, error) {
	parsed, err := RentalContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RentalContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RentalContract{RentalContractCaller: RentalContractCaller{contract: contract}, RentalContractTransactor: RentalContractTransactor{contract: contract}, RentalContractFilterer: RentalContractFilterer{contract: contract}}, nil
}

// RentalContract is an auto generated Go binding around an Ethereum contract.
type RentalContract struct {
	RentalContractCaller     // Read-only binding to the contract
	RentalContractTransactor // Write-only binding to the contract
	RentalContractFilterer   // Log filterer for contract events
}

// RentalContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RentalContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RentalContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RentalContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RentalContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RentalContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RentalContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RentalContractSession struct {
	Contract     *RentalContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RentalContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RentalContractCallerSession struct {
	Contract *RentalContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RentalContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RentalContractTransactorSession struct {
	Contract     *RentalContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RentalContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RentalContractRaw struct {
	Contract *RentalContract // Generic contract binding to access the raw methods on
}

// RentalContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RentalContractCallerRaw struct {
	Contract *RentalContractCaller // Generic read-only contract binding to access the raw methods on
}

// RentalContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RentalContractTransactorRaw struct {
	Contract *RentalContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRentalContract creates a new instance of RentalContract, bound to a specific deployed contract.
func NewRentalContract(address common.Address, backend bind.ContractBackend) (*RentalContract, error) {
	contract, err := bindRentalContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RentalContract{RentalContractCaller: RentalContractCaller{contract: contract}, RentalContractTransactor: RentalContractTransactor{contract: contract}, RentalContractFilterer: RentalContractFilterer{contract: contract}}, nil
}

// NewRentalContractCaller creates a new read-only instance of RentalContract, bound to a specific deployed contract.
func NewRentalContractCaller(address common.Address, caller bind.ContractCaller) (*RentalContractCaller, error) {
	contract, err := bindRentalContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RentalContractCaller{contract: contract}, nil
}

// NewRentalContractTransactor creates a new write-only instance of RentalContract, bound to a specific deployed contract.
func NewRentalContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RentalContractTransactor, error) {
	contract, err := bindRentalContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RentalContractTransactor{contract: contract}, nil
}

// NewRentalContractFilterer creates a new log filterer instance of RentalContract, bound to a specific deployed contract.
func NewRentalContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RentalContractFilterer, error) {
	contract, err := bindRentalContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RentalContractFilterer{contract: contract}, nil
}

// bindRentalContract binds a generic wrapper to an already deployed contract.
func bindRentalContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RentalContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RentalContract *RentalContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RentalContract.Contract.RentalContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RentalContract *RentalContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RentalContract.Contract.RentalContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RentalContract *RentalContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RentalContract.Contract.RentalContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RentalContract *RentalContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RentalContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RentalContract *RentalContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RentalContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RentalContract *RentalContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RentalContract.Contract.contract.Transact(opts, method, params...)
}

// TimeCall is a free data retrieval call binding the contract method 0x03261030.
//
// Solidity: function Time_call() view returns(uint256)
func (_RentalContract *RentalContractCaller) TimeCall(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RentalContract.contract.Call(opts, &out, "Time_call")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeCall is a free data retrieval call binding the contract method 0x03261030.
//
// Solidity: function Time_call() view returns(uint256)
func (_RentalContract *RentalContractSession) TimeCall() (*big.Int, error) {
	return _RentalContract.Contract.TimeCall(&_RentalContract.CallOpts)
}

// TimeCall is a free data retrieval call binding the contract method 0x03261030.
//
// Solidity: function Time_call() view returns(uint256)
func (_RentalContract *RentalContractCallerSession) TimeCall() (*big.Int, error) {
	return _RentalContract.Contract.TimeCall(&_RentalContract.CallOpts)
}

// Properties is a free data retrieval call binding the contract method 0x14bfcdd7.
//
// Solidity: function properties(address , uint256 ) view returns(address owner, uint256 id, uint256 perNight, bool available, bool payInFull, uint256 securityDeposit)
func (_RentalContract *RentalContractCaller) Properties(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Owner           common.Address
	Id              *big.Int
	PerNight        *big.Int
	Available       bool
	PayInFull       bool
	SecurityDeposit *big.Int
}, error) {
	var out []interface{}
	err := _RentalContract.contract.Call(opts, &out, "properties", arg0, arg1)

	outstruct := new(struct {
		Owner           common.Address
		Id              *big.Int
		PerNight        *big.Int
		Available       bool
		PayInFull       bool
		SecurityDeposit *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PerNight = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Available = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.PayInFull = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.SecurityDeposit = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Properties is a free data retrieval call binding the contract method 0x14bfcdd7.
//
// Solidity: function properties(address , uint256 ) view returns(address owner, uint256 id, uint256 perNight, bool available, bool payInFull, uint256 securityDeposit)
func (_RentalContract *RentalContractSession) Properties(arg0 common.Address, arg1 *big.Int) (struct {
	Owner           common.Address
	Id              *big.Int
	PerNight        *big.Int
	Available       bool
	PayInFull       bool
	SecurityDeposit *big.Int
}, error) {
	return _RentalContract.Contract.Properties(&_RentalContract.CallOpts, arg0, arg1)
}

// Properties is a free data retrieval call binding the contract method 0x14bfcdd7.
//
// Solidity: function properties(address , uint256 ) view returns(address owner, uint256 id, uint256 perNight, bool available, bool payInFull, uint256 securityDeposit)
func (_RentalContract *RentalContractCallerSession) Properties(arg0 common.Address, arg1 *big.Int) (struct {
	Owner           common.Address
	Id              *big.Int
	PerNight        *big.Int
	Available       bool
	PayInFull       bool
	SecurityDeposit *big.Int
}, error) {
	return _RentalContract.Contract.Properties(&_RentalContract.CallOpts, arg0, arg1)
}

// Rentals is a free data retrieval call binding the contract method 0x090fc58b.
//
// Solidity: function rentals(address , uint256 ) view returns(address owner, uint256 id, address renter, uint256 t1, uint256 t2)
func (_RentalContract *RentalContractCaller) Rentals(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Owner  common.Address
	Id     *big.Int
	Renter common.Address
	T1     *big.Int
	T2     *big.Int
}, error) {
	var out []interface{}
	err := _RentalContract.contract.Call(opts, &out, "rentals", arg0, arg1)

	outstruct := new(struct {
		Owner  common.Address
		Id     *big.Int
		Renter common.Address
		T1     *big.Int
		T2     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Renter = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.T1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.T2 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rentals is a free data retrieval call binding the contract method 0x090fc58b.
//
// Solidity: function rentals(address , uint256 ) view returns(address owner, uint256 id, address renter, uint256 t1, uint256 t2)
func (_RentalContract *RentalContractSession) Rentals(arg0 common.Address, arg1 *big.Int) (struct {
	Owner  common.Address
	Id     *big.Int
	Renter common.Address
	T1     *big.Int
	T2     *big.Int
}, error) {
	return _RentalContract.Contract.Rentals(&_RentalContract.CallOpts, arg0, arg1)
}

// Rentals is a free data retrieval call binding the contract method 0x090fc58b.
//
// Solidity: function rentals(address , uint256 ) view returns(address owner, uint256 id, address renter, uint256 t1, uint256 t2)
func (_RentalContract *RentalContractCallerSession) Rentals(arg0 common.Address, arg1 *big.Int) (struct {
	Owner  common.Address
	Id     *big.Int
	Renter common.Address
	T1     *big.Int
	T2     *big.Int
}, error) {
	return _RentalContract.Contract.Rentals(&_RentalContract.CallOpts, arg0, arg1)
}

// SecurityDeposits is a free data retrieval call binding the contract method 0xf12fdc4e.
//
// Solidity: function securityDeposits(address , uint256 ) view returns(address owner, uint256 id, address renter, uint256 amount, uint256 timestamp, bool dispute)
func (_RentalContract *RentalContractCaller) SecurityDeposits(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Owner     common.Address
	Id        *big.Int
	Renter    common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Dispute   bool
}, error) {
	var out []interface{}
	err := _RentalContract.contract.Call(opts, &out, "securityDeposits", arg0, arg1)

	outstruct := new(struct {
		Owner     common.Address
		Id        *big.Int
		Renter    common.Address
		Amount    *big.Int
		Timestamp *big.Int
		Dispute   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Renter = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Dispute = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// SecurityDeposits is a free data retrieval call binding the contract method 0xf12fdc4e.
//
// Solidity: function securityDeposits(address , uint256 ) view returns(address owner, uint256 id, address renter, uint256 amount, uint256 timestamp, bool dispute)
func (_RentalContract *RentalContractSession) SecurityDeposits(arg0 common.Address, arg1 *big.Int) (struct {
	Owner     common.Address
	Id        *big.Int
	Renter    common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Dispute   bool
}, error) {
	return _RentalContract.Contract.SecurityDeposits(&_RentalContract.CallOpts, arg0, arg1)
}

// SecurityDeposits is a free data retrieval call binding the contract method 0xf12fdc4e.
//
// Solidity: function securityDeposits(address , uint256 ) view returns(address owner, uint256 id, address renter, uint256 amount, uint256 timestamp, bool dispute)
func (_RentalContract *RentalContractCallerSession) SecurityDeposits(arg0 common.Address, arg1 *big.Int) (struct {
	Owner     common.Address
	Id        *big.Int
	Renter    common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Dispute   bool
}, error) {
	return _RentalContract.Contract.SecurityDeposits(&_RentalContract.CallOpts, arg0, arg1)
}

// UpdatePropertyAvailability is a paid mutator transaction binding the contract method 0xf913a4bc.
//
// Solidity: function UpdatePropertyAvailability(uint256 id, bool available) returns()
func (_RentalContract *RentalContractTransactor) UpdatePropertyAvailability(opts *bind.TransactOpts, id *big.Int, available bool) (*types.Transaction, error) {
	return _RentalContract.contract.Transact(opts, "UpdatePropertyAvailability", id, available)
}

// UpdatePropertyAvailability is a paid mutator transaction binding the contract method 0xf913a4bc.
//
// Solidity: function UpdatePropertyAvailability(uint256 id, bool available) returns()
func (_RentalContract *RentalContractSession) UpdatePropertyAvailability(id *big.Int, available bool) (*types.Transaction, error) {
	return _RentalContract.Contract.UpdatePropertyAvailability(&_RentalContract.TransactOpts, id, available)
}

// UpdatePropertyAvailability is a paid mutator transaction binding the contract method 0xf913a4bc.
//
// Solidity: function UpdatePropertyAvailability(uint256 id, bool available) returns()
func (_RentalContract *RentalContractTransactorSession) UpdatePropertyAvailability(id *big.Int, available bool) (*types.Transaction, error) {
	return _RentalContract.Contract.UpdatePropertyAvailability(&_RentalContract.TransactOpts, id, available)
}

// FileDispute is a paid mutator transaction binding the contract method 0xf40d0602.
//
// Solidity: function fileDispute(address owner, uint256 id, bool dispute) returns()
func (_RentalContract *RentalContractTransactor) FileDispute(opts *bind.TransactOpts, owner common.Address, id *big.Int, dispute bool) (*types.Transaction, error) {
	return _RentalContract.contract.Transact(opts, "fileDispute", owner, id, dispute)
}

// FileDispute is a paid mutator transaction binding the contract method 0xf40d0602.
//
// Solidity: function fileDispute(address owner, uint256 id, bool dispute) returns()
func (_RentalContract *RentalContractSession) FileDispute(owner common.Address, id *big.Int, dispute bool) (*types.Transaction, error) {
	return _RentalContract.Contract.FileDispute(&_RentalContract.TransactOpts, owner, id, dispute)
}

// FileDispute is a paid mutator transaction binding the contract method 0xf40d0602.
//
// Solidity: function fileDispute(address owner, uint256 id, bool dispute) returns()
func (_RentalContract *RentalContractTransactorSession) FileDispute(owner common.Address, id *big.Int, dispute bool) (*types.Transaction, error) {
	return _RentalContract.Contract.FileDispute(&_RentalContract.TransactOpts, owner, id, dispute)
}

// ListProperty is a paid mutator transaction binding the contract method 0xa8d659bb.
//
// Solidity: function listProperty(uint256 perNight, bool payInFull, uint256 securityDeposit) returns()
func (_RentalContract *RentalContractTransactor) ListProperty(opts *bind.TransactOpts, perNight *big.Int, payInFull bool, securityDeposit *big.Int) (*types.Transaction, error) {
	return _RentalContract.contract.Transact(opts, "listProperty", perNight, payInFull, securityDeposit)
}

// ListProperty is a paid mutator transaction binding the contract method 0xa8d659bb.
//
// Solidity: function listProperty(uint256 perNight, bool payInFull, uint256 securityDeposit) returns()
func (_RentalContract *RentalContractSession) ListProperty(perNight *big.Int, payInFull bool, securityDeposit *big.Int) (*types.Transaction, error) {
	return _RentalContract.Contract.ListProperty(&_RentalContract.TransactOpts, perNight, payInFull, securityDeposit)
}

// ListProperty is a paid mutator transaction binding the contract method 0xa8d659bb.
//
// Solidity: function listProperty(uint256 perNight, bool payInFull, uint256 securityDeposit) returns()
func (_RentalContract *RentalContractTransactorSession) ListProperty(perNight *big.Int, payInFull bool, securityDeposit *big.Int) (*types.Transaction, error) {
	return _RentalContract.Contract.ListProperty(&_RentalContract.TransactOpts, perNight, payInFull, securityDeposit)
}

// ReleaseDeposit is a paid mutator transaction binding the contract method 0x524dd059.
//
// Solidity: function releaseDeposit(address renter, uint256 id) returns()
func (_RentalContract *RentalContractTransactor) ReleaseDeposit(opts *bind.TransactOpts, renter common.Address, id *big.Int) (*types.Transaction, error) {
	return _RentalContract.contract.Transact(opts, "releaseDeposit", renter, id)
}

// ReleaseDeposit is a paid mutator transaction binding the contract method 0x524dd059.
//
// Solidity: function releaseDeposit(address renter, uint256 id) returns()
func (_RentalContract *RentalContractSession) ReleaseDeposit(renter common.Address, id *big.Int) (*types.Transaction, error) {
	return _RentalContract.Contract.ReleaseDeposit(&_RentalContract.TransactOpts, renter, id)
}

// ReleaseDeposit is a paid mutator transaction binding the contract method 0x524dd059.
//
// Solidity: function releaseDeposit(address renter, uint256 id) returns()
func (_RentalContract *RentalContractTransactorSession) ReleaseDeposit(renter common.Address, id *big.Int) (*types.Transaction, error) {
	return _RentalContract.Contract.ReleaseDeposit(&_RentalContract.TransactOpts, renter, id)
}

// RentProperty is a paid mutator transaction binding the contract method 0x3b88fa21.
//
// Solidity: function rentProperty(address owner, uint256 id, uint256 t1, uint256 t2) payable returns()
func (_RentalContract *RentalContractTransactor) RentProperty(opts *bind.TransactOpts, owner common.Address, id *big.Int, t1 *big.Int, t2 *big.Int) (*types.Transaction, error) {
	return _RentalContract.contract.Transact(opts, "rentProperty", owner, id, t1, t2)
}

// RentProperty is a paid mutator transaction binding the contract method 0x3b88fa21.
//
// Solidity: function rentProperty(address owner, uint256 id, uint256 t1, uint256 t2) payable returns()
func (_RentalContract *RentalContractSession) RentProperty(owner common.Address, id *big.Int, t1 *big.Int, t2 *big.Int) (*types.Transaction, error) {
	return _RentalContract.Contract.RentProperty(&_RentalContract.TransactOpts, owner, id, t1, t2)
}

// RentProperty is a paid mutator transaction binding the contract method 0x3b88fa21.
//
// Solidity: function rentProperty(address owner, uint256 id, uint256 t1, uint256 t2) payable returns()
func (_RentalContract *RentalContractTransactorSession) RentProperty(owner common.Address, id *big.Int, t1 *big.Int, t2 *big.Int) (*types.Transaction, error) {
	return _RentalContract.Contract.RentProperty(&_RentalContract.TransactOpts, owner, id, t1, t2)
}

// UpdateRentalPrice is a paid mutator transaction binding the contract method 0x90abcd98.
//
// Solidity: function updateRentalPrice(uint256 newPrice, uint256 id) returns()
func (_RentalContract *RentalContractTransactor) UpdateRentalPrice(opts *bind.TransactOpts, newPrice *big.Int, id *big.Int) (*types.Transaction, error) {
	return _RentalContract.contract.Transact(opts, "updateRentalPrice", newPrice, id)
}

// UpdateRentalPrice is a paid mutator transaction binding the contract method 0x90abcd98.
//
// Solidity: function updateRentalPrice(uint256 newPrice, uint256 id) returns()
func (_RentalContract *RentalContractSession) UpdateRentalPrice(newPrice *big.Int, id *big.Int) (*types.Transaction, error) {
	return _RentalContract.Contract.UpdateRentalPrice(&_RentalContract.TransactOpts, newPrice, id)
}

// UpdateRentalPrice is a paid mutator transaction binding the contract method 0x90abcd98.
//
// Solidity: function updateRentalPrice(uint256 newPrice, uint256 id) returns()
func (_RentalContract *RentalContractTransactorSession) UpdateRentalPrice(newPrice *big.Int, id *big.Int) (*types.Transaction, error) {
	return _RentalContract.Contract.UpdateRentalPrice(&_RentalContract.TransactOpts, newPrice, id)
}

// RentalContractPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the RentalContract contract.
type RentalContractPriceUpdatedIterator struct {
	Event *RentalContractPriceUpdated // Event containing the contract specifics and raw log

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
func (it *RentalContractPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RentalContractPriceUpdated)
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
		it.Event = new(RentalContractPriceUpdated)
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
func (it *RentalContractPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RentalContractPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RentalContractPriceUpdated represents a PriceUpdated event raised by the RentalContract contract.
type RentalContractPriceUpdated struct {
	Owner    common.Address
	NewPrice *big.Int
	ID       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0xb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c0.
//
// Solidity: event PriceUpdated(address owner, uint256 newPrice, uint256 ID)
func (_RentalContract *RentalContractFilterer) FilterPriceUpdated(opts *bind.FilterOpts) (*RentalContractPriceUpdatedIterator, error) {

	logs, sub, err := _RentalContract.contract.FilterLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return &RentalContractPriceUpdatedIterator{contract: _RentalContract.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0xb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c0.
//
// Solidity: event PriceUpdated(address owner, uint256 newPrice, uint256 ID)
func (_RentalContract *RentalContractFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *RentalContractPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _RentalContract.contract.WatchLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RentalContractPriceUpdated)
				if err := _RentalContract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
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

// ParsePriceUpdated is a log parse operation binding the contract event 0xb556fac599c3c70efb9ab1fa725ecace6c81cc48d1455f886607def065f3e0c0.
//
// Solidity: event PriceUpdated(address owner, uint256 newPrice, uint256 ID)
func (_RentalContract *RentalContractFilterer) ParsePriceUpdated(log types.Log) (*RentalContractPriceUpdated, error) {
	event := new(RentalContractPriceUpdated)
	if err := _RentalContract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RentalContractPropertyListedIterator is returned from FilterPropertyListed and is used to iterate over the raw logs and unpacked data for PropertyListed events raised by the RentalContract contract.
type RentalContractPropertyListedIterator struct {
	Event *RentalContractPropertyListed // Event containing the contract specifics and raw log

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
func (it *RentalContractPropertyListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RentalContractPropertyListed)
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
		it.Event = new(RentalContractPropertyListed)
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
func (it *RentalContractPropertyListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RentalContractPropertyListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RentalContractPropertyListed represents a PropertyListed event raised by the RentalContract contract.
type RentalContractPropertyListed struct {
	Owner common.Address
	ID    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPropertyListed is a free log retrieval operation binding the contract event 0xd45ae8d02a95e46a15b86aa93911a32c7e0bf35515df8de4945f3afb1ba11039.
//
// Solidity: event PropertyListed(address owner, uint256 ID)
func (_RentalContract *RentalContractFilterer) FilterPropertyListed(opts *bind.FilterOpts) (*RentalContractPropertyListedIterator, error) {

	logs, sub, err := _RentalContract.contract.FilterLogs(opts, "PropertyListed")
	if err != nil {
		return nil, err
	}
	return &RentalContractPropertyListedIterator{contract: _RentalContract.contract, event: "PropertyListed", logs: logs, sub: sub}, nil
}

// WatchPropertyListed is a free log subscription operation binding the contract event 0xd45ae8d02a95e46a15b86aa93911a32c7e0bf35515df8de4945f3afb1ba11039.
//
// Solidity: event PropertyListed(address owner, uint256 ID)
func (_RentalContract *RentalContractFilterer) WatchPropertyListed(opts *bind.WatchOpts, sink chan<- *RentalContractPropertyListed) (event.Subscription, error) {

	logs, sub, err := _RentalContract.contract.WatchLogs(opts, "PropertyListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RentalContractPropertyListed)
				if err := _RentalContract.contract.UnpackLog(event, "PropertyListed", log); err != nil {
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

// ParsePropertyListed is a log parse operation binding the contract event 0xd45ae8d02a95e46a15b86aa93911a32c7e0bf35515df8de4945f3afb1ba11039.
//
// Solidity: event PropertyListed(address owner, uint256 ID)
func (_RentalContract *RentalContractFilterer) ParsePropertyListed(log types.Log) (*RentalContractPropertyListed, error) {
	event := new(RentalContractPropertyListed)
	if err := _RentalContract.contract.UnpackLog(event, "PropertyListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RentalContractPropertyRentedIterator is returned from FilterPropertyRented and is used to iterate over the raw logs and unpacked data for PropertyRented events raised by the RentalContract contract.
type RentalContractPropertyRentedIterator struct {
	Event *RentalContractPropertyRented // Event containing the contract specifics and raw log

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
func (it *RentalContractPropertyRentedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RentalContractPropertyRented)
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
		it.Event = new(RentalContractPropertyRented)
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
func (it *RentalContractPropertyRentedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RentalContractPropertyRentedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RentalContractPropertyRented represents a PropertyRented event raised by the RentalContract contract.
type RentalContractPropertyRented struct {
	Owner    common.Address
	ID       *big.Int
	Renter   common.Address
	Checkin  *big.Int
	Checkout *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPropertyRented is a free log retrieval operation binding the contract event 0xda21554a181142009a8ac60c8d97feeae9dcd4433f8d31bbe4ee4080692a4f3b.
//
// Solidity: event PropertyRented(address owner, uint256 ID, address renter, uint256 checkin, uint256 checkout)
func (_RentalContract *RentalContractFilterer) FilterPropertyRented(opts *bind.FilterOpts) (*RentalContractPropertyRentedIterator, error) {

	logs, sub, err := _RentalContract.contract.FilterLogs(opts, "PropertyRented")
	if err != nil {
		return nil, err
	}
	return &RentalContractPropertyRentedIterator{contract: _RentalContract.contract, event: "PropertyRented", logs: logs, sub: sub}, nil
}

// WatchPropertyRented is a free log subscription operation binding the contract event 0xda21554a181142009a8ac60c8d97feeae9dcd4433f8d31bbe4ee4080692a4f3b.
//
// Solidity: event PropertyRented(address owner, uint256 ID, address renter, uint256 checkin, uint256 checkout)
func (_RentalContract *RentalContractFilterer) WatchPropertyRented(opts *bind.WatchOpts, sink chan<- *RentalContractPropertyRented) (event.Subscription, error) {

	logs, sub, err := _RentalContract.contract.WatchLogs(opts, "PropertyRented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RentalContractPropertyRented)
				if err := _RentalContract.contract.UnpackLog(event, "PropertyRented", log); err != nil {
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

// ParsePropertyRented is a log parse operation binding the contract event 0xda21554a181142009a8ac60c8d97feeae9dcd4433f8d31bbe4ee4080692a4f3b.
//
// Solidity: event PropertyRented(address owner, uint256 ID, address renter, uint256 checkin, uint256 checkout)
func (_RentalContract *RentalContractFilterer) ParsePropertyRented(log types.Log) (*RentalContractPropertyRented, error) {
	event := new(RentalContractPropertyRented)
	if err := _RentalContract.contract.UnpackLog(event, "PropertyRented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
