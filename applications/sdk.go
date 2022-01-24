package application

import (
	"github.com/steve-care-software/steve/applications/blockchains"
	"github.com/steve-care-software/steve/applications/identities"
)

// Application represents a blockchain application
type Application interface {
	Blockchain() blockchains.Application
	Identity() identities.Application
}
