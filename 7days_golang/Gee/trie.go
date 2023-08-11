package gee

//支持动态路径/xxxx/:name形式

import (
	"fmt"
	"regexp"
	"strings"
)

type node struct {
	uri string
	// support /xxx/:name
	value         string
	childrenNodes []*node
	// support /xx/*
	iswild bool
}

const WildUri = "/:name"

func NewNode() *node {
	return &node{
		uri:           "",
		value:         "head",
		childrenNodes: make([]*node, 0),
		iswild:        false,
	}
}

func newNode(uri, value string, iswild bool) *node {
	return &node{
		uri:           uri,
		value:         value,
		childrenNodes: make([]*node, 0),
		iswild:        iswild,
	}
}

func (head *node) Add(uri string) {
	//fmt.Println("URI:", uri)
	head.insert(uri)
	//head.Print()
}
func (head *node) Get(uri string) (*node, string) {
	return head.search(uri)
}

func splitUri(uri string) []string {
	uri = strings.TrimSpace(uri)
	parts := strings.Split(uri, "/")
	//fmt.Println(parts)
	for i := 0; i < len(parts)-1; i++ {
		parts[i] += "/"
	}
	if parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}
	//fmt.Println(parts)
	return parts
}

func (head *node) search(uri string) (*node, string) {
	searchNode := head
	parts := splitUri(uri)
	key := ""

	for i, p := range parts {
		node := searchNode.searchChild(p)
		if node == nil {
			return nil, key
		}
		key = key + node.uri
		searchNode = node
		if searchNode.iswild {
			node.value = strings.Join(parts[i:], "")
			break
		}
	}
	if len(searchNode.childrenNodes) > 0 {
		return nil, ""
	}
	return searchNode, key
}

func (n *node) searchChild(part string) *node {
	if len(n.childrenNodes) == 0 {
		return nil
	}
	for _, node := range n.childrenNodes {
		//fmt.Printf("search:%p:%+v, [%s:%s]\n", node, node, node.value, part)
		if node.uri == part || node.iswild {
			return node
		}
	}
	return nil
}

func (head *node) insert(uri string) {
	n := head
	parts := splitUri(uri)
	matchReg, _ := regexp.Compile(":(.*)")
	for _, p := range parts {
		ChildNode := n.searchChild(p)
		//fmt.Printf("%p -> %p\n", n, ChildNode)
		if ChildNode == nil {
			match := matchReg.MatchString(p)
			if match {
				ChildNode = newNode(p, p, true)
			} else {
				ChildNode = newNode(p, "", false)
			}
			n.childrenNodes = append(n.childrenNodes, ChildNode)
		}
		//fmt.Printf("parent:%p, %p:%+v\n", n, ChildNode, ChildNode)
		n = ChildNode
	}
}

func (head *node) Print() {
	head.printChildren()
}

func (n *node) printChildren() *node {
	var child *node = nil
	fmt.Println("======")
	fmt.Printf("%p:%+v\n", n, n)
	for _, one := range n.childrenNodes {
		child = one.printChildren()
	}
	return child
}
