pragma solidity ^0.4.18;

contract FileSystem {

  mapping (uint64 => Node) idToNode;
  mapping (uint64 => Attrs) idToAttrs;

  struct Head {
    uint32 rootNode;
    uint64 nextInd;
  }

  struct DNode {
    bytes32 name;
    uint64 version;

    uint64[] children;
    uint64 parent;

    bytes32[] data;

  }

  // https://godoc.org/bazil.org/fuse#Attr.String
  struct Attrs {
    uint64 Inode;
    uint64 Size;
    uint64 Blocks; //512 byte units? should be 32

    uint Atime;
    uint Mtime;
    uint Ctime;
    uint Mode; // os.FileMode

    uint32 uid;
    uint32 gid;  
  }

  address owner;
}
