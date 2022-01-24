package identities

import (
	"github.com/steve-care-software/digital-diamonds/domain/receipts"
	"github.com/steve-care-software/digital-diamonds/domain/transactions/secrets"
	"github.com/steve-care-software/digital-diamonds/domain/cryptography/keys/signature"
)

// Application represents an identity application
type Application interface {
	Request() (signature.PublicKey, error)
	Transact(trx secrets.Secret) (receipts.Receipt, error)
}
