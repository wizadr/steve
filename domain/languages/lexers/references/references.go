package references

import "github.com/steve-care-software/steve/domain/languages/lexers/references/links"

type references struct {
	references   links.Links
	replacements links.Links
}

func createReferencesWithReferencesAndReplacements(
	references links.Links,
	replacements links.Links,
) References {
	return createReferencesInternally(references, replacements)
}

func createReferencesWithReferences(
	references links.Links,
) References {
	return createReferencesInternally(references, nil)
}

func createReferencesWithReplacements(
	replacements links.Links,
) References {
	return createReferencesInternally(nil, replacements)
}

func createReferences() References {
	return createReferencesInternally(nil, nil)
}

func createReferencesInternally(
	repl links.Links,
	replacements links.Links,
) References {
	out := references{
		references:   repl,
		replacements: replacements,
	}

	return &out
}

// HasReferences returns true if there is references, false otherwise
func (obj *references) HasReferences() bool {
	return obj.references != nil
}

// References returns the references, if any
func (obj *references) References() links.Links {
	return obj.references
}

// HasReplacements returns true if there is replacements, false otherwise
func (obj *references) HasReplacements() bool {
	return obj.replacements != nil
}

// Replacements returns the replacements, if any
func (obj *references) Replacements() links.Links {
	return obj.replacements
}
