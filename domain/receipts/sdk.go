package receipts

import (
	"github.com/steve-care-software/steve/domain/cryptography/keys/signature"
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/transactions/secrets"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a receipt builder
type Builder interface {
	Create() Builder
	WithSecret(secret secrets.Secret) Builder
	WithSignature(signature signature.Signature) Builder
	Now() (Receipt, error)
}

// Receipt represents a transaction receipt
type Receipt interface {
	Hash() hash.Hash
	Secret() secrets.Secret
	Signature() signature.Signature
}
