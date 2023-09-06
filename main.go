package main

import (
	"bitbc/bitbc"
)

func main() {
	bc := bitbc.NewBlockchain()
	defer bc.Db.Close()

	cli := bitbc.CLI{BC: bc}
	cli.Run()

}
