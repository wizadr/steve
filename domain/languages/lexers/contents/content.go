package contents

type content struct {
	content       string
	isPrefixLegal bool
	isSuffixLegal bool
}

func createContent(
	str string,
) Content {
	return createContentInternally(str, false, false)
}

func createContentWithPrefix(
	str string,
) Content {
	return createContentInternally(str, true, false)
}

func createContentWithSuffix(
	str string,
) Content {
	return createContentInternally(str, false, true)
}

func createContentWithPrefixAndSuffix(
	str string,
) Content {
	return createContentInternally(str, true, true)
}

func createContentInternally(
	str string,
	isPrefixLegal bool,
	isSuffixLegal bool,
) Content {
	out := content{
		content:       str,
		isPrefixLegal: isPrefixLegal,
		isSuffixLegal: isSuffixLegal,
	}

	return &out
}

// Content returns the content
func (obj *content) Content() string {
	return obj.content
}

// IsPrefixLegal returns true if the prefix is legal, false otherwise
func (obj *content) IsPrefixLegal() bool {
	return obj.isPrefixLegal
}

// IsSuffixLegal returns true if the suffix is legal, false otherwise
func (obj *content) IsSuffixLegal() bool {
	return obj.isSuffixLegal
}
