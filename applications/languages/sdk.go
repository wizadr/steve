package languages

import (
	"github.com/steve-care-software/steve/applications/languages/lexers"
	"github.com/steve-care-software/steve/applications/languages/parsers"
)

// Application represents the language application
type Application interface {
	Lexer() lexers.Application
	Parser() parsers.Application
}
