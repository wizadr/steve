package genesis

import (
	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

type unit struct {
	hash    hash.Hash
	content UnitContent
	sig     signature.Signature
}

func createUnit(
	hash hash.Hash,
	content UnitContent,
	sig signature.Signature,
) Unit {
	out := unit{
		hash:    hash,
		content: content,
		sig:     sig,
	}

	return &out
}

// Hash returns the hash
func (obj *unit) Hash() hash.Hash {
	return obj.hash
}

// Content returns the content
func (obj *unit) Content() UnitContent {
	return obj.content
}

// Signature returns the signature
func (obj *unit) Signature() signature.Signature {
	return obj.sig
}
