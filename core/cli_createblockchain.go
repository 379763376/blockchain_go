package core

import (
	"blockchain_go"
	"fmt"
	"log"
)

func (cli *CLI) createBlockchain(address string) {
	if !blockchain_go.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := CreateBlockchain(address)
	defer bc.db.Close()

	UTXOSet := blockchain_go.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
