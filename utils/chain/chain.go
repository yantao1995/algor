package chain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 碰撞
func MiningCrash() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 10)
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}
		//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
		//fmt.Println(hexutil.Encode(publicKeyBytes)[4:])
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

		balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
		fmt.Println("次数:", i)
		fmt.Println("私钥:", hexutil.Encode(privateKeyBytes)[2:])
		fmt.Println("地址:", address)
		fmt.Println("余额:", balance)
		if balance.Cmp(big.NewInt(0)) > 0 {
			break
		}
		fmt.Printf("\033[%dA\033[K", 5)
	}

	fmt.Println("按任意键退出")
	j := ""
	fmt.Scanln(&j)
}
