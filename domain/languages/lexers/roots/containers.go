package roots

import (
	"errors"
	"fmt"
	"strings"
)

type containers struct {
	list []Container
	mp   map[string]Container
}

func createContainers(
	list []Container,
	mp map[string]Container,
) Containers {
	out := containers{
		list: list,
		mp:   mp,
	}

	return &out
}

// Content returns the content
func (obj *containers) Content() string {
	content := []string{}
	for _, oneContainer := range obj.list {
		content = append(content, oneContainer.Content())
	}

	return strings.Join(content, "")
}

// List returns the containers list
func (obj *containers) List() []Container {
	return obj.list
}

// Fetch fetches a container by name
func (obj *containers) Fetch(name string) (Container, error) {
	if container, ok := obj.mp[name]; ok {
		return container, nil
	}

	str := fmt.Sprintf("the container (name: %s) does not exists", name)
	return nil, errors.New(str)
}

// Scan scans the containers
func (obj *containers) Scan(name string) (Root, Node, Container, string, error) {
	for _, container := range obj.list {
		root, node, subContainer, content, err := container.Scan(name)
		if err != nil {
			continue
		}

		if root != nil {
			return root, nil, nil, "", nil
		}

		if node != nil {
			return nil, node, nil, "", nil
		}

		if subContainer != nil {
			return nil, nil, subContainer, "", nil
		}

		if content != "" {
			return nil, nil, nil, content, nil
		}
	}

	str := fmt.Sprintf("the element (name: %s) could not be found in the current Containers instance", name)
	return nil, nil, nil, "", errors.New(str)
}
