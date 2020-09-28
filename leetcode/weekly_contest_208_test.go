package dance

import (
	"fmt"
	"testing"
)

func minOperations(logs []string) int {
	stack := make([]string, 0)
	for _, log := range logs {
		switch log {
		case "../":
			length := len(stack)
			if length > 0 {
				stack = stack[:length-1]
			}
		case "./":
			fmt.Println("./")
		default:
			stack = append(stack, log)
		}
	}
	return len(stack)
}

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	if len(customers) == 0 {
		return 0
	}
	left := 0
	cost := 0
	earn := 0
	max := 0
	maxIndex := -1
	for i := 0; ; i++ {
		if i < len(customers) {
			earn += min(left+customers[i], 4) * boardingCost
			left = left + customers[i] - min(left+customers[i], 4)
		} else if left == 0 {
			break
		} else {
			earn += min(left, 4) * boardingCost
			left -= min(left, 4)
		}
		cost += runningCost
		if earn-cost > max {
			max = earn - cost
			maxIndex = i
		}
	}
	if maxIndex == -1 {
		return -1
	}
	return maxIndex + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestMinOperationsMaxProfit(t *testing.T) {
	t.Log(minOperationsMaxProfit([]int{8, 3}, 5, 6))
}

type Node struct {
	Dead     bool
	Name     string
	Children []*Node
}

type ThroneInheritance struct {
	First *Node
	Graph map[string]string
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		&Node{
			Name:     kingName,
			Children: make([]*Node, 0),
		},
		make(map[string]string),
	}
}

func getNodeByGraph(graph map[string]string, root *Node, name string) *Node {
	king := root.Name
	stack := make([]string, 0)
	var child = name
	for child != king {
		stack = append(stack, child)
		child = graph[child]
	}
	node := root
	for i := len(stack) - 1; i >= 0; i-- {
		for j := len(node.Children) - 1; j >= 0; j-- {
			if node.Children[j].Name == stack[i] {
				node = node.Children[j]
				break
			}
		}
	}
	return node
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.Graph[childName] = parentName
	parent := getNodeByGraph(this.Graph, this.First, parentName)
	parent.Children = append(parent.Children, &Node{
		Name:     childName,
		Children: make([]*Node, 0),
	})
}

func (this *ThroneInheritance) Death(name string) {
	node := getNodeByGraph(this.Graph, this.First, name)
	node.Dead = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	order := make([]string, 0)
	var leftFirst func(node *Node)
	leftFirst = func(node *Node) {
		if node == nil {
			panic("node is nil")
		}
		if !node.Dead {
			order = append(order, node.Name)
		}
		for _, child := range node.Children {
			leftFirst(child)
		}
	}
	leftFirst(this.First)
	return order
}
