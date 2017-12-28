package main

import (
	"os"
	"fmt"
	"strings"
	"log"
	"math/big"
	"time"

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
	case "add_product":
		hexAddr := os.Args[2]
		data := os.Args[3]
		AddProduct(auth, conn, common.HexToAddress(hexAddr), data)
	case "get_product":
		hexAddr := os.Args[2]
		uid := os.Args[3]
		GetProduct(auth, conn, common.HexToAddress(hexAddr), uid)
	}

}

func DoDeploy(auth *bind.TransactOpts, conn bind.ContractBackend) (common.Address, *SupplyChain, error) {
	//deploy contract
	addr, _, contract, err := DeploySupplyChain(auth, conn, "Mobile Phones", "Mobile Phones Supply Chain")
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}

	fmt.Printf("Contract address: %v\n", addr.Hex())
	return addr, contract, err
}

func AddProduct(auth *bind.TransactOpts, conn bind.ContractBackend, addr common.Address, data string) (* big.Int, error) {
	// Instantiate the contract and display its greeting
	sc, err := NewSupplyChain(addr, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Greeter contract: %v", err)
	}


	tr, err := sc.AddProduct(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
	}, data) //"{\"name\":\"iPhone 6S\", \"id\": \"7532134\"}"
	if err != nil {
		log.Fatalf("Failed to add product: %v", err)
	}

	fmt.Printf("Transaction: %v\n", tr)

	fmt.Println("Waitting for miners")
	time.Sleep(10 * time.Second)

	productId, _ := sc.NumberOfProducts(nil)

	fmt.Printf("New product UID: %v\n", productId)
	return productId, err
}

func GetProduct(auth *bind.TransactOpts, conn bind.ContractBackend, addr common.Address, uid string) (struct {
	Uid     *big.Int
	State   *big.Int
	Data    string
	Owner   common.Address
	Creator common.Address
}) {
	// Instantiate the contract and display its greeting
	sc, err := NewSupplyChain(addr, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Greeter contract: %v", err)
	}

	n := new(big.Int)
	n, ok := n.SetString(uid, 10)
	if !ok {
		fmt.Println("WRONG UID!")
	}

	info, _ := sc.GetProductInfo(nil, n)
	fmt.Printf("Product Info: %v\n", info)

	return info
}