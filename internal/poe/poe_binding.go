// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poe

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

// ProofOfExistenceMetaData contains all meta data concerning the ProofOfExistence contract.
var ProofOfExistenceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getRecord\",\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"recorder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordHash\",\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"HashRecorded\",\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recorder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// ProofOfExistenceABI is the input ABI used to generate the binding from.
// Deprecated: Use ProofOfExistenceMetaData.ABI instead.
var ProofOfExistenceABI = ProofOfExistenceMetaData.ABI

// ProofOfExistence is an auto generated Go binding around an Ethereum contract.
type ProofOfExistence struct {
	ProofOfExistenceCaller     // Read-only binding to the contract
	ProofOfExistenceTransactor // Write-only binding to the contract
	ProofOfExistenceFilterer   // Log filterer for contract events
}

// ProofOfExistenceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProofOfExistenceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofOfExistenceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProofOfExistenceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofOfExistenceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProofOfExistenceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofOfExistenceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProofOfExistenceSession struct {
	Contract     *ProofOfExistence // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofOfExistenceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProofOfExistenceCallerSession struct {
	Contract *ProofOfExistenceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ProofOfExistenceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProofOfExistenceTransactorSession struct {
	Contract     *ProofOfExistenceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ProofOfExistenceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProofOfExistenceRaw struct {
	Contract *ProofOfExistence // Generic contract binding to access the raw methods on
}

// ProofOfExistenceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProofOfExistenceCallerRaw struct {
	Contract *ProofOfExistenceCaller // Generic read-only contract binding to access the raw methods on
}

// ProofOfExistenceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProofOfExistenceTransactorRaw struct {
	Contract *ProofOfExistenceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProofOfExistence creates a new instance of ProofOfExistence, bound to a specific deployed contract.
func NewProofOfExistence(address common.Address, backend bind.ContractBackend) (*ProofOfExistence, error) {
	contract, err := bindProofOfExistence(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProofOfExistence{ProofOfExistenceCaller: ProofOfExistenceCaller{contract: contract}, ProofOfExistenceTransactor: ProofOfExistenceTransactor{contract: contract}, ProofOfExistenceFilterer: ProofOfExistenceFilterer{contract: contract}}, nil
}

// NewProofOfExistenceCaller creates a new read-only instance of ProofOfExistence, bound to a specific deployed contract.
func NewProofOfExistenceCaller(address common.Address, caller bind.ContractCaller) (*ProofOfExistenceCaller, error) {
	contract, err := bindProofOfExistence(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProofOfExistenceCaller{contract: contract}, nil
}

// NewProofOfExistenceTransactor creates a new write-only instance of ProofOfExistence, bound to a specific deployed contract.
func NewProofOfExistenceTransactor(address common.Address, transactor bind.ContractTransactor) (*ProofOfExistenceTransactor, error) {
	contract, err := bindProofOfExistence(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProofOfExistenceTransactor{contract: contract}, nil
}

// NewProofOfExistenceFilterer creates a new log filterer instance of ProofOfExistence, bound to a specific deployed contract.
func NewProofOfExistenceFilterer(address common.Address, filterer bind.ContractFilterer) (*ProofOfExistenceFilterer, error) {
	contract, err := bindProofOfExistence(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProofOfExistenceFilterer{contract: contract}, nil
}

// bindProofOfExistence binds a generic wrapper to an already deployed contract.
func bindProofOfExistence(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProofOfExistenceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofOfExistence *ProofOfExistenceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProofOfExistence.Contract.ProofOfExistenceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofOfExistence *ProofOfExistenceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofOfExistence.Contract.ProofOfExistenceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofOfExistence *ProofOfExistenceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofOfExistence.Contract.ProofOfExistenceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofOfExistence *ProofOfExistenceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProofOfExistence.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofOfExistence *ProofOfExistenceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofOfExistence.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofOfExistence *ProofOfExistenceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofOfExistence.Contract.contract.Transact(opts, method, params...)
}

// GetRecord is a free data retrieval call binding the contract method 0x213681cd.
//
// Solidity: function getRecord(bytes32 hash) view returns(address recorder, uint256 timestamp)
func (_ProofOfExistence *ProofOfExistenceCaller) GetRecord(opts *bind.CallOpts, hash [32]byte) (struct {
	Recorder  common.Address
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _ProofOfExistence.contract.Call(opts, &out, "getRecord", hash)

	outstruct := new(struct {
		Recorder  common.Address
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recorder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRecord is a free data retrieval call binding the contract method 0x213681cd.
//
// Solidity: function getRecord(bytes32 hash) view returns(address recorder, uint256 timestamp)
func (_ProofOfExistence *ProofOfExistenceSession) GetRecord(hash [32]byte) (struct {
	Recorder  common.Address
	Timestamp *big.Int
}, error) {
	return _ProofOfExistence.Contract.GetRecord(&_ProofOfExistence.CallOpts, hash)
}

// GetRecord is a free data retrieval call binding the contract method 0x213681cd.
//
// Solidity: function getRecord(bytes32 hash) view returns(address recorder, uint256 timestamp)
func (_ProofOfExistence *ProofOfExistenceCallerSession) GetRecord(hash [32]byte) (struct {
	Recorder  common.Address
	Timestamp *big.Int
}, error) {
	return _ProofOfExistence.Contract.GetRecord(&_ProofOfExistence.CallOpts, hash)
}

// RecordHash is a paid mutator transaction binding the contract method 0xece11836.
//
// Solidity: function recordHash(bytes32 hash) returns()
func (_ProofOfExistence *ProofOfExistenceTransactor) RecordHash(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _ProofOfExistence.contract.Transact(opts, "recordHash", hash)
}

// RecordHash is a paid mutator transaction binding the contract method 0xece11836.
//
// Solidity: function recordHash(bytes32 hash) returns()
func (_ProofOfExistence *ProofOfExistenceSession) RecordHash(hash [32]byte) (*types.Transaction, error) {
	return _ProofOfExistence.Contract.RecordHash(&_ProofOfExistence.TransactOpts, hash)
}

// RecordHash is a paid mutator transaction binding the contract method 0xece11836.
//
// Solidity: function recordHash(bytes32 hash) returns()
func (_ProofOfExistence *ProofOfExistenceTransactorSession) RecordHash(hash [32]byte) (*types.Transaction, error) {
	return _ProofOfExistence.Contract.RecordHash(&_ProofOfExistence.TransactOpts, hash)
}

// ProofOfExistenceHashRecordedIterator is returned from FilterHashRecorded and is used to iterate over the raw logs and unpacked data for HashRecorded events raised by the ProofOfExistence contract.
type ProofOfExistenceHashRecordedIterator struct {
	Event *ProofOfExistenceHashRecorded // Event containing the contract specifics and raw log

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
func (it *ProofOfExistenceHashRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofOfExistenceHashRecorded)
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
		it.Event = new(ProofOfExistenceHashRecorded)
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
func (it *ProofOfExistenceHashRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofOfExistenceHashRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofOfExistenceHashRecorded represents a HashRecorded event raised by the ProofOfExistence contract.
type ProofOfExistenceHashRecorded struct {
	Hash      [32]byte
	Recorder  common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHashRecorded is a free log retrieval operation binding the contract event 0x1b5599beed06add0dbaf4353ea33bc551212b9814a5625a0bbbc33ccdb932605.
//
// Solidity: event HashRecorded(bytes32 indexed hash, address indexed recorder, uint256 timestamp)
func (_ProofOfExistence *ProofOfExistenceFilterer) FilterHashRecorded(opts *bind.FilterOpts, hash [][32]byte, recorder []common.Address) (*ProofOfExistenceHashRecordedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var recorderRule []interface{}
	for _, recorderItem := range recorder {
		recorderRule = append(recorderRule, recorderItem)
	}

	logs, sub, err := _ProofOfExistence.contract.FilterLogs(opts, "HashRecorded", hashRule, recorderRule)
	if err != nil {
		return nil, err
	}
	return &ProofOfExistenceHashRecordedIterator{contract: _ProofOfExistence.contract, event: "HashRecorded", logs: logs, sub: sub}, nil
}

// WatchHashRecorded is a free log subscription operation binding the contract event 0x1b5599beed06add0dbaf4353ea33bc551212b9814a5625a0bbbc33ccdb932605.
//
// Solidity: event HashRecorded(bytes32 indexed hash, address indexed recorder, uint256 timestamp)
func (_ProofOfExistence *ProofOfExistenceFilterer) WatchHashRecorded(opts *bind.WatchOpts, sink chan<- *ProofOfExistenceHashRecorded, hash [][32]byte, recorder []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var recorderRule []interface{}
	for _, recorderItem := range recorder {
		recorderRule = append(recorderRule, recorderItem)
	}

	logs, sub, err := _ProofOfExistence.contract.WatchLogs(opts, "HashRecorded", hashRule, recorderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofOfExistenceHashRecorded)
				if err := _ProofOfExistence.contract.UnpackLog(event, "HashRecorded", log); err != nil {
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

// ParseHashRecorded is a log parse operation binding the contract event 0x1b5599beed06add0dbaf4353ea33bc551212b9814a5625a0bbbc33ccdb932605.
//
// Solidity: event HashRecorded(bytes32 indexed hash, address indexed recorder, uint256 timestamp)
func (_ProofOfExistence *ProofOfExistenceFilterer) ParseHashRecorded(log types.Log) (*ProofOfExistenceHashRecorded, error) {
	event := new(ProofOfExistenceHashRecorded)
	if err := _ProofOfExistence.contract.UnpackLog(event, "HashRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
