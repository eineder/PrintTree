package print

import (
	"io/ioutil"
	"strings"
	"testing"
)

type TreeNode struct {
	Content  string
	Children []*TreeNode
}

func getChildren(node *TreeNode) []*TreeNode {
	return node.Children
}

func getNodeContent(node *TreeNode) string {
	return node.Content
}

func getExpectedOutput(testName string) string {
	data, err := ioutil.ReadFile("print_test_expected.txt")
	if err != nil {
		panic(err)
	}
	content := string(data)
	sections := strings.Split(content, "\n\n")
	for _, section := range sections {
		lines := strings.SplitN(section, ":\n", 2)
		if len(lines) == 2 && lines[0] == testName {
			return lines[1] + "\n"
		}
	}
	return ""
}

func TestFormat(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
	}{
		{
			name: "Single node",
			root: &TreeNode{Content: "root"},
		},
		{
			name: "Two levels",
			root: &TreeNode{
				Content: "root",
				Children: []*TreeNode{
					{Content: "child1"},
					{Content: "child2"},
				},
			},
		},
		{
			name: "Three levels",
			root: &TreeNode{
				Content: "root",
				Children: []*TreeNode{
					{
						Content: "child1",
						Children: []*TreeNode{
							{Content: "grandchild1"},
							{Content: "grandchild2"},
						},
					},
					{Content: "child2"},
				},
			},
		},
		{
			name: "Unbalanced tree",
			root: &TreeNode{
				Content: "root",
				Children: []*TreeNode{
					{
						Content: "child1",
						Children: []*TreeNode{
							{Content: "grandchild1"},
						},
					},
					{
						Content: "child2",
						Children: []*TreeNode{
							{Content: "grandchild2"},
							{Content: "grandchild3"},
						},
					},
				},
			},
		},
		{
			name: "Wide tree",
			root: &TreeNode{
				Content: "root",
				Children: []*TreeNode{
					{Content: "child1"},
					{Content: "child2"},
					{Content: "child3"},
					{Content: "child4"},
				},
			},
		},
		{
			name: "Deep tree",
			root: &TreeNode{
				Content: "root",
				Children: []*TreeNode{
					{
						Content: "child1",
						Children: []*TreeNode{
							{
								Content: "child2",
								Children: []*TreeNode{
									{
										Content: "child3",
										Children: []*TreeNode{
											{
												Content: "child4",
												Children: []*TreeNode{
													{
														Content: "child5",
														Children: []*TreeNode{
															{
																Content: "child6",
																Children: []*TreeNode{
																	{
																		Content: "child7",
																		Children: []*TreeNode{
																			{
																				Content: "child8",
																				Children: []*TreeNode{
																					{
																						Content: "child9",
																						Children: []*TreeNode{
																							{Content: "child10"},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Complex tree",
			root: &TreeNode{
				Content: "root",
				Children: []*TreeNode{
					{
						Content: "child1",
						Children: []*TreeNode{
							{Content: "grandchild1"},
							{
								Content: "grandchild2",
								Children: []*TreeNode{
									{Content: "greatgrandchild1"},
									{Content: "greatgrandchild2"},
								},
							},
						},
					},
					{
						Content: "child2",
						Children: []*TreeNode{
							{Content: "grandchild3"},
							{
								Content: "grandchild4",
								Children: []*TreeNode{
									{Content: "greatgrandchild3"},
									{Content: "greatgrandchild4"},
								},
							},
						},
					},
					{Content: "child3"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := getExpectedOutput(tt.name)
			output := Format(tt.root, getChildren, getNodeContent)
			if output != expected {
				t.Errorf("expected %q, got %q", expected, output)
			}
		})
	}
}
