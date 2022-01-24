package asts

import "github.com/steve-care-software/steve/domain/languages/lexers/containers/patterns"

type result struct {
	input       string
	discoveries patterns.Discoveries
	prefix      string
	suffix      string
}

func createResult(
	input string,
	discoveries patterns.Discoveries,
) Result {
	return createResultInternally(input, discoveries, "", "")
}

func createResultWithPrefix(
	input string,
	discoveries patterns.Discoveries,
	prefix string,
) Result {
	return createResultInternally(input, discoveries, prefix, "")
}

func createResultWithSuffix(
	input string,
	discoveries patterns.Discoveries,
	suffix string,
) Result {
	return createResultInternally(input, discoveries, "", suffix)
}

func createResultWithPrefixAndSuffix(
	input string,
	discoveries patterns.Discoveries,
	prefix string,
	suffix string,
) Result {
	return createResultInternally(input, discoveries, prefix, suffix)
}

func createResultInternally(
	input string,
	discoveries patterns.Discoveries,
	prefix string,
	suffix string,
) Result {
	out := result{
		input:       input,
		discoveries: discoveries,
		prefix:      prefix,
		suffix:      suffix,
	}

	return &out
}

// Input returns the input
func (obj *result) Input() string {
	return obj.input
}

// Discoveries returns the discoveries
func (obj *result) Discoveries() patterns.Discoveries {
	return obj.discoveries
}

// HasPrefix returns true if there is a prefix, false otherwise
func (obj *result) HasPrefix() bool {
	return obj.prefix != ""
}

// Prefix returns the prefix, if any
func (obj *result) Prefix() string {
	return obj.prefix
}

// HasSuffix returns true if there is a suffix, false otherwise
func (obj *result) HasSuffix() bool {
	return obj.suffix != ""
}

// Suffix returns the suffix, if any
func (obj *result) Suffix() string {
	return obj.suffix
}
