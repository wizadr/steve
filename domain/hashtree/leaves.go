package hashtree

type leaves struct {
	list []Leaf
}

func createLeaves(list []Leaf) Leaves {
	out := leaves{
		list: list,
	}

	return &out
}

// Leaves returns the leaves
func (obj *leaves) Leaves() []Leaf {
	out := []Leaf{}
	for _, oneLeaf := range obj.list {
		out = append(out, oneLeaf)
	}

	return out
}

// Merge merge Leaves instances
func (obj *leaves) Merge(lves Leaves) Leaves {
	for _, oneLeaf := range lves.Leaves() {
		obj.list = append(obj.list, oneLeaf.(*leaf))
	}

	return obj
}

// HashTree returns the hashtree
func (obj *leaves) HashTree() (HashTree, error) {
	length := len(obj.list)
	if length == 2 {
		left := obj.list[0]
		right := obj.list[1]
		parent := createParentLeaf(left, right)
		return parent.HashTree()
	}

	childrenLeaves, err := obj.createChildrenLeaves()
	if err != nil {
		return nil, err
	}

	return childrenLeaves.HashTree()
}

func (obj *leaves) createChildrenLeaves() (Leaves, error) {
	childrenLeaves := []Leaf{}
	for index, oneLeaf := range obj.list {

		if index%2 != 0 {
			continue
		}

		left := oneLeaf
		right := obj.list[index+1]
		child, err := createChildLeaf(left, right)
		if err != nil {
			return nil, err
		}

		parent := createParentLeaf(left, right)
		childWithParent := createLeafWithParent(child.Head(), parent)
		childrenLeaves = append(childrenLeaves, childWithParent)
	}

	return createLeaves(childrenLeaves), nil
}
