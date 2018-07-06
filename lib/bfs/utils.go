package bfs

import (
	"log"
	"os"
)

var Debug = false

func p_err(s string, args ...interface{}) {
	log.Printf(s, args...)
	os.Exit(1)
}

func p_out(s string, args ...interface{}) {
	if !Debug {
		return
	}
	log.Printf(s, args...)
}
