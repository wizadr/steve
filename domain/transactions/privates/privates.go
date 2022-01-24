package privates

import "github.com/steve-care-software/steve/domain/hash"

type privates struct {
	hash hash.Hash
	list []Private
}

func createPrivates(
	hash hash.Hash,
	list []Private,
) Privates {
	out := privates{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *privates) Hash() hash.Hash {
	return obj.hash
}

// List returns the list of privates
func (obj *privates) List() []Private {
	return obj.list
}
