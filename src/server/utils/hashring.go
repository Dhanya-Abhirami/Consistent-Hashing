package utils

import (
	"log"
	"hash/crc32"
	"errors"
	"strconv"
	// "reflect"
	// "github.com/rogpeppe/go-internal/fmtsort"
)

type HashRing struct {
	Nodes map[uint32]string
	KeysMap map[uint32]string
	Keys int
}

var ErrNodeNotFound = errors.New("Node not found")

func NewHashRing(keys int) *HashRing{
	hashRing := new(HashRing)
	hashRing.Nodes = make(map[uint32]string)
	hashRing.KeysMap = make(map[uint32]string)
	hashRing.Keys = keys
	return hashRing 
}

func getHash(id string) uint32{
	return crc32.ChecksumIEEE([]byte(id))
}

func (hashRing *HashRing) AddNode(id string) int{
	hash := getHash(id)
	hashRing.Nodes[hash]=id
	remap := hashRing.Remap()
	return remap
}

func (hashRing *HashRing) RemoveNode(id string) (int,error){
	hash := getHash(id)
	if _, found := hashRing.Nodes[hash]; found {
		delete(hashRing.Nodes, hash)
		remap := hashRing.Remap()
		return remap,nil
	} else{
		return 0,ErrNodeNotFound
	}
}


func (hashRing *HashRing) GetMapping(id string) (string,error){
	hash := getHash(id)
	if _, found := hashRing.Nodes[hash]; found {
		return hashRing.Nodes[hash],nil
	} else{
		return "",ErrNodeNotFound
	}
}

func (hashRing *HashRing) Remap() int{
	remap := 0
	for n := 1; n <= hashRing.Keys; n++ {
		key, _ := strconv.Atoi(n)
		hash := getHash(key)
		newId := hashRing.GetMapping(hash)
		if hashRing.KeysMap[hash]!=newId {
			remap += 1
			hashRing.KeysMap[hash]=newId
		} 
		log.Print(key,hash,hashRing)
    } 
	return remap
}
