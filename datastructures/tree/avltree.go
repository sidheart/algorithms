package tree

import (
	"fmt"
	"github.com/sidheart/algorithms/util"
)

// An AVLTreeNode represents a leaf or internal node in an AVL tree
type AVLTreeNode struct {
	key			  					util.Comparable
	balanceFactor		 			int8
	data 							interface{}
	leftChild, rightChild, parent 	*AVLTreeNode
}

// An AVLTree is a kind of self-balancing binary tree https://en.wikipedia.org/wiki/AVL_tree
type AVLTree struct {
	root *AVLTreeNode
}

// rotateLeft makes the given root the left subtree of its right child, effectively rotating the tree
// this operation may be better explained by the visualization here https://en.wikipedia.org/wiki/AVL_tree#Rebalancing
func (root *AVLTreeNode) rotateLeft() {
	if root == nil || root.rightChild == nil {
		return
	}
	newRoot := root.rightChild
	newRightSubtree := newRoot.leftChild
	newRoot.leftChild = root
	newRoot.parent = root.parent
	root.parent = newRoot
	if newRightSubtree != nil {
		root.rightChild = newRightSubtree
		newRightSubtree.parent = root
	}
}

// rotateRight makes the given root the right subtree of its left child, effectively rotating the tree
// this operation may be better explained by the visualization here https://en.wikipedia.org/wiki/AVL_tree#Rebalancing
func (root *AVLTreeNode) rotateRight() {
	if root == nil || root.leftChild == nil {
		return
	}
	newRoot := root.leftChild
	newLeftSubtree := newRoot.rightChild
	newRoot.rightChild = root
	newRoot.parent = root.parent
	root.parent = newRoot
	if newLeftSubtree != nil {
		root.leftChild = newLeftSubtree
		newLeftSubtree.parent = root
	}
}

// updateBalanceFactorsAfterDoubleRotate updates the balance factors of the relevant nodes which are modified during a
// left-right or right-left rotation. The root receiver should be the new node at the location where the AVL tree
// invariant was previously violated.
func (root *AVLTreeNode) updateBalanceFactorsAfterDoubleRotate() {
	if root.balanceFactor == -1 {
		root.leftChild.balanceFactor = 0
		root.rightChild.balanceFactor = 1
	} else if root.balanceFactor == 0 {
		root.leftChild.balanceFactor = 0
		root.rightChild.balanceFactor = 0
	} else {
		root.leftChild.balanceFactor = -1
		root.rightChild.balanceFactor = 0
	}
	root.balanceFactor = 0
}

// rotate fixes a violation of the AVL tree invariant at the given root
// the behavior is undefined if the given root does not violate the AVL tree invariant
func (root *AVLTreeNode) rotate() *AVLTreeNode {
	if root == nil {
		return nil
	}
	if root.balanceFactor == -1 { // Because rotate() was called, this means that root is doubly left-heavy
		if root.leftChild.balanceFactor == 1 { // Left child is right-heavy, requires a double rotation
			root.leftChild.rotateLeft()
			root.rotateRight()
			root.parent.updateBalanceFactorsAfterDoubleRotate()
			return nil
		} else { // A simple right rotation will fix this violation
			root.rotateRight()
			// Update balance factors
			newRoot := root.parent
			if newRoot.balanceFactor == 0 {
				newRoot.balanceFactor = 1
				return newRoot
			} else { // The previous balance factor must have been -1, if it were 1, we would have called rotateLeftRight()
				newRoot.balanceFactor = 0
				root.balanceFactor = 0
				return nil
			}
		}
	} else if root.balanceFactor == 1 { // Because rotate() was called, this means that root is doubly right-heavy
		if root.rightChild.balanceFactor == -1 { // Right child is left-heavy, requires a double rotation
			root.rightChild.rotateRight()
			root.rotateLeft()
			root.parent.updateBalanceFactorsAfterDoubleRotate()
			return nil // The AVL tree invariant now holds for the entire tree
		} else {  // A simple left rotation will fix this violation
			root.rotateLeft()
			newRoot := root.parent
			// Update balance factors
			if newRoot.balanceFactor == 0 {
				newRoot.balanceFactor = -1
				return newRoot
			} else { // The previous balance factor must have been 1, if it were -1, we would have called rotateRightLeft()
				newRoot.balanceFactor = 0
				root.balanceFactor = 0
				return nil
			}
		}
	} else { // root does not violate the AVL tree invariant, rotate() should not have been called
		fmt.Println("rotate() called on non-violating receiver")
		return nil
	}
}

// retrace updates the balance factors of nodes along a leaf to root path and performs the necessary rotations to
// ensure that the AVL tree invariant holds
func (root *AVLTreeNode) retrace(addend int8) {
	if root == nil {
		return
	}
	var newRoot *AVLTreeNode
	tempBalance := root.balanceFactor + addend
	if tempBalance > 1 || tempBalance < -1 { // Violation of the invariant, rotation is necessary
		newRoot = root.rotate()
	} else {
		newRoot = root
		newRoot.balanceFactor = tempBalance
	}
	// Continue retracing up to the root, or until a particular subtree's height does not change
	if newRoot == nil || newRoot.balanceFactor == 0 || newRoot.parent == nil {
		return
	} else if newRoot.parent.leftChild == newRoot {
		newRoot.parent.retrace(-1)
	} else {
		newRoot.parent.retrace(1)
	}
}

// search performs a binary search for data matching the key in O(log(n)) time
func (root *AVLTreeNode) search(key util.Comparable) (data interface{}, ok bool) {
	if root == nil {
		return
	}
	rootKey := root.key
	if key.Compare(rootKey) == 0 {
		return root.data, true
	} else if key.Compare(rootKey) < 0 {
		return root.leftChild.search(key)
	} else {
		return root.rightChild.search(key)
	}
}

// traverse applies f to every element of the tree in order
func (root *AVLTreeNode) traverse(f func(interface{})) {
	if root == nil {
		return
	}
	if root.leftChild != nil {
		root.leftChild.traverse(f)
	}
	f(root.data)
	if root.rightChild != nil {
		root.rightChild.traverse(f)
	}
}

// insert inserts the given key and data into the tree in O(log(n)) time
func (root *AVLTreeNode) insert(key util.Comparable, data interface{}) {
	if root == nil {
		return
	}
	rootKey := root.key
	if key.Compare(rootKey) <= 0 {
		if root.leftChild == nil {
			root.leftChild = &AVLTreeNode{key, 0, data, nil, nil, root}
			root.retrace(-1)
		} else {
			root.leftChild.insert(key, data)
		}
	} else {
		if root.rightChild == nil {
			root.rightChild = &AVLTreeNode{key, 0, data, nil, nil, root}
			root.retrace(1)
		} else {
			root.rightChild.insert(key, data)
		}
	}
}

// delete deletes the first instance of data matching the given key from the tree in O(log(n)) time
func (root *AVLTreeNode) delete(key util.Comparable) (deleted interface{}, ok bool) {
	if root == nil {
		return
	}
	rootKey := root.key
	if key.Compare(rootKey) == 0 {
		parent := root.parent
		if root.leftChild == nil && root.rightChild == nil { // This node is a leaf, fairly simple to delete
			if parent != nil {
				if parent.leftChild == root {
					parent.leftChild = nil
				} else {
					parent.rightChild = nil
				}
			}
		} else if root.leftChild != nil { // Replace the node with its left child
			newRoot := root.leftChild
			root.rightChild.parent = newRoot
			newRoot.parent = parent
			if parent != nil {
				if parent.leftChild == root {
					parent.leftChild = newRoot
				} else {
					parent.rightChild = newRoot
				}
			}
		} else { // Replace the node with its right child
			newRoot := root.rightChild
			root.leftChild.parent = newRoot
			newRoot.parent = parent
			if parent != nil {
				if parent.leftChild == root {
					parent.leftChild = newRoot
				} else {
					parent.rightChild = newRoot
				}
			}
		}
		return root.data, true
	} else if key.Compare(rootKey) < 0 {
		return root.leftChild.delete(key)
	} else {
		return root.rightChild.delete(key)
	}
}

// Search performs a binary search for data matching the key in O(log(n)) time
func (tree AVLTree) Search(key util.Comparable) (data interface{}, ok bool) {
	return tree.root.search(key)
}

func (tree AVLTree) Traverse(f func(interface{})) {
	tree.root.traverse(f)
}

func (tree AVLTree) Insert(key util.Comparable, data interface{}) {
	if tree.root == nil {
		tree.root = &AVLTreeNode{key, 0, data, nil, nil, nil}
	} else {
		tree.root.insert(key, data)
	}
}
