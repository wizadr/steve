package rings

import "github.com/steve-care-software/digital-diamonds/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder(min uint, max uint) Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter, min, max)
}

// Builder represents a ring builder
type Builder interface {
	Create() Builder
	WithList(list []hash.Hash) Builder
	Now() (Ring, error)
}

// Ring represents an hash ring
type Ring interface {
	Hash() hash.Hash
	List() []hash.Hash
}
