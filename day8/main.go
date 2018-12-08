package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	childnodes int
	metadata   int
	value      int
	metadatas  *list.List
	children   *list.List
}

func parseNodes(nodelist *list.List, nodes *list.List, metadatas int) int {
	if nodelist == nil {
		return metadatas
	} else {
		node := new(Node)

		e := nodelist.Front()
		e.Next()
		node.childnodes = e.Value.(int)
		nodelist.Remove(e)

		e = nodelist.Front()
		e.Next()
		node.metadata = e.Value.(int)
		nodelist.Remove(e)

		node.children = list.New()
		node.metadatas = list.New()
		nodes.PushBack(node)

		for i := 0; i < node.childnodes; i++ {
			metadatas = parseNodes(nodelist, node.children, metadatas)
		}

		for m := 0; m < node.metadata; m++ {
			e = nodelist.Front()
			e.Next()
			node.metadatas.PushBack(e.Value.(int))
			metadatas += e.Value.(int)

			// Calculate the value, if this is a terminal node
			if node.childnodes > 0 {
				node.value += e.Value.(int)
			}
			nodelist.Remove(e)
		}
		return metadatas
	}
}

// Wrote this for debugging nodes
func printNode(tmpNode *Node) {
	fmt.Printf("Node childs[%d] metadatas[%d] value[%d]\n", tmpNode.childnodes, tmpNode.metadata, tmpNode.value)
	for m := tmpNode.metadatas.Front(); m != nil; m = m.Next() {
		fmt.Printf("%d ", m.Value.(int))
	}
	fmt.Println()
}

// Get the answer for Part 2
func calculateValue(node *Node) int {
	if node == nil {
		return 0
	} else if node.childnodes == 0 {
		var value int = 0

		for m := node.metadatas.Front(); m != nil; m = m.Next() {
			metadatanumber := m.Value.(int)
			value += metadatanumber
		}

		return value
	} else {
		var value int = 0

		for m := node.metadatas.Front(); m != nil; m = m.Next() {
			metadata := m.Value.(int)
			f := node.children.Front()

			// Look for the case like C, where the metadata references a node
			// that doesn't exist
			if metadata <= node.children.Len() {
				for i := 1; i < metadata; i++ {
					f = f.Next()
				}
				referencedNode := f.Value.(*Node)
				if referencedNode != nil {
					value += calculateValue(referencedNode)
				}
			}
		}
		return value
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read in all of the dependencies
	var tree string
	for scanner.Scan() {
		tree = scanner.Text()
	}

	s := strings.Split(tree, " ")
	l := list.New()
	for i := 0; i < len(s); i++ {
		num, _ := strconv.Atoi(s[i])

		l.PushBack(num)
	}

	nodes := list.New()
	metadata := parseNodes(l, nodes, 0)
	// Answer for part 1
	fmt.Printf("Metadata %d\n", metadata)

	// part 2 : find the value of the root node
	e := nodes.Front()
	rootNode := e.Value.(*Node)

	value := calculateValue(rootNode)
	fmt.Printf("Value of rootNode : %d\n", value)
}
