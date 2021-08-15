package web3

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"web3-cli/utils"
)

func AuthSet(c *ethclient.Client, privKey *ecdsa.PrivateKey) (*bind.TransactOpts) {
	publicKeyECDSA := privKey.PublicKey


	fromAddress := crypto.PubkeyToAddress(publicKeyECDSA)
	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	utils.Handle(err, "Не удалось получить Nonce")

	//Задаем настройки транзакции
	auth := bind.NewKeyedTransactor(privKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = big.NewInt(0)

	return auth
}
