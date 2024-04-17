package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 校验地址格式
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Println(re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d"))

	// 验证是一个智能合约 还是一个标准的以太坊账户
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	codeByte, err := client.CodeAt(context.Background(), common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d"), nil)
	if err != nil {
		log.Fatal(err)
	}
	// 大于0存在字节码，则为智能合约；当地址上没有字节码时，则为一个标准的以太坊账户
	isContract := len(codeByte) > 0
	fmt.Println(isContract)
}
