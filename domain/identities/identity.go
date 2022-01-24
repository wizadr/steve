package identities

import (
	"github.com/steve-care-software/digital-diamonds/domain/transactions/privates"
	"github.com/steve-care-software/digital-diamonds/domain/transactions/secrets"
)

type identity struct {
	name     string
	outgoing privates.Privates
	incoming secrets.Secrets
}

func createIdentity(
	name string,
) Identity {
	return createIdentityInternally(name, nil, nil)
}

func createIdentityWithOutgoing(
	name string,
	outgoing privates.Privates,
) Identity {
	return createIdentityInternally(name, outgoing, nil)
}

func createIdentityWithIncoming(
	name string,
	incoming secrets.Secrets,
) Identity {
	return createIdentityInternally(name, nil, incoming)
}

func createIdentityWithOutgoingAndIncoming(
	name string,
	outgoing privates.Privates,
	incoming secrets.Secrets,
) Identity {
	return createIdentityInternally(name, outgoing, incoming)
}

func createIdentityInternally(
	name string,
	outgoing privates.Privates,
	incoming secrets.Secrets,
) Identity {
	out := identity{
		name:     name,
		outgoing: outgoing,
		incoming: incoming,
	}

	return &out
}

// Name returns the name
func (obj *identity) Name() string {
	return obj.name
}

// HasOutgoing returns true if there is outgoing transactions, false otherwise
func (obj *identity) HasOutgoing() bool {
	return obj.outgoing != nil
}

// Outgoing returns the outgoing transactions
func (obj *identity) Outgoing() privates.Privates {
	return obj.outgoing
}

// HasIncoming returns true if there is incoming transactions, false otherwise
func (obj *identity) HasIncoming() bool {
	return obj.incoming != nil
}

// Incoming returns the incoming transactions
func (obj *identity) Incoming() secrets.Secrets {
	return obj.incoming
}
