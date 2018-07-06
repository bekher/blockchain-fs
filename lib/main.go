package main

import (
	"fmt"

	"blockchain-fs/lib/bfs"
)

func main() {
	fmt.Printf("EVM filesystem initializing...")
	bfs.Init(true, "testdir", "todo get acct")
}
