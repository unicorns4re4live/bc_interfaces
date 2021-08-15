package cli

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/ethclient"
	contract "web3-cli/web3/contracts"
)



type commandID int

const (
	CMD_AUTH commandID = iota
	CMD_QUIT
)

type command struct {
	id string
	args []string
}

type cli struct {
	commands chan command
	ethClient *ethclient.Client
	contract *contract.Storage
	userKey *ecdsa.PrivateKey
}