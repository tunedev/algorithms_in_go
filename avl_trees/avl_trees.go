package avltrees

import (
	"fmt"
	"math"
)

// In a balanced tree the difference between the height of the left tree and right tree should be less 1,
// 			height(left) - height(right) <= 1

// AVL Trees is an example of self balancing trees, below are the list of self balancing trees inview
// AVL ==> (Adelson-Vesky and Landis)(Easiest to understand)
// Red-black trees, B-trees, Splay Trees, 2-3 Trees

// AVL trees are a special type of BST that automatically rebalance every time we add or remove nodes
// They do this by ensuring the defference between the height of the left and right sub trees are never more than one(1)

// Rotations:
//  ===>> Left(LL)
//  ===>> Right(RR)
//  ===>> Left-Right(LR)
//  ===>> Right-Left(RL)

type AVLNode struct {
	val int
	right *AVLNode
	left *AVLNode
	height int
}

type AVLTree struct {
	root *AVLNode
}

func (n *AVLNode) String() string {
	return fmt.Sprintf("Value: %v", n.val)
}

func NewAVLTree() *AVLTree{
	return &AVLTree{}
}

func (t *AVLTree) Insert(val int) {
	t.root = t.insert(t.root, val)
}

func (t *AVLTree) insert(root *AVLNode, val int) *AVLNode {
	if root == nil {
		return &AVLNode{val: val}
	}

	if val < root.val {
		root.left = t.insert(root.left, val)
	}else {
		root.right = t.insert(root.right, val)
	}

	t.setHeight(root)
	
	return t.balance(root)
}

func (t *AVLTree) balance(root *AVLNode) *AVLNode {
	if t.isLeftHeavy(root) {
		if t.getBalanceFactor(root.left) < 0 {
			fmt.Printf("Left Rotate Node: %v\n", root.left)
			root.left = t.leftRotate(root.left)
		}
		fmt.Printf("Right Rotate Node: %v\n", root)
		root = t.rightRotate(root)
		} else if t.isRightHeavy(root) {
			if t.getBalanceFactor(root.right) > 0 {
				fmt.Printf("Right Rotate Node: %v\n", root.right)
				root.right = t.rightRotate(root.right)
			}
		fmt.Printf("Left Rotate Node: %v\n", root)
		root = t.leftRotate(root)
	}

	return root
}

func (t *AVLTree) leftRotate(root *AVLNode) *AVLNode {
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root

	t.setHeight(root)
	t.setHeight(newRoot)
	return newRoot
}

func (t *AVLTree) rightRotate(root *AVLNode) *AVLNode {
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = root

	t.setHeight(root)
	t.setHeight(newRoot)

	return newRoot
}

func (t *AVLTree) isLeftHeavy(node *AVLNode) bool {
	return t.getBalanceFactor(node) > 1
}

func (t *AVLTree) isRightHeavy(node *AVLNode) bool {
	return t.getBalanceFactor(node) < -1
}

func (t *AVLTree) getBalanceFactor(node *AVLNode) int {
	return t.getHeight(node.left) - t.getHeight(node.right)
}

func (t *AVLTree) setHeight(node *AVLNode) {
	node.height = int(math.Max(float64(t.getHeight(node)), float64(t.getHeight(node)))) + 1
}

func (t *AVLTree) getHeight(node *AVLNode) int {
	if node == nil {
		return -1
	}
	return node.height
}

func (t *AVLTree) string(node *AVLNode, prefix string, isTail bool, result *string) {
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

func (t *AVLTree) String() string {
    var result string
    t.string(t.root, "", true, &result)
    return result
}
