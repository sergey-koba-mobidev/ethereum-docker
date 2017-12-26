package main

import (
	"os"
	"fmt"
	"strings"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const key = "{\"address\":\"007ccffb7916f37f7aeef05e8096ecfbe55afc2f\",\"Crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"44600de3be08153de243f0de344b9841c774a2fcd5b7f3406b9ea9a42c8be97b\",\"cipherparams\":{\"iv\":\"9ca680df904bc0138d2d52f9c432a797\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"a6524c4ef60060476f2e47dea7f5945aac8159019d3bf5275a683c1970d03c40\"},\"mac\":\"4daf73bc6da7f75beb9a8ea28896c5f1f074fb8cd54880fa9a4e34700dc0ee00\"},\"id\":\"2615dd6a-6d73-4b05-a02d-257b0fd28d3a\",\"version\":3}"

func main() {
	conn, err := ethclient.Dial("http://bootstrap:8545")

	if err != nil {
		fmt.Printf("Cannot connect to node %v", err)
		os.Exit(1)
	}

	auth, errAuth := bind.NewTransactor(strings.NewReader(key), "")
	if errAuth != nil {
		log.Fatalf("Failed to create new transactor : %v", errAuth)
	}

	action := os.Args[1]
	fmt.Printf("Executing action: %v\n", action)

	switch action {
	case "deploy":
		DoDeploy(auth, conn)
	case "greet":
		hexAddr := os.Args[2]
		DoGreet(conn, common.HexToAddress(hexAddr))
	}

}

func DoDeploy(auth *bind.TransactOpts, conn bind.ContractBackend) (common.Address, *Greeter, error) {
	//deploy contract
	addr, _, contract, err := DeployGreeter(auth, conn, "Hello I am first deployed SC! Gratz!")
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}

	fmt.Printf("Contract address: %v\n", addr.Hex())
	fmt.Printf("Contract: %v\n", contract)
	return addr, contract, err
}

func DoGreet(conn bind.ContractBackend, addr common.Address) (string, error) {
	// Instantiate the contract and display its greeting
	greeter, err := NewGreeter(addr, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Greeter contract: %v", err)
	}
	greeting, err := greeter.Greet(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve greeter greeting: %v", err)
	}
	fmt.Println("Contract Greeting:", greeting)
	return greeting, err
}