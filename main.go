package main

import (
	"fmt"
	"orchlab.com/consistent_hash/chash"
)

const (
	node1 = "192.168.1.1"
	node2 = "192.168.1.2"
	node3 = "192.168.1.3"
)

func getNodesCount(nodes chash.NodesArray) (int, int, int) {
	node1Count := 0
	node2Count := 0
	node3Count := 0

	for _, node := range nodes {
		if node.NodeKey == node1 {
			node1Count += 1
		}
		if node.NodeKey == node2 {
			node2Count += 1

		}
		if node.NodeKey == node3 {
			node3Count += 1

		}
	}
	return node1Count, node2Count, node3Count
}

func main() {
	nodeWeight := make(map[string]int)
	nodeWeight[node1] = 2
	nodeWeight[node2] = 2
	nodeWeight[node3] = 3

	vitualSpots := 100

	hash := chash.NewHashRing(vitualSpots)

	hash.AddNodes(nodeWeight)
	if hash.GetNode("1") != node3 {
		fmt.Println("expetcd %v got %v", node3, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node3 {
		fmt.Println("expetcd %v got %v", node3, hash.GetNode("2"))
	}
	if hash.GetNode("3") != node2 {
		fmt.Println("expetcd %v got %v", node2, hash.GetNode("3"))
	}
	c1, c2, c3 := getNodesCount(hash.Nodes)
	fmt.Println("len of nodes is %v after AddNodes node1:%v, node2:%v, node3:%v", len(hash.Nodes), c1, c2, c3)

	hash.RemoveNode(node3)
	if hash.GetNode("1") != node1 {
		fmt.Println("expetcd %v got %v", node1, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node2 {
		fmt.Println("expetcd %v got %v", node1, hash.GetNode("2"))
	}
	if hash.GetNode("3") != node2 {
		fmt.Println("expetcd %v got %v", node2, hash.GetNode("3"))
	}
	c1, c2, c3 = getNodesCount(hash.Nodes)
	fmt.Println("len of nodes is %v after RemoveNode node1:%v, node2:%v, node3:%v", len(hash.Nodes), c1, c2, c3)

	hash.AddNode(node3, 3)
	if hash.GetNode("1") != node3 {
		fmt.Println("expetcd %v got %v", node3, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node3 {
		fmt.Println("expetcd %v got %v", node3, hash.GetNode("2"))
	}
	if hash.GetNode("3") != node2 {
		fmt.Println("expetcd %v got %v", node2, hash.GetNode("3"))
	}
	c1, c2, c3 = getNodesCount(hash.Nodes)
	fmt.Println("len of nodes is %v after AddNode node1:%v, node2:%v, node3:%v", len(hash.Nodes), c1, c2, c3)

}

