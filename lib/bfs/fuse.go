package bfs

import (
	"context"
	"os"
	"time"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

type Dfs struct{}

func getNextInode() uint64 {
	nextInd++
	return nextInd
}

var _ fs.Node = (*DNode)(nil)
var _ fs.FS = (*Dfs)(nil)

func (Dfs) Root() (n fs.Node, err error) {
	return root, nil
}

func (n *DNode) init(name string, fmode os.FileMode) {
	n.Name = name
	n.kids = make(map[string]*DNode)
	n.Attrs = fuse.Attr{
		Inode: getNextInode(),
		Mode:  fmode,
		Mtime: time.Now(),
		Atime: time.Now(),
	}
	n.Version = 1
}

func (n *DNode) Attr(ctx context.Context, Attr *fuse.Attr) error {
	*Attr = n.Attrs
	return nil
}
