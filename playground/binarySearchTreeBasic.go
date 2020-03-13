package main

/* from https://www.bogotobogo.com/GoLang/GoLang_Binary_Search_Tree.php
	This tutorial needed some editing as it is broken the way it is written.

 LEARN: how methods (look like functions) can take receivers
https://mail.google.com/mail/u/0/?zx=shklwl6pk5zk#inbox
*/

import (
	"fmt"
)

type Tree struct {
	root *Node
}

type Node struct {
	key   byte
	left  *Node
	right *Node
}

// Tree
func (t *Tree) insert(data byte) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

// Node
func (n *Node) insert(data byte) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func main() {
	var t Tree

	t.insert('F')
	t.insert('B')
	t.insert('A')
	t.insert('D')
	t.insert('C')
	t.insert('E')
	t.insert('G')
	t.insert('I')
	t.insert('H')
	fmt.Println("Print out the pre-ordered binary search tree nodes")
	printPreOrder(t.root)
	fmt.Printf("\n")
	fmt.Println("Now print out the post order BST nodes")
	printPostOrder(t.root)
	fmt.Printf("\n")
	fmt.Println("Now print out the nodes in order")
	printInOrder(t.root)
	fmt.Printf("\n")

}

func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		// tut. had printf format of %d, which just prints out the integer hex value of the ASCII character
		fmt.Printf("%c ", n.key) // also tut. had n.data
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%c ", n.key)
	}
}

func printInOrder(n *Node) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%c ", n.key)
		printInOrder(n.right)
	}
}
