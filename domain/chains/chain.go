package chains

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/steve/domain/hash"
)

type chain struct {
	id        uuid.UUID
	hash      hash.Hash
	root      hash.Hash
	createdOn time.Time
	head      Block
}

func createChain(
	id uuid.UUID,
	hash hash.Hash,
	root hash.Hash,
	createdOn time.Time,
) Chain {
	return createChainInternally(id, hash, root, createdOn, nil)
}

func createChainWithHead(
	id uuid.UUID,
	hash hash.Hash,
	root hash.Hash,
	createdOn time.Time,
	head Block,
) Chain {
	return createChainInternally(id, hash, root, createdOn, head)
}

func createChainInternally(
	id uuid.UUID,
	hash hash.Hash,
	root hash.Hash,
	createdOn time.Time,
	head Block,
) Chain {
	out := chain{
		id:        id,
		hash:      hash,
		root:      root,
		createdOn: createdOn,
		head:      head,
	}

	return &out
}

// ID returns the id
func (obj *chain) ID() uuid.UUID {
	return obj.id
}

// Hash returns the hash
func (obj *chain) Hash() hash.Hash {
	return obj.hash
}

// Root returns the root
func (obj *chain) Root() hash.Hash {
	return obj.root
}

// CreatedOn returns the createdOn
func (obj *chain) CreatedOn() time.Time {
	return obj.createdOn
}

// HasHead returns true if there is a link, false otherwise
func (obj *chain) HasHead() bool {
	return obj.head != nil
}

// Head returns the head, if any
func (obj *chain) Head() Block {
	return obj.head
}
