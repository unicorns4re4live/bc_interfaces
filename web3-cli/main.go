package main

import (
	"fmt"
	"web3-cli/cli"
)

func main() {
	s := cli.NewCli()

	fmt.Println(s)

	go s.ReadCommands()

	s.PushCommands()

}