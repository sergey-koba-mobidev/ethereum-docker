# Running samples

- `dc run builder go run /contracts/samples/get_balance.go`

# Generating GO versions of contracts

- read this [https://zupzup.org/eth-smart-contracts-go/](https://zupzup.org/eth-smart-contracts-go/)
- run `dc run builder abigen --sol=/contracts/samples/greeter/greeter.sol --pkg=main --out=/contracts/samples/greeter/greeter.go`