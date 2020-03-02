package core

import (
	"blockchain_go"
	"fmt"
	"log"
)

func (cli *CLI) listAddresses() {
	wallets, err := blockchain_go.NewWallets()
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
