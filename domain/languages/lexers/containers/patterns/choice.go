package patterns

type choice struct {
	serie     Serie
	isReverse bool
}

func createChoice(
	serie Serie,
	isReverse bool,
) Choice {
	out := choice{
		serie:     serie,
		isReverse: isReverse,
	}

	return &out
}

// Serie returns the serie
func (obj *choice) Serie() Serie {
	return obj.serie
}

// IsReverse returns true if reverse, false otherwise
func (obj *choice) IsReverse() bool {
	return obj.isReverse
}
