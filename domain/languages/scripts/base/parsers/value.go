package parsers

type value struct {
	typ     Type
	element interface{}
}

func createValue(
	typ Type,
	element interface{},
) Value {
	out := value{
		typ:     typ,
		element: element,
	}

	return &out
}

// Type returns the type
func (obj *value) Type() Type {
	return obj.typ
}

// Element returns the element
func (obj *value) Element() interface{} {
	return obj.element
}
