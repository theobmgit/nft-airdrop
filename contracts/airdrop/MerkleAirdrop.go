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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"week\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"RemovedFunder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tranche\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"TrancheAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tranche\",\"type\":\"uint256\"}],\"name\":\"TrancheExpired\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tranche\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimWeek\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityProvider\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tranches\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_balances\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_merkleProofs\",\"type\":\"bytes32[][]\"}],\"name\":\"claimWeeks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_trancheId\",\"type\":\"uint256\"}],\"name\":\"expireTranche\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_totalAllocation\",\"type\":\"uint256\"}],\"name\":\"seedNewAllocations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"trancheId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tranches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tranche\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"verifyClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620021d5380380620021d58339818101604052810190620000379190620001e9565b620000576200004b6200009f60201b60201c565b620000a760201b60201c565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200021b565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200019d8262000170565b9050919050565b6000620001b18262000190565b9050919050565b620001c381620001a4565b8114620001cf57600080fd5b50565b600081519050620001e381620001b8565b92915050565b6000602082840312156200020257620002016200016b565b5b60006200021284828501620001d2565b91505092915050565b611faa806200022b6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063ba59371411610071578063ba5937141461018d578063d5efd20a146101a9578063eb0d07f5146101c5578063ebf04917146101f5578063f2fde38b14610213578063fc0c546a1461022f576100b4565b8063120aa877146100b957806358b4e4b4146100e957806365ef53b114610105578063715018a61461013557806371c5ecb11461013f5780638da5cb5b1461016f575b600080fd5b6100d360048036038101906100ce9190611045565b61024d565b6040516100e091906110a0565b60405180910390f35b61010360048036038101906100fe919061124a565b61027c565b005b61011f600480360381019061011a91906112cd565b610298565b60405161012c919061131c565b60405180910390f35b61013d6103df565b005b61015960048036038101906101549190611337565b610467565b6040516101669190611373565b60405180910390f35b61017761047f565b604051610184919061139d565b60405180910390f35b6101a760048036038101906101a2919061155c565b6104a8565b005b6101c360048036038101906101be9190611337565b6105b8565b005b6101df60048036038101906101da919061124a565b610689565b6040516101ec91906110a0565b60405180910390f35b6101fd6106a1565b60405161020a919061131c565b60405180910390f35b61022d60048036038101906102289190611617565b6106a7565b005b61023761079e565b60405161024491906116a3565b60405180910390f35b60036020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b610288848484846107c4565b610292848361099b565b50505050565b60006102a2610a35565b73ffffffffffffffffffffffffffffffffffffffff166102c061047f565b73ffffffffffffffffffffffffffffffffffffffff1614610316576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030d9061171b565b60405180910390fd5b610365333084600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610a3d909392919063ffffffff16565b60045490508260026000838152602001908152602001600020819055506103986001600454610ac690919063ffffffff16565b6004819055507f5c8770684b8f82e9ade880fb05ccfb53c969170cd40e9746a3703f241c9023ec8184846040516103d19392919061173b565b60405180910390a192915050565b6103e7610a35565b73ffffffffffffffffffffffffffffffffffffffff1661040561047f565b73ffffffffffffffffffffffffffffffffffffffff161461045b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104529061171b565b60405180910390fd5b6104656000610adc565b565b60026020528060005260406000206000915090505481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000835190508251811480156104be5750815181145b6104fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104f4906117be565b60405180910390fd5b6000805b828110156105a557610563878783815181106105205761051f6117de565b5b602002602001015187848151811061053b5761053a6117de565b5b6020026020010151878581518110610556576105556117de565b5b60200260200101516107c4565b610590858281518110610579576105786117de565b5b602002602001015183610ac690919063ffffffff16565b9150808061059d9061183c565b915050610501565b506105b0868261099b565b505050505050565b6105c0610a35565b73ffffffffffffffffffffffffffffffffffffffff166105de61047f565b73ffffffffffffffffffffffffffffffffffffffff1614610634576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062b9061171b565b60405180910390fd5b6000801b60026000838152602001908152602001600020819055507fcc071cbd9ae50a4c78d1153b76bd2d46ba8d4c7662842718ec3de1d67a144daf8160405161067e919061131c565b60405180910390a150565b600061069785858585610ba0565b9050949350505050565b60045481565b6106af610a35565b73ffffffffffffffffffffffffffffffffffffffff166106cd61047f565b73ffffffffffffffffffffffffffffffffffffffff1614610723576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071a9061171b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610792576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610789906118f6565b60405180910390fd5b61079b81610adc565b50565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6004548310610808576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ff90611962565b60405180910390fd5b6003600084815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156108a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089d906119ce565b60405180910390fd5b6108b284848484610ba0565b6108f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e890611a3a565b60405180910390fd5b60016003600085815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a84848460405161098d93929190611a5a565b60405180910390a150505050565b60008111156109f6576109f18282600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610bf79092919063ffffffff16565b610a31565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2890611b03565b60405180910390fd5b5050565b600033905090565b610ac0846323b872dd60e01b858585604051602401610a5e93929190611b23565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610c7d565b50505050565b60008183610ad49190611b5a565b905092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808584604051602001610bb6929190611c19565b604051602081830303815290604052805190602001209050610bec83600260008881526020019081526020016000205483610d44565b915050949350505050565b610c788363a9059cbb60e01b8484604051602401610c16929190611c45565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610c7d565b505050565b6000610cdf826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16610d5b9092919063ffffffff16565b9050600081511115610d3f5780806020019051810190610cff9190611c9a565b610d3e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3590611d39565b60405180910390fd5b5b505050565b600082610d518584610d73565b1490509392505050565b6060610d6a8484600085610de8565b90509392505050565b60008082905060005b8451811015610ddd576000858281518110610d9a57610d996117de565b5b60200260200101519050808311610dbc57610db58382610efc565b9250610dc9565b610dc68184610efc565b92505b508080610dd59061183c565b915050610d7c565b508091505092915050565b606082471015610e2d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e2490611dcb565b60405180910390fd5b610e3685610f13565b610e75576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e6c90611e37565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610e9e9190611ed1565b60006040518083038185875af1925050503d8060008114610edb576040519150601f19603f3d011682016040523d82523d6000602084013e610ee0565b606091505b5091509150610ef0828286610f36565b92505050949350505050565b600082600052816020526040600020905092915050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60608315610f4657829050610f96565b600083511115610f595782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f8d9190611f2c565b60405180910390fd5b9392505050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610fc481610fb1565b8114610fcf57600080fd5b50565b600081359050610fe181610fbb565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061101282610fe7565b9050919050565b61102281611007565b811461102d57600080fd5b50565b60008135905061103f81611019565b92915050565b6000806040838503121561105c5761105b610fa7565b5b600061106a85828601610fd2565b925050602061107b85828601611030565b9150509250929050565b60008115159050919050565b61109a81611085565b82525050565b60006020820190506110b56000830184611091565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611109826110c0565b810181811067ffffffffffffffff82111715611128576111276110d1565b5b80604052505050565b600061113b610f9d565b90506111478282611100565b919050565b600067ffffffffffffffff821115611167576111666110d1565b5b602082029050602081019050919050565b600080fd5b6000819050919050565b6111908161117d565b811461119b57600080fd5b50565b6000813590506111ad81611187565b92915050565b60006111c66111c18461114c565b611131565b905080838252602082019050602084028301858111156111e9576111e8611178565b5b835b8181101561121257806111fe888261119e565b8452602084019350506020810190506111eb565b5050509392505050565b600082601f830112611231576112306110bb565b5b81356112418482602086016111b3565b91505092915050565b6000806000806080858703121561126457611263610fa7565b5b600061127287828801611030565b945050602061128387828801610fd2565b935050604061129487828801610fd2565b925050606085013567ffffffffffffffff8111156112b5576112b4610fac565b5b6112c18782880161121c565b91505092959194509250565b600080604083850312156112e4576112e3610fa7565b5b60006112f28582860161119e565b925050602061130385828601610fd2565b9150509250929050565b61131681610fb1565b82525050565b6000602082019050611331600083018461130d565b92915050565b60006020828403121561134d5761134c610fa7565b5b600061135b84828501610fd2565b91505092915050565b61136d8161117d565b82525050565b60006020820190506113886000830184611364565b92915050565b61139781611007565b82525050565b60006020820190506113b2600083018461138e565b92915050565b600067ffffffffffffffff8211156113d3576113d26110d1565b5b602082029050602081019050919050565b60006113f76113f2846113b8565b611131565b9050808382526020820190506020840283018581111561141a57611419611178565b5b835b81811015611443578061142f8882610fd2565b84526020840193505060208101905061141c565b5050509392505050565b600082601f830112611462576114616110bb565b5b81356114728482602086016113e4565b91505092915050565b600067ffffffffffffffff821115611496576114956110d1565b5b602082029050602081019050919050565b60006114ba6114b58461147b565b611131565b905080838252602082019050602084028301858111156114dd576114dc611178565b5b835b8181101561152457803567ffffffffffffffff811115611502576115016110bb565b5b80860161150f898261121c565b855260208501945050506020810190506114df565b5050509392505050565b600082601f830112611543576115426110bb565b5b81356115538482602086016114a7565b91505092915050565b6000806000806080858703121561157657611575610fa7565b5b600061158487828801611030565b945050602085013567ffffffffffffffff8111156115a5576115a4610fac565b5b6115b18782880161144d565b935050604085013567ffffffffffffffff8111156115d2576115d1610fac565b5b6115de8782880161144d565b925050606085013567ffffffffffffffff8111156115ff576115fe610fac565b5b61160b8782880161152e565b91505092959194509250565b60006020828403121561162d5761162c610fa7565b5b600061163b84828501611030565b91505092915050565b6000819050919050565b600061166961166461165f84610fe7565b611644565b610fe7565b9050919050565b600061167b8261164e565b9050919050565b600061168d82611670565b9050919050565b61169d81611682565b82525050565b60006020820190506116b86000830184611694565b92915050565b600082825260208201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006117056020836116be565b9150611710826116cf565b602082019050919050565b60006020820190508181036000830152611734816116f8565b9050919050565b6000606082019050611750600083018661130d565b61175d6020830185611364565b61176a604083018461130d565b949350505050565b7f4d69736d61746368696e6720696e707574730000000000000000000000000000600082015250565b60006117a86012836116be565b91506117b382611772565b602082019050919050565b600060208201905081810360008301526117d78161179b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061184782610fb1565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036118795761187861180d565b5b600182019050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006118e06026836116be565b91506118eb82611884565b604082019050919050565b6000602082019050818103600083015261190f816118d3565b9050919050565b7f5765656b2063616e6e6f7420626520696e207468652066757475726500000000600082015250565b600061194c601c836116be565b915061195782611916565b602082019050919050565b6000602082019050818103600083015261197b8161193f565b9050919050565b7f4c502068617320616c726561647920636c61696d656400000000000000000000600082015250565b60006119b86016836116be565b91506119c382611982565b602082019050919050565b600060208201905081810360008301526119e7816119ab565b9050919050565b7f496e636f7272656374206d65726b6c652070726f6f6600000000000000000000600082015250565b6000611a246016836116be565b9150611a2f826119ee565b602082019050919050565b60006020820190508181036000830152611a5381611a17565b9050919050565b6000606082019050611a6f600083018661138e565b611a7c602083018561130d565b611a89604083018461130d565b949350505050565b7f4e6f2062616c616e636520776f756c64206265207472616e736665727265642060008201527f2d206e6f7420676f696e6720746f20776173746520796f757220676173000000602082015250565b6000611aed603d836116be565b9150611af882611a91565b604082019050919050565b60006020820190508181036000830152611b1c81611ae0565b9050919050565b6000606082019050611b38600083018661138e565b611b45602083018561138e565b611b52604083018461130d565b949350505050565b6000611b6582610fb1565b9150611b7083610fb1565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611ba557611ba461180d565b5b828201905092915050565b60008160601b9050919050565b6000611bc882611bb0565b9050919050565b6000611bda82611bbd565b9050919050565b611bf2611bed82611007565b611bcf565b82525050565b6000819050919050565b611c13611c0e82610fb1565b611bf8565b82525050565b6000611c258285611be1565b601482019150611c358284611c02565b6020820191508190509392505050565b6000604082019050611c5a600083018561138e565b611c67602083018461130d565b9392505050565b611c7781611085565b8114611c8257600080fd5b50565b600081519050611c9481611c6e565b92915050565b600060208284031215611cb057611caf610fa7565b5b6000611cbe84828501611c85565b91505092915050565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b6000611d23602a836116be565b9150611d2e82611cc7565b604082019050919050565b60006020820190508181036000830152611d5281611d16565b9050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b6000611db56026836116be565b9150611dc082611d59565b604082019050919050565b60006020820190508181036000830152611de481611da8565b9050919050565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b6000611e21601d836116be565b9150611e2c82611deb565b602082019050919050565b60006020820190508181036000830152611e5081611e14565b9050919050565b600081519050919050565b600081905092915050565b60005b83811015611e8b578082015181840152602081019050611e70565b83811115611e9a576000848401525b50505050565b6000611eab82611e57565b611eb58185611e62565b9350611ec5818560208601611e6d565b80840191505092915050565b6000611edd8284611ea0565b915081905092915050565b600081519050919050565b6000611efe82611ee8565b611f0881856116be565b9350611f18818560208601611e6d565b611f21816110c0565b840191505092915050565b60006020820190508181036000830152611f468184611ef3565b90509291505056fea264697066735822122094e824851149a2bbff96e0570e095424ea5282a834cee8c0fb8df22feb1b81d964736f6c637828302e382e31342d646576656c6f702e323032322e342e32362b636f6d6d69742e66626563646265370059",
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

// VerifyClaim is a free data retrieval call binding the contract method 0xeb0d07f5.
//
// Solidity: function verifyClaim(address _liquidityProvider, uint256 _tranche, uint256 _balance, bytes32[] _merkleProof) view returns(bool valid)
func (_Contracts *ContractsCaller) VerifyClaim(opts *bind.CallOpts, _liquidityProvider common.Address, _tranche *big.Int, _balance *big.Int, _merkleProof [][32]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "verifyClaim", _liquidityProvider, _tranche, _balance, _merkleProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClaim is a free data retrieval call binding the contract method 0xeb0d07f5.
//
// Solidity: function verifyClaim(address _liquidityProvider, uint256 _tranche, uint256 _balance, bytes32[] _merkleProof) view returns(bool valid)
func (_Contracts *ContractsSession) VerifyClaim(_liquidityProvider common.Address, _tranche *big.Int, _balance *big.Int, _merkleProof [][32]byte) (bool, error) {
	return _Contracts.Contract.VerifyClaim(&_Contracts.CallOpts, _liquidityProvider, _tranche, _balance, _merkleProof)
}

// VerifyClaim is a free data retrieval call binding the contract method 0xeb0d07f5.
//
// Solidity: function verifyClaim(address _liquidityProvider, uint256 _tranche, uint256 _balance, bytes32[] _merkleProof) view returns(bool valid)
func (_Contracts *ContractsCallerSession) VerifyClaim(_liquidityProvider common.Address, _tranche *big.Int, _balance *big.Int, _merkleProof [][32]byte) (bool, error) {
	return _Contracts.Contract.VerifyClaim(&_Contracts.CallOpts, _liquidityProvider, _tranche, _balance, _merkleProof)
}

// ClaimWeek is a paid mutator transaction binding the contract method 0x58b4e4b4.
//
// Solidity: function claimWeek(address _liquidityProvider, uint256 _tranche, uint256 _balance, bytes32[] _merkleProof) returns()
func (_Contracts *ContractsTransactor) ClaimWeek(opts *bind.TransactOpts, _liquidityProvider common.Address, _tranche *big.Int, _balance *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "claimWeek", _liquidityProvider, _tranche, _balance, _merkleProof)
}

// ClaimWeek is a paid mutator transaction binding the contract method 0x58b4e4b4.
//
// Solidity: function claimWeek(address _liquidityProvider, uint256 _tranche, uint256 _balance, bytes32[] _merkleProof) returns()
func (_Contracts *ContractsSession) ClaimWeek(_liquidityProvider common.Address, _tranche *big.Int, _balance *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.ClaimWeek(&_Contracts.TransactOpts, _liquidityProvider, _tranche, _balance, _merkleProof)
}

// ClaimWeek is a paid mutator transaction binding the contract method 0x58b4e4b4.
//
// Solidity: function claimWeek(address _liquidityProvider, uint256 _tranche, uint256 _balance, bytes32[] _merkleProof) returns()
func (_Contracts *ContractsTransactorSession) ClaimWeek(_liquidityProvider common.Address, _tranche *big.Int, _balance *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.ClaimWeek(&_Contracts.TransactOpts, _liquidityProvider, _tranche, _balance, _merkleProof)
}

// ClaimWeeks is a paid mutator transaction binding the contract method 0xba593714.
//
// Solidity: function claimWeeks(address _liquidityProvider, uint256[] _tranches, uint256[] _balances, bytes32[][] _merkleProofs) returns()
func (_Contracts *ContractsTransactor) ClaimWeeks(opts *bind.TransactOpts, _liquidityProvider common.Address, _tranches []*big.Int, _balances []*big.Int, _merkleProofs [][][32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "claimWeeks", _liquidityProvider, _tranches, _balances, _merkleProofs)
}

// ClaimWeeks is a paid mutator transaction binding the contract method 0xba593714.
//
// Solidity: function claimWeeks(address _liquidityProvider, uint256[] _tranches, uint256[] _balances, bytes32[][] _merkleProofs) returns()
func (_Contracts *ContractsSession) ClaimWeeks(_liquidityProvider common.Address, _tranches []*big.Int, _balances []*big.Int, _merkleProofs [][][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.ClaimWeeks(&_Contracts.TransactOpts, _liquidityProvider, _tranches, _balances, _merkleProofs)
}

// ClaimWeeks is a paid mutator transaction binding the contract method 0xba593714.
//
// Solidity: function claimWeeks(address _liquidityProvider, uint256[] _tranches, uint256[] _balances, bytes32[][] _merkleProofs) returns()
func (_Contracts *ContractsTransactorSession) ClaimWeeks(_liquidityProvider common.Address, _tranches []*big.Int, _balances []*big.Int, _merkleProofs [][][32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.ClaimWeeks(&_Contracts.TransactOpts, _liquidityProvider, _tranches, _balances, _merkleProofs)
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

// SeedNewAllocations is a paid mutator transaction binding the contract method 0x65ef53b1.
//
// Solidity: function seedNewAllocations(bytes32 _merkleRoot, uint256 _totalAllocation) returns(uint256 trancheId)
func (_Contracts *ContractsTransactor) SeedNewAllocations(opts *bind.TransactOpts, _merkleRoot [32]byte, _totalAllocation *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "seedNewAllocations", _merkleRoot, _totalAllocation)
}

// SeedNewAllocations is a paid mutator transaction binding the contract method 0x65ef53b1.
//
// Solidity: function seedNewAllocations(bytes32 _merkleRoot, uint256 _totalAllocation) returns(uint256 trancheId)
func (_Contracts *ContractsSession) SeedNewAllocations(_merkleRoot [32]byte, _totalAllocation *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SeedNewAllocations(&_Contracts.TransactOpts, _merkleRoot, _totalAllocation)
}

// SeedNewAllocations is a paid mutator transaction binding the contract method 0x65ef53b1.
//
// Solidity: function seedNewAllocations(bytes32 _merkleRoot, uint256 _totalAllocation) returns(uint256 trancheId)
func (_Contracts *ContractsTransactorSession) SeedNewAllocations(_merkleRoot [32]byte, _totalAllocation *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SeedNewAllocations(&_Contracts.TransactOpts, _merkleRoot, _totalAllocation)
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

// ContractsRemovedFunderIterator is returned from FilterRemovedFunder and is used to iterate over the raw logs and unpacked data for RemovedFunder events raised by the Contracts contract.
type ContractsRemovedFunderIterator struct {
	Event *ContractsRemovedFunder // Event containing the contract specifics and raw log

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
func (it *ContractsRemovedFunderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRemovedFunder)
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
		it.Event = new(ContractsRemovedFunder)
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
func (it *ContractsRemovedFunderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRemovedFunderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRemovedFunder represents a RemovedFunder event raised by the Contracts contract.
type ContractsRemovedFunder struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemovedFunder is a free log retrieval operation binding the contract event 0x716dbcc83324f9116e347cd75111c8f466ede05562baa322dabd8368834f22a7.
//
// Solidity: event RemovedFunder(address indexed _address)
func (_Contracts *ContractsFilterer) FilterRemovedFunder(opts *bind.FilterOpts, _address []common.Address) (*ContractsRemovedFunderIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RemovedFunder", _addressRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRemovedFunderIterator{contract: _Contracts.contract, event: "RemovedFunder", logs: logs, sub: sub}, nil
}

// WatchRemovedFunder is a free log subscription operation binding the contract event 0x716dbcc83324f9116e347cd75111c8f466ede05562baa322dabd8368834f22a7.
//
// Solidity: event RemovedFunder(address indexed _address)
func (_Contracts *ContractsFilterer) WatchRemovedFunder(opts *bind.WatchOpts, sink chan<- *ContractsRemovedFunder, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RemovedFunder", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRemovedFunder)
				if err := _Contracts.contract.UnpackLog(event, "RemovedFunder", log); err != nil {
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

// ParseRemovedFunder is a log parse operation binding the contract event 0x716dbcc83324f9116e347cd75111c8f466ede05562baa322dabd8368834f22a7.
//
// Solidity: event RemovedFunder(address indexed _address)
func (_Contracts *ContractsFilterer) ParseRemovedFunder(log types.Log) (*ContractsRemovedFunder, error) {
	event := new(ContractsRemovedFunder)
	if err := _Contracts.contract.UnpackLog(event, "RemovedFunder", log); err != nil {
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
	Tranche     *big.Int
	MerkleRoot  [32]byte
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTrancheAdded is a free log retrieval operation binding the contract event 0x5c8770684b8f82e9ade880fb05ccfb53c969170cd40e9746a3703f241c9023ec.
//
// Solidity: event TrancheAdded(uint256 tranche, bytes32 merkleRoot, uint256 totalAmount)
func (_Contracts *ContractsFilterer) FilterTrancheAdded(opts *bind.FilterOpts) (*ContractsTrancheAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "TrancheAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsTrancheAddedIterator{contract: _Contracts.contract, event: "TrancheAdded", logs: logs, sub: sub}, nil
}

// WatchTrancheAdded is a free log subscription operation binding the contract event 0x5c8770684b8f82e9ade880fb05ccfb53c969170cd40e9746a3703f241c9023ec.
//
// Solidity: event TrancheAdded(uint256 tranche, bytes32 merkleRoot, uint256 totalAmount)
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

// ParseTrancheAdded is a log parse operation binding the contract event 0x5c8770684b8f82e9ade880fb05ccfb53c969170cd40e9746a3703f241c9023ec.
//
// Solidity: event TrancheAdded(uint256 tranche, bytes32 merkleRoot, uint256 totalAmount)
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
