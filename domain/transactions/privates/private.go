package privates

import (
	"github.com/steve-care-software/steve/domain/cryptography/keys/signature"
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/receipts"
)

type private struct {
	hash    hash.Hash
	pk      signature.PrivateKey
	receipt receipts.Receipt
}

func createPrivateKey(
	hash hash.Hash,
	pk signature.PrivateKey,
	receipt receipts.Receipt,
) Private {
	out := private{
		hash:    hash,
		pk:      pk,
		receipt: receipt,
	}

	return &out
}

// Hash returns the hash
func (obj *private) Hash() hash.Hash {
	return obj.hash
}

// PrivateKey returns the private key
func (obj *private) PrivateKey() signature.PrivateKey {
	return obj.pk
}

// Receipt returns the receipt
func (obj *private) Receipt() receipts.Receipt {
	return obj.receipt
}
