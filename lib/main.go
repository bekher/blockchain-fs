package main

import (
	"fmt"

	"blockchain-fs/lib/bfs"
)

func main() {
	fmt.Printf("EVM FS Initializing")
	bfs.Init(true, "testdir", "todo get acct")
}
