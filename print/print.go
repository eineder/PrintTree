package print

import (
	"fmt"
	"strings"
)

type pathNode[T any] struct {
	Node            *T
	HasRightSibling bool
}

func Print[T any](node *T, getChildren func(*T) []*T, getNodeContent func(*T) string) {
	var path []*pathNode[T] = []*pathNode[T]{}
	printNode(node, path, getChildren, getNodeContent)
}

func printNode[T any](node *T, path []*pathNode[T], getChildren func(*T) []*T, getNodeContent func(*T) string) {
	line := getLine(node, path)
	isRoot := len(path) == 0
	if !isRoot {
		fmt.Println(line)
		fmt.Println(line)
		fmt.Println(line)
		fmt.Println(line)
	}
	content := getNodeContent(node)
	if isRoot {
		fmt.Println("#### " + content)
	} else {
		fmt.Println(line + "#### " + content)
	}

	children := getChildren(node)
	isLeaf := len(children) == 0
	if isLeaf {
		return
	}

	hasRightHandSibling := false
	if len(path) > 0 {
		parent := path[len(path)-1]
		parentNode := parent.Node
		hasRightHandSibling = hasRightSibling(node, parentNode, getChildren)
	}
	pathNode := &pathNode[T]{Node: node, HasRightSibling: hasRightHandSibling}
	newPath := append(path, pathNode)
	for childIndex := 0; childIndex < len(children); childIndex++ {
		child := children[childIndex]
		printNode(child, newPath, getChildren, getNodeContent)
	}
}

func getLine[T any](node *T, path []*pathNode[T]) string {
	line := ""
	for i := 0; i < len(path); i++ {
		pathNode := path[i]
		lineFragment := getLineFragment(pathNode)
		line += lineFragment
	}
	line += "   #"
	return line
}

func getLineFragment[T any](node *pathNode[T]) string {
	if node.HasRightSibling {
		return "   #"
	}

	return "    "
}

func lpad(s string, padChar string, n int) string {
	return strings.Repeat(padChar, n) + s
}

func hasRightSibling[T any](node *T, parentNode *T, getChildren func(*T) []*T) bool {
	if parentNode == nil {
		// The root has no sibling
		return false
	}
	siblings := getChildren(parentNode)
	rightMostChild := siblings[len(siblings)-1]
	hasRightSibling := rightMostChild != node
	return hasRightSibling
}
