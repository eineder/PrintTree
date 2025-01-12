package print

import (
	"fmt"
	"strings"
)

const indent int = 4

func Print[T any](node *T, parentNode *T, getChildren func(*T) []*T, getNodeContent func(*T) string) {
	printNode(0, node, parentNode, getChildren, getNodeContent)
}

func printNode[T any](level int, node *T, parentNode *T, getChildren func(*T) []*T, getNodeContent func(*T) string) {

	hasRightSibling := hasRightSibling(node, parentNode, getChildren)

	padding := ""
	if level != 0 {
		padding = lpad("", " ", indent)
	}
	horizontalBranchLine := padding

	for l := 0; l <= level-1; l++ {
		if l == level-1 {
			horizontalBranchLine += "#####"
		} else {

			if hasRightSibling {
				horizontalBranchLine += "#   "
			} else {
				horizontalBranchLine += "    "
			}
		}
	}

	content := getNodeContent(node)
	fmt.Println(horizontalBranchLine + " " + content)

	children := getChildren(node)
	if len(children) == 0 {
		return
	}

	for i := 0; i < len(children); i++ {
		verticalBranchesLine := lpad("", " ", indent)
		for l := 0; l < level; l++ {
			verticalBranchesLine += "#   "
		}
		for i := 0; i < 4; i++ {
			fmt.Println(verticalBranchesLine + "#")
		}
		child := children[i]
		printNode(level+1, child, node, getChildren, getNodeContent)
	}
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
