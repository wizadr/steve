package genesis

import (
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/owners"
)

type unit struct {
	hash    hash.Hash
	content UnitContent
	owner   owners.Owner
}

func createUnit(
	hash hash.Hash,
	content UnitContent,
	owner owners.Owner,
) Unit {
	out := unit{
		hash:    hash,
		content: content,
		owner:   owner,
	}

	return &out
}

// Hash returns the hash
func (obj *unit) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *unit) Content() UnitContent {
	return obj.content
}

// Owner returns the owner
func (obj *unit) Owner() owners.Owner {
	return obj.owner
}
