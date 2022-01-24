package parsers

type sliceDelimiter struct {
	index  IndexValue
	amount IndexValue
}

func createSliceDelimiter(
	index IndexValue,
) SliceDelimiter {
	return createSliceDelimiterInternally(index, nil)
}

func createSliceDelimiterWithAmount(
	index IndexValue,
	amount IndexValue,
) SliceDelimiter {
	return createSliceDelimiterInternally(index, amount)
}

func createSliceDelimiterInternally(
	index IndexValue,
	amount IndexValue,
) SliceDelimiter {
	out := sliceDelimiter{
		index:  index,
		amount: amount,
	}

	return &out
}

// Index returns the index, if any
func (obj *sliceDelimiter) Index() IndexValue {
	return obj.index
}

// HasAmount returns true if there is an amount, false otherwise
func (obj *sliceDelimiter) HasAmount() bool {
	return obj.amount != nil
}

// Amount returns the amount, if any
func (obj *sliceDelimiter) Amount() IndexValue {
	return obj.amount
}
