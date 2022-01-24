package secrets

import "github.com/steve-care-software/digital-diamonds/domain/hash"

type secrets struct {
	hash hash.Hash
	list []Secret
}

func createSecrets(
	hash hash.Hash,
	list []Secret,
) Secrets {
	out := secrets{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *secrets) Hash() hash.Hash {
	return obj.hash
}

// List returns the list of secrets
func (obj *secrets) List() []Secret {
	return obj.list
}
