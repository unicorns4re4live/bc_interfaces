package cli

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"os"
	"strings"
	"time"
	"unsafe"
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

//GET METHODS
//func retrieve(c *cli) {
//	fmt.Println("RETRIEVE", c.contract)
//	result, err := c.contract.Retrieve(nil)
//	utils.Handle(err, "Проблема при вызове метода retrieve")
//	fmt.Println("Число в контракте: ", result)
//}
func admin_amount(c *cli) {
	result, err := c.contract.AdminAmount(nil)
	utils.Handle(err, "Проблема при вызове метода admin_amount")
	fmt.Println("Число админов в контракте: ", result)
}

func get_template_names(c *cli) {
	result, err := c.contract.GetTemplateNames(nil)
	utils.Handle(err, "Проблема при вызове метода get_template_names")
	fmt.Println("Шаблоны перевода: ", strings.Join(result, ","))
}

func get_template_by_name(c *cli, args []string) {
	templateName := strings.TrimSpace(strings.Join(args, " "))
	category, value, err := c.contract.GetTemplateByName(nil, templateName)
	utils.Handle(err, "Проблема при вызове метода get_template_by_name")
	fmt.Println("Айди категории: ", category, ", Значение шаблона: ", value)
}

func get_userlist(c *cli) {
	userList, err := c.contract.GetUserlist(nil)
	utils.Handle(err, "Проблема при вызове метода get_userlist")
	fmt.Println("Список юзеров: ", userList)
}

func get_user(c *cli, args []string) {
	if len(args) != 1 {
		fmt.Println("Некоректные аргументы")
		return
	}
	userAddress := common.HexToAddress(strings.TrimSpace(args[0]))
	user, err := c.contract.Users(nil, userAddress)
	utils.Handle(err, "Проблема при вызове метода get_user")
	fmt.Println("In bytes:", user.Pwhash[:])
	fmt.Println("Баланс: ", user.Balance, ", Хэш пароля", common.BytesToHash(user.Pwhash[:]), ", Админ?", user.Role)
}

//SET
func create_user(c *cli, args []string) {
	if len(args) != 1 {
		fmt.Println("Некоректные аргументы")
		return
	}
	strForPw := strings.TrimSpace(args[0])

	pwHash := crypto.Keccak256([]byte(strForPw))

	fmt.Println(common.BytesToHash(pwHash[:]))

	arr32 := *(*[32]byte)(unsafe.Pointer(&pwHash))

	fmt.Println("HASH", arr32)
	txOptions := web3.AuthSet(c.ethClient, c.userKey)

	tx, err := c.contract.CreateUser(txOptions, arr32)

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

func create_transfer(c *cli, args []string) {

	txOptions := web3.AuthSet(c.ethClient, c.userKey)
	userAddress := common.HexToAddress(strings.TrimSpace(args[0]))

	big1 := new(big.Int)
	categoryId, _ := big1.SetString(strings.TrimSpace(args[1]) , 10)

	big2 := new(big.Int)
	val, _ := big2.SetString(strings.TrimSpace(args[2]) , 10)

	description := args[3]

	strForPw := strings.TrimSpace(args[4])

	pwHash := crypto.Keccak256([]byte(strForPw))



	arr32 := *(*[32]byte)(unsafe.Pointer(&pwHash))
	//arr32 := [32]byte{100 230 4 120 124 191 25 72 65 231 182 141 124 210 135 134 246 201 160 163 171 159 139 10 14 135 203 67 135 171 1 7]
	fmt.Println("AAAAAAAAAAAAAAAA", arr32)

	tx, err := c.contract.CreateTransfer(txOptions, userAddress, categoryId, val, description, arr32)

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

//func store(c *cli, args []string) {
//	if len(args) != 1 {
//		fmt.Println("Некоректные аргументы")
//		return
//	}
//	hehe := new(big.Int)
//	num, ok := hehe.SetString(strings.TrimSpace(args[0]) , 10)
//	if !ok {
//		fmt.Println("Аргумент не является числом")
//		return
//	}
//	txOptions := web3.AuthSet(c.ethClient, c.userKey)
//
//
//	tx, err := c.contract.Store(txOptions, num)
//	utils.Handle(err, "Проблема при вызове метода store")
//	fmt.Println("Транзакция отправлена: ", tx.Hash())
//	_, isPending, err := c.ethClient.TransactionByHash(context.Background(), tx.Hash())
//	if err != nil {
//		fmt.Println("Проблемы с транзакцией:", err)
//		return
//	}
//
//	for isPending {
//		_, isPending, err = c.ethClient.TransactionByHash(context.Background(), tx.Hash())
//		time.Sleep(3)
//	}
//
//	fmt.Println("Транзакция выполнена.")
//}
