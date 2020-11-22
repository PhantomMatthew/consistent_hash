package main

import (
	"fmt"
	"math"
	"math/rand"
	"orchlab.com/consistent_hash/chash"
)

//const (
//	node1 = "192.168.1.1"
//	node2 = "192.168.1.2"
//	node3 = "192.168.1.3"
////	node4 = "192.168.1.4"
////	node5 = "192.168.1.5"
////	node6 = "192.168.1.6"
////	node7 = "192.168.1.7"
////	node8 = "192.168.1.8"
////	node9 = "192.168.1.9"
////	node10 = "192.168.1.10"
//)
//
////var server_nodes = []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4", "192.168.1.5", "192.168.1.6", "192.168.1.7", "192.168.1.8", "192.168.1.9", "192.168.1.10"}
//
//func getNodesCount(nodes chash.NodesArray) (int, int, int) {
////func getNodesCount(nodes chash.NodesArray) (map[string]int) {
//	//nodeCount := make(map[string]int)
//	//for _, server_node := range server_nodes {
//	//	nodeCount[server_node] = 0
//	//}
//	//
//	//for _, node := range nodes {
//	//	for _, server_node := range server_nodes {
//	//		if node.NodeKey == server_node {
//	//			nodeCount[server_node] += 1
//	//		}
//	//	}
//	//}
//	//
//	//return nodeCount
//
//	node1Count := 0
//	node2Count := 0
//	node3Count := 0
//
//	for _, node := range nodes {
//		if node.NodeKey == node1 {
//			node1Count += 1
//		}
//		if node.NodeKey == node2 {
//			node2Count += 1
//
//		}
//		if node.NodeKey == node3 {
//			node3Count += 1
//
//		}
//	}
//	return node1Count, node2Count, node3Count
//}
//
//
//
//
//func main() {
//	nodeWeight := make(map[string]int)
//
//	//for _, server_node := range server_nodes {
//	//	nodeWeight[server_node] = 1
//	//}
//	nodeWeight[node1] = 2
//	nodeWeight[node2] = 2
//	nodeWeight[node3] = 3
//
//	vitualSpots := 150
//
//	hash := chash.NewHashRing(vitualSpots)
//
//	hash.AddNodes(nodeWeight)
//
//	fmt.Println("hello got ", hash.GetNode("hello"))
//
//	//testCounts := getNodesCount(hash.Nodes)
//	//for _, server_node := range server_nodes {
//	//	fmt.Println(server_node, testCounts[server_node])
//	//}
//
//	if hash.GetNode("1") != node3 {
//		fmt.Println("expetcd %v got %v", node3, hash.GetNode("1"))
//	}
//	if hash.GetNode("2") != node3 {
//		fmt.Println("expetcd %v got %v", node3, hash.GetNode("2"))
//	}
//	if hash.GetNode("3") != node2 {
//		fmt.Println("expetcd %v got %v", node2, hash.GetNode("3"))
//	}
//	c1, c2, c3 := getNodesCount(hash.Nodes)
//	fmt.Println("len of nodes is %v after AddNodes node1:%v, node2:%v, node3:%v", len(hash.Nodes), c1, c2, c3)
//
//	hash.RemoveNode(node3)
//	if hash.GetNode("1") != node1 {
//		fmt.Println("expetcd %v got %v", node1, hash.GetNode("1"))
//	}
//	if hash.GetNode("2") != node2 {
//		fmt.Println("expetcd %v got %v", node1, hash.GetNode("2"))
//	}
//	if hash.GetNode("3") != node2 {
//		fmt.Println("expetcd %v got %v", node2, hash.GetNode("3"))
//	}
//	c1, c2, c3 = getNodesCount(hash.Nodes)
//	fmt.Println("len of nodes is %v after RemoveNode node1:%v, node2:%v, node3:%v", len(hash.Nodes), c1, c2, c3)
//
//	hash.AddNode(node3, 3)
//	if hash.GetNode("1") != node3 {
//		fmt.Println("expetcd %v got %v", node3, hash.GetNode("1"))
//	}
//	if hash.GetNode("2") != node3 {
//		fmt.Println("expetcd %v got %v", node3, hash.GetNode("2"))
//	}
//	if hash.GetNode("3") != node2 {
//		fmt.Println("expetcd %v got %v", node2, hash.GetNode("3"))
//	}
//	c1, c2, c3 = getNodesCount(hash.Nodes)
//	fmt.Println("len of nodes is %v after AddNode node1:%v, node2:%v, node3:%v", len(hash.Nodes), c1, c2, c3)
//
//}


var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var server_nodes = []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4", "192.168.1.5", "192.168.1.6", "192.168.1.7", "192.168.1.8", "192.168.1.9", "192.168.1.10"}

func getNodesCount(nodes chash.NodesArray) (map[string]int) {
	nodeCount := make(map[string]int)
	for _, server_node := range server_nodes {
		nodeCount[server_node] = 0
	}

	//for _, node := range nodes {
	//	for _, server_node := range server_nodes {
	//		if node.NodeKey == server_node {
	//			nodeCount[server_node] += 1
	//		}
	//	}
	//}




	return nodeCount
}




func main() {
	nodeWeight := make(map[string]int)
	nodeCount := make(map[string]int)

	for _, server_node := range server_nodes {
		nodeWeight[server_node] = 1
		nodeCount[server_node] = 0

	}

	virtualSpots := 500

	hash := chash.NewHashRing(virtualSpots)

	hash.AddNodes(nodeWeight)

	for i := 0; i < 1000000; i++ {
		nodeString := hash.GetNode(randSeq(4096))
		fmt.Println(nodeString)
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