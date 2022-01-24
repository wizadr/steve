package genesis

import (
	"time"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/rings"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewUnitBuilder creates a new unit builder
func NewUnitBuilder() UnitBuilder {
	hashAdapter := hash.NewAdapter()
	return createUnitBuilder(hashAdapter)
}

// NewUnitContentBuilder creates a new unit content builder
func NewUnitContentBuilder() UnitContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createUnitContentBuilder(hashAdapter)
}

// Builder represents a genesis builder
type Builder interface {
	Create() Builder
	WithUnits(units Units) Builder
	WithFees(fees uint) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Genesis, error)
}

// Genesis represents a genesis
type Genesis interface {
	Hash() hash.Hash
	Units() Units
	Fees() uint
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
	WithOwner(owner rings.Ring) UnitBuilder
	Now() (Unit, error)
}

// Unit represents a genesis unit
type Unit interface {
	Hash() hash.Hash
	Content() UnitContent
	Owner() rings.Ring
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
