package cardinality

import (
	"strings"
)

type adapter struct {
	specificAdapter          SpecificAdapter
	builder                  Builder
	nonZeroMultipleCharacter string
	multipleCharacter        string
	optionalCharacter        string
}

func createAdapter(
	specificAdapter SpecificAdapter,
	builder Builder,
	nonZeroMultipleCharacter string,
	multipleCharacter string,
	optionalCharacter string,
) Adapter {
	out := adapter{
		specificAdapter:          specificAdapter,
		builder:                  builder,
		nonZeroMultipleCharacter: nonZeroMultipleCharacter,
		multipleCharacter:        multipleCharacter,
		optionalCharacter:        optionalCharacter,
	}

	return &out
}

// ToCardinality converts content to a cardinality instance
func (app *adapter) ToCardinality(content string) (Cardinality, error) {
	builder := app.builder.Create()
	trimmedContent := strings.TrimSpace(content)
	if trimmedContent == app.nonZeroMultipleCharacter {
		return builder.IsNonZeroMultiple().Now()
	}

	if trimmedContent == app.multipleCharacter {
		return builder.IsZeroMultiple().Now()
	}

	if trimmedContent == app.optionalCharacter {
		return builder.IsOptional().Now()
	}

	specific, err := app.specificAdapter.ToSpecific(trimmedContent)
	if err != nil {
		return nil, err
	}

	return builder.WithSpecific(specific).Now()
}

// ToPatternString converts a cardinality to pattern string
func (app *adapter) ToPatternString(cardinality Cardinality) (string, error) {
	if cardinality.IsOptional() {
		return app.optionalCharacter, nil
	}

	if cardinality.IsNonZeroMultiple() {
		return app.nonZeroMultipleCharacter, nil
	}

	if cardinality.IsZeroMultiple() {
		return app.multipleCharacter, nil
	}

	specific := cardinality.Specific()
	return app.specificAdapter.ToPatternString(specific)
}
