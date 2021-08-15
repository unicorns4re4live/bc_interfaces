package cli

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"
	"web3-cli/utils"
	"web3-cli/web3"
)

func exit() {
	fmt.Println("Bye")
	os.Exit(1)
}

func auth(c *cli, args []string) {
	if len(args) != 1 {
		fmt.Println("Некоректные аргументы")
		return
	}
	key := args[0]
	fmt.Println("Private key", key)
	client, privKey, instance := web3.InitEthereumData(key)
	c.contract = instance
	c.ethClient = client
	c.userKey = privKey
}

func retrieve(c *cli) {
	fmt.Println("RETRIEVE", c.contract)
	result, err := c.contract.Retrieve(nil)
	utils.Handle(err, "Проблема при вызове метода retrieve")
	fmt.Println("Число в контракте: ", result)
}

func store(c *cli, args []string) {
	if len(args) != 1 {
		fmt.Println("Некоректные аргументы")
		return
	}
	hehe := new(big.Int)
	num, ok := hehe.SetString(strings.TrimSpace(args[0]) , 10)
	if !ok {
		fmt.Println("Аргумент не является числом")
		return
	}
	txOptions := web3.AuthSet(c.ethClient, c.userKey)


	tx, err := c.contract.Store(txOptions, num)
	utils.Handle(err, "Проблема при вызове метода store")
	fmt.Println("Транзакция отправлена: ", tx.Hash())
	_, isPending, err := c.ethClient.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		fmt.Println("Проблемы с транзакцией:", err)
		return
	}

	for isPending {
		_, isPending, err = c.ethClient.TransactionByHash(context.Background(), tx.Hash())
		time.Sleep(3)
	}

	fmt.Println("Транзакция выполнена.")
}