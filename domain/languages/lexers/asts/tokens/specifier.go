package tokens

import "github.com/steve-care-software/steve/domain/languages/lexers/cardinality"

type specifier struct {
	containerName string
	cardinality   cardinality.Specific
	result        TokenResult
}

func createSpecifier(
	containerName string,
	cardinality cardinality.Specific,
	result TokenResult,
) Specifier {
	out := specifier{
		containerName: containerName,
		cardinality:   cardinality,
		result:        result,
	}

	return &out
}

// ContainerName returns the container name
func (obj *specifier) ContainerName() string {
	return obj.containerName
}

// Cardinality returns the cardinality
func (obj *specifier) Cardinality() cardinality.Specific {
	return obj.cardinality
}

// Result returns the result
func (obj *specifier) Result() TokenResult {
	return obj.result
}

// IsValid returns true if the specifier is valid, false otherwise
func (obj *specifier) IsValid() bool {
	amount := obj.Amount()
	return obj.cardinality.IsValid(amount)
}

// Amount returns the amount of results based on the containerName
func (obj *specifier) Amount() uint {
	return obj.fetchAmountResultsFromContainerName(obj.containerName, obj.result)
}

func (obj *specifier) fetchAmountResultsFromContainerName(containerName string, result TokenResult) uint {
	if !result.HasMatches() {
		return 0
	}

	total := uint(0)
	matches := result.Matches().All()
	for _, oneMatch := range matches {
		if !oneMatch.IsValid() {
			continue
		}

		if !oneMatch.HasContent() {
			continue
		}

		content := oneMatch.Content()
		blockMatches := content.Must().Matches().BlockMatches()
		for _, oneBlockMatch := range blockMatches {
			if !oneBlockMatch.IsValid() {
				continue
			}

			total += obj.fetchAmountResultsFromContainerNameAndBlockMatch(containerName, oneBlockMatch)
		}
	}

	return total
}

func (obj *specifier) fetchAmountResultsFromContainerNameAndBlockMatch(containerName string, blockMatch BlockMatch) uint {
	blockMatchContent := blockMatch.Content()
	if blockMatchContent.IsContainer() {
		container := blockMatchContent.Container()
		return obj.fetchAmountResultsFromContainerNameAndMatch(containerName, container)
	}

	if blockMatchContent.IsBlock() {
		blockMatch := blockMatchContent.Block()
		return obj.fetchAmountResultsFromContainerNameAndBlockMatch(containerName, blockMatch)
	}

	return 0
}

func (obj *specifier) fetchAmountResultsFromContainerNameAndMatch(containerName string, match Match) uint {
	if match.Content().IsToken() {
		token := match.Content().Token()
		if token.Path().Element().Name() == containerName {
			return token.Result().Amount()
		}

		return 0
	}

	if match.Content().IsRule() {
		if match.Content().Rule().Rule().Base().Name() == containerName {
			return match.Content().Rule().Result().Result().Amount()
		}

		return 0
	}

	return 0
}
