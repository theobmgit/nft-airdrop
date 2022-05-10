package main

import (
	"encoding/hex"
	"github.com/cbergoon/merkletree"
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"log"
	"math/big"
)

type TokenReceiver struct {
	tokenId *big.Int
	address common.Address
}

//CalculateHash hashes the values of a TokenReceiver
func (t TokenReceiver) CalculateHash() ([]byte, error) {
	return solsha3.SoliditySHA3(
		[]string{"uint256", "address"},

		[]interface{}{
			t.tokenId,
			t.address,
		},
	), nil
}

//Equals tests for equality of two Contents
func (t TokenReceiver) Equals(other merkletree.Content) (bool, error) {
	return t.tokenId.Cmp(other.(TokenReceiver).tokenId) == 0, nil
}

func main() {
	//Build list of Content to build tree
	var list []merkletree.Content
	list = append(list, TokenReceiver{
		tokenId: big.NewInt(4),
		address: common.HexToAddress("0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"),
	})
	list = append(list, TokenReceiver{
		tokenId: big.NewInt(5),
		address: common.HexToAddress("0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"),
	})
	list = append(list, TokenReceiver{
		tokenId: big.NewInt(6),
		address: common.HexToAddress("0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB"),
	})

	//Create a new Merkle Tree from the list of Content
	t, err := merkletree.NewTree(list)
	if err != nil {
		log.Fatal(err)
	}

	//Get the Merkle Root of the tree
	mr := t.MerkleRoot()
	log.Println(hex.EncodeToString(mr))
}
