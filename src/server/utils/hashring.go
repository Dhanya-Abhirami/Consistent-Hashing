package utils

import (
	"hash/crc32"
	"errors"
	"sort"
	"strconv"
)

type HashRing struct {
	Servers map[uint32]string
	Keys int
	KeysMap map[string]KeysMapping
}

type KeysMapping struct {
	KeyHash uint32 
	ServerMapped string
}

var ErrServerNotFound = errors.New("Server not found")

func NewHashRing(keys int) *HashRing{
	hashRing := new(HashRing)
	hashRing.Servers = make(map[uint32]string)
	hashRing.Keys = keys
	hashRing.KeysMap = make(map[string]KeysMapping)
	hashRing.InitializeMap()
	return hashRing 
}

func getHash(id string) uint32{
	return crc32.ChecksumIEEE([]byte(id))
}

func (hashRing *HashRing) AddServer(id string) int{
	hash := getHash(id)
	hashRing.Servers[hash]=id
	remap := hashRing.Remap()
	return remap
}

func (hashRing *HashRing) RemoveServer(id string) (int,error){
	hash := getHash(id)
	if _, found := hashRing.Servers[hash]; found {
		delete(hashRing.Servers, hash)
		remap := hashRing.Remap()
		return remap,nil
	} else{
		return 0,ErrServerNotFound
	}
}


func (hashRing *HashRing) GetMapping(id string) string{
	return hashRing.KeysMap[id].ServerMapped
}

func (hashRing *HashRing) InitializeMap(){
	for n := 1; n <= hashRing.Keys; n++ {
		key := strconv.Itoa(n)
		hash := getHash(key)
		hashRing.KeysMap[key] = KeysMapping{KeyHash : hash, ServerMapped : key}
    } 
}

func (hashRing *HashRing) Remap() int{
	remap := 0
	keys := make([]uint32, 0, hashRing.Keys)
	for k := range hashRing.Servers {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	for n := 1; n <= hashRing.Keys; n++ {
		key := strconv.Itoa(n)
		hash := hashRing.KeysMap[key].KeyHash
		idx := sort.Search(len(keys), func(idx int) bool { return keys[idx] >= hash })
		if idx >= len(keys) {
			idx = 0
		} 
		newServerMapped := hashRing.Servers[keys[idx]]
		if entry, ok := hashRing.KeysMap[key]; ok {
			if entry.ServerMapped != newServerMapped {
				remap += 1
				entry.ServerMapped = newServerMapped
				hashRing.KeysMap[key] = entry
			}
		}
    } 
	return remap
}
