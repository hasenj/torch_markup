// Code generated by "stringer -type=Token"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EOF-0]
	_ = x[LINEFEED-1]
	_ = x[INDENT-2]
	_ = x[DEDENT-3]
	_ = x[COMMENT-4]
	_ = x[VLINE-5]
	_ = x[GT-6]
	_ = x[TEXT-7]
}

const _Token_name = "EOFLINEFEEDINDENTDEDENTCOMMENTVLINEGTTEXT"

var _Token_index = [...]uint8{0, 3, 11, 17, 23, 30, 35, 37, 41}

func (i Token) String() string {
	if i >= Token(len(_Token_index)-1) {
		return "Token(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Token_name[_Token_index[i]:_Token_index[i+1]]
}