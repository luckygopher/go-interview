package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	// keystore 是一个包含了经过加密了的钱包私钥，每个文件只能包含一个钱包私钥对
	createKs()
	//importKs()
}

func createKs() {
	ks := keystore.NewKeyStore("./geth/keystore/tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "123456"
	// 创建新的钱包
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}

func importKs() {
	file := "./geth/keystore/tmp/UTC--2024-04-17T09-37-33.337636000Z--91884f85b457a89e18e5fd38a1cadffaeaa430e2"
	ks := keystore.NewKeyStore("./geth/keystore/tmp/import", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	password := "123456"
	newPassword := "123123"
	account, err := ks.Import(jsonBytes, password, newPassword)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
