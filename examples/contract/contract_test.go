package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/chenzhijie/go-web3/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
	"time"
)

func TestDeployContract(t *testing.T) {
	// change to your rpc provider
	client, err := ethclient.Dial(eth.RPC_URL)
	if err != nil {
		panic(err)
	}
	privateKey, err := crypto.HexToECDSA("c1f14e4132c1858b390ff169dd045082b9ba1ca022de641e7aa24b0322510499")
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	chainId, err := client.ChainID(context.Background())
	//set chainID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	address, tx, instance, err := DeployMain(auth, client)
	if err != nil {
		panic(err)
	}
	fmt.Println("contract:" + address.Hex())
	fmt.Println("deploy hash:" + tx.Hash().String())
	_, _ = instance, tx
}

func TestCallContract(t *testing.T) {

	// change to your rpc provider
	client, err := ethclient.Dial(eth.RPC_URL)
	if err != nil {
		panic(err)
	}
	//load contract 0x771B03a85843907B89B66dbbCD90E1b8761a4fa3
	main, err := NewMain(common.HexToAddress("0x797C88225547126879aA4EF444A66D627E402366"), client)
	if err != nil {
		panic(err)
	}
	num, err := main.Retrieve(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("num:%d", num)
	fmt.Println("")

	//write contract
	privateKey, err := crypto.HexToECDSA("c1f14e4132c1858b390ff169dd045082b9ba1ca022de641e7aa24b0322510499")
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	chainId, err := client.ChainID(context.Background())
	//set chainID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	main.Store(auth, new(big.Int).SetInt64(1000))

	//sleep
	time.Sleep(time.Duration(5) * time.Second)

	newNum, err := main.Retrieve(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("newNum:%d", newNum)
	fmt.Println("")

}
