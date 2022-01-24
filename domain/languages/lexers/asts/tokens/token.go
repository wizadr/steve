package tokens

import (
	"github.com/steve-care-software/steve/domain/languages/lexers/asts/paths"
)

type token struct {
	name    string
	path    paths.Path
	content TokenContent
}

func createToken(
	name string,
	path paths.Path,
) Token {
	return createTokenInternally(name, path, nil)
}

func createTokenWithContent(
	name string,
	path paths.Path,
	content TokenContent,
) Token {
	return createTokenInternally(name, path, content)
}

func createTokenInternally(
	name string,
	path paths.Path,
	content TokenContent,
) Token {
	out := token{
		name:    name,
		path:    path,
		content: content,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}

// IsValid returns true if the token is valid, false otherwise
func (obj *token) IsValid() bool {
	if obj.HasContent() {
		return obj.Content().IsValid()
	}

	return false
}

// IsExact returns true if the token is exact, false otherwise
func (obj *token) IsExact() bool {
	if obj.HasContent() {
		return obj.Content().IsExact()
	}

	return false
}

// Discoveries returns the discoveries
func (obj *token) Discoveries() string {
	if obj.HasContent() {
		return obj.Content().Discoveries()
	}

	return ""
}

// Path returns the path
func (obj *token) Path() paths.Path {
	return obj.path
}

// HasContent returns true if there is content, false otherwise
func (obj *token) HasContent() bool {
	return obj.content != nil
}

// Content returns the content, if any
func (obj *token) Content() TokenContent {
	return obj.content
}
