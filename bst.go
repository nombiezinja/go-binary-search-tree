package gobst

type Node struct {
	left  *Node
	right *Node
	value float64
	// Add whatever you want here
	//...
	//...
}

type Bst struct {
	root *Node
	size float64
}

/* Get size of the tree */
func (bst *Bst) Size() float64 {
	return bst.size
}

/* Get size of the tree */
func (bst *Bst) Root() *Node {
	return bst.root
}

// Construtor
func NewBst() *Bst {
	bst := new(Bst)
	bst.size = 0
	return bst
}
