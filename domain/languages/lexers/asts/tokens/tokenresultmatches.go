package tokens

import (
	"strings"
)

type tokenResultMatches struct {
	list []Token
}

func createTokenResultMatches(
	list []Token,
) TokenResultMatches {
	out := tokenResultMatches{
		list: list,
	}

	return &out
}

// All returns the tokens
func (obj *tokenResultMatches) All() []Token {
	return obj.list
}

// Amount returns the tokens amount
func (obj *tokenResultMatches) Amount() uint {
	return uint(len(obj.list))
}

// IsValid returns true if the matches is valid, false otherwise
func (obj *tokenResultMatches) IsValid() bool {
	for _, oneToken := range obj.list {
		if !oneToken.IsValid() {
			return false
		}
	}

	return true
}

// IsExact returns true if the result is exact, false otherwise
func (obj *tokenResultMatches) IsExact() bool {
	for _, oneToken := range obj.list {
		if !oneToken.IsExact() {
			return false
		}
	}

	return true
}

// Discoveries returns the discoveries
func (obj *tokenResultMatches) Discoveries() string {
	list := []string{}
	for _, oneToken := range obj.list {
		list = append(list, oneToken.Discoveries())
	}

	return strings.Join(list, "")
}

// IsAmount returns true if there is this exact amount of tokens
func (obj *tokenResultMatches) IsAmount(amount uint) bool {
	listAmount := len(obj.list)
	return listAmount == int(amount)
}

// IsInRange returns true if the amount of tokens is within the given range
func (obj *tokenResultMatches) IsInRange(min uint, max uint) bool {
	listAmount := len(obj.list)
	return obj.IsAtLeast(min) && int(max) > listAmount
}

// IsAtLeast returns true if there is at least this amount of tokens
func (obj *tokenResultMatches) IsAtLeast(min uint) bool {
	listAmount := len(obj.list)
	return int(min) <= listAmount
}
