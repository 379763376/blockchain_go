package core

import (
	"blockchain_go"
	"fmt"
)

func (cli *CLI) createWallet() {
	wallets, _ := blockchain_go.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("Your new address: %s\n", address)
}
