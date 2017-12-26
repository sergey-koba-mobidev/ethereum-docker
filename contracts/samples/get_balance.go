package main

import (
	"fmt"
	//"log"
	//"math/big"
	//"time"
	"context"
	"os"
	"github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	//"github.com/ethereum/go-ethereum/core"
	//"github.com/ethereum/go-ethereum/crypto"
)

func main() {
    client, err := ethclient.Dial("http://bootstrap:8545")

	if err != nil {
		fmt.Printf("Cannot connect to node %v", err)
		os.Exit(1)
	}

	var address = common.HexToAddress("0x007ccffb7916f37f7aeef05e8096ecfbe55afc2f")

	var context = context.Background()
    balance, err := client.BalanceAt(context, address, nil)

    if (err != nil) {
    	fmt.Printf("%v", err)
    	os.Exit(1)
    }

    fmt.Printf("Balance : %v\n", balance)
}