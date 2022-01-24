package transactions

import (
	"github.com/steve-care-software/steve/domain/cryptography/keys/signature"
	"github.com/steve-care-software/steve/domain/hash"
)

type transaction struct {
	hash    hash.Hash
	content Content
	auth    signature.RingSignature
}

func createTransaction(
	hash hash.Hash,
	content Content,
	auth signature.RingSignature,
) Transaction {
	out := transaction{
		hash:    hash,
		content: content,
		auth:    auth,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *transaction) Content() Content {
	return obj.content
}

// Authorization returns the authorization ring signature
func (obj *transaction) Authorization() signature.RingSignature {
	return obj.auth
}
