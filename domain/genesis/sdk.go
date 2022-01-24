package genesis

import (
	"time"

	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

// Builder represents a genesis builder
type Builder interface {
	Create() Builder
	WithUnits(units Units) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Genesis, error)
}

// Genesis represents a genesis
type Genesis interface {
	Hash() hash.Hash
	Units() Units
	CreatedOn() time.Time
}

// UnitsBuilder represents a units builder
type UnitsBuilder interface {
	Create() UnitsBuilder
	WithList(list []Unit) UnitsBuilder
	Now() (Units, error)
}

// Units represents genesis units
type Units interface {
	Hash() hash.Hash
	All() []Unit
}

// UnitBuilder represents a unit builder
type UnitBuilder interface {
	Create() UnitBuilder
	WithContent(content UnitContent) UnitBuilder
	WithSignature(sig signature.Signature) UnitBuilder
	Now() (Unit, error)
}

// Unit represents a genesis unit
type Unit interface {
	Hash() hash.Hash
	Content() UnitContent
	Signature() signature.Signature
}

// UnitContentBuilder represents the unit content builder
type UnitContentBuilder interface {
	Create() UnitContentBuilder
	WithAmount(amount uint) UnitContentBuilder
	WithNonce(nonce string) UnitContentBuilder
	ActivatedOn(activatedOn time.Time) UnitContentBuilder
	CreatedOn(createdOn time.Time) UnitContentBuilder
	Now() (UnitContent, error)
}

// UnitContent represents a unit content
type UnitContent interface {
	Hash() hash.Hash
	Amount() uint
	Nonce() string
	ActivatedOn() time.Time
	CreatedOn() time.Time
}
