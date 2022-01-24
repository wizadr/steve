package hashtree

import (
	"bytes"

	"github.com/steve-care-software/steve/domain/hash"
)

type parentLeaf struct {
	left  Leaf
	right Leaf
}

func createParentLeaf(left Leaf, right Leaf) ParentLeaf {
	out := parentLeaf{
		left:  left,
		right: right,
	}

	return &out
}

// HashTree returns the hashtree
func (obj *parentLeaf) HashTree() (HashTree, error) {
	data := bytes.Join([][]byte{
		obj.Left().Head().Bytes(),
		obj.Right().Head().Bytes(),
	}, []byte{})

	hsh, err := hash.NewAdapter().FromBytes(data)
	if err != nil {
		return nil, err
	}

	out := createHashTree(*hsh, obj)
	return out, nil
}

// BlockLeaves returns the block leaves
func (obj *parentLeaf) BlockLeaves() Leaves {
	left := obj.Left()
	right := obj.Right()
	leftLeaves := left.Leaves()
	rightLeaves := right.Leaves()
	return leftLeaves.Merge(rightLeaves)
}

// Left returns the left leaf
func (obj *parentLeaf) Left() Leaf {
	return obj.left
}

// Right returns the right leaf
func (obj *parentLeaf) Right() Leaf {
	return obj.right
}
