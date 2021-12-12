package utils

import(
	"sort"
)

type HashRing struct {
	Nodes Nodes
}

type Nodes []*Node

func (n Nodes) Len() int { 
	return len(n) 
}

func (n Nodes) Less(i int, j int) bool { 
	return n[i].hashId < n[j].hashId 
}

func (n Nodes) Swap(i int, j int) { 
	n[i], n[j] = n[j], n[i]
}

func NewHashRing() *HashRing{
	hashRing := new(HashRing)
	hashRing.Nodes = Nodes{}
	return hashRing 
}

func (hashRing *HashRing) AddNode(id string){
	node := newNode(id)
	hashRing.Nodes = append(hashRing.Nodes,node)
	sort.Sort(hashRing.Nodes)
}

func (hashRing *HashRing) RemoveNode(id string){

}
