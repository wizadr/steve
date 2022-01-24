package hashtree

import (
	"bytes"

	"github.com/steve-care-software/steve/domain/hash"
)

type leaf struct {
	head       hash.Hash
	parentLeaf ParentLeaf
}

func createLeafWithParent(head hash.Hash, parent ParentLeaf) Leaf {
	out := leaf{
		head:       head,
		parentLeaf: parent,
	}

	return &out
}

func createLeaf(head hash.Hash) Leaf {
	out := leaf{
		head:       head,
		parentLeaf: nil,
	}

	return &out
}

func createChildLeaf(left Leaf, right Leaf) (Leaf, error) {
	data := bytes.Join([][]byte{
		left.Head().Bytes(),
		right.Head().Bytes(),
	}, []byte{})

	h, err := hash.NewAdapter().FromBytes(data)
	if err != nil {
		return nil, err
	}

	out := createLeaf(*h)
	return out, nil
}

// Head returns the head hash
func (obj *leaf) Head() hash.Hash {
	return obj.head
}

// HasParent returns true if there is a parent, false otherwise
func (obj *leaf) HasParent() bool {
	return obj.parentLeaf != nil
}

// Parent returns the parent, if any
func (obj *leaf) Parent() ParentLeaf {
	return obj.parentLeaf
}

// Leaves returns the leaves
func (obj *leaf) Leaves() Leaves {
	if obj.HasParent() {
		return obj.Parent().BlockLeaves()
	}

	leaves := []Leaf{
		obj,
	}

	output := createLeaves(leaves)
	return output
}

// Height returns the leaf height
func (obj *leaf) Height() int {
	cpt := 0
	var oneLeaf Leaf
	for {

		if oneLeaf == nil {
			oneLeaf = obj
		}

		if !oneLeaf.HasParent() {
			return cpt
		}

		cpt++
		oneLeaf = oneLeaf.Parent().Left()
	}
}
