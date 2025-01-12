package main

import print "github.com/eineder/printtree/print"

type Node struct {
	Content  string
	Children []*Node
}

func main() {
	// Example usage of Print function
	root := Node{
		Content: "root",
		Children: []*Node{
			{
				Content: "child1",
				Children: []*Node{
					{Content: "child1.1",
						Children: []*Node{
							{Content: "child1.1.1"},
							{Content: "child1.1.2"},
						}},
					{Content: "child1.2",
						Children: []*Node{
							{Content: "child1.2.1"},
							{Content: "child1.2.2"},
						}},
				},
			},
			// {
			// 	Content: "child2",
			// 	Children: []*Node{
			// 		{Content: "child2.1"},
			// 		{Content: "child2.2"},
			// 	},
			// },
		},
	}

	getChildren := func(node *Node) []*Node {
		return node.Children
	}

	getNodeContent := func(node *Node) string {
		return node.Content
	}

	print.Print(&root, nil, getChildren, getNodeContent)
}
