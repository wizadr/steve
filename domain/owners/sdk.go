package owners

import "github.com/steve-care-software/digital-diamonds/domain/hash"

// Owners represents owners
type Owners interface {
	Hash() hash.Hash
	List() []Owner
}

// Owner represents an owner
type Owner interface {
	Hash() hash.Hash
	Hashes() []hash.Hash
}
