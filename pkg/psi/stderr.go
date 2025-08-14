package psi

import "fmt"

/* ------------------------------------------------------------------------ */

// ErrMethodOnNil is an assert()-level error (that does not panic) where a
// method was called on a nil instance of a thing.
var ErrMethodOnNil = fmt.Errorf("method called on a nil instance")
