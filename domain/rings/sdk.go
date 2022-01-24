package rings

import "github.com/steve-care-software/steve/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewRingBuilder creates a new ring builder instance
func NewRingBuilder(min uint, max uint) RingBuilder {
	hashAdapter := hash.NewAdapter()
	return createRingBuilder(hashAdapter, min, max)
}

// Builder represents a ring builder
type Builder interface {
	Create() Builder
	WithList(list []Ring) Builder
	Now() (Rings, error)
}

// Rings represents rings
type Rings interface {
	Hash() hash.Hash
	List() []Ring
}

// RingBuilder represents a ring builder
type RingBuilder interface {
	Create() RingBuilder
	WithList(list []hash.Hash) RingBuilder
	Now() (Ring, error)
}

// Ring represents an hash ring
type Ring interface {
	Hash() hash.Hash
	List() []hash.Hash
}
