// Code generated by "stringer -type=ErrorCode -output errors_generated.go"; DO NOT EDIT.

package ga

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnknownError-0]
	_ = x[InvalidArgument-1]
	_ = x[CannotCrossover-2]
	_ = x[EvalError-3]
	_ = x[CannotMakeNextGeneration-4]
	_ = x[CannotBuildGeneration-5]
	_ = x[CannotSelectIndividual-6]
}

const _ErrorCode_name = "UnknownErrorInvalidArgumentCannotCrossoverEvalErrorCannotMakeNextGenerationCannotBuildGenerationCannotSelectIndividual"

var _ErrorCode_index = [...]uint8{0, 12, 27, 42, 51, 75, 96, 118}

func (i ErrorCode) String() string {
	if i < 0 || i >= ErrorCode(len(_ErrorCode_index)-1) {
		return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorCode_name[_ErrorCode_index[i]:_ErrorCode_index[i+1]]
}
