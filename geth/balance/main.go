package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// infura网关，初始化客户端
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatalf("infura网关%v", err)
	}
	fmt.Println(client)

	// 连接本地ganache RPC主机
	ganacheClient, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatalf("本地ganache%v", err)
	}
	fmt.Println(ganacheClient)

	// 账户地址转换为common.Address类型 以太坊的账户要么是钱包地址 要么是智能合约地址
	address := common.HexToAddress("0x388C818CA8B9251b393131C08a736A67ccB19297")
	fmt.Printf("%s\n", address.Hex())

	// 查询账户余额
	// @区块要用bigint, nil则对应最新余额
	// @以太坊的数字是使用尽可能小的单位处理的，是定点精度。ETH值 = wei/10^18
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("query balance err %v", err)
	}
	fmt.Println(balance)

	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	ethVal := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethVal)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), address)
	if err != nil {
		log.Fatalf("query pending balance err %v", err)
	}
	fPendingBalance := new(big.Float)
	fPendingBalance.SetString(pendingBalance.String())
	pendingEthVal := new(big.Float).Quo(fPendingBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(pendingEthVal)
}
