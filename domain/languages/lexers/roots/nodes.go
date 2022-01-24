package roots

import (
	"errors"
	"fmt"
)

type nodes struct {
	list map[string]Node
}

func createNodes(
	list map[string]Node,
) Nodes {
	out := nodes{
		list: list,
	}

	return &out
}

// All returns the nodes
func (obj *nodes) All() map[string]Node {
	return obj.list
}

// Amount returns the amount of nodes
func (obj *nodes) Amount() uint {
	return uint(len(obj.list))
}

// Fetch fetches the node by name
func (obj *nodes) Fetch(name string) (Node, error) {
	if node, ok := obj.list[name]; ok {
		return node, nil
	}

	str := fmt.Sprintf("the node (name: %s) does not exists", name)
	return nil, errors.New(str)
}

// Scan scans the node
func (obj *nodes) Scan(name string) (Root, Node, Container, string, error) {
	for _, node := range obj.list {
		root, subNode, container, content, err := node.Scan(name)
		if err != nil {
			continue
		}

		if root != nil {
			return root, nil, nil, "", nil
		}

		if subNode != nil {
			return nil, subNode, nil, "", nil
		}

		if container != nil {
			return nil, nil, container, "", nil
		}

		if content != "" {
			return nil, nil, nil, content, nil
		}
	}

	str := fmt.Sprintf("the element (name: %s) could not be found in the current Nodes instance", name)
	return nil, nil, nil, "", errors.New(str)
}

// Exists returns true if the propertyexists, false otherwise
func (obj *nodes) Exists(name string) bool {
	if _, ok := obj.list[name]; ok {
		return true
	}

	return false
}
