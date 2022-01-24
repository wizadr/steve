package identities

import "github.com/steve-care-software/digital-diamonds/domain/transactions/privates"

type identity struct {
	name    string
	trx     Transactions
	genesis privates.Privates
}

func createIdentity(
	name string,
	trx Transactions,
) Identity {
	return createIdentityInternally(name, trx, nil)
}

func createIdentityWithGenesis(
	name string,
	trx Transactions,
	genesis privates.Privates,
) Identity {
	return createIdentityInternally(name, trx, genesis)
}

func createIdentityInternally(
	name string,
	trx Transactions,
	genesis privates.Privates,
) Identity {
	out := identity{
		name:    name,
		trx:     trx,
		genesis: genesis,
	}

	return &out
}

// Name returns the name
func (obj *identity) Name() string {
	return obj.name
}

// Transactions returns the transactions
func (obj *identity) Transactions() Transactions {
	return obj.trx
}

// HasGenesis returns true if there is genesis transactions, false otherwise
func (obj *identity) HasGenesis() bool {
	return obj.genesis != nil
}

// Genesis returns the genesis transactions, if any
func (obj *identity) Genesis() privates.Privates {
	return obj.genesis
}
