package utils

import (
	"hash/crc32"
)

type Node struct {
	id string 
	hashId uint32
}

func newNode(id string) *Node {
	node := new(Node)
	node.id = id
	node.hashId = crc32.ChecksumIEEE([]byte(id))
	return node
}

func (node *Node) getHashId(id string) uint32{
	return node.hashId
}