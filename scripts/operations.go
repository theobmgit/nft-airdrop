package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	nftToken "nft-airdrop/contracts/token"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/ee09af9e74a243c4a7ee638fcf0a9d21")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("fd95304a7c242e89305064a3cdee4fdcd91762d10f54d4f0cd38a9d0b6b1e0bc")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	address := common.HexToAddress("0xeBEC5C95f26Eb92d91e89D1A3B1083D00bc0D25f")
	instance, err := nftToken.NewContracts(address, client)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(instance.Name(nil))

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatalln(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = 300000
	auth.GasPrice = gasPrice

	tx, err := instance.SafeMint(auth, fromAddress, big.NewInt(0))

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
