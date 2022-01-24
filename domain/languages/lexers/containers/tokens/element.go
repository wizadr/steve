package tokens

import "github.com/steve-care-software/steve/domain/languages/lexers/cardinality"

type element struct {
	content     Content
	code        string
	cardinality cardinality.Cardinality
}

func createElement(
	content Content,
	code string,
) Element {
	return createElementInternally(content, code, nil)
}

func createElementWithCardinality(
	content Content,
	code string,
	cardinality cardinality.Cardinality,
) Element {
	return createElementInternally(content, code, cardinality)
}

func createElementWithSubElementsAndCardinality(
	content Content,
	code string,
	cardinality cardinality.Cardinality,
) Element {
	return createElementInternally(content, code, cardinality)
}

func createElementInternally(
	content Content,
	code string,
	cardinality cardinality.Cardinality,
) Element {
	out := element{
		content:     content,
		code:        code,
		cardinality: cardinality,
	}

	return &out
}

// Content returns the content
func (obj *element) Content() Content {
	return obj.content
}

// Code returns the code
func (obj *element) Code() string {
	return obj.code
}

// HasCardinality returns true if there is cardinality, false otherwise
func (obj *element) HasCardinality() bool {
	return obj.cardinality != nil
}

// Cardinality returns the cardinality, if any
func (obj *element) Cardinality() cardinality.Cardinality {
	return obj.cardinality
}
