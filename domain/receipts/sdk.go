package receipts

import (
	"github.com/steve-care-software/digital-diamonds/domain/transactions/secrets"
	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

// Receipt represents a transaction receipt
type Receipt interface {
	Hash() hash.Hash
	Secret() secrets.Secret
	Signature() signature.Signature
}
