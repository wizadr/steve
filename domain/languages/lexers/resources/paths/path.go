package paths

import "github.com/steve-care-software/steve/domain/languages/lexers/resources/containers"

type path struct {
	rootToken string
	container containers.Container
}

func createPath(
	rootToken string,
	container containers.Container,
) Path {
	out := path{
		rootToken: rootToken,
		container: container,
	}

	return &out
}

// RootToken returns the root token
func (obj *path) RootToken() string {
	return obj.rootToken
}

// Container returns the container
func (obj *path) Container() containers.Container {
	return obj.container
}
