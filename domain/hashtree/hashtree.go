package hashtree

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

// HTree represents a concrete HashTree implementation
type hashtree struct {
	head       hash.Hash
	parentLeaf ParentLeaf
}

func createHashTree(head hash.Hash, parent ParentLeaf) HashTree {
	out := hashtree{
		head:       head,
		parentLeaf: parent,
	}

	return &out
}

func createHashTreeFromBlocks(blocks [][]byte) (HashTree, error) {
	blockHashes, blockHashesErr := createBlockFromData(blocks)
	if blockHashesErr != nil {
		return nil, blockHashesErr
	}

	return blockHashes.HashTree()
}

// Height returns the hashtree height
func (obj *hashtree) Height() int {
	left := obj.parentLeaf.Left()
	return left.Height() + 2
}

// Length returns the hashtree length
func (obj *hashtree) Length() int {
	blockLeaves := obj.parentLeaf.BlockLeaves()
	return len(blockLeaves.Leaves())
}

// Head returns the head hash
func (obj *hashtree) Head() hash.Hash {
	return obj.head
}

// Parent returns the parent leaf
func (obj *hashtree) Parent() ParentLeaf {
	return obj.parentLeaf
}

// Compact returns the compact version of the hashtree
func (obj *hashtree) Compact() Compact {
	blockLeaves := obj.parentLeaf.BlockLeaves()
	return createCompact(obj.head, blockLeaves)
}

// Order orders data that matches the leafs of the HashTree
func (obj *hashtree) Order(data [][]byte) ([][]byte, error) {
	hashAdapter := hash.NewAdapter()
	hashed := map[string][]byte{}
	for _, oneData := range data {
		hsh, err := hashAdapter.FromBytes(oneData)
		if err != nil {
			return nil, err
		}

		hashAsString := hsh.String()
		hashed[hashAsString] = oneData
	}

	out := [][]byte{}
	leaves := obj.parentLeaf.BlockLeaves().Leaves()
	for _, oneLeaf := range leaves {
		leafHashAsString := oneLeaf.Head().String()
		if oneData, ok := hashed[leafHashAsString]; ok {
			out = append(out, oneData)
			continue
		}

		//must be a filling Leaf, so continue:
		continue
	}

	if len(out) != len(data) {
		str := fmt.Sprintf("the length of the input data (%d) does not match the length of the output (%d), therefore, some data blocks could not be found in the hash leaves", len(data), len(out))
		return nil, errors.New(str)
	}

	return out, nil
}
