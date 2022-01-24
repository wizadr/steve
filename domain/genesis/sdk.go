package genesis

import (
	"time"

	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

// Genesis represents a genesis
type Genesis interface {
	Hash() hash.Hash
	Units() Units
	CreatedOn() time.Time
}

// Units represents genesis units
type Units interface {
	Hash() hash.Hash
	All() []Unit
}

// Unit represents a genesis unit
type Unit interface {
	Hash() hash.Hash
	Content() UnitContent
	Signature() signature.Signature
}

// UnitContent represents a unit content
type UnitContent interface {
	Hash() hash.Hash
	Amount() uint
	Nonce() string
	CreatedOn() time.Time
}
