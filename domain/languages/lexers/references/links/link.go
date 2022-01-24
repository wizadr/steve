package links

type link struct {
	localToken     string
	referenceToken string
	include        Include
}

func createReference(
	localToken string,
	referenceToken string,
	include Include,
) Link {
	out := link{
		localToken:     localToken,
		referenceToken: referenceToken,
		include:        include,
	}

	return &out
}

// LocalToken returns the local token
func (obj *link) LocalToken() string {
	return obj.localToken
}

// ReferenceToken returns the reference token
func (obj *link) ReferenceToken() string {
	return obj.referenceToken
}

// Include returns the include
func (obj *link) Include() Include {
	return obj.include
}
