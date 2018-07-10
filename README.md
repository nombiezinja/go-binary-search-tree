# Binary search trees in Go

Simple implementations of a binary search tree written in GoLang. For more information: <a href="http://en.wikipedia.org/wiki/Binary_search_tree" target="_blank">basic</a>

<p align="center">
  <img src="https://github.com/m2omou/go-binary-search-tree/raw/master/bst.png" alt="Her" />
</p>

Forked from https://github.com/m2omou/go-binary-search-tree/ (thanks for the good work), changes made to allow importing as a package:
* remove `main.go`
* change package name `from` main to `gobst`



## Installation

```cmd
$>go get github.com/nombiezinja/go-binary-search-tree
```
```go
import "github.com/nombiiezinja/go-binary-search-tree"
```

## Methods

### insert(item)
> Inserts the item into the tree. Returns true if inserted, false if duplicate.

### remove(item)
> Removes the item from the tree. Returns true if removed, false if not found.

### size()
> Number of nodes in the tree.

### Exists()
> Returns true if the key exists or return false.

### showInOrder()
> Print the tree in-order.

### showPreOrder()
> Print the tree pre-order.

### showPostOrder()
> Print the tree post-order.

### Tree traversal

```go
	// Print the tree in-order
	// Traverse the left sub-tree, root, right sub-tree
	showInOrder(test.root)

	// Print the tree pre-order
	// Traverse the root, left sub-tree, right sub-tree
	showPreOrder(test.root)

	// Print the tree post-order
	// Traverse left sub-tree, right sub-tree, root
	showPostOrder(test.root)
```

<p align="center">
  <img src="https://github.com/m2omou/go-binary-search-tree/raw/master/pre_post_in_order.png" alt="Her" />
</p>
