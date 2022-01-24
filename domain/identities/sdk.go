package identities

import (
	"github.com/steve-care-software/digital-diamonds/domain/transactions/privates"
	"github.com/steve-care-software/digital-diamonds/domain/transactions/secrets"
)

// Identity represents the identity
type Identity interface {
	Name() string
	Outgoing() privates.Privates
	Incoming() secrets.Secrets
}
