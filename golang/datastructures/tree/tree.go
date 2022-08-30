package tree

import "github.com/sidheart/algorithms/util"

// A Tree simulates a hierarchical tree structure, with a root value and subtrees of children with a parent node
type Tree interface {
	// Inserts data into the tree with a particular ordering key. It is recommended to use unique keys for each node.
	Insert(key util.Comparable, data interface{})
	// Removes the first element matching the key from the tree, and returns the deleted data or (nil, false) if none
	// is found.
	Delete(key util.Comparable) (interface{}, bool)
	// Returns the first found element matching the given key, or (nil, false) if none is found.
	Search(key util.Comparable) (interface{}, bool)
	// Applies the function f to every element in the tree. Visitation order is dependent on implementation.
	Traverse(f func(interface{}))
}