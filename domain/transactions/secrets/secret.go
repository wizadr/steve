package secrets

import (
	"github.com/steve-care-software/digital-diamonds/domain/hash"
	"github.com/steve-care-software/digital-diamonds/domain/transactions"
)

type secret struct {
	hash   hash.Hash
	amount uint
	nonce  string
	public transactions.Transaction
	origin Secret
	sides  Secrets
}

func createSecret(
	hash hash.Hash,
	amount uint,
	nonce string,
	public transactions.Transaction,
) Secret {
	return createSecretInternally(hash, amount, nonce, public, nil, nil)
}

func createSecretWithOrigin(
	hash hash.Hash,
	amount uint,
	nonce string,
	public transactions.Transaction,
	origin Secret,
) Secret {
	return createSecretInternally(hash, amount, nonce, public, origin, nil)
}

func createSecretWithSides(
	hash hash.Hash,
	amount uint,
	nonce string,
	public transactions.Transaction,
	sides Secrets,
) Secret {
	return createSecretInternally(hash, amount, nonce, public, nil, sides)
}

func createSecretWithOriginAndSides(
	hash hash.Hash,
	amount uint,
	nonce string,
	public transactions.Transaction,
	origin Secret,
	sides Secrets,
) Secret {
	return createSecretInternally(hash, amount, nonce, public, origin, sides)
}

func createSecretInternally(
	hash hash.Hash,
	amount uint,
	nonce string,
	public transactions.Transaction,
	origin Secret,
	sides Secrets,
) Secret {
	out := secret{
		hash:   hash,
		amount: amount,
		nonce:  nonce,
		public: public,
		origin: origin,
		sides:  sides,
	}

	return &out
}

// Hash returns the hash
func (obj *secret) Hash() hash.Hash {
	return obj.hash
}

// Amount returns the amount
func (obj *secret) Amount() uint {
	return obj.amount
}

// Nonce returns the nonce
func (obj *secret) Nonce() string {
	return obj.nonce
}

// Public returns the public transaction
func (obj *secret) Public() transactions.Transaction {
	return obj.public
}

// HasOrigin returns true if there is an origin, false otherwise
func (obj *secret) HasOrigin() bool {
	return obj.origin != nil
}

// Origin returns the origin, if any
func (obj *secret) Origin() Secret {
	return obj.origin
}

// HasSides returns true if there is sides, false otherwise
func (obj *secret) HasSides() bool {
	return obj.sides != nil
}

// Sides returns the sides, if any
func (obj *secret) Sides() Secrets {
	return obj.sides
}
