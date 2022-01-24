package owners

import (
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/rings"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewOwnerBuilder creates a new owner builder
func NewOwnerBuilder() OwnerBuilder {
	hashAdapter := hash.NewAdapter()
	return createOwnerBuilder(hashAdapter)
}

// Builder represents the owners builder
type Builder interface {
	Create() Builder
	WithList(list []Owner) Builder
	Now() (Owners, error)
}

// Owners represents owners
type Owners interface {
	Hash() hash.Hash
	List() []Owner
}

// OwnerBuilder represents the owner builder
type OwnerBuilder interface {
	Create() OwnerBuilder
	WithRings(rings []rings.Ring) OwnerBuilder
	Now() (Owner, error)
}

// Owner represents an owner
type Owner interface {
	Hash() hash.Hash
	Rings() []rings.Ring
}
