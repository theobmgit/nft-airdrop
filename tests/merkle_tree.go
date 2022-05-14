package main

import (
	"bytes"
	"encoding/hex"
	"github.com/cbergoon/merkletree"
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
	"sort"
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
	list = append(list, TokenReceiver{
		tokenId: big.NewInt(7),
		address: common.HexToAddress("0x617F2E2fD72FD9D5503197092aC168c91465E7f2"),
	})
	list = append(list, TokenReceiver{
		tokenId: big.NewInt(8),
		address: common.HexToAddress("0x17F6AD8Ef982297579C203069C1DbfFE4348c372"),
	})
	list = append(list, TokenReceiver{
		tokenId: big.NewInt(9),
		address: common.HexToAddress("0x5c6B0f7Bf3E7ce046039Bd8FABdfD3f9F5021678"),
	})
	list = append(list, TokenReceiver{
		tokenId: big.NewInt(10),
		address: common.HexToAddress("0x03C6FcED478cBbC9a4FAB34eF9f40767739D1Ff7"),
	})
	list = append(list, TokenReceiver{
		tokenId: big.NewInt(11),
		address: common.HexToAddress("0x1aE0EA34a72D944a8C7603FfB3eC30a6669E454C"),
	})

	sort.Slice(list, func(i, j int) bool {
		hashI, _ := list[i].CalculateHash()
		hashJ, _ := list[j].CalculateHash()
		return bytes.Compare(hashI, hashJ) == -1
	})

	//Create a new Merkle Tree from the list of Content
	t, err := merkletree.NewTreeWithHashStrategy(list, sha3.NewLegacyKeccak256)
	if err != nil {
		log.Fatal(err)
	}

	//Get the Merkle Root of the tree
	mr := t.MerkleRoot()
	log.Println("Root:", hex.EncodeToString(mr))

	for i, leaf := range t.Leafs {
		log.Println("Leaf:", i)
		log.Println(hex.EncodeToString(leaf.Hash))
	}

	for i, _ := range list {
		path, idx, err := t.GetMerklePath(list[i])
		if err != nil {
			return
		}
		log.Println("Index:", idx)
		for _, hash := range path {
			log.Println("Path: ", hex.EncodeToString(hash))
		}
	}

	log.Println()
	log.Println("Address")
	path, idx, err := t.GetMerklePath(TokenReceiver{
		tokenId: big.NewInt(9),
		address: common.HexToAddress("0x5c6B0f7Bf3E7ce046039Bd8FABdfD3f9F5021678"),
	})
	if err != nil {
		return
	}
	log.Println("Index:", idx)
	for _, hash := range path {
		log.Println("Path: ", hex.EncodeToString(hash))
	}
}
