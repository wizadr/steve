package genesis

import (
	"time"

	"github.com/steve-care-software/digital-diamonds/domain/hash"
)

type genesis struct {
	hash      hash.Hash
	units     Units
	createdOn time.Time
}

func createGenesis(
	hash hash.Hash,
	units Units,
	createdOn time.Time,
) Genesis {
	out := genesis{
		hash:      hash,
		units:     units,
		createdOn: createdOn,
	}

	return &out
}

// Hash returns the hash
func (obj *genesis) Hash() hash.Hash {
	return obj.hash
}

// Units returns the units
func (obj *genesis) Units() Units {
	return obj.units
}

// CreatedOn returns the creation time
func (obj *genesis) CreatedOn() time.Time {
	return obj.createdOn
}
