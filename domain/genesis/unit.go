package genesis

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/rings"
)

type unit struct {
	hash    hash.Hash
	content UnitContent
	owner   rings.Ring
}

func createUnit(
	hash hash.Hash,
	content UnitContent,
	owner rings.Ring,
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
func (obj *unit) Owner() rings.Ring {
	return obj.owner
}
