package parsers

type headerVariable struct {
	isMandatory bool
	isInput     bool
	declaration VariableDeclaration
	keyname     string
}

func createHeaderVariable(
	isMandatory bool,
	isInput bool,
	declaration VariableDeclaration,
	keyname string,
) HeaderVariable {
	out := headerVariable{
		isMandatory: isMandatory,
		isInput:     isInput,
		declaration: declaration,
		keyname:     keyname,
	}

	return &out
}

// IsMandatory returns true if mandatory, false otherwise
func (obj *headerVariable) IsMandatory() bool {
	return obj.isMandatory
}

// IsInput returns true if input, false otherwise
func (obj *headerVariable) IsInput() bool {
	return obj.isInput
}

// Declaration returns the VariableDeclaration
func (obj *headerVariable) Declaration() VariableDeclaration {
	return obj.declaration
}

// Keyname returns the keyname
func (obj *headerVariable) Keyname() string {
	return obj.keyname
}
