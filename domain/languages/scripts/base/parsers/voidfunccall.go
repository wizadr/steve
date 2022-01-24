package parsers

type voidFuncCall struct {
	log Assignable
}

func createVoidFuncCallWithLog(
	log Assignable,
) VoidFuncCall {
	return createVoidFuncCallInternally(log)
}

func createVoidFuncCallInternally(
	log Assignable,
) VoidFuncCall {
	out := voidFuncCall{
		log: log,
	}

	return &out
}

// IsLog returns true if there is a log, false otherwise
func (obj *voidFuncCall) IsLog() bool {
	return obj.log != nil
}

// Log returns the log, if any
func (obj *voidFuncCall) Log() Assignable {
	return obj.log
}
