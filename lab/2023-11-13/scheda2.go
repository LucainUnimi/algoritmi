package main

import "fmt"

type Node struct {
	value    string
	parent   *Node
	children []*Node
	level    int
}

type Tree = *Node

type Trees = []Tree

func newNode(root Tree, value string, level int) *Node {
	return &Node{
		value:    value,
		parent:   root,
		children: nil,
		level:    level,
	}
}

func initTree(value string) Tree {
	return &Node{
		value:    value,
		parent:   nil,
		children: nil,
		level:    0,
	}
}

func addChild(root Tree, value string) {
	root.children = append(root.children, newNode(root, value, root.level+1))
}

func printChild(root Tree) {
	for _, sub := range root.children {
		fmt.Printf("%s ", sub.value)
	}
}

func search(root Tree, name string) Tree {
	if root.value == name {
		return root
	}

	for _, child := range root.children {
		v := search(child, name)
		if v != nil {
			return v
		}
	}

	return nil
}

func stampaSubordinati(roots Trees, name string) {
	for _, root := range roots {
		v := search(root, name)
		if v != nil {
			printChild(v)
		}
	}
}

func countLeaf(root Tree) int {
	var c int = 0
	if root.children == nil {
		return 1
	}
	for _, child := range root.children {
		c += countLeaf(child)
	}
	return c
}

func quantiSenzaSubordinati(roots Trees) int {
	var c int = 0
	for _, root := range roots {
		c += countLeaf(root)
	}
	return c
}

func supervisore(roots Trees, name string) string {
	for _, root := range roots {
		v := search(root, name)
		if v != nil {
			return v.parent.value
		}
	}
	return ""
}

func stampaSuperiori(roots Trees, name string) {
	var v Tree
	for _, root := range roots {
		v = search(root, name)
		if v != nil {
			break
		}
	}
	for v.parent != nil {
		v = v.parent
		fmt.Printf("%v ", v.value)
	}
}

func main() {

	var Gerarchie Trees

	Gerarchie = append(Gerarchie, initTree("Anna"))
	Gerarchie = append(Gerarchie, initTree("Gianni"))

	addChild(Gerarchie[0], "Bruno")
	addChild(Gerarchie[0], "Carlo")
	addChild(Gerarchie[0], "Daniela")

	addChild(Gerarchie[0].children[0], "Enrico")
	addChild(Gerarchie[0].children[0], "Francesco")

	addChild(Gerarchie[0].children[0].children[1], "Irene")

	addChild(Gerarchie[1], "Harry")

	stampaSubordinati(Gerarchie, "Francesco")
	fmt.Println()
	fmt.Printf("quantiSenzaSubordinati(Gerarchie): %v\n", quantiSenzaSubordinati(Gerarchie))
	fmt.Printf("supervisore(Gerarchie, \"Enrico\"): %v\n", supervisore(Gerarchie, "Enrico"))
	stampaSuperiori(Gerarchie, "Irene")
	fmt.Println()

}
