package bfs

import (
	"log"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)


type DNode struct {
	Name    string
	Attrs   fuse.Attr
	Version int

	dirty    bool // flushed to eth ledger
	parent   *DNode
	children map[string]*DNode
	data     []byte
}

type Head struct {
	Root    string
	NextInd uint64
	Replica uint64
}

var root *DNode
var nextInd uint64 = 1

var sem chan int

type FS struct{}

// Init account=etherum account that stores the db head
func Init(debug bool, mountPoint string, acct string) {
	//nodeID := uint64(rand.Int63())
	if err := os.MkdirAll(mountPoint, 0755); err != nil {
		pErr("mount pt creation fail\n")
	}

	fuse.Unmount(mountPoint)

	var c *fuse.Conn = nil
	var err error = nil

	c, err = fuse.Mount(mountPoint)

	if err != nil {
		log.Fatal(err)
	}

	sem = make(chan int, 1)
	// go Flusher(sem)

	err = fs.Serve(c, FS{})
	if err != nil {
		log.Fatal(err)
	}

	// check if the mount process has an error to report
	<-c.Ready

  if err := c.MountError; err != nil {
    log.Fatal(err)
  }
}

// semaphore stuff
func in() {
	sem <- 1
}

func out() {
	<-sem
}

// gets the head element from the blockchain
func getHead(acct string) (*Head, uint64) {
	// TODO: call blockchain
	return nil, 0
}

func pErr(s string, args ...interface{}) {
	log.Printf(s, args...)
	os.Exit(1)
}

func pOut(s string, args ...interface{}) {
	if !debug {
		return
	}
	log.Printf(s, args...)
}

// FS funcs
type (FS) Root() (fs.Node, error) {
  p_out("root returns as %d\n", int(root.Attrs.Inode))
  return root, nil
}
