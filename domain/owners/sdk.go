package owners

import "github.com/steve-care-software/digital-diamonds/domain/hash"

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
	WithHashes(hashes []hash.Hash) OwnerBuilder
	Now() (Owner, error)
}

// Owner represents an owner
type Owner interface {
	Hash() hash.Hash
	Hashes() []hash.Hash
}
