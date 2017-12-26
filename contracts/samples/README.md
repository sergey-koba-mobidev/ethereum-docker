# Running samples

- `dc run builder go run /contracts/samples/get_balance.go`

# Generating GO versions of contracts

- read this [https://zupzup.org/eth-smart-contracts-go/](https://zupzup.org/eth-smart-contracts-go/)
- run `dc run builder abigen --sol=/contracts/samples/greeter/greeter.sol --pkg=main --out=/contracts/samples/greeter/greeter.go`

# Deploy contract
- run `dc run builder bash -c "cd /contracts/samples/greeter/ && go build . && ./greeter deploy"`

# Greeter
- run `dc run builder abigen --sol=/contracts/samples/greeter/greeter.sol --pkg=main --out=/contracts/samples/greeter/greeter.go` to build go version of contract
- run `dc up -d` to launch Ethereum blockchain
- run `docker exec -it bootstrap geth --datadir=~/.ethereum/devchain attach` to attach to geth on bootstrap node, start mining by `miners.start()` and exit `exit`
- deploy smart contract `dc run builder bash -c "cd /contracts/samples/greeter/ && go build . && ./greeter deploy"`, note the contract's address
- get greeting message from smart contract `dc run builder bash -c "cd /contracts/samples/greeter/ && go build . && ./greeter greet 0xFF37a57B8D373518aBE222Db1077eD9A968a5FDf"` replace `0xFF37a57B8D373518aBE222Db1077eD9A968a5FDf` with your contract address