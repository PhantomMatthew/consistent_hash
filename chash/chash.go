package chash

import (
	"crypto/sha1"
	jump "github.com/lithammer/go-jump-consistent-hash"
	"hash"
	"hash/fnv"
	"math"
	"sort"
	"strconv"
	"sync"
)

const (
	//DefaultVirualSpots default virual spots
	DefaultVirualSpots = 400
)

type Node struct {
	NodeKey   string
	SpotValue uint32
}

type NodesArray []Node

func (p NodesArray) Len() int           { return len(p) }
func (p NodesArray) Less(i, j int) bool { return p[i].SpotValue < p[j].SpotValue }
func (p NodesArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p NodesArray) Sort()              { sort.Sort(p) }

//HashRing store nodes and weigths
type HashRing struct {
	VirualSpots int
	Nodes       NodesArray
	Weights     map[string]int
	mu          sync.RWMutex
	HashFuncName string
	HashFunc 	hash.Hash
}

//NewHashRing create a hash ring with virual spots
func NewHashRing(spots int, hashName string) *HashRing {
	if spots == 0 {
		spots = DefaultVirualSpots
	}

	h := &HashRing{
		VirualSpots: spots,
		Weights:     make(map[string]int),
		HashFuncName:  hashName,
	}
	return h
}

type HashFunc func()

//AddNodes add nodes to hash ring
func (h *HashRing) AddNodes(nodeWeight map[string]int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for nodeKey, w := range nodeWeight {
		h.Weights[nodeKey] = w
	}
	h.generate()
}

//AddNode add node to hash ring
func (h *HashRing) AddNode(nodeKey string, weight int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.Weights[nodeKey] = weight
	h.generate()
}

//RemoveNode remove node
func (h *HashRing) RemoveNode(nodeKey string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.Weights, nodeKey)
	h.generate()
}

//UpdateNode update node with weight
func (h *HashRing) UpdateNode(nodeKey string, weight int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.Weights[nodeKey] = weight
	h.generate()
}

func (h *HashRing) generate() {
	var totalW int
	for _, w := range h.Weights {
		totalW += w
	}

	totalVirtualSpots := h.VirualSpots * len(h.Weights)
	h.Nodes = NodesArray{}

	for nodeKey, w := range h.Weights {
		spots := int(math.Floor(float64(w) / float64(totalW) * float64(totalVirtualSpots)))
		for i := 1; i <= spots; i++ {
			n := Node{
				NodeKey:   nodeKey,
				SpotValue: h.HashCalculation((nodeKey + ":" + strconv.Itoa(i))),
			}
			h.Nodes = append(h.Nodes, n)
			h.HashFunc.Reset()
		}
	}
	h.Nodes.Sort()
}

func genValue(bs []byte) uint32 {
	if len(bs) < 4 {
		return 0
	}
	v := (uint32(bs[3]) << 24) | (uint32(bs[2]) << 16) | (uint32(bs[1]) << 8) | (uint32(bs[0]))
	return v
}

//GetNode get node with key
func (h *HashRing) GetNode(s string) string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if len(h.Nodes) == 0 {
		return ""
	}

	v := h.HashCalculation(s)
	i := sort.Search(len(h.Nodes), func(i int) bool { return h.Nodes[i].SpotValue >= v })

	if i == len(h.Nodes) {
		i = 0
	}
	return h.Nodes[i].NodeKey
}

func (h *HashRing) HashCalculation(content string) uint32 {
	var v uint32
	if h.HashFuncName == "sha1" {
		h.HashFunc = sha1.New()
		h.HashFunc.Write([]byte(content))
		hashBytes := h.HashFunc.Sum(nil)
		v = genValue(hashBytes[6:10])
	}

	if h.HashFuncName == "fnv" {
		h.HashFunc = fnv.New32()
		h.HashFunc .Write([]byte(content))
		hashBytes := h.HashFunc .Sum(nil)
		v = genValue(hashBytes)
	}

	if h.HashFuncName == "jump" {
		h.HashFunc = jump.NewCRC32()
		v = (uint32)(jump.New(4, jump.NewCRC32()).Hash(content))
		//h.HashFunc .Write([]byte(content))
		//hashBytes := h.HashFunc .Sum(nil)
		//v = genValue(hashBytes)
	}
	return v
}