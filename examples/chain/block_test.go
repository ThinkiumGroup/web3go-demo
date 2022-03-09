package main

import (
	"fmt"
	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/eth"
	"testing"
)

func TestGetBlockByNumber(t *testing.T) {
	// change to your rpc provider
	web3, err := web3.NewWeb3(eth.RPC_URL)
	if err != nil {
		panic(err)
	}
	blockNumber, err := web3.Eth.GetBlockNumber()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current block number: ", blockNumber)
}
