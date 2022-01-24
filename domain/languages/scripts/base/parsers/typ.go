package parsers

import (
	"errors"
	"fmt"
)

type typ struct {
	dept uint
	flag uint8
}

func createType(
	dept uint,
	flag uint8,
) Type {
	out := typ{
		dept: dept,
		flag: flag,
	}

	return &out
}

// Dept returns the dept
func (obj *typ) Dept() uint {
	return obj.dept
}

// Flag returns the flag
func (obj *typ) Flag() uint8 {
	return obj.flag
}

// Compare compares the current type with the input type, returns nil if the same, an error otherwise
func (obj *typ) Compare(input Type) error {
	if obj.dept != input.Dept() {
		str := fmt.Sprintf("the dept is incompatible, current: %d, input: %d", obj.dept, input.Dept())
		return errors.New(str)
	}

	if obj.flag != input.Flag() {
		str := fmt.Sprintf("the flag is incompatible, current: %d, input: %d", obj.flag, input.Flag())
		return errors.New(str)
	}

	return nil
}
