// Code generated by "stringer -type CredentialsLocation"; DO NOT EDIT.

package cliconfig

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CredentialsNotAvailable-0]
	_ = x[CredentialsInPrimaryFile-80]
	_ = x[CredentialsInOtherFile-79]
	_ = x[CredentialsViaHelper-72]
	_ = x[CredentialsFromEnvironment-69]
}

const (
	_CredentialsLocation_name_0 = "CredentialsNotAvailable"
	_CredentialsLocation_name_1 = "CredentialsFromEnvironment"
	_CredentialsLocation_name_2 = "CredentialsViaHelper"
	_CredentialsLocation_name_3 = "CredentialsInOtherFileCredentialsInPrimaryFile"
)

var (
	_CredentialsLocation_index_3 = [...]uint8{0, 22, 46}
)

func (i CredentialsLocation) String() string {
	switch {
	case i == 0:
		return _CredentialsLocation_name_0
	case i == 69:
		return _CredentialsLocation_name_1
	case i == 72:
		return _CredentialsLocation_name_2
	case 79 <= i && i <= 80:
		i -= 79
		return _CredentialsLocation_name_3[_CredentialsLocation_index_3[i]:_CredentialsLocation_index_3[i+1]]
	default:
		return "CredentialsLocation(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
