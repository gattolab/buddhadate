package buddha

import "errors"

// ErrInvalidFormat is returned when a value cannot be parsed using the
// supplied layout.
var ErrInvalidFormat = errors.New("buddha: invalid date format")
