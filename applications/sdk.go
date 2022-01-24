package application

import (
	"github.com/steve-care-software/digital-diamonds/applications/blockchains"
	"github.com/steve-care-software/digital-diamonds/applications/identities"
)

// Application represents a blockchain application
type Application interface {
	Blockchain() blockchains.Application
	Identity() identities.Application
}
