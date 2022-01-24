package receipts

import (
	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/transactions/secrets"
)

type receipt struct {
	hash   hash.Hash
	secret secrets.Secret
	sig    signature.Signature
}

func createReceipt(
	hash hash.Hash,
	secret secrets.Secret,
	sig signature.Signature,
) Receipt {
	out := receipt{
		hash:   hash,
		secret: secret,
		sig:    sig,
	}

	return &out
}

// Hash returns the hash
func (obj *receipt) Hash() hash.Hash {
	return obj.hash
}

// Secret returns the secret
func (obj *receipt) Secret() secrets.Secret {
	return obj.secret
}

// Signature returns the signature
func (obj *receipt) Signature() signature.Signature {
	return obj.sig
}
