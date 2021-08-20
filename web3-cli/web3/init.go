package web3

import (
	"crypto/ecdsa"
	"encoding/json"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"strings"
	"web3-cli/utils"
	contract "web3-cli/web3/contracts"
)

//Получаем наш instance, сеть и публичный ключ
func InitEthereumData(privKey string) (*ethclient.Client,*ecdsa.PrivateKey, *contract.Change){
	var configs EthereumConfigs
	plan, _:= ioutil.ReadFile("EthereumConfig.json") // JSON с конфигами
	err := json.Unmarshal(plan, &configs)
	utils.Handle(err, "Жсон с конфиг эфириума не заанмаршалился")


	client, err := ethclient.Dial(configs.Network) //Задаем сеть, в которой всё осуществляется
	instance := LoadContract(configs.Contract)

	utils.Handle(err, "не подключились к сети")
	privateKey, err := crypto.HexToECDSA(strings.TrimSpace(privKey))

	utils.Handle(err, "чето не то с прив.ключом")

	return client, privateKey, instance
}
