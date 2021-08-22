// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ChangeMetaData contains all meta data concerning the Change contract.
var ChangeMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"tr_id\",\"type\":\"uint256\"},{\"name\":\"pw\",\"type\":\"string\"}],\"name\":\"get_my_transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"adr_to\",\"type\":\"address\"},{\"name\":\"template\",\"type\":\"string\"},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"pwhash\",\"type\":\"bytes32\"}],\"name\":\"use_template\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"offer_amount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"adr_to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"category\",\"type\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"pwhash\",\"type\":\"bytes32\"}],\"name\":\"create_transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get_users_amount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"create_category\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offer_id\",\"type\":\"uint256\"}],\"name\":\"vote_against\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"boosts\",\"outputs\":[{\"name\":\"to_boost\",\"type\":\"address\"},{\"name\":\"cons\",\"type\":\"address\"},{\"name\":\"finished\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get_userlist\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offer_id\",\"type\":\"uint256\"}],\"name\":\"check_offer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"offer_id\",\"type\":\"uint256\"}],\"name\":\"get_proc\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"category\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"create_template\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user_adr\",\"type\":\"address\"}],\"name\":\"offer_user_to_boost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"get_template_by_name\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transfers\",\"outputs\":[{\"name\":\"adr_from\",\"type\":\"address\"},{\"name\":\"adr_to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"category\",\"type\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\"},{\"name\":\"pwhash\",\"type\":\"bytes32\"},{\"name\":\"time\",\"type\":\"uint256\"},{\"name\":\"finished\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"offer_id\",\"type\":\"uint256\"}],\"name\":\"vote_for\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"},{\"name\":\"pwhash\",\"type\":\"bytes32\"},{\"name\":\"role\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"template_names\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transfer_amount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"categories\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin_amount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userlist\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pwhash\",\"type\":\"bytes32\"}],\"name\":\"create_user\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tr_id\",\"type\":\"uint256\"}],\"name\":\"stop_my_transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get_template_names\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get_categories\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
}

// ChangeABI is the input ABI used to generate the binding from.
// Deprecated: Use ChangeMetaData.ABI instead.
var ChangeABI = ChangeMetaData.ABI

// Change is an auto generated Go binding around an Ethereum contract.
type Change struct {
	ChangeCaller     // Read-only binding to the contract
	ChangeTransactor // Write-only binding to the contract
	ChangeFilterer   // Log filterer for contract events
}

// ChangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChangeSession struct {
	Contract     *Change           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChangeCallerSession struct {
	Contract *ChangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ChangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChangeTransactorSession struct {
	Contract     *ChangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChangeRaw struct {
	Contract *Change // Generic contract binding to access the raw methods on
}

// ChangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChangeCallerRaw struct {
	Contract *ChangeCaller // Generic read-only contract binding to access the raw methods on
}

// ChangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChangeTransactorRaw struct {
	Contract *ChangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChange creates a new instance of Change, bound to a specific deployed contract.
func NewChange(address common.Address, backend bind.ContractBackend) (*Change, error) {
	contract, err := bindChange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Change{ChangeCaller: ChangeCaller{contract: contract}, ChangeTransactor: ChangeTransactor{contract: contract}, ChangeFilterer: ChangeFilterer{contract: contract}}, nil
}

// NewChangeCaller creates a new read-only instance of Change, bound to a specific deployed contract.
func NewChangeCaller(address common.Address, caller bind.ContractCaller) (*ChangeCaller, error) {
	contract, err := bindChange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChangeCaller{contract: contract}, nil
}

// NewChangeTransactor creates a new write-only instance of Change, bound to a specific deployed contract.
func NewChangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChangeTransactor, error) {
	contract, err := bindChange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChangeTransactor{contract: contract}, nil
}

// NewChangeFilterer creates a new log filterer instance of Change, bound to a specific deployed contract.
func NewChangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChangeFilterer, error) {
	contract, err := bindChange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChangeFilterer{contract: contract}, nil
}

// bindChange binds a generic wrapper to an already deployed contract.
func bindChange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Change *ChangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Change.Contract.ChangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Change *ChangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Change.Contract.ChangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Change *ChangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Change.Contract.ChangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Change *ChangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Change.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Change *ChangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Change.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Change *ChangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Change.Contract.contract.Transact(opts, method, params...)
}

// AdminAmount is a free data retrieval call binding the contract method 0xcce7f919.
//
// Solidity: function admin_amount() view returns(uint256)
func (_Change *ChangeCaller) AdminAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "admin_amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminAmount is a free data retrieval call binding the contract method 0xcce7f919.
//
// Solidity: function admin_amount() view returns(uint256)
func (_Change *ChangeSession) AdminAmount() (*big.Int, error) {
	return _Change.Contract.AdminAmount(&_Change.CallOpts)
}

// AdminAmount is a free data retrieval call binding the contract method 0xcce7f919.
//
// Solidity: function admin_amount() view returns(uint256)
func (_Change *ChangeCallerSession) AdminAmount() (*big.Int, error) {
	return _Change.Contract.AdminAmount(&_Change.CallOpts)
}

// Boosts is a free data retrieval call binding the contract method 0x4afd82e7.
//
// Solidity: function boosts(uint256 ) view returns(address to_boost, address cons, bool finished)
func (_Change *ChangeCaller) Boosts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ToBoost  common.Address
	Cons     common.Address
	Finished bool
}, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "boosts", arg0)

	outstruct := new(struct {
		ToBoost  common.Address
		Cons     common.Address
		Finished bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ToBoost = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Cons = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Finished = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Boosts is a free data retrieval call binding the contract method 0x4afd82e7.
//
// Solidity: function boosts(uint256 ) view returns(address to_boost, address cons, bool finished)
func (_Change *ChangeSession) Boosts(arg0 *big.Int) (struct {
	ToBoost  common.Address
	Cons     common.Address
	Finished bool
}, error) {
	return _Change.Contract.Boosts(&_Change.CallOpts, arg0)
}

// Boosts is a free data retrieval call binding the contract method 0x4afd82e7.
//
// Solidity: function boosts(uint256 ) view returns(address to_boost, address cons, bool finished)
func (_Change *ChangeCallerSession) Boosts(arg0 *big.Int) (struct {
	ToBoost  common.Address
	Cons     common.Address
	Finished bool
}, error) {
	return _Change.Contract.Boosts(&_Change.CallOpts, arg0)
}

// Categories is a free data retrieval call binding the contract method 0xc6cdbe5e.
//
// Solidity: function categories(uint256 ) view returns(string)
func (_Change *ChangeCaller) Categories(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "categories", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Categories is a free data retrieval call binding the contract method 0xc6cdbe5e.
//
// Solidity: function categories(uint256 ) view returns(string)
func (_Change *ChangeSession) Categories(arg0 *big.Int) (string, error) {
	return _Change.Contract.Categories(&_Change.CallOpts, arg0)
}

// Categories is a free data retrieval call binding the contract method 0xc6cdbe5e.
//
// Solidity: function categories(uint256 ) view returns(string)
func (_Change *ChangeCallerSession) Categories(arg0 *big.Int) (string, error) {
	return _Change.Contract.Categories(&_Change.CallOpts, arg0)
}

// GetCategories is a free data retrieval call binding the contract method 0xfd4ec427.
//
// Solidity: function get_categories() view returns(string[])
func (_Change *ChangeCaller) GetCategories(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "get_categories")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetCategories is a free data retrieval call binding the contract method 0xfd4ec427.
//
// Solidity: function get_categories() view returns(string[])
func (_Change *ChangeSession) GetCategories() ([]string, error) {
	return _Change.Contract.GetCategories(&_Change.CallOpts)
}

// GetCategories is a free data retrieval call binding the contract method 0xfd4ec427.
//
// Solidity: function get_categories() view returns(string[])
func (_Change *ChangeCallerSession) GetCategories() ([]string, error) {
	return _Change.Contract.GetCategories(&_Change.CallOpts)
}

// GetProc is a free data retrieval call binding the contract method 0x71a28423.
//
// Solidity: function get_proc(uint256 offer_id) view returns(address[])
func (_Change *ChangeCaller) GetProc(opts *bind.CallOpts, offer_id *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "get_proc", offer_id)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetProc is a free data retrieval call binding the contract method 0x71a28423.
//
// Solidity: function get_proc(uint256 offer_id) view returns(address[])
func (_Change *ChangeSession) GetProc(offer_id *big.Int) ([]common.Address, error) {
	return _Change.Contract.GetProc(&_Change.CallOpts, offer_id)
}

// GetProc is a free data retrieval call binding the contract method 0x71a28423.
//
// Solidity: function get_proc(uint256 offer_id) view returns(address[])
func (_Change *ChangeCallerSession) GetProc(offer_id *big.Int) ([]common.Address, error) {
	return _Change.Contract.GetProc(&_Change.CallOpts, offer_id)
}

// GetTemplateByName is a free data retrieval call binding the contract method 0x82d26d90.
//
// Solidity: function get_template_by_name(string name) view returns(uint256, uint256)
func (_Change *ChangeCaller) GetTemplateByName(opts *bind.CallOpts, name string) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "get_template_by_name", name)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTemplateByName is a free data retrieval call binding the contract method 0x82d26d90.
//
// Solidity: function get_template_by_name(string name) view returns(uint256, uint256)
func (_Change *ChangeSession) GetTemplateByName(name string) (*big.Int, *big.Int, error) {
	return _Change.Contract.GetTemplateByName(&_Change.CallOpts, name)
}

// GetTemplateByName is a free data retrieval call binding the contract method 0x82d26d90.
//
// Solidity: function get_template_by_name(string name) view returns(uint256, uint256)
func (_Change *ChangeCallerSession) GetTemplateByName(name string) (*big.Int, *big.Int, error) {
	return _Change.Contract.GetTemplateByName(&_Change.CallOpts, name)
}

// GetTemplateNames is a free data retrieval call binding the contract method 0xea913927.
//
// Solidity: function get_template_names() view returns(string[])
func (_Change *ChangeCaller) GetTemplateNames(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "get_template_names")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetTemplateNames is a free data retrieval call binding the contract method 0xea913927.
//
// Solidity: function get_template_names() view returns(string[])
func (_Change *ChangeSession) GetTemplateNames() ([]string, error) {
	return _Change.Contract.GetTemplateNames(&_Change.CallOpts)
}

// GetTemplateNames is a free data retrieval call binding the contract method 0xea913927.
//
// Solidity: function get_template_names() view returns(string[])
func (_Change *ChangeCallerSession) GetTemplateNames() ([]string, error) {
	return _Change.Contract.GetTemplateNames(&_Change.CallOpts)
}

// GetUserlist is a free data retrieval call binding the contract method 0x5f44a335.
//
// Solidity: function get_userlist() view returns(address[])
func (_Change *ChangeCaller) GetUserlist(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "get_userlist")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUserlist is a free data retrieval call binding the contract method 0x5f44a335.
//
// Solidity: function get_userlist() view returns(address[])
func (_Change *ChangeSession) GetUserlist() ([]common.Address, error) {
	return _Change.Contract.GetUserlist(&_Change.CallOpts)
}

// GetUserlist is a free data retrieval call binding the contract method 0x5f44a335.
//
// Solidity: function get_userlist() view returns(address[])
func (_Change *ChangeCallerSession) GetUserlist() ([]common.Address, error) {
	return _Change.Contract.GetUserlist(&_Change.CallOpts)
}

// GetUsersAmount is a free data retrieval call binding the contract method 0x2b7f9649.
//
// Solidity: function get_users_amount() view returns(uint256)
func (_Change *ChangeCaller) GetUsersAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "get_users_amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsersAmount is a free data retrieval call binding the contract method 0x2b7f9649.
//
// Solidity: function get_users_amount() view returns(uint256)
func (_Change *ChangeSession) GetUsersAmount() (*big.Int, error) {
	return _Change.Contract.GetUsersAmount(&_Change.CallOpts)
}

// GetUsersAmount is a free data retrieval call binding the contract method 0x2b7f9649.
//
// Solidity: function get_users_amount() view returns(uint256)
func (_Change *ChangeCallerSession) GetUsersAmount() (*big.Int, error) {
	return _Change.Contract.GetUsersAmount(&_Change.CallOpts)
}

// OfferAmount is a free data retrieval call binding the contract method 0x12a31b00.
//
// Solidity: function offer_amount() view returns(uint256)
func (_Change *ChangeCaller) OfferAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "offer_amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OfferAmount is a free data retrieval call binding the contract method 0x12a31b00.
//
// Solidity: function offer_amount() view returns(uint256)
func (_Change *ChangeSession) OfferAmount() (*big.Int, error) {
	return _Change.Contract.OfferAmount(&_Change.CallOpts)
}

// OfferAmount is a free data retrieval call binding the contract method 0x12a31b00.
//
// Solidity: function offer_amount() view returns(uint256)
func (_Change *ChangeCallerSession) OfferAmount() (*big.Int, error) {
	return _Change.Contract.OfferAmount(&_Change.CallOpts)
}

// TemplateNames is a free data retrieval call binding the contract method 0xaef3d2cf.
//
// Solidity: function template_names(uint256 ) view returns(string)
func (_Change *ChangeCaller) TemplateNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "template_names", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TemplateNames is a free data retrieval call binding the contract method 0xaef3d2cf.
//
// Solidity: function template_names(uint256 ) view returns(string)
func (_Change *ChangeSession) TemplateNames(arg0 *big.Int) (string, error) {
	return _Change.Contract.TemplateNames(&_Change.CallOpts, arg0)
}

// TemplateNames is a free data retrieval call binding the contract method 0xaef3d2cf.
//
// Solidity: function template_names(uint256 ) view returns(string)
func (_Change *ChangeCallerSession) TemplateNames(arg0 *big.Int) (string, error) {
	return _Change.Contract.TemplateNames(&_Change.CallOpts, arg0)
}

// TransferAmount is a free data retrieval call binding the contract method 0xc61ffd29.
//
// Solidity: function transfer_amount() view returns(uint256)
func (_Change *ChangeCaller) TransferAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "transfer_amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferAmount is a free data retrieval call binding the contract method 0xc61ffd29.
//
// Solidity: function transfer_amount() view returns(uint256)
func (_Change *ChangeSession) TransferAmount() (*big.Int, error) {
	return _Change.Contract.TransferAmount(&_Change.CallOpts)
}

// TransferAmount is a free data retrieval call binding the contract method 0xc61ffd29.
//
// Solidity: function transfer_amount() view returns(uint256)
func (_Change *ChangeCallerSession) TransferAmount() (*big.Int, error) {
	return _Change.Contract.TransferAmount(&_Change.CallOpts)
}

// Transfers is a free data retrieval call binding the contract method 0x9377d711.
//
// Solidity: function transfers(uint256 ) view returns(address adr_from, address adr_to, uint256 value, uint256 category, string description, bytes32 pwhash, uint256 time, bool finished)
func (_Change *ChangeCaller) Transfers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	AdrFrom     common.Address
	AdrTo       common.Address
	Value       *big.Int
	Category    *big.Int
	Description string
	Pwhash      [32]byte
	Time        *big.Int
	Finished    bool
}, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "transfers", arg0)

	outstruct := new(struct {
		AdrFrom     common.Address
		AdrTo       common.Address
		Value       *big.Int
		Category    *big.Int
		Description string
		Pwhash      [32]byte
		Time        *big.Int
		Finished    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AdrFrom = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AdrTo = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Category = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Description = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Pwhash = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Time = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Finished = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Transfers is a free data retrieval call binding the contract method 0x9377d711.
//
// Solidity: function transfers(uint256 ) view returns(address adr_from, address adr_to, uint256 value, uint256 category, string description, bytes32 pwhash, uint256 time, bool finished)
func (_Change *ChangeSession) Transfers(arg0 *big.Int) (struct {
	AdrFrom     common.Address
	AdrTo       common.Address
	Value       *big.Int
	Category    *big.Int
	Description string
	Pwhash      [32]byte
	Time        *big.Int
	Finished    bool
}, error) {
	return _Change.Contract.Transfers(&_Change.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x9377d711.
//
// Solidity: function transfers(uint256 ) view returns(address adr_from, address adr_to, uint256 value, uint256 category, string description, bytes32 pwhash, uint256 time, bool finished)
func (_Change *ChangeCallerSession) Transfers(arg0 *big.Int) (struct {
	AdrFrom     common.Address
	AdrTo       common.Address
	Value       *big.Int
	Category    *big.Int
	Description string
	Pwhash      [32]byte
	Time        *big.Int
	Finished    bool
}, error) {
	return _Change.Contract.Transfers(&_Change.CallOpts, arg0)
}

// Userlist is a free data retrieval call binding the contract method 0xd8e136e2.
//
// Solidity: function userlist(uint256 ) view returns(address)
func (_Change *ChangeCaller) Userlist(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "userlist", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Userlist is a free data retrieval call binding the contract method 0xd8e136e2.
//
// Solidity: function userlist(uint256 ) view returns(address)
func (_Change *ChangeSession) Userlist(arg0 *big.Int) (common.Address, error) {
	return _Change.Contract.Userlist(&_Change.CallOpts, arg0)
}

// Userlist is a free data retrieval call binding the contract method 0xd8e136e2.
//
// Solidity: function userlist(uint256 ) view returns(address)
func (_Change *ChangeCallerSession) Userlist(arg0 *big.Int) (common.Address, error) {
	return _Change.Contract.Userlist(&_Change.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(uint256 balance, bytes32 pwhash, bool role)
func (_Change *ChangeCaller) Users(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance *big.Int
	Pwhash  [32]byte
	Role    bool
}, error) {
	var out []interface{}
	err := _Change.contract.Call(opts, &out, "users", arg0)

	outstruct := new(struct {
		Balance *big.Int
		Pwhash  [32]byte
		Role    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Pwhash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Role = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(uint256 balance, bytes32 pwhash, bool role)
func (_Change *ChangeSession) Users(arg0 common.Address) (struct {
	Balance *big.Int
	Pwhash  [32]byte
	Role    bool
}, error) {
	return _Change.Contract.Users(&_Change.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(uint256 balance, bytes32 pwhash, bool role)
func (_Change *ChangeCallerSession) Users(arg0 common.Address) (struct {
	Balance *big.Int
	Pwhash  [32]byte
	Role    bool
}, error) {
	return _Change.Contract.Users(&_Change.CallOpts, arg0)
}

// CheckOffer is a paid mutator transaction binding the contract method 0x626933fc.
//
// Solidity: function check_offer(uint256 offer_id) returns()
func (_Change *ChangeTransactor) CheckOffer(opts *bind.TransactOpts, offer_id *big.Int) (*types.Transaction, error) {
	return _Change.contract.Transact(opts, "check_offer", offer_id)
}

// CheckOffer is a paid mutator transaction binding the contract method 0x626933fc.
//
// Solidity: function check_offer(uint256 offer_id) returns()
func (_Change *ChangeSession) CheckOffer(offer_id *big.Int) (*types.Transaction, error) {
	return _Change.Contract.CheckOffer(&_Change.TransactOpts, offer_id)
}

// CheckOffer is a paid mutator transaction binding the contract method 0x626933fc.
//
// Solidity: function check_offer(uint256 offer_id) returns()
func (_Change *ChangeTransactorSession) CheckOffer(offer_id *big.Int) (*types.Transaction, error) {
	return _Change.Contract.CheckOffer(&_Change.TransactOpts, offer_id)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x3effbb6a.
//
// Solidity: function create_category(string name) returns()
func (_Change *ChangeTransactor) CreateCategory(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Change.contract.Transact(opts, "create_category", name)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x3effbb6a.
//
// Solidity: function create_category(string name) returns()
func (_Change *ChangeSession) CreateCategory(name string) (*types.Transaction, error) {
	return _Change.Contract.CreateCategory(&_Change.TransactOpts, name)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x3effbb6a.
//
// Solidity: function create_category(string name) returns()
func (_Change *ChangeTransactorSession) CreateCategory(name string) (*types.Transaction, error) {
	return _Change.Contract.CreateCategory(&_Change.TransactOpts, name)
}

// CreateTemplate is a paid mutator transaction binding the contract method 0x74500045.
//
// Solidity: function create_template(string name, uint256 category, uint256 value) returns()
func (_Change *ChangeTransactor) CreateTemplate(opts *bind.TransactOpts, name string, category *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Change.contract.Transact(opts, "create_template", name, category, value)
}

// CreateTemplate is a paid mutator transaction binding the contract method 0x74500045.
//
// Solidity: function create_template(string name, uint256 category, uint256 value) returns()
func (_Change *ChangeSession) CreateTemplate(name string, category *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Change.Contract.CreateTemplate(&_Change.TransactOpts, name, category, value)
}

// CreateTemplate is a paid mutator transaction binding the contract method 0x74500045.
//
// Solidity: function create_template(string name, uint256 category, uint256 value) returns()
func (_Change *ChangeTransactorSession) CreateTemplate(name string, category *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Change.Contract.CreateTemplate(&_Change.TransactOpts, name, category, value)
}

// CreateTransfer is a paid mutator transaction binding the contract method 0x223afbe0.
//
// Solidity: function create_transfer(address adr_to, uint256 value, uint256 category, string description, bytes32 pwhash) returns()
func (_Change *ChangeTransactor) CreateTransfer(opts *bind.TransactOpts, adr_to common.Address, value *big.Int, category *big.Int, description string, pwhash [32]byte) (*types.Transaction, error) {
	return _Change.contract.Transact(opts, "create_transfer", adr_to, value, category, description, pwhash)
}

// CreateTransfer is a paid mutator transaction binding the contract method 0x223afbe0.
//
// Solidity: function create_transfer(address adr_to, uint256 value, uint256 category, string description, bytes32 pwhash) returns()
func (_Change *ChangeSession) CreateTransfer(adr_to common.Address, value *big.Int, category *big.Int, description string, pwhash [32]byte) (*types.Transaction, error) {
	return _Change.Contract.CreateTransfer(&_Change.TransactOpts, adr_to, value, category, description, pwhash)
}

// CreateTransfer is a paid mutator transaction binding the contract method 0x223afbe0.
//
// Solidity: function create_transfer(address adr_to, uint256 value, uint256 category, string description, bytes32 pwhash) returns()
func (_Change *ChangeTransactorSession) CreateTransfer(adr_to common.Address, value *big.Int, category *big.Int, description string, pwhash [32]byte) (*types.Transaction, error) {
	return _Change.Contract.CreateTransfer(&_Change.TransactOpts, adr_to, value, category, description, pwhash)
}

// CreateUser is a paid mutator transaction binding the contract method 0xd945d59f.
//
// Solidity: function create_user(bytes32 pwhash) returns()
func (_Change *ChangeTransactor) CreateUser(opts *bind.TransactOpts, pwhash [32]byte) (*types.Transaction, error) {
	return _Change.contract.Transact(opts, "create_user", pwhash)
}

// CreateUser is a paid mutator transaction binding the contract method 0xd945d59f.
//
// Solidity: function create_user(bytes32 pwhash) returns()
func (_Change *ChangeSession) CreateUser(pwhash [32]byte) (*types.Transaction, error) {
	return _Change.Contract.CreateUser(&_Change.TransactOpts, pwhash)
}

// CreateUser is a paid mutator transaction binding the contract method 0xd945d59f.
//
// Solidity: function create_user(bytes32 pwhash) returns()
func (_Change *ChangeTransactorSession) CreateUser(pwhash [32]byte) (*types.Transaction, error) {
	return _Change.Contract.CreateUser(&_Change.TransactOpts, pwhash)
}

// GetMyTransfer is a paid mutator transaction binding the contract method 0x009076d1.
//
// Solidity: function get_my_transfer(uint256 tr_id, string pw) returns()
func (_Change *ChangeTransactor) GetMyTransfer(opts *bind.TransactOpts, tr_id *big.Int, pw string) (*types.Transaction, error) {
	return _Change.contract.Transact(opts, "get_my_transfer", tr_id, pw)
}

// GetMyTransfer is a paid mutator transaction binding the contract method 0x009076d1.
//
// Solidity: function get_my_transfer(uint256 tr_id, string pw) returns()
func (_Change *ChangeSession) GetMyTransfer(tr_id *big.Int, pw string) (*types.Transaction, error) {
	return _Change.Contract.GetMyTransfer(&_Change.TransactOpts, tr_id, pw)
}

// GetMyTransfer is a paid mutator transaction binding the contract method 0x009076d1.
//
// Solidity: function get_my_transfer(uint256 tr_id, string pw) returns()
func (_Change *ChangeTransactorSession) GetMyTransfer(tr_id *big.Int, pw string) (*types.Transaction, error) {
	return _Change.Contract.GetMyTransfer(&_Change.TransactOpts, tr_id, pw)
}

// OfferUserToBoost is a paid mutator transaction binding the contract method 0x782bb01c.
//
// Solidity: function offer_user_to_boost(address user_adr) returns()
func (_Change *ChangeTransactor) OfferUserToBoost(opts *bind.TransactOpts, user_adr common.Address) (*types.Transaction, error) {
	return _Change.contract.Transact(opts, "offer_user_to_boost", user_adr)
}

// OfferUserToBoost is a paid mutator transaction binding the contract method 0x782bb01c.
//
// Solidity: function offer_user_to_boost(address user_adr) returns()
func (_Change *ChangeSession) OfferUserToBoost(user_adr common.Address) (*types.Transaction, error) {
	return _Change.Contract.OfferUserToBoost(&_Change.TransactOpts, user_adr)
}

// OfferUserToBoost is a paid mutator transaction binding the contract method 0x782bb01c.
//
// Solidity: function offer_user_to_boost(address user_adr) returns()
func (_Change *ChangeTransactorSession) OfferUserToBoost(user_adr common.Address) (*types.Transaction, error) {
	return _Change.Contract.OfferUserToBoost(&_Change.TransactOpts, user_adr)
}

// StopMyTransfer is a paid mutator transaction binding the contract method 0xe38eeb71.
//
// Solidity: function stop_my_transfer(uint256 tr_id) returns()
func (_Change *ChangeTransactor) StopMyTransfer(opts *bind.TransactOpts, tr_id *big.Int) (*types.Transaction, error) {
	return _Change.contract.Transact(opts, "stop_my_transfer", tr_id)
}

// StopMyTransfer is a paid mutator transaction binding the contract method 0xe38eeb71.
//
// Solidity: function stop_my_transfer(uint256 tr_id) returns()
func (_Change *ChangeSession) StopMyTransfer(tr_id *big.Int) (*types.Transaction, error) {
	return _Change.Contract.StopMyTransfer(&_Change.TransactOpts, tr_id)
}

// StopMyTransfer is a paid mutator transaction binding the contract method 0xe38eeb71.
//
// Solidity: function stop_my_transfer(uint256 tr_id) returns()
func (_Change *ChangeTransactorSession) StopMyTransfer(tr_id *big.Int) (*types.Transaction, error) {
	return _Change.Contract.StopMyTransfer(&_Change.TransactOpts, tr_id)
}

// UseTemplate is a paid mutator transaction binding the contract method 0x031a0ba8.
//
// Solidity: function use_template(address adr_to, string template, string description, bytes32 pwhash) returns()
func (_Change *ChangeTransactor) UseTemplate(opts *bind.TransactOpts, adr_to common.Address, template string, description string, pwhash [32]byte) (*types.Transaction, error) {
	return _Change.contract.Transact(opts, "use_template", adr_to, template, description, pwhash)
}

// UseTemplate is a paid mutator transaction binding the contract method 0x031a0ba8.
//
// Solidity: function use_template(address adr_to, string template, string description, bytes32 pwhash) returns()
func (_Change *ChangeSession) UseTemplate(adr_to common.Address, template string, description string, pwhash [32]byte) (*types.Transaction, error) {
	return _Change.Contract.UseTemplate(&_Change.TransactOpts, adr_to, template, description, pwhash)
}

// UseTemplate is a paid mutator transaction binding the contract method 0x031a0ba8.
//
// Solidity: function use_template(address adr_to, string template, string description, bytes32 pwhash) returns()
func (_Change *ChangeTransactorSession) UseTemplate(adr_to common.Address, template string, description string, pwhash [32]byte) (*types.Transaction, error) {
	return _Change.Contract.UseTemplate(&_Change.TransactOpts, adr_to, template, description, pwhash)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x4acadfbb.
//
// Solidity: function vote_against(uint256 offer_id) returns()
func (_Change *ChangeTransactor) VoteAgainst(opts *bind.TransactOpts, offer_id *big.Int) (*types.Transaction, error) {
	return _Change.contract.Transact(opts, "vote_against", offer_id)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x4acadfbb.
//
// Solidity: function vote_against(uint256 offer_id) returns()
func (_Change *ChangeSession) VoteAgainst(offer_id *big.Int) (*types.Transaction, error) {
	return _Change.Contract.VoteAgainst(&_Change.TransactOpts, offer_id)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x4acadfbb.
//
// Solidity: function vote_against(uint256 offer_id) returns()
func (_Change *ChangeTransactorSession) VoteAgainst(offer_id *big.Int) (*types.Transaction, error) {
	return _Change.Contract.VoteAgainst(&_Change.TransactOpts, offer_id)
}

// VoteFor is a paid mutator transaction binding the contract method 0xa31cb00f.
//
// Solidity: function vote_for(uint256 offer_id) returns()
func (_Change *ChangeTransactor) VoteFor(opts *bind.TransactOpts, offer_id *big.Int) (*types.Transaction, error) {
	return _Change.contract.Transact(opts, "vote_for", offer_id)
}

// VoteFor is a paid mutator transaction binding the contract method 0xa31cb00f.
//
// Solidity: function vote_for(uint256 offer_id) returns()
func (_Change *ChangeSession) VoteFor(offer_id *big.Int) (*types.Transaction, error) {
	return _Change.Contract.VoteFor(&_Change.TransactOpts, offer_id)
}

// VoteFor is a paid mutator transaction binding the contract method 0xa31cb00f.
//
// Solidity: function vote_for(uint256 offer_id) returns()
func (_Change *ChangeTransactorSession) VoteFor(offer_id *big.Int) (*types.Transaction, error) {
	return _Change.Contract.VoteFor(&_Change.TransactOpts, offer_id)
}
