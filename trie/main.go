package main

import (
	"fmt"
)

type node struct {
	data        map[byte]*node
	isFinalNode bool
}

func (n *node) addchar(a byte) {
	n.data[a] = &node{make(map[byte]*node), false}
	n.isFinalNode = true
}

func (n *node) add(a string) {
	if child, ok := n.data[a[0]]; ok {
		child.addchar(a[0])
		if len(a) > 1 {
			child.add(a[1 : len(a)-1])
		}
	}
	n.data[a[0]] = &node{make(map[byte]*node), false}
	if len(a) > 1 {
		n.data[a[0]].add(a[1:])
	} else {
		n.data[a[0]].addchar(a[0])
	}
}

func search(n *node, s string) bool {
	if n == nil {
		return false
	}

	if len(s) == 0 {
		return n.isFinalNode
	}

	if child, ok := n.data[s[0]]; ok {
		return search(child, s[1:])
	}
	return false
}

func main() {
	root := &node{make(map[byte]*node), false}
	root.add("abc")
	r := search(root, "ab")
	fmt.Println(r)
}
