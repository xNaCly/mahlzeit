// assert provides common assertions to crash the api if an unknown case was hit
package assert

import "log/slog"

// True prints an error via slog if predicate is false, accepts args in the
// format slog accepts args
func True(predicate bool, args ...any) {
	if !predicate {
		if len(args) > 0 {
			slog.Error("assertion failed", args...)
		} else {
			slog.Error("assertion failed")
		}
		panic("assertion failed")
	}
}

// NoError prints an error with the error content via slog if error is not
// null, accepts args in the format slog accepts args
func NoError(error error, args ...any) {
	if error != nil {
		a := []any{"err", error}
		a = append(a, args...)
		slog.Error("assertion failed", a...)
		panic("assertion failed")
	}
}
