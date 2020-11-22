package main

import (
	"fmt"
	"math"
	"math/rand"
	"orchlab.com/consistent_hash/chash"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var server_nodes = []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4", "192.168.1.5", "192.168.1.6", "192.168.1.7", "192.168.1.8", "192.168.1.9", "192.168.1.10"}

func main() {
	nodeWeight := make(map[string]int)
	nodeCount := make(map[string]int)

	for _, server_node := range server_nodes {
		nodeWeight[server_node] = 1
		nodeCount[server_node] = 0

	}

	virtualSpots := 500

	hash := chash.NewHashRing(virtualSpots, "sha1")

	hash.AddNodes(nodeWeight)

	for i := 0; i < 1000000; i++ {
		nodeString := hash.GetNode(randSeq(4096))
		for _, server_node := range server_nodes {
			if  nodeString == server_node {
				nodeCount[server_node] += 1
			}
		}
	}

	var sd float64

	for _, server_node := range server_nodes {
		fmt.Println(server_node, nodeCount[server_node])
		sd += math.Pow(float64(nodeCount[server_node] - 100000), 2)
	}

	sd = math.Sqrt(sd/float64(len(server_nodes)))

	fmt.Println("sd is", sd)
}