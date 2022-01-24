package cardinality

import (
	"fmt"
	"strconv"
	"strings"
)

type specificAdapter struct {
	builder         Builder
	specificBuilder SpecificBuilder
	rangeBuilder    RangeBuilder
	prefix          string
	suffix          string
	rangeSeparator  string
}

func createSpecificAdapter(
	builder Builder,
	specificBuilder SpecificBuilder,
	rangeBuilder RangeBuilder,
	prefix string,
	suffix string,
	rangeSeparator string,
) SpecificAdapter {
	out := specificAdapter{
		builder:         builder,
		specificBuilder: specificBuilder,
		rangeBuilder:    rangeBuilder,
		prefix:          prefix,
		suffix:          suffix,
		rangeSeparator:  rangeSeparator,
	}

	return &out
}

// ToSpecific converts content to a specific cardinality
func (app *specificAdapter) ToSpecific(content string) (Specific, error) {
	trimmedContent := strings.TrimSpace(content)
	if strings.HasPrefix(content, app.prefix) {
		trimmedContent = trimmedContent[1:]
	}

	if strings.HasSuffix(content, app.suffix) {
		trimmedContent = trimmedContent[:len(trimmedContent)-1]
	}

	builder := app.specificBuilder.Create()
	sections := strings.Split(trimmedContent, app.rangeSeparator)
	if len(sections) == 1 {
		amount, err := strconv.Atoi(sections[0])
		if err != nil {
			return nil, err
		}

		builder.WithAmount(uint(amount))
	}

	if len(sections) == 2 {
		min, err := strconv.Atoi(sections[0])
		if err != nil {
			return nil, err
		}

		rangeBuilder := app.rangeBuilder.Create().WithMinimum(uint(min))
		trimmedSection := strings.TrimSpace(sections[1])
		if trimmedSection != "" {
			max, err := strconv.Atoi(trimmedSection)
			if err != nil {
				return nil, err
			}

			rangeBuilder.WithMaximum(uint(max))
		}

		rnge, err := rangeBuilder.Now()
		if err != nil {
			return nil, err
		}

		builder.WithRange(rnge)
	}

	return builder.Now()
}

// ToPatternString converts a specific instance to a pattern string
func (app *specificAdapter) ToPatternString(specific Specific) (string, error) {
	if specific.IsAmount() {
		amount := specific.Amount()
		return fmt.Sprintf("%s%d%s", app.prefix, *amount, app.suffix), nil
	}

	rnge := specific.Range()
	min := rnge.Min()
	if rnge.HasMax() {
		max := rnge.Max()
		return fmt.Sprintf("%s%d%s%d%s", app.prefix, min, app.rangeSeparator, *max, app.suffix), nil
	}

	return fmt.Sprintf("%s%d%s%s", app.prefix, min, app.rangeSeparator, app.suffix), nil
}
