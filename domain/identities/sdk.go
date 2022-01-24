package identities

import (
	"github.com/steve-care-software/digital-diamonds/domain/transactions/privates"
	"github.com/steve-care-software/digital-diamonds/domain/transactions/secrets"
)

// Builder represents the identity builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithOutgoing(outgoing privates.Privates) Builder
	WithIncoming(incoming secrets.Secrets) Builder
	Now() (Identity, error)
}

// Identity represents the identity
type Identity interface {
	Name() string
	HasOutgoing() bool
	Outgoing() privates.Privates
	HasIncoming() bool
	Incoming() secrets.Secrets
}
