// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// DataTypesBurnWithSigData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesBurnWithSigData struct {
	TokenId *big.Int
	Sig     DataTypesEIP712Signature
}

// DataTypesEIP712Signature is an auto generated low-level Go binding around an user-defined struct.
type DataTypesEIP712Signature struct {
	V        uint8
	R        [32]byte
	S        [32]byte
	Deadline *big.Int
}

// DataTypesPermitData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesPermitData struct {
	Spender common.Address
	TokenId *big.Int
	Sig     DataTypesEIP712Signature
}

// DataTypesPermitForAllData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesPermitForAllData struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Sig      DataTypesEIP712Signature
}

// ICollectionV1Spec is an auto generated low-level Go binding around an user-defined struct.
type ICollectionV1Spec struct {
	IsMintable     bool
	MintPrice      *big.Int
	MintCurrency   common.Address
	MaxPerWallet   *big.Int
	TotalMaxSupply *big.Int
	TotalSupply    *big.Int
	BaseUrl        string
}

// SCollectionV1MetaData contains all meta data concerning the SCollectionV1 contract.
var SCollectionV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApproveToCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"InvalidAmountPaid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentSupply\",\"type\":\"uint256\"}],\"name\":\"InvalidMaxSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxPerWalletExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxSupplyExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintERC2309QuantityExceedsLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMintable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnershipNotInitializedForExtraData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureInvalidS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureInvalidV\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValue\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURN_WITH_SIG_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CM\",\"outputs\":[{\"internalType\":\"contractCollectionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUrl\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.BurnWithSigData\",\"name\":\"_vars\",\"type\":\"tuple\"}],\"name\":\"burnWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSpec\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isMintable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"mintPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintCurrency\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"maxPerWallet\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"totalMaxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseUrl\",\"type\":\"string\"}],\"internalType\":\"structICollectionV1.Spec\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_vars\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPerWallet\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintCurrency\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.PermitData\",\"name\":\"_vars\",\"type\":\"tuple\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.PermitForAllData\",\"name\":\"_vars\",\"type\":\"tuple\"}],\"name\":\"permitForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roles\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseUrl\",\"type\":\"string\"}],\"name\":\"setBaseUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isMintable\",\"type\":\"bool\"}],\"name\":\"setIsMintable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_maxPerWallet\",\"type\":\"uint128\"}],\"name\":\"setMaxPerWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"_mintCurrency\",\"type\":\"address\"}],\"name\":\"setMintCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"setMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalMaxSupply\",\"type\":\"uint256\"}],\"name\":\"setTotalMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sigNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMaxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SCollectionV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use SCollectionV1MetaData.ABI instead.
var SCollectionV1ABI = SCollectionV1MetaData.ABI

// SCollectionV1 is an auto generated Go binding around an Ethereum contract.
type SCollectionV1 struct {
	SCollectionV1Caller     // Read-only binding to the contract
	SCollectionV1Transactor // Write-only binding to the contract
	SCollectionV1Filterer   // Log filterer for contract events
}

// SCollectionV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type SCollectionV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCollectionV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SCollectionV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCollectionV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SCollectionV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SCollectionV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SCollectionV1Session struct {
	Contract     *SCollectionV1    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SCollectionV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SCollectionV1CallerSession struct {
	Contract *SCollectionV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SCollectionV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SCollectionV1TransactorSession struct {
	Contract     *SCollectionV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SCollectionV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type SCollectionV1Raw struct {
	Contract *SCollectionV1 // Generic contract binding to access the raw methods on
}

// SCollectionV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SCollectionV1CallerRaw struct {
	Contract *SCollectionV1Caller // Generic read-only contract binding to access the raw methods on
}

// SCollectionV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SCollectionV1TransactorRaw struct {
	Contract *SCollectionV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSCollectionV1 creates a new instance of SCollectionV1, bound to a specific deployed contract.
func NewSCollectionV1(address common.Address, backend bind.ContractBackend) (*SCollectionV1, error) {
	contract, err := bindSCollectionV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SCollectionV1{SCollectionV1Caller: SCollectionV1Caller{contract: contract}, SCollectionV1Transactor: SCollectionV1Transactor{contract: contract}, SCollectionV1Filterer: SCollectionV1Filterer{contract: contract}}, nil
}

// NewSCollectionV1Caller creates a new read-only instance of SCollectionV1, bound to a specific deployed contract.
func NewSCollectionV1Caller(address common.Address, caller bind.ContractCaller) (*SCollectionV1Caller, error) {
	contract, err := bindSCollectionV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SCollectionV1Caller{contract: contract}, nil
}

// NewSCollectionV1Transactor creates a new write-only instance of SCollectionV1, bound to a specific deployed contract.
func NewSCollectionV1Transactor(address common.Address, transactor bind.ContractTransactor) (*SCollectionV1Transactor, error) {
	contract, err := bindSCollectionV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SCollectionV1Transactor{contract: contract}, nil
}

// NewSCollectionV1Filterer creates a new log filterer instance of SCollectionV1, bound to a specific deployed contract.
func NewSCollectionV1Filterer(address common.Address, filterer bind.ContractFilterer) (*SCollectionV1Filterer, error) {
	contract, err := bindSCollectionV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SCollectionV1Filterer{contract: contract}, nil
}

// bindSCollectionV1 binds a generic wrapper to an already deployed contract.
func bindSCollectionV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SCollectionV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SCollectionV1 *SCollectionV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SCollectionV1.Contract.SCollectionV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SCollectionV1 *SCollectionV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SCollectionV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SCollectionV1 *SCollectionV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SCollectionV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SCollectionV1 *SCollectionV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SCollectionV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SCollectionV1 *SCollectionV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCollectionV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SCollectionV1 *SCollectionV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SCollectionV1.Contract.contract.Transact(opts, method, params...)
}

// BURNWITHSIGTYPEHASH is a free data retrieval call binding the contract method 0x335c927f.
//
// Solidity: function BURN_WITH_SIG_TYPEHASH() view returns(bytes32)
func (_SCollectionV1 *SCollectionV1Caller) BURNWITHSIGTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "BURN_WITH_SIG_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BURNWITHSIGTYPEHASH is a free data retrieval call binding the contract method 0x335c927f.
//
// Solidity: function BURN_WITH_SIG_TYPEHASH() view returns(bytes32)
func (_SCollectionV1 *SCollectionV1Session) BURNWITHSIGTYPEHASH() ([32]byte, error) {
	return _SCollectionV1.Contract.BURNWITHSIGTYPEHASH(&_SCollectionV1.CallOpts)
}

// BURNWITHSIGTYPEHASH is a free data retrieval call binding the contract method 0x335c927f.
//
// Solidity: function BURN_WITH_SIG_TYPEHASH() view returns(bytes32)
func (_SCollectionV1 *SCollectionV1CallerSession) BURNWITHSIGTYPEHASH() ([32]byte, error) {
	return _SCollectionV1.Contract.BURNWITHSIGTYPEHASH(&_SCollectionV1.CallOpts)
}

// CM is a free data retrieval call binding the contract method 0xc16ac06f.
//
// Solidity: function CM() view returns(address)
func (_SCollectionV1 *SCollectionV1Caller) CM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "CM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CM is a free data retrieval call binding the contract method 0xc16ac06f.
//
// Solidity: function CM() view returns(address)
func (_SCollectionV1 *SCollectionV1Session) CM() (common.Address, error) {
	return _SCollectionV1.Contract.CM(&_SCollectionV1.CallOpts)
}

// CM is a free data retrieval call binding the contract method 0xc16ac06f.
//
// Solidity: function CM() view returns(address)
func (_SCollectionV1 *SCollectionV1CallerSession) CM() (common.Address, error) {
	return _SCollectionV1.Contract.CM(&_SCollectionV1.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SCollectionV1 *SCollectionV1Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SCollectionV1 *SCollectionV1Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SCollectionV1.Contract.BalanceOf(&_SCollectionV1.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SCollectionV1 *SCollectionV1CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SCollectionV1.Contract.BalanceOf(&_SCollectionV1.CallOpts, owner)
}

// BaseUrl is a free data retrieval call binding the contract method 0x5bcabf04.
//
// Solidity: function baseUrl() view returns(string)
func (_SCollectionV1 *SCollectionV1Caller) BaseUrl(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "baseUrl")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseUrl is a free data retrieval call binding the contract method 0x5bcabf04.
//
// Solidity: function baseUrl() view returns(string)
func (_SCollectionV1 *SCollectionV1Session) BaseUrl() (string, error) {
	return _SCollectionV1.Contract.BaseUrl(&_SCollectionV1.CallOpts)
}

// BaseUrl is a free data retrieval call binding the contract method 0x5bcabf04.
//
// Solidity: function baseUrl() view returns(string)
func (_SCollectionV1 *SCollectionV1CallerSession) BaseUrl() (string, error) {
	return _SCollectionV1.Contract.BaseUrl(&_SCollectionV1.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SCollectionV1 *SCollectionV1Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SCollectionV1 *SCollectionV1Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SCollectionV1.Contract.GetApproved(&_SCollectionV1.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SCollectionV1 *SCollectionV1CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SCollectionV1.Contract.GetApproved(&_SCollectionV1.CallOpts, tokenId)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_SCollectionV1 *SCollectionV1Caller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_SCollectionV1 *SCollectionV1Session) GetDomainSeparator() ([32]byte, error) {
	return _SCollectionV1.Contract.GetDomainSeparator(&_SCollectionV1.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_SCollectionV1 *SCollectionV1CallerSession) GetDomainSeparator() ([32]byte, error) {
	return _SCollectionV1.Contract.GetDomainSeparator(&_SCollectionV1.CallOpts)
}

// GetSpec is a free data retrieval call binding the contract method 0x1d2164a1.
//
// Solidity: function getSpec() view returns((bool,uint256,address,uint128,uint256,uint256,string))
func (_SCollectionV1 *SCollectionV1Caller) GetSpec(opts *bind.CallOpts) (ICollectionV1Spec, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "getSpec")

	if err != nil {
		return *new(ICollectionV1Spec), err
	}

	out0 := *abi.ConvertType(out[0], new(ICollectionV1Spec)).(*ICollectionV1Spec)

	return out0, err

}

// GetSpec is a free data retrieval call binding the contract method 0x1d2164a1.
//
// Solidity: function getSpec() view returns((bool,uint256,address,uint128,uint256,uint256,string))
func (_SCollectionV1 *SCollectionV1Session) GetSpec() (ICollectionV1Spec, error) {
	return _SCollectionV1.Contract.GetSpec(&_SCollectionV1.CallOpts)
}

// GetSpec is a free data retrieval call binding the contract method 0x1d2164a1.
//
// Solidity: function getSpec() view returns((bool,uint256,address,uint128,uint256,uint256,string))
func (_SCollectionV1 *SCollectionV1CallerSession) GetSpec() (ICollectionV1Spec, error) {
	return _SCollectionV1.Contract.GetSpec(&_SCollectionV1.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SCollectionV1 *SCollectionV1Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SCollectionV1 *SCollectionV1Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SCollectionV1.Contract.IsApprovedForAll(&_SCollectionV1.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SCollectionV1 *SCollectionV1CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SCollectionV1.Contract.IsApprovedForAll(&_SCollectionV1.CallOpts, owner, operator)
}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() view returns(bool)
func (_SCollectionV1 *SCollectionV1Caller) IsMintable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "isMintable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() view returns(bool)
func (_SCollectionV1 *SCollectionV1Session) IsMintable() (bool, error) {
	return _SCollectionV1.Contract.IsMintable(&_SCollectionV1.CallOpts)
}

// IsMintable is a free data retrieval call binding the contract method 0x46b45af7.
//
// Solidity: function isMintable() view returns(bool)
func (_SCollectionV1 *SCollectionV1CallerSession) IsMintable() (bool, error) {
	return _SCollectionV1.Contract.IsMintable(&_SCollectionV1.CallOpts)
}

// MaxPerWallet is a free data retrieval call binding the contract method 0x453c2310.
//
// Solidity: function maxPerWallet() view returns(uint128)
func (_SCollectionV1 *SCollectionV1Caller) MaxPerWallet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "maxPerWallet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPerWallet is a free data retrieval call binding the contract method 0x453c2310.
//
// Solidity: function maxPerWallet() view returns(uint128)
func (_SCollectionV1 *SCollectionV1Session) MaxPerWallet() (*big.Int, error) {
	return _SCollectionV1.Contract.MaxPerWallet(&_SCollectionV1.CallOpts)
}

// MaxPerWallet is a free data retrieval call binding the contract method 0x453c2310.
//
// Solidity: function maxPerWallet() view returns(uint128)
func (_SCollectionV1 *SCollectionV1CallerSession) MaxPerWallet() (*big.Int, error) {
	return _SCollectionV1.Contract.MaxPerWallet(&_SCollectionV1.CallOpts)
}

// MintCurrency is a free data retrieval call binding the contract method 0xc2d08583.
//
// Solidity: function mintCurrency() view returns(address)
func (_SCollectionV1 *SCollectionV1Caller) MintCurrency(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "mintCurrency")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MintCurrency is a free data retrieval call binding the contract method 0xc2d08583.
//
// Solidity: function mintCurrency() view returns(address)
func (_SCollectionV1 *SCollectionV1Session) MintCurrency() (common.Address, error) {
	return _SCollectionV1.Contract.MintCurrency(&_SCollectionV1.CallOpts)
}

// MintCurrency is a free data retrieval call binding the contract method 0xc2d08583.
//
// Solidity: function mintCurrency() view returns(address)
func (_SCollectionV1 *SCollectionV1CallerSession) MintCurrency() (common.Address, error) {
	return _SCollectionV1.Contract.MintCurrency(&_SCollectionV1.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_SCollectionV1 *SCollectionV1Caller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_SCollectionV1 *SCollectionV1Session) MintPrice() (*big.Int, error) {
	return _SCollectionV1.Contract.MintPrice(&_SCollectionV1.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_SCollectionV1 *SCollectionV1CallerSession) MintPrice() (*big.Int, error) {
	return _SCollectionV1.Contract.MintPrice(&_SCollectionV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SCollectionV1 *SCollectionV1Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SCollectionV1 *SCollectionV1Session) Name() (string, error) {
	return _SCollectionV1.Contract.Name(&_SCollectionV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SCollectionV1 *SCollectionV1CallerSession) Name() (string, error) {
	return _SCollectionV1.Contract.Name(&_SCollectionV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SCollectionV1 *SCollectionV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SCollectionV1 *SCollectionV1Session) Owner() (common.Address, error) {
	return _SCollectionV1.Contract.Owner(&_SCollectionV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SCollectionV1 *SCollectionV1CallerSession) Owner() (common.Address, error) {
	return _SCollectionV1.Contract.Owner(&_SCollectionV1.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SCollectionV1 *SCollectionV1Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SCollectionV1 *SCollectionV1Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SCollectionV1.Contract.OwnerOf(&_SCollectionV1.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SCollectionV1 *SCollectionV1CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SCollectionV1.Contract.OwnerOf(&_SCollectionV1.CallOpts, tokenId)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() view returns(address)
func (_SCollectionV1 *SCollectionV1Caller) Roles(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "roles")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() view returns(address)
func (_SCollectionV1 *SCollectionV1Session) Roles() (common.Address, error) {
	return _SCollectionV1.Contract.Roles(&_SCollectionV1.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() view returns(address)
func (_SCollectionV1 *SCollectionV1CallerSession) Roles() (common.Address, error) {
	return _SCollectionV1.Contract.Roles(&_SCollectionV1.CallOpts)
}

// SigNonce is a free data retrieval call binding the contract method 0xb4cbf1cb.
//
// Solidity: function sigNonce(address ) view returns(uint256)
func (_SCollectionV1 *SCollectionV1Caller) SigNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "sigNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SigNonce is a free data retrieval call binding the contract method 0xb4cbf1cb.
//
// Solidity: function sigNonce(address ) view returns(uint256)
func (_SCollectionV1 *SCollectionV1Session) SigNonce(arg0 common.Address) (*big.Int, error) {
	return _SCollectionV1.Contract.SigNonce(&_SCollectionV1.CallOpts, arg0)
}

// SigNonce is a free data retrieval call binding the contract method 0xb4cbf1cb.
//
// Solidity: function sigNonce(address ) view returns(uint256)
func (_SCollectionV1 *SCollectionV1CallerSession) SigNonce(arg0 common.Address) (*big.Int, error) {
	return _SCollectionV1.Contract.SigNonce(&_SCollectionV1.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_SCollectionV1 *SCollectionV1Caller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_SCollectionV1 *SCollectionV1Session) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SCollectionV1.Contract.SupportsInterface(&_SCollectionV1.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_SCollectionV1 *SCollectionV1CallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SCollectionV1.Contract.SupportsInterface(&_SCollectionV1.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SCollectionV1 *SCollectionV1Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SCollectionV1 *SCollectionV1Session) Symbol() (string, error) {
	return _SCollectionV1.Contract.Symbol(&_SCollectionV1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SCollectionV1 *SCollectionV1CallerSession) Symbol() (string, error) {
	return _SCollectionV1.Contract.Symbol(&_SCollectionV1.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SCollectionV1 *SCollectionV1Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SCollectionV1 *SCollectionV1Session) TokenURI(tokenId *big.Int) (string, error) {
	return _SCollectionV1.Contract.TokenURI(&_SCollectionV1.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SCollectionV1 *SCollectionV1CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SCollectionV1.Contract.TokenURI(&_SCollectionV1.CallOpts, tokenId)
}

// TotalMaxSupply is a free data retrieval call binding the contract method 0xa0617ad0.
//
// Solidity: function totalMaxSupply() view returns(uint256)
func (_SCollectionV1 *SCollectionV1Caller) TotalMaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "totalMaxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMaxSupply is a free data retrieval call binding the contract method 0xa0617ad0.
//
// Solidity: function totalMaxSupply() view returns(uint256)
func (_SCollectionV1 *SCollectionV1Session) TotalMaxSupply() (*big.Int, error) {
	return _SCollectionV1.Contract.TotalMaxSupply(&_SCollectionV1.CallOpts)
}

// TotalMaxSupply is a free data retrieval call binding the contract method 0xa0617ad0.
//
// Solidity: function totalMaxSupply() view returns(uint256)
func (_SCollectionV1 *SCollectionV1CallerSession) TotalMaxSupply() (*big.Int, error) {
	return _SCollectionV1.Contract.TotalMaxSupply(&_SCollectionV1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SCollectionV1 *SCollectionV1Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SCollectionV1.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SCollectionV1 *SCollectionV1Session) TotalSupply() (*big.Int, error) {
	return _SCollectionV1.Contract.TotalSupply(&_SCollectionV1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SCollectionV1 *SCollectionV1CallerSession) TotalSupply() (*big.Int, error) {
	return _SCollectionV1.Contract.TotalSupply(&_SCollectionV1.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SCollectionV1 *SCollectionV1Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SCollectionV1 *SCollectionV1Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.Approve(&_SCollectionV1.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.Approve(&_SCollectionV1.TransactOpts, to, tokenId)
}

// BurnWithSig is a paid mutator transaction binding the contract method 0x1fee7e3c.
//
// Solidity: function burnWithSig((uint256,(uint8,bytes32,bytes32,uint256)) _vars) returns()
func (_SCollectionV1 *SCollectionV1Transactor) BurnWithSig(opts *bind.TransactOpts, _vars DataTypesBurnWithSigData) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "burnWithSig", _vars)
}

// BurnWithSig is a paid mutator transaction binding the contract method 0x1fee7e3c.
//
// Solidity: function burnWithSig((uint256,(uint8,bytes32,bytes32,uint256)) _vars) returns()
func (_SCollectionV1 *SCollectionV1Session) BurnWithSig(_vars DataTypesBurnWithSigData) (*types.Transaction, error) {
	return _SCollectionV1.Contract.BurnWithSig(&_SCollectionV1.TransactOpts, _vars)
}

// BurnWithSig is a paid mutator transaction binding the contract method 0x1fee7e3c.
//
// Solidity: function burnWithSig((uint256,(uint8,bytes32,bytes32,uint256)) _vars) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) BurnWithSig(_vars DataTypesBurnWithSigData) (*types.Transaction, error) {
	return _SCollectionV1.Contract.BurnWithSig(&_SCollectionV1.TransactOpts, _vars)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes _vars) returns()
func (_SCollectionV1 *SCollectionV1Transactor) Initialize(opts *bind.TransactOpts, _vars []byte) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "initialize", _vars)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes _vars) returns()
func (_SCollectionV1 *SCollectionV1Session) Initialize(_vars []byte) (*types.Transaction, error) {
	return _SCollectionV1.Contract.Initialize(&_SCollectionV1.TransactOpts, _vars)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes _vars) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) Initialize(_vars []byte) (*types.Transaction, error) {
	return _SCollectionV1.Contract.Initialize(&_SCollectionV1.TransactOpts, _vars)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _quantity) payable returns(uint256)
func (_SCollectionV1 *SCollectionV1Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "mint", _to, _quantity)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _quantity) payable returns(uint256)
func (_SCollectionV1 *SCollectionV1Session) Mint(_to common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.Mint(&_SCollectionV1.TransactOpts, _to, _quantity)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _quantity) payable returns(uint256)
func (_SCollectionV1 *SCollectionV1TransactorSession) Mint(_to common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.Mint(&_SCollectionV1.TransactOpts, _to, _quantity)
}

// Permit is a paid mutator transaction binding the contract method 0x045b96b6.
//
// Solidity: function permit((address,uint256,(uint8,bytes32,bytes32,uint256)) _vars) returns()
func (_SCollectionV1 *SCollectionV1Transactor) Permit(opts *bind.TransactOpts, _vars DataTypesPermitData) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "permit", _vars)
}

// Permit is a paid mutator transaction binding the contract method 0x045b96b6.
//
// Solidity: function permit((address,uint256,(uint8,bytes32,bytes32,uint256)) _vars) returns()
func (_SCollectionV1 *SCollectionV1Session) Permit(_vars DataTypesPermitData) (*types.Transaction, error) {
	return _SCollectionV1.Contract.Permit(&_SCollectionV1.TransactOpts, _vars)
}

// Permit is a paid mutator transaction binding the contract method 0x045b96b6.
//
// Solidity: function permit((address,uint256,(uint8,bytes32,bytes32,uint256)) _vars) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) Permit(_vars DataTypesPermitData) (*types.Transaction, error) {
	return _SCollectionV1.Contract.Permit(&_SCollectionV1.TransactOpts, _vars)
}

// PermitForAll is a paid mutator transaction binding the contract method 0x241eda5f.
//
// Solidity: function permitForAll((address,address,bool,(uint8,bytes32,bytes32,uint256)) _vars) returns()
func (_SCollectionV1 *SCollectionV1Transactor) PermitForAll(opts *bind.TransactOpts, _vars DataTypesPermitForAllData) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "permitForAll", _vars)
}

// PermitForAll is a paid mutator transaction binding the contract method 0x241eda5f.
//
// Solidity: function permitForAll((address,address,bool,(uint8,bytes32,bytes32,uint256)) _vars) returns()
func (_SCollectionV1 *SCollectionV1Session) PermitForAll(_vars DataTypesPermitForAllData) (*types.Transaction, error) {
	return _SCollectionV1.Contract.PermitForAll(&_SCollectionV1.TransactOpts, _vars)
}

// PermitForAll is a paid mutator transaction binding the contract method 0x241eda5f.
//
// Solidity: function permitForAll((address,address,bool,(uint8,bytes32,bytes32,uint256)) _vars) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) PermitForAll(_vars DataTypesPermitForAllData) (*types.Transaction, error) {
	return _SCollectionV1.Contract.PermitForAll(&_SCollectionV1.TransactOpts, _vars)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SCollectionV1 *SCollectionV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SCollectionV1 *SCollectionV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _SCollectionV1.Contract.RenounceOwnership(&_SCollectionV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SCollectionV1.Contract.RenounceOwnership(&_SCollectionV1.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SCollectionV1 *SCollectionV1Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SCollectionV1 *SCollectionV1Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SafeTransferFrom(&_SCollectionV1.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SafeTransferFrom(&_SCollectionV1.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_SCollectionV1 *SCollectionV1Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_SCollectionV1 *SCollectionV1Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SafeTransferFrom0(&_SCollectionV1.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SafeTransferFrom0(&_SCollectionV1.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SCollectionV1 *SCollectionV1Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SCollectionV1 *SCollectionV1Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetApprovalForAll(&_SCollectionV1.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetApprovalForAll(&_SCollectionV1.TransactOpts, operator, approved)
}

// SetBaseUrl is a paid mutator transaction binding the contract method 0xc7c3268b.
//
// Solidity: function setBaseUrl(string _baseUrl) returns()
func (_SCollectionV1 *SCollectionV1Transactor) SetBaseUrl(opts *bind.TransactOpts, _baseUrl string) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "setBaseUrl", _baseUrl)
}

// SetBaseUrl is a paid mutator transaction binding the contract method 0xc7c3268b.
//
// Solidity: function setBaseUrl(string _baseUrl) returns()
func (_SCollectionV1 *SCollectionV1Session) SetBaseUrl(_baseUrl string) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetBaseUrl(&_SCollectionV1.TransactOpts, _baseUrl)
}

// SetBaseUrl is a paid mutator transaction binding the contract method 0xc7c3268b.
//
// Solidity: function setBaseUrl(string _baseUrl) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) SetBaseUrl(_baseUrl string) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetBaseUrl(&_SCollectionV1.TransactOpts, _baseUrl)
}

// SetIsMintable is a paid mutator transaction binding the contract method 0x88459f9a.
//
// Solidity: function setIsMintable(bool _isMintable) returns()
func (_SCollectionV1 *SCollectionV1Transactor) SetIsMintable(opts *bind.TransactOpts, _isMintable bool) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "setIsMintable", _isMintable)
}

// SetIsMintable is a paid mutator transaction binding the contract method 0x88459f9a.
//
// Solidity: function setIsMintable(bool _isMintable) returns()
func (_SCollectionV1 *SCollectionV1Session) SetIsMintable(_isMintable bool) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetIsMintable(&_SCollectionV1.TransactOpts, _isMintable)
}

// SetIsMintable is a paid mutator transaction binding the contract method 0x88459f9a.
//
// Solidity: function setIsMintable(bool _isMintable) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) SetIsMintable(_isMintable bool) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetIsMintable(&_SCollectionV1.TransactOpts, _isMintable)
}

// SetMaxPerWallet is a paid mutator transaction binding the contract method 0xe35807d5.
//
// Solidity: function setMaxPerWallet(uint128 _maxPerWallet) returns()
func (_SCollectionV1 *SCollectionV1Transactor) SetMaxPerWallet(opts *bind.TransactOpts, _maxPerWallet *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "setMaxPerWallet", _maxPerWallet)
}

// SetMaxPerWallet is a paid mutator transaction binding the contract method 0xe35807d5.
//
// Solidity: function setMaxPerWallet(uint128 _maxPerWallet) returns()
func (_SCollectionV1 *SCollectionV1Session) SetMaxPerWallet(_maxPerWallet *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetMaxPerWallet(&_SCollectionV1.TransactOpts, _maxPerWallet)
}

// SetMaxPerWallet is a paid mutator transaction binding the contract method 0xe35807d5.
//
// Solidity: function setMaxPerWallet(uint128 _maxPerWallet) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) SetMaxPerWallet(_maxPerWallet *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetMaxPerWallet(&_SCollectionV1.TransactOpts, _maxPerWallet)
}

// SetMintCurrency is a paid mutator transaction binding the contract method 0xae04a62f.
//
// Solidity: function setMintCurrency(address _mintCurrency) returns()
func (_SCollectionV1 *SCollectionV1Transactor) SetMintCurrency(opts *bind.TransactOpts, _mintCurrency common.Address) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "setMintCurrency", _mintCurrency)
}

// SetMintCurrency is a paid mutator transaction binding the contract method 0xae04a62f.
//
// Solidity: function setMintCurrency(address _mintCurrency) returns()
func (_SCollectionV1 *SCollectionV1Session) SetMintCurrency(_mintCurrency common.Address) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetMintCurrency(&_SCollectionV1.TransactOpts, _mintCurrency)
}

// SetMintCurrency is a paid mutator transaction binding the contract method 0xae04a62f.
//
// Solidity: function setMintCurrency(address _mintCurrency) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) SetMintCurrency(_mintCurrency common.Address) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetMintCurrency(&_SCollectionV1.TransactOpts, _mintCurrency)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_SCollectionV1 *SCollectionV1Transactor) SetMintPrice(opts *bind.TransactOpts, _mintPrice *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "setMintPrice", _mintPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_SCollectionV1 *SCollectionV1Session) SetMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetMintPrice(&_SCollectionV1.TransactOpts, _mintPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) SetMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetMintPrice(&_SCollectionV1.TransactOpts, _mintPrice)
}

// SetTotalMaxSupply is a paid mutator transaction binding the contract method 0x9d9e3c47.
//
// Solidity: function setTotalMaxSupply(uint256 _totalMaxSupply) returns()
func (_SCollectionV1 *SCollectionV1Transactor) SetTotalMaxSupply(opts *bind.TransactOpts, _totalMaxSupply *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "setTotalMaxSupply", _totalMaxSupply)
}

// SetTotalMaxSupply is a paid mutator transaction binding the contract method 0x9d9e3c47.
//
// Solidity: function setTotalMaxSupply(uint256 _totalMaxSupply) returns()
func (_SCollectionV1 *SCollectionV1Session) SetTotalMaxSupply(_totalMaxSupply *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetTotalMaxSupply(&_SCollectionV1.TransactOpts, _totalMaxSupply)
}

// SetTotalMaxSupply is a paid mutator transaction binding the contract method 0x9d9e3c47.
//
// Solidity: function setTotalMaxSupply(uint256 _totalMaxSupply) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) SetTotalMaxSupply(_totalMaxSupply *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.SetTotalMaxSupply(&_SCollectionV1.TransactOpts, _totalMaxSupply)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SCollectionV1 *SCollectionV1Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SCollectionV1 *SCollectionV1Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.TransferFrom(&_SCollectionV1.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SCollectionV1.Contract.TransferFrom(&_SCollectionV1.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SCollectionV1 *SCollectionV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SCollectionV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SCollectionV1 *SCollectionV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SCollectionV1.Contract.TransferOwnership(&_SCollectionV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SCollectionV1 *SCollectionV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SCollectionV1.Contract.TransferOwnership(&_SCollectionV1.TransactOpts, newOwner)
}

// SCollectionV1ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SCollectionV1 contract.
type SCollectionV1ApprovalIterator struct {
	Event *SCollectionV1Approval // Event containing the contract specifics and raw log

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
func (it *SCollectionV1ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCollectionV1Approval)
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
		it.Event = new(SCollectionV1Approval)
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
func (it *SCollectionV1ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCollectionV1ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCollectionV1Approval represents a Approval event raised by the SCollectionV1 contract.
type SCollectionV1Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SCollectionV1 *SCollectionV1Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SCollectionV1ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SCollectionV1.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SCollectionV1ApprovalIterator{contract: _SCollectionV1.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SCollectionV1 *SCollectionV1Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SCollectionV1Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SCollectionV1.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCollectionV1Approval)
				if err := _SCollectionV1.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SCollectionV1 *SCollectionV1Filterer) ParseApproval(log types.Log) (*SCollectionV1Approval, error) {
	event := new(SCollectionV1Approval)
	if err := _SCollectionV1.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCollectionV1ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SCollectionV1 contract.
type SCollectionV1ApprovalForAllIterator struct {
	Event *SCollectionV1ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SCollectionV1ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCollectionV1ApprovalForAll)
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
		it.Event = new(SCollectionV1ApprovalForAll)
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
func (it *SCollectionV1ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCollectionV1ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCollectionV1ApprovalForAll represents a ApprovalForAll event raised by the SCollectionV1 contract.
type SCollectionV1ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SCollectionV1 *SCollectionV1Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SCollectionV1ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SCollectionV1.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SCollectionV1ApprovalForAllIterator{contract: _SCollectionV1.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SCollectionV1 *SCollectionV1Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SCollectionV1ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SCollectionV1.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCollectionV1ApprovalForAll)
				if err := _SCollectionV1.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SCollectionV1 *SCollectionV1Filterer) ParseApprovalForAll(log types.Log) (*SCollectionV1ApprovalForAll, error) {
	event := new(SCollectionV1ApprovalForAll)
	if err := _SCollectionV1.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCollectionV1ConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the SCollectionV1 contract.
type SCollectionV1ConsecutiveTransferIterator struct {
	Event *SCollectionV1ConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *SCollectionV1ConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCollectionV1ConsecutiveTransfer)
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
		it.Event = new(SCollectionV1ConsecutiveTransfer)
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
func (it *SCollectionV1ConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCollectionV1ConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCollectionV1ConsecutiveTransfer represents a ConsecutiveTransfer event raised by the SCollectionV1 contract.
type SCollectionV1ConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_SCollectionV1 *SCollectionV1Filterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*SCollectionV1ConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SCollectionV1.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SCollectionV1ConsecutiveTransferIterator{contract: _SCollectionV1.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_SCollectionV1 *SCollectionV1Filterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *SCollectionV1ConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SCollectionV1.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCollectionV1ConsecutiveTransfer)
				if err := _SCollectionV1.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_SCollectionV1 *SCollectionV1Filterer) ParseConsecutiveTransfer(log types.Log) (*SCollectionV1ConsecutiveTransfer, error) {
	event := new(SCollectionV1ConsecutiveTransfer)
	if err := _SCollectionV1.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCollectionV1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SCollectionV1 contract.
type SCollectionV1InitializedIterator struct {
	Event *SCollectionV1Initialized // Event containing the contract specifics and raw log

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
func (it *SCollectionV1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCollectionV1Initialized)
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
		it.Event = new(SCollectionV1Initialized)
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
func (it *SCollectionV1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCollectionV1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCollectionV1Initialized represents a Initialized event raised by the SCollectionV1 contract.
type SCollectionV1Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SCollectionV1 *SCollectionV1Filterer) FilterInitialized(opts *bind.FilterOpts) (*SCollectionV1InitializedIterator, error) {

	logs, sub, err := _SCollectionV1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SCollectionV1InitializedIterator{contract: _SCollectionV1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SCollectionV1 *SCollectionV1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SCollectionV1Initialized) (event.Subscription, error) {

	logs, sub, err := _SCollectionV1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCollectionV1Initialized)
				if err := _SCollectionV1.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SCollectionV1 *SCollectionV1Filterer) ParseInitialized(log types.Log) (*SCollectionV1Initialized, error) {
	event := new(SCollectionV1Initialized)
	if err := _SCollectionV1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCollectionV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SCollectionV1 contract.
type SCollectionV1OwnershipTransferredIterator struct {
	Event *SCollectionV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SCollectionV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCollectionV1OwnershipTransferred)
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
		it.Event = new(SCollectionV1OwnershipTransferred)
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
func (it *SCollectionV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCollectionV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCollectionV1OwnershipTransferred represents a OwnershipTransferred event raised by the SCollectionV1 contract.
type SCollectionV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SCollectionV1 *SCollectionV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SCollectionV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SCollectionV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SCollectionV1OwnershipTransferredIterator{contract: _SCollectionV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SCollectionV1 *SCollectionV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SCollectionV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SCollectionV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCollectionV1OwnershipTransferred)
				if err := _SCollectionV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SCollectionV1 *SCollectionV1Filterer) ParseOwnershipTransferred(log types.Log) (*SCollectionV1OwnershipTransferred, error) {
	event := new(SCollectionV1OwnershipTransferred)
	if err := _SCollectionV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SCollectionV1TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SCollectionV1 contract.
type SCollectionV1TransferIterator struct {
	Event *SCollectionV1Transfer // Event containing the contract specifics and raw log

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
func (it *SCollectionV1TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SCollectionV1Transfer)
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
		it.Event = new(SCollectionV1Transfer)
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
func (it *SCollectionV1TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SCollectionV1TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SCollectionV1Transfer represents a Transfer event raised by the SCollectionV1 contract.
type SCollectionV1Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SCollectionV1 *SCollectionV1Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SCollectionV1TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SCollectionV1.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SCollectionV1TransferIterator{contract: _SCollectionV1.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SCollectionV1 *SCollectionV1Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SCollectionV1Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SCollectionV1.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SCollectionV1Transfer)
				if err := _SCollectionV1.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SCollectionV1 *SCollectionV1Filterer) ParseTransfer(log types.Log) (*SCollectionV1Transfer, error) {
	event := new(SCollectionV1Transfer)
	if err := _SCollectionV1.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
