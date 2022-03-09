package account

import (
	"encoding/hex"
	"fmt"
	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

func TestAccount(t *testing.T) {
	//get account
	pv, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	privateKey := hex.EncodeToString(crypto.FromECDSA(pv))
	//c1f14e4132c1858b390ff169dd045082b9ba1ca022de641e7aa24b0322510499
	fmt.Println("privateKey: ", privateKey)
	// change to your rpc provider
	web3, err := web3.NewWeb3(eth.RPC_URL)
	if err != nil {
		panic(err)
	}
	err = web3.Eth.SetAccount(privateKey)
	if err != nil {
		panic(err)
	}
	//get address
	//0x901F21a0a09536F24feF8f5565eBcacB7aC5DE31
	fmt.Println("address: ", web3.Eth.Address())
}

//get tkm https://www.thinkiumdev.net/DApp%20Development/Faucet.html

func TestSendTransaction(t *testing.T) {
	// change to your rpc provider
	web3, err := web3.NewWeb3(eth.RPC_URL)
	if err != nil {
		panic(err)
	}
	web3.Eth.SetAccount("c1f14e4132c1858b390ff169dd045082b9ba1ca022de641e7aa24b0322510499")
	balance, err := web3.Eth.GetBalance(web3.Eth.Address(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("balance: ", balance)
	nonce, err := web3.Eth.GetNonce(web3.Eth.Address(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Latest nonce: ", nonce)
	//transfer
	//set chainId
	web3.Eth.SetChainId(50001)
	// transfer 1 tkm from 0x901F21a0a09536F24feF8f5565eBcacB7aC5DE31 to 0x90021a2CcA84a0611363ec1f9ccb36B4761DE08E
	hash, err := web3.Eth.SendRawEIP1559Transaction(common.HexToAddress("0x90021a2CcA84a0611363ec1f9ccb36B4761DE08E"),
		new(big.Int).SetInt64(1000000000000000000), 300000, new(big.Int).SetInt64(300000), new(big.Int).SetInt64(300000), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("hash: ", hash)
}
