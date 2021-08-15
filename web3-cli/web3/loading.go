package web3

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	contract "web3-cli/web3/contracts"
)

//Цепляемся к контракту
func LoadContract(address string) *contract.Storage {
	var configs EthereumConfigs
	plan, _:= ioutil.ReadFile("EthereumConfig.json") // JSON с конфигами
	err := json.Unmarshal(plan, &configs)

	client, err := ethclient.Dial(configs.Network) //Задаем сеть, в которой всё осуществляется
	if err != nil {
		log.Fatal("не подключились к сети")
	}

	contractAddr := common.HexToAddress(address)

	instance, err := contract.NewStorage(contractAddr, client)

	if err != nil {
		log.Fatal("Чёто пошло не так при создании инстанса")
	}

	return instance

}
