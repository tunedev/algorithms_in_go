package binarytree

import (
	"fmt"
	"math"
)

// Tree is a non linear datastructure that has many applications
// DB, graphical user interfaces, websites are example use cases
// It stores elements in an hierarchical structure, the elements are refered to as nodes, and the conection between elements are called edges
// The initial element is refered to as the root node, while the elements are the bottom of the hierarchy without children are refered to as leaf nodes
// in a Binary tree each nodes can only have two children

// Below are the application of trees considerations
// - Represent hierarchical data
// - Databases - indexing for data quick lookup
// - Autocompletion
// - Compilers - syntax tree to parse expressions
// - Compression algorithms(JPEG,MP3)

// Tree operations and their runtime
// Initial considerations for binarySearchTree
// **** Left < Node < Right

// Lookup O(log n)
// Insert O(log n)
// Delete O(log n)

type node struct {
	val int
	left *node
	right *node
}

func (n *node) String() string {
	return fmt.Sprintf("Node: %v", n.val)
}

type BinarySearchTree struct {
	root *node
}

func NewBST() *BinarySearchTree{
	return &BinarySearchTree{}
}

func (t *BinarySearchTree) Insert(val int) {
	newNode := &node{
		val: val,
	}
	if t.root == nil {
		t.root = newNode
		return
	}

	current := t.root
	for {
		if val < current.val{
			if current.left == nil {
				current.left = newNode
				break
			}
			current = current.left
		}else {
			if current.right == nil {
				current.right = newNode
				break
			}
			current = current.right
		}
	}
}

func (t *BinarySearchTree) Find(val int) bool {
	current := t.root
	for current != nil {
		if val < current.val {
			current = current.left
		}else if val > current.val {
			current = current.right
		}else {
			return true
		}
	}
	return false
}

func (t *BinarySearchTree) PrintTreeElementPreOrder() {
	t.printTreeElementPreOrder(t.root)
}

func (t *BinarySearchTree) printTreeElementPreOrder(root *node) {
	if root == nil {
		return
	}

	fmt.Println(root.val)
	t.printTreeElementPreOrder(root.left)
	t.printTreeElementPreOrder(root.right)
}

func (t *BinarySearchTree) PrintTreeElementPostOrder() {
	t.printTreeElementPostOrder(t.root)
}

func (t *BinarySearchTree) printTreeElementPostOrder(root *node) {
	if root == nil {
		return
	}

	t.printTreeElementPreOrder(root.left)
	t.printTreeElementPreOrder(root.right)
	fmt.Println(root.val)
}

func (t *BinarySearchTree) PrintTreeElementInOrder() {
	t.printTreeElementInOrder(t.root)
}

func (t *BinarySearchTree) printTreeElementInOrder(root *node) {
	if root == nil {
		return
	}

	t.printTreeElementInOrder(root.left)
	fmt.Println(root.val)
	t.printTreeElementInOrder(root.right)
}

func (t *BinarySearchTree) height(root *node) int {
	if root == nil {
		return -1
	}else if t.isLeaf(root) {
		return 0
	}
	return 1 + int(math.Max(float64((t.height(root.left))), float64(t.height(root.right))))
}

func (t *BinarySearchTree) Height() int {
	return t.height(t.root)
}

func (t *BinarySearchTree) isLeaf(root *node) bool {
	return root.left == nil && root.right == nil
}

// O(N)
func (t *BinarySearchTree) min(root *node) int {
	if t.isLeaf(root) {
		return root.val
	}

	left := float64(t.min(root.left))
	right := float64(t.min(root.right))

	return int(math.Min(float64(root.val), math.Min(left, right)))
}

// O(log N)
func (t *BinarySearchTree) Min() int {
	if t.root == nil {
		return math.MinInt
	}

	cur := t.root
	last := cur

	for cur != nil {
		last = cur
		cur = cur.left
	}

	return last.val
}

func (t *BinarySearchTree) equals(first *node, second *node) bool {
	if first == nil && second == nil {
		return true
	}

	if first != nil && second != nil {
		return first.val == second.val && 
			t.equals(first.left, second.left) && 
			t.equals(first.right, second.right)
	}

	return false
}

func (t *BinarySearchTree) Equals(second *BinarySearchTree) bool {
	if second == nil {
		return false
	}
	return t.equals(t.root, second.root)
}

func (t *BinarySearchTree) IsBST() bool {
	return t.isBST(t.root, math.MinInt, math.MaxInt)
}

func (t *BinarySearchTree) isBST(root *node, min, max int) bool {
	if root == nil {
		return true
	}

	if root.val < min || root.val > max {
		return false
	}

	return t.isBST(root.left, min, root.val - 1) && t.isBST(root.right, root.val + 1, max)
}

func (t *BinarySearchTree) PrintNodeAt(kth int) {
	t.printNodeAt(t.root, kth)
}

func (t *BinarySearchTree) printNodeAt(root *node, kth int) {
	if root == nil {
		return
	}
	if kth == 0 {
		fmt.Println(root)
		return
	}

	t.printNodeAt(root.left, kth - 1)
	t.printNodeAt(root.right, kth - 1)
}

func (t *BinarySearchTree) PrintLevelOrderTraversal() {
	len := t.Height()
	for i := 0; i <= len; i++ {
		t.PrintNodeAt(i)
	}
}

// func (t *BinarySearchTree) string(root *node, space int, result *string) {
//     if root == nil {
//         return
//     }

//     // Increase distance between levels
//     space += 10

//     // Process right child first
//     t.string(root.right, space, result)

//     // Print current node after space count
//     *result += fmt.Sprintln()
//     for i := 10; i < space; i++ {
//         *result += " "
//     }
//     *result += fmt.Sprintln(root.val)

//     // Process left child
//     t.string(root.left, space, result)
// }

// func (t *BinarySearchTree) String() string {
// 	var result string
//   t.string(t.root, 0, &result)
// 	return result
// }

func (t *BinarySearchTree) string(node *node, prefix string, isTail bool, result *string) {
    if node == nil {
        return
    }

    newPrefix := prefix
    if isTail {
        newPrefix += "    "
    } else {
        newPrefix += "│   "
    }

    if node.right != nil {
        t.string(node.right, newPrefix, false, result)
    }

    *result += prefix
    if isTail {
        *result += "└── "
    } else {
        *result += "├── "
    }
    *result += fmt.Sprintln(node.val)

    if node.left != nil {
        t.string(node.left, newPrefix, true, result)
    }
}

func (t *BinarySearchTree) String() string {
    var result string
    t.string(t.root, "", true, &result)
    return result
}
