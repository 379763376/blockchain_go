package core

import (
	"blockchain_go"
	"fmt"
	"log"
)

func (cli *CLI) send(from, to string, amount int) {
	if !blockchain_go.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !blockchain_go.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := NewBlockchain()
	UTXOSet := blockchain_go.UTXOSet{bc}
	defer bc.db.Close()

	tx := blockchain_go.NewUTXOTransaction(from, to, amount, &UTXOSet)
	cbTx := blockchain_go.NewCoinbaseTX(from, "")
	txs := []*blockchain_go.Transaction{cbTx, tx}

	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)
	fmt.Println("Success!")
}
