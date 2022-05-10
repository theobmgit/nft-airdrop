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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"week\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tranche\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"TrancheAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tranche\",\"type\":\"uint256\"}],\"name\":\"TrancheExpired\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_trancheId\",\"type\":\"uint256\"}],\"name\":\"expireTranche\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tranche\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tranches\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_merkleProofs\",\"type\":\"bytes32[][]\"}],\"name\":\"redeemMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"seedNewTranche\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"trancheId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tranches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tranche\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001c1738038062001c178339818101604052810190620000379190620001e9565b620000576200004b6200009f60201b60201c565b620000a760201b60201c565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200021b565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200019d8262000170565b9050919050565b6000620001b18262000190565b9050919050565b620001c381620001a4565b8114620001cf57600080fd5b50565b600081519050620001e381620001b8565b92915050565b6000602082840312156200020257620002016200016b565b5b60006200021284828501620001d2565b91505092915050565b6119ec806200022b6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063d58ea88e11610071578063d58ea88e1461018d578063d5efd20a146101a9578063e329e18b146101c5578063ebf04917146101f5578063f2fde38b14610213578063fc0c546a1461022f576100b4565b80630fc376ff146100b9578063120aa877146100d5578063715018a61461010557806371c5ecb11461010f578063857f219c1461013f5780638da5cb5b1461016f575b600080fd5b6100d360048036038101906100ce9190610eff565b61024d565b005b6100ef60048036038101906100ea9190610f82565b610269565b6040516100fc9190610fdd565b60405180910390f35b61010d610298565b005b61012960048036038101906101249190610ff8565b610320565b6040516101369190611034565b60405180910390f35b6101596004803603810190610154919061104f565b610338565b604051610166919061108b565b60405180910390f35b61017761042d565b60405161018491906110b5565b60405180910390f35b6101a760048036038101906101a29190611274565b610456565b005b6101c360048036038101906101be9190610ff8565b61054f565b005b6101df60048036038101906101da9190610eff565b610620565b6040516101ec9190610fdd565b60405180910390f35b6101fd610638565b60405161020a919061108b565b60405180910390f35b61022d6004803603810190610228919061132f565b61063e565b005b610237610735565b60405161024491906113bb565b60405180910390f35b6102598484848461075b565b6102638483610932565b50505050565b60036020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b6102a0610ad8565b73ffffffffffffffffffffffffffffffffffffffff166102be61042d565b73ffffffffffffffffffffffffffffffffffffffff1614610314576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030b90611433565b60405180910390fd5b61031e6000610ae0565b565b60026020528060005260406000206000915090505481565b6000610342610ad8565b73ffffffffffffffffffffffffffffffffffffffff1661036061042d565b73ffffffffffffffffffffffffffffffffffffffff16146103b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ad90611433565b60405180910390fd5b60045490508160026000838152602001908152602001600020819055506103e96001600454610ba490919063ffffffff16565b6004819055507f899089a7b39272321d9f5253a8c187ea3b7507701ebb523db8ae97763f0f7db88183604051610420929190611453565b60405180910390a1919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008351905082518114801561046c5750815181145b6104ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a2906114c8565b60405180910390fd5b60005b8181101561054757610510868683815181106104cd576104cc6114e8565b5b60200260200101518684815181106104e8576104e76114e8565b5b6020026020010151868581518110610503576105026114e8565b5b602002602001015161075b565b61053486858381518110610527576105266114e8565b5b6020026020010151610932565b808061053f90611546565b9150506104ae565b505050505050565b610557610ad8565b73ffffffffffffffffffffffffffffffffffffffff1661057561042d565b73ffffffffffffffffffffffffffffffffffffffff16146105cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c290611433565b60405180910390fd5b6000801b60026000838152602001908152602001600020819055507fcc071cbd9ae50a4c78d1153b76bd2d46ba8d4c7662842718ec3de1d67a144daf81604051610615919061108b565b60405180910390a150565b600061062e85858585610bba565b9050949350505050565b60045481565b610646610ad8565b73ffffffffffffffffffffffffffffffffffffffff1661066461042d565b73ffffffffffffffffffffffffffffffffffffffff16146106ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b190611433565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610729576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072090611600565b60405180910390fd5b61073281610ae0565b50565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600454831061079f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107969061166c565b60405180910390fd5b6003600084815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561083d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610834906116d8565b60405180910390fd5b61084984848484610bba565b610888576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161087f90611744565b60405180910390fd5b60016003600085815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a84848460405161092493929190611764565b60405180910390a150505050565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e836040518263ffffffff1660e01b81526004016109a5919061108b565b602060405180830381865afa1580156109c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e691906117b0565b73ffffffffffffffffffffffffffffffffffffffff1603610a3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a339061184f565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342842e0e610a82610ad8565b84846040518463ffffffff1660e01b8152600401610aa29392919061186f565b600060405180830381600087803b158015610abc57600080fd5b505af1158015610ad0573d6000803e3d6000fd5b505050505050565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60008183610bb291906118a6565b905092915050565b600080610bc78685610bf2565b9050610be783600260008881526020019081526020016000205483610c25565b915050949350505050565b60008183604051602001610c07929190611965565b60405160208183030381529060405280519060200120905092915050565b600082610c328584610c3c565b1490509392505050565b60008082905060005b8451811015610ca6576000858281518110610c6357610c626114e8565b5b60200260200101519050808311610c8557610c7e8382610cb1565b9250610c92565b610c8f8184610cb1565b92505b508080610c9e90611546565b915050610c45565b508091505092915050565b600082600052816020526040600020905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610d0782610cdc565b9050919050565b610d1781610cfc565b8114610d2257600080fd5b50565b600081359050610d3481610d0e565b92915050565b6000819050919050565b610d4d81610d3a565b8114610d5857600080fd5b50565b600081359050610d6a81610d44565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610dbe82610d75565b810181811067ffffffffffffffff82111715610ddd57610ddc610d86565b5b80604052505050565b6000610df0610cc8565b9050610dfc8282610db5565b919050565b600067ffffffffffffffff821115610e1c57610e1b610d86565b5b602082029050602081019050919050565b600080fd5b6000819050919050565b610e4581610e32565b8114610e5057600080fd5b50565b600081359050610e6281610e3c565b92915050565b6000610e7b610e7684610e01565b610de6565b90508083825260208201905060208402830185811115610e9e57610e9d610e2d565b5b835b81811015610ec75780610eb38882610e53565b845260208401935050602081019050610ea0565b5050509392505050565b600082601f830112610ee657610ee5610d70565b5b8135610ef6848260208601610e68565b91505092915050565b60008060008060808587031215610f1957610f18610cd2565b5b6000610f2787828801610d25565b9450506020610f3887828801610d5b565b9350506040610f4987828801610d5b565b925050606085013567ffffffffffffffff811115610f6a57610f69610cd7565b5b610f7687828801610ed1565b91505092959194509250565b60008060408385031215610f9957610f98610cd2565b5b6000610fa785828601610d5b565b9250506020610fb885828601610d25565b9150509250929050565b60008115159050919050565b610fd781610fc2565b82525050565b6000602082019050610ff26000830184610fce565b92915050565b60006020828403121561100e5761100d610cd2565b5b600061101c84828501610d5b565b91505092915050565b61102e81610e32565b82525050565b60006020820190506110496000830184611025565b92915050565b60006020828403121561106557611064610cd2565b5b600061107384828501610e53565b91505092915050565b61108581610d3a565b82525050565b60006020820190506110a0600083018461107c565b92915050565b6110af81610cfc565b82525050565b60006020820190506110ca60008301846110a6565b92915050565b600067ffffffffffffffff8211156110eb576110ea610d86565b5b602082029050602081019050919050565b600061110f61110a846110d0565b610de6565b9050808382526020820190506020840283018581111561113257611131610e2d565b5b835b8181101561115b57806111478882610d5b565b845260208401935050602081019050611134565b5050509392505050565b600082601f83011261117a57611179610d70565b5b813561118a8482602086016110fc565b91505092915050565b600067ffffffffffffffff8211156111ae576111ad610d86565b5b602082029050602081019050919050565b60006111d26111cd84611193565b610de6565b905080838252602082019050602084028301858111156111f5576111f4610e2d565b5b835b8181101561123c57803567ffffffffffffffff81111561121a57611219610d70565b5b8086016112278982610ed1565b855260208501945050506020810190506111f7565b5050509392505050565b600082601f83011261125b5761125a610d70565b5b813561126b8482602086016111bf565b91505092915050565b6000806000806080858703121561128e5761128d610cd2565b5b600061129c87828801610d25565b945050602085013567ffffffffffffffff8111156112bd576112bc610cd7565b5b6112c987828801611165565b935050604085013567ffffffffffffffff8111156112ea576112e9610cd7565b5b6112f687828801611165565b925050606085013567ffffffffffffffff81111561131757611316610cd7565b5b61132387828801611246565b91505092959194509250565b60006020828403121561134557611344610cd2565b5b600061135384828501610d25565b91505092915050565b6000819050919050565b600061138161137c61137784610cdc565b61135c565b610cdc565b9050919050565b600061139382611366565b9050919050565b60006113a582611388565b9050919050565b6113b58161139a565b82525050565b60006020820190506113d060008301846113ac565b92915050565b600082825260208201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b600061141d6020836113d6565b9150611428826113e7565b602082019050919050565b6000602082019050818103600083015261144c81611410565b9050919050565b6000604082019050611468600083018561107c565b6114756020830184611025565b9392505050565b7f4d69736d61746368696e6720696e707574730000000000000000000000000000600082015250565b60006114b26012836113d6565b91506114bd8261147c565b602082019050919050565b600060208201905081810360008301526114e1816114a5565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061155182610d3a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361158357611582611517565b5b600182019050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006115ea6026836113d6565b91506115f58261158e565b604082019050919050565b60006020820190508181036000830152611619816115dd565b9050919050565b7f5472616e6368652063616e6e6f7420626520696e207468652066757475726500600082015250565b6000611656601f836113d6565b915061166182611620565b602082019050919050565b6000602082019050818103600083015261168581611649565b9050919050565b7f4163636f756e742068617320616c726561647920636c61696d65640000000000600082015250565b60006116c2601b836113d6565b91506116cd8261168c565b602082019050919050565b600060208201905081810360008301526116f1816116b5565b9050919050565b7f496e636f7272656374206d65726b6c652070726f6f6600000000000000000000600082015250565b600061172e6016836113d6565b9150611739826116f8565b602082019050919050565b6000602082019050818103600083015261175d81611721565b9050919050565b600060608201905061177960008301866110a6565b611786602083018561107c565b611793604083018461107c565b949350505050565b6000815190506117aa81610d0e565b92915050565b6000602082840312156117c6576117c5610cd2565b5b60006117d48482850161179b565b91505092915050565b7f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b6000611839602c836113d6565b9150611844826117dd565b604082019050919050565b600060208201905081810360008301526118688161182c565b9050919050565b600060608201905061188460008301866110a6565b61189160208301856110a6565b61189e604083018461107c565b949350505050565b60006118b182610d3a565b91506118bc83610d3a565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156118f1576118f0611517565b5b828201905092915050565b6000819050919050565b61191761191282610d3a565b6118fc565b82525050565b60008160601b9050919050565b60006119358261191d565b9050919050565b60006119478261192a565b9050919050565b61195f61195a82610cfc565b61193c565b82525050565b60006119718285611906565b602082019150611981828461194e565b601482019150819050939250505056fea2646970667358221220691456430f7583603b9fef10fed84dd90aca7e37963188a3c0c289041f8c245764736f6c637827302e382e31342d646576656c6f702e323032322e352e392b636f6d6d69742e34343135376161360058",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_Contracts *ContractsCaller) Claimed(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "claimed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_Contracts *ContractsSession) Claimed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Contracts.Contract.Claimed(&_Contracts.CallOpts, arg0, arg1)
}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_Contracts *ContractsCallerSession) Claimed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Contracts.Contract.Claimed(&_Contracts.CallOpts, arg0, arg1)
}

// MerkleRoots is a free data retrieval call binding the contract method 0x71c5ecb1.
//
// Solidity: function merkleRoots(uint256 ) view returns(bytes32)
func (_Contracts *ContractsCaller) MerkleRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "merkleRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoots is a free data retrieval call binding the contract method 0x71c5ecb1.
//
// Solidity: function merkleRoots(uint256 ) view returns(bytes32)
func (_Contracts *ContractsSession) MerkleRoots(arg0 *big.Int) ([32]byte, error) {
	return _Contracts.Contract.MerkleRoots(&_Contracts.CallOpts, arg0)
}

// MerkleRoots is a free data retrieval call binding the contract method 0x71c5ecb1.
//
// Solidity: function merkleRoots(uint256 ) view returns(bytes32)
func (_Contracts *ContractsCallerSession) MerkleRoots(arg0 *big.Int) ([32]byte, error) {
	return _Contracts.Contract.MerkleRoots(&_Contracts.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contracts *ContractsCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contracts *ContractsSession) Token() (common.Address, error) {
	return _Contracts.Contract.Token(&_Contracts.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Contracts *ContractsCallerSession) Token() (common.Address, error) {
	return _Contracts.Contract.Token(&_Contracts.CallOpts)
}

// Tranches is a free data retrieval call binding the contract method 0xebf04917.
//
// Solidity: function tranches() view returns(uint256)
func (_Contracts *ContractsCaller) Tranches(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "tranches")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tranches is a free data retrieval call binding the contract method 0xebf04917.
//
// Solidity: function tranches() view returns(uint256)
func (_Contracts *ContractsSession) Tranches() (*big.Int, error) {
	return _Contracts.Contract.Tranches(&_Contracts.CallOpts)
}

// Tranches is a free data retrieval call binding the contract method 0xebf04917.
//
// Solidity: function tranches() view returns(uint256)
func (_Contracts *ContractsCallerSession) Tranches() (*big.Int, error) {
	return _Contracts.Contract.Tranches(&_Contracts.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0xe329e18b.
//
// Solidity: function verify(address _account, uint256 _tranche, uint256 _tokenId, bytes32[] _merkleProof) view returns(bool valid)
func (_Contracts *ContractsCaller) Verify(opts *bind.CallOpts, _account common.Address, _tranche *big.Int, _tokenId *big.Int, _merkleProof [][32]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "verify", _account, _tranche, _tokenId, _merkleProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xe329e18b.
//
// Solidity: function verify(address _account, uint256 _tranche, uint256 _tokenId, bytes32[] _merkleProof) view returns(bool valid)
func (_Contracts *ContractsSession) Verify(_account common.Address, _tranche *big.Int, _tokenId *big.Int, _merkleProof [][32]byte) (bool, error) {
	return _Contracts.Contract.Verify(&_Contracts.CallOpts, _account, _tranche, _tokenId, _merkleProof)
}

// Verify is a free data retrieval call binding the contract method 0xe329e18b.
//
// Solidity: function verify(address _account, uint256 _tranche, uint256 _tokenId, bytes32[] _merkleProof) view returns(bool valid)
func (_Contracts *ContractsCallerSession) Verify(_account common.Address, _tranche *big.Int, _tokenId *big.Int, _merkleProof [][32]byte) (bool, error) {
	return _Contracts.Contract.Verify(&_Contracts.CallOpts, _account, _tranche, _tokenId, _merkleProof)
}

// ExpireTranche is a paid mutator transaction binding the contract method 0xd5efd20a.
//
// Solidity: function expireTranche(uint256 _trancheId) returns()
func (_Contracts *ContractsTransactor) ExpireTranche(opts *bind.TransactOpts, _trancheId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "expireTranche", _trancheId)
}

// ExpireTranche is a paid mutator transaction binding the contract method 0xd5efd20a.
//
// Solidity: function expireTranche(uint256 _trancheId) returns()
func (_Contracts *ContractsSession) ExpireTranche(_trancheId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ExpireTranche(&_Contracts.TransactOpts, _trancheId)
}

// ExpireTranche is a paid mutator transaction binding the contract method 0xd5efd20a.
//
// Solidity: function expireTranche(uint256 _trancheId) returns()
func (_Contracts *ContractsTransactorSession) ExpireTranche(_trancheId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ExpireTranche(&_Contracts.TransactOpts, _trancheId)
}

// Redeem is a paid mutator transaction binding the contract method 0x0fc376ff.
//
// Solidity: function redeem(address _account, uint256 _tranche, uint256 _tokenId, bytes32[] _merkleProof) returns()
func (_Contracts *ContractsTransactor) Redeem(opts *bind.TransactOpts, _account common.Address, _tranche *big.Int, _tokenId *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "redeem", _account, _tranche, _tokenId, _merkleProof)
}

// Redeem is a paid mutator transaction binding the contract method 0x0fc376ff.
//
// Solidity: function redeem(address _account, uint256 _tranche, uint256 _tokenId, bytes32[] _merkleProof) returns()
func (_Contracts *ContractsSession) Redeem(_account common.Address, _tranche *big.Int, _tokenId *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Redeem(&_Contracts.TransactOpts, _account, _tranche, _tokenId, _merkleProof)
}

// Redeem is a paid mutator transaction binding the contract method 0x0fc376ff.
//
// Solidity: function redeem(address _account, uint256 _tranche, uint256 _tokenId, bytes32[] _merkleProof) returns()
func (_Contracts *ContractsTransactorSession) Redeem(_account common.Address, _tranche *big.Int, _tokenId *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Redeem(&_Contracts.TransactOpts, _account, _tranche, _tokenId, _merkleProof)
}

// RedeemMany is a paid mutator transaction binding the contract method 0xd58ea88e.
//
// Solidity: function redeemMany(address _account, uint256[] _tranches, uint256[] _tokenIds, bytes32[][] _merkleProofs) returns()
func (_Contracts *ContractsTransactor) RedeemMany(opts *bind.TransactOpts, _account common.Address, _tranches []*big.Int, _tokenIds []*big.Int, _merkleProofs [][][32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "redeemMany", _account, _tranches, _tokenIds, _merkleProofs)
}

// RedeemMany is a paid mutator transaction binding the contract method 0xd58ea88e.
//
// Solidity: function redeemMany(address _account, uint256[] _tranches, uint256[] _tokenIds, bytes32[][] _merkleProofs) returns()
func (_Contracts *ContractsSession) RedeemMany(_account common.Address, _tranches []*big.Int, _tokenIds []*big.Int, _merkleProofs [][][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.RedeemMany(&_Contracts.TransactOpts, _account, _tranches, _tokenIds, _merkleProofs)
}

// RedeemMany is a paid mutator transaction binding the contract method 0xd58ea88e.
//
// Solidity: function redeemMany(address _account, uint256[] _tranches, uint256[] _tokenIds, bytes32[][] _merkleProofs) returns()
func (_Contracts *ContractsTransactorSession) RedeemMany(_account common.Address, _tranches []*big.Int, _tokenIds []*big.Int, _merkleProofs [][][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.RedeemMany(&_Contracts.TransactOpts, _account, _tranches, _tokenIds, _merkleProofs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// SeedNewTranche is a paid mutator transaction binding the contract method 0x857f219c.
//
// Solidity: function seedNewTranche(bytes32 _merkleRoot) returns(uint256 trancheId)
func (_Contracts *ContractsTransactor) SeedNewTranche(opts *bind.TransactOpts, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "seedNewTranche", _merkleRoot)
}

// SeedNewTranche is a paid mutator transaction binding the contract method 0x857f219c.
//
// Solidity: function seedNewTranche(bytes32 _merkleRoot) returns(uint256 trancheId)
func (_Contracts *ContractsSession) SeedNewTranche(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.SeedNewTranche(&_Contracts.TransactOpts, _merkleRoot)
}

// SeedNewTranche is a paid mutator transaction binding the contract method 0x857f219c.
//
// Solidity: function seedNewTranche(bytes32 _merkleRoot) returns(uint256 trancheId)
func (_Contracts *ContractsTransactorSession) SeedNewTranche(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.SeedNewTranche(&_Contracts.TransactOpts, _merkleRoot)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// ContractsClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Contracts contract.
type ContractsClaimedIterator struct {
	Event *ContractsClaimed // Event containing the contract specifics and raw log

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
func (it *ContractsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsClaimed)
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
		it.Event = new(ContractsClaimed)
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
func (it *ContractsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsClaimed represents a Claimed event raised by the Contracts contract.
type ContractsClaimed struct {
	Claimant common.Address
	Week     *big.Int
	Balance  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: event Claimed(address claimant, uint256 week, uint256 balance)
func (_Contracts *ContractsFilterer) FilterClaimed(opts *bind.FilterOpts) (*ContractsClaimedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &ContractsClaimedIterator{contract: _Contracts.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: event Claimed(address claimant, uint256 week, uint256 balance)
func (_Contracts *ContractsFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *ContractsClaimed) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsClaimed)
				if err := _Contracts.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: event Claimed(address claimant, uint256 week, uint256 balance)
func (_Contracts *ContractsFilterer) ParseClaimed(log types.Log) (*ContractsClaimed, error) {
	event := new(ContractsClaimed)
	if err := _Contracts.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
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
		it.Event = new(ContractsOwnershipTransferred)
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
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTrancheAddedIterator is returned from FilterTrancheAdded and is used to iterate over the raw logs and unpacked data for TrancheAdded events raised by the Contracts contract.
type ContractsTrancheAddedIterator struct {
	Event *ContractsTrancheAdded // Event containing the contract specifics and raw log

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
func (it *ContractsTrancheAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTrancheAdded)
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
		it.Event = new(ContractsTrancheAdded)
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
func (it *ContractsTrancheAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTrancheAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTrancheAdded represents a TrancheAdded event raised by the Contracts contract.
type ContractsTrancheAdded struct {
	Tranche    *big.Int
	MerkleRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTrancheAdded is a free log retrieval operation binding the contract event 0x899089a7b39272321d9f5253a8c187ea3b7507701ebb523db8ae97763f0f7db8.
//
// Solidity: event TrancheAdded(uint256 tranche, bytes32 merkleRoot)
func (_Contracts *ContractsFilterer) FilterTrancheAdded(opts *bind.FilterOpts) (*ContractsTrancheAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "TrancheAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsTrancheAddedIterator{contract: _Contracts.contract, event: "TrancheAdded", logs: logs, sub: sub}, nil
}

// WatchTrancheAdded is a free log subscription operation binding the contract event 0x899089a7b39272321d9f5253a8c187ea3b7507701ebb523db8ae97763f0f7db8.
//
// Solidity: event TrancheAdded(uint256 tranche, bytes32 merkleRoot)
func (_Contracts *ContractsFilterer) WatchTrancheAdded(opts *bind.WatchOpts, sink chan<- *ContractsTrancheAdded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "TrancheAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTrancheAdded)
				if err := _Contracts.contract.UnpackLog(event, "TrancheAdded", log); err != nil {
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

// ParseTrancheAdded is a log parse operation binding the contract event 0x899089a7b39272321d9f5253a8c187ea3b7507701ebb523db8ae97763f0f7db8.
//
// Solidity: event TrancheAdded(uint256 tranche, bytes32 merkleRoot)
func (_Contracts *ContractsFilterer) ParseTrancheAdded(log types.Log) (*ContractsTrancheAdded, error) {
	event := new(ContractsTrancheAdded)
	if err := _Contracts.contract.UnpackLog(event, "TrancheAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTrancheExpiredIterator is returned from FilterTrancheExpired and is used to iterate over the raw logs and unpacked data for TrancheExpired events raised by the Contracts contract.
type ContractsTrancheExpiredIterator struct {
	Event *ContractsTrancheExpired // Event containing the contract specifics and raw log

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
func (it *ContractsTrancheExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTrancheExpired)
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
		it.Event = new(ContractsTrancheExpired)
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
func (it *ContractsTrancheExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTrancheExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTrancheExpired represents a TrancheExpired event raised by the Contracts contract.
type ContractsTrancheExpired struct {
	Tranche *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTrancheExpired is a free log retrieval operation binding the contract event 0xcc071cbd9ae50a4c78d1153b76bd2d46ba8d4c7662842718ec3de1d67a144daf.
//
// Solidity: event TrancheExpired(uint256 tranche)
func (_Contracts *ContractsFilterer) FilterTrancheExpired(opts *bind.FilterOpts) (*ContractsTrancheExpiredIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "TrancheExpired")
	if err != nil {
		return nil, err
	}
	return &ContractsTrancheExpiredIterator{contract: _Contracts.contract, event: "TrancheExpired", logs: logs, sub: sub}, nil
}

// WatchTrancheExpired is a free log subscription operation binding the contract event 0xcc071cbd9ae50a4c78d1153b76bd2d46ba8d4c7662842718ec3de1d67a144daf.
//
// Solidity: event TrancheExpired(uint256 tranche)
func (_Contracts *ContractsFilterer) WatchTrancheExpired(opts *bind.WatchOpts, sink chan<- *ContractsTrancheExpired) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "TrancheExpired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTrancheExpired)
				if err := _Contracts.contract.UnpackLog(event, "TrancheExpired", log); err != nil {
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

// ParseTrancheExpired is a log parse operation binding the contract event 0xcc071cbd9ae50a4c78d1153b76bd2d46ba8d4c7662842718ec3de1d67a144daf.
//
// Solidity: event TrancheExpired(uint256 tranche)
func (_Contracts *ContractsFilterer) ParseTrancheExpired(log types.Log) (*ContractsTrancheExpired, error) {
	event := new(ContractsTrancheExpired)
	if err := _Contracts.contract.UnpackLog(event, "TrancheExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
