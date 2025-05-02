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
//	useBranchSymbols - a boolean indicating whether to use branch symbols in the output
//
// Returns:
//
//	A string representation of the tree structure.
func Format[T any](node *T, getChildren func(*T) []*T, getNodeContent func(*T) string, useBranchSymbols bool) string {
	var path []*pathNode[T] = []*pathNode[T]{}
	return printNode(node, path, getChildren, getNodeContent, useBranchSymbols)
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
//	useBranchSymbols - a boolean indicating whether to use branch symbols in the output
func Print[T any](node *T, getChildren func(*T) []*T, getNodeContent func(*T) string, useBranchSymbols bool) {
	output := Format(node, getChildren, getNodeContent, useBranchSymbols)
	fmt.Print(output)
}

func printNode[T any](node *T, path []*pathNode[T], getChildren func(*T) []*T, getNodeContent func(*T) string, useBranchSymbols bool) string {
	var result strings.Builder
	line := getLine(node, path, useBranchSymbols)
	isRoot := len(path) == 0
	content := getNodeContent(node)
	if isRoot {
		result.WriteString(content + "\n")
	} else {
		result.WriteString(line + content + "\n")
	}

	children := getChildren(node)
	isLeaf := len(children) == 0
	if isLeaf {
		return result.String()
	}

	for childIndex := 0; childIndex < len(children); childIndex++ {
		child := children[childIndex]
		hasRightHandSibling := childIndex < len(children)-1
		pathNode := &pathNode[T]{Node: node, HasRightSibling: hasRightHandSibling}
		newPath := append(path, pathNode)
		result.WriteString(printNode(child, newPath, getChildren, getNodeContent, useBranchSymbols))
	}
	return result.String()
}

func getLine[T any](node *T, path []*pathNode[T], useBranchSymbols bool) string {
	if !useBranchSymbols {
		return strings.Repeat("    ", len(path))
	}

	line := ""
	for i := 0; i < len(path); i++ {
		pathNode := path[i]
		lineFragment := getLineFragment(pathNode, i == len(path)-1)
		line += lineFragment
	}
	return line
}

func getLineFragment[T any](node *pathNode[T], isLast bool) string {
	if node.HasRightSibling {
		if isLast {
			return "├── "
		}
		return "│   "
	}
	if isLast {
		return "└── "
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
