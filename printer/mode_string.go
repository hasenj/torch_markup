// Code generated by "stringer -type=Mode"; DO NOT EDIT.

package printer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[KeepNewlines-1]
}

const _Mode_name = "KeepNewlines"

var _Mode_index = [...]uint8{0, 12}

func (i Mode) String() string {
	i -= 1
	if i < 0 || i >= Mode(len(_Mode_index)-1) {
		return "Mode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Mode_name[_Mode_index[i]:_Mode_index[i+1]]
}
