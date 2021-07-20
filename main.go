package main

import "fmt"

func main() {

	bc := NewBlockChain()
	bc.AddBlock("第一个")
	bc.AddBlock("第二个")

	for i, block := range bc.Blocks {
		fmt.Printf("+++++++++++++++ %d ++++++++++++++\n", i)
		fmt.Printf("PrevBlockHash : %x\n", block.PrevBlockHash)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("Data : %s\n", block.Data)
	}
}
