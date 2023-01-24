// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// bdnsRegistryDomain is an auto generated low-level Go binding around an user-defined struct.
type bdnsRegistryDomain struct {
	IpAddress  string
	DeviceType string
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ip_address\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"seclevel\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"device_type\",\"type\":\"string\"}],\"name\":\"createRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"seclevel\",\"type\":\"int256\"}],\"name\":\"getRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"ip_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"device_type\",\"type\":\"string\"}],\"internalType\":\"structbdnsRegistry.Domain[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60806040526000805534801561001457600080fd5b5061097c806100246000396000f3fe608060405234801561001057600080fd5b506004361061004b5760003560e01c80625fcd5e146100505780630d8e6e2c14610079578063f4b518bc1461008a578063fa5441611461009f575b600080fd5b61006361005e36600461070d565b6100c8565b604051610070919061081a565b60405180910390f35b60015b604051908152602001610070565b61009d610098366004610669565b61026d565b005b61007c6100ad366004610640565b6001600160a01b031660009081526002602052604090205490565b60606005836040516100da91906107fe565b9081526040805191829003602090810183206000868152908252828120805480840286018401909452838552929184015b82821015610261578382906000526020600020906002020160405180604001604052908160008201805461013e906108f5565b80601f016020809104026020016040519081016040528092919081815260200182805461016a906108f5565b80156101b75780601f1061018c576101008083540402835291602001916101b7565b820191906000526020600020905b81548152906001019060200180831161019a57829003601f168201915b505050505081526020016001820180546101d0906108f5565b80601f01602080910402602001604051908101604052809291908181526020018280546101fc906108f5565b80156102495780601f1061021e57610100808354040283529160200191610249565b820191906000526020600020905b81548152906001019060200180831161022c57829003601f168201915b5050505050815250508152602001906001019061010b565b50505050905092915050565b61027a600054600161039a565b60008181559081526004602052604090206102969088886104ea565b50600587876040516102a99291906107ee565b90815260408051918290036020908101832060008781529082528290206060601f890183900490920284018201835291830187815291929182918990899081908501838280828437600092019190915250505090825250604080516020601f870181900481028201810190925285815291810191908690869081908401838280828437600092018290525093909452505083546001810185559381526020908190208351805194956002029091019361036993508492919091019061056e565b506020828101518051610382926001850192019061056e565b50505061039133600054610404565b50505050505050565b6000806103a7838561089f565b9050838110156103fd5760405162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015260640160405180910390fd5b9392505050565b6001600160a01b03821661041757600080fd5b610421828261045d565b60405181906001600160a01b038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6000818152600160205260409020546001600160a01b03161561047f57600080fd5b600081815260016020818152604080842080546001600160a01b0319166001600160a01b0388169081179091558452600282528084208590556003909152909120546104ca9161039a565b6001600160a01b0390921660009081526003602052604090209190915550565b8280546104f6906108f5565b90600052602060002090601f016020900481019282610518576000855561055e565b82601f106105315782800160ff1982351617855561055e565b8280016001018555821561055e579182015b8281111561055e578235825591602001919060010190610543565b5061056a9291506105e2565b5090565b82805461057a906108f5565b90600052602060002090601f01602090048101928261059c576000855561055e565b82601f106105b557805160ff191683800117855561055e565b8280016001018555821561055e579182015b8281111561055e5782518255916020019190600101906105c7565b5b8082111561056a57600081556001016105e3565b60008083601f84011261060957600080fd5b50813567ffffffffffffffff81111561062157600080fd5b60208301915083602082850101111561063957600080fd5b9250929050565b60006020828403121561065257600080fd5b81356001600160a01b03811681146103fd57600080fd5b60008060008060008060006080888a03121561068457600080fd5b873567ffffffffffffffff8082111561069c57600080fd5b6106a88b838c016105f7565b909950975060208a01359150808211156106c157600080fd5b6106cd8b838c016105f7565b909750955060408a0135945060608a01359150808211156106ed57600080fd5b506106fa8a828b016105f7565b989b979a50959850939692959293505050565b6000806040838503121561072057600080fd5b823567ffffffffffffffff8082111561073857600080fd5b818501915085601f83011261074c57600080fd5b81358181111561075e5761075e610930565b604051601f8201601f19908116603f0116810190838211818310171561078657610786610930565b8160405282815288602084870101111561079f57600080fd5b826020860160208301376000602093820184015298969091013596505050505050565b600081518084526107da8160208601602086016108c5565b601f01601f19169290920160200192915050565b8183823760009101908152919050565b600082516108108184602087016108c5565b9190910192915050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561089157888303603f1901855281518051878552610865888601826107c2565b91890151858303868b015291905061087d81836107c2565b968901969450505090860190600101610841565b509098975050505050505050565b600082198211156108c057634e487b7160e01b600052601160045260246000fd5b500190565b60005b838110156108e05781810151838201526020016108c8565b838111156108ef576000848401525b50505050565b600181811c9082168061090957607f821691505b6020821081141561092a57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220dc1cd6072b1168c8b199c48d4235ae3b02d588d7da11a88af9a8f9dc9ebde36964736f6c63430008070033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0xfa544161.
//
// Solidity: function getOwner(address token) view returns(uint256)
func (_Api *ApiCaller) GetOwner(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getOwner", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0xfa544161.
//
// Solidity: function getOwner(address token) view returns(uint256)
func (_Api *ApiSession) GetOwner(token common.Address) (*big.Int, error) {
	return _Api.Contract.GetOwner(&_Api.CallOpts, token)
}

// GetOwner is a free data retrieval call binding the contract method 0xfa544161.
//
// Solidity: function getOwner(address token) view returns(uint256)
func (_Api *ApiCallerSession) GetOwner(token common.Address) (*big.Int, error) {
	return _Api.Contract.GetOwner(&_Api.CallOpts, token)
}

// GetRecord is a free data retrieval call binding the contract method 0x005fcd5e.
//
// Solidity: function getRecord(string domain, int256 seclevel) view returns((string,string)[])
func (_Api *ApiCaller) GetRecord(opts *bind.CallOpts, domain string, seclevel *big.Int) ([]bdnsRegistryDomain, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getRecord", domain, seclevel)

	if err != nil {
		return *new([]bdnsRegistryDomain), err
	}

	out0 := *abi.ConvertType(out[0], new([]bdnsRegistryDomain)).(*[]bdnsRegistryDomain)

	return out0, err

}

// GetRecord is a free data retrieval call binding the contract method 0x005fcd5e.
//
// Solidity: function getRecord(string domain, int256 seclevel) view returns((string,string)[])
func (_Api *ApiSession) GetRecord(domain string, seclevel *big.Int) ([]bdnsRegistryDomain, error) {
	return _Api.Contract.GetRecord(&_Api.CallOpts, domain, seclevel)
}

// GetRecord is a free data retrieval call binding the contract method 0x005fcd5e.
//
// Solidity: function getRecord(string domain, int256 seclevel) view returns((string,string)[])
func (_Api *ApiCallerSession) GetRecord(domain string, seclevel *big.Int) ([]bdnsRegistryDomain, error) {
	return _Api.Contract.GetRecord(&_Api.CallOpts, domain, seclevel)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint256)
func (_Api *ApiCaller) GetVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint256)
func (_Api *ApiSession) GetVersion() (*big.Int, error) {
	return _Api.Contract.GetVersion(&_Api.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint256)
func (_Api *ApiCallerSession) GetVersion() (*big.Int, error) {
	return _Api.Contract.GetVersion(&_Api.CallOpts)
}

// CreateRecord is a paid mutator transaction binding the contract method 0xf4b518bc.
//
// Solidity: function createRecord(string domain, string ip_address, int256 seclevel, string device_type) returns()
func (_Api *ApiTransactor) CreateRecord(opts *bind.TransactOpts, domain string, ip_address string, seclevel *big.Int, device_type string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "createRecord", domain, ip_address, seclevel, device_type)
}

// CreateRecord is a paid mutator transaction binding the contract method 0xf4b518bc.
//
// Solidity: function createRecord(string domain, string ip_address, int256 seclevel, string device_type) returns()
func (_Api *ApiSession) CreateRecord(domain string, ip_address string, seclevel *big.Int, device_type string) (*types.Transaction, error) {
	return _Api.Contract.CreateRecord(&_Api.TransactOpts, domain, ip_address, seclevel, device_type)
}

// CreateRecord is a paid mutator transaction binding the contract method 0xf4b518bc.
//
// Solidity: function createRecord(string domain, string ip_address, int256 seclevel, string device_type) returns()
func (_Api *ApiTransactorSession) CreateRecord(domain string, ip_address string, seclevel *big.Int, device_type string) (*types.Transaction, error) {
	return _Api.Contract.CreateRecord(&_Api.TransactOpts, domain, ip_address, seclevel, device_type)
}

// ApiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Api contract.
type ApiTransferIterator struct {
	Event *ApiTransfer // Event containing the contract specifics and raw log

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
func (it *ApiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiTransfer)
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
		it.Event = new(ApiTransfer)
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
func (it *ApiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiTransfer represents a Transfer event raised by the Api contract.
type ApiTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Api *ApiFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ApiTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ApiTransferIterator{contract: _Api.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Api *ApiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ApiTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiTransfer)
				if err := _Api.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Api *ApiFilterer) ParseTransfer(log types.Log) (*ApiTransfer, error) {
	event := new(ApiTransfer)
	if err := _Api.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
