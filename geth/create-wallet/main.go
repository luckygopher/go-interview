package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// 16进制编码 剥离0x = 用于签署交易的私钥，将被视为密码，谁拥有它就可以访问你的所有资产
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	// 派生公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// 16进制编码 剥离0x和前2个字符 = 公钥
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	// 生成公共地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	println(address)
}
