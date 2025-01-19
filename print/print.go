package print

import (
	"fmt"
	"strings"
)

type pathNode[T any] struct {
	Node            *T
	HasRightSibling bool
}

// Generates a string representation of a tree structure starting from the given node.
// It uses the provided getChildren function to retrieve the children of a node and the
// getNodeContent function to get the content of a node.
//
// Type Parameters:
//
//	T - the type of the node
//
// Parameters:
//
//	node - the root node of the tree to format
//	getChildren - a function that returns the children of a given node
//	getNodeContent - a function that returns the content of a given node as a string
//
// Returns:
//
//	A string representation of the tree structure.
func Format[T any](node *T, getChildren func(*T) []*T, getNodeContent func(*T) string) string {
	var path []*pathNode[T] = []*pathNode[T]{}
	return printNode(node, path, getChildren, getNodeContent)
}

// Prints a string representation of a tree structure starting from the given node.
// It uses the provided getChildren function to retrieve the children of a node and the
// getNodeContent function to get the content of a node.
//
// Type Parameters:
//
//	T - the type of the node
//
// Parameters:
//
//	node - the root node of the tree to format
//	getChildren - a function that returns the children of a given node
//	getNodeContent - a function that returns the content of a given node as a string
//
// Returns:
//
//	A string representation of the tree structure.
func Print[T any](node *T, getChildren func(*T) []*T, getNodeContent func(*T) string) {
	output := Format(node, getChildren, getNodeContent)
	fmt.Print(output)
}

func printNode[T any](node *T, path []*pathNode[T], getChildren func(*T) []*T, getNodeContent func(*T) string) string {
	var result strings.Builder
	line := getLine(node, path)
	isRoot := len(path) == 0
	if !isRoot {
		result.WriteString(line + "\n")
		result.WriteString(line + "\n")
		result.WriteString(line + "\n")
		result.WriteString(line + "\n")
	}
	content := getNodeContent(node)
	if isRoot {
		result.WriteString("#### " + content + "\n")
	} else {
		result.WriteString(line + "#### " + content + "\n")
	}

	children := getChildren(node)
	isLeaf := len(children) == 0
	if isLeaf {
		return result.String()
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
		result.WriteString(printNode(child, newPath, getChildren, getNodeContent))
	}
	return result.String()
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
