package cli

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/ethclient"
	contract "web3-cli/web3/contracts"
)



type command struct {
	id string
	args []string
}

type cli struct {
	commands chan command
	ethClient *ethclient.Client
	contract *contract.Change
	userKey *ecdsa.PrivateKey
}
