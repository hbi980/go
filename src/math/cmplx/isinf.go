// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmplx

import "math"

// IsInf returns true if either real(x) or imag(x) is an infinity.

// IsInf 在 real(x) 或 imag(x) 为无限大值时返回 true。
func IsInf(x complex128) bool {
	if math.IsInf(real(x), 0) || math.IsInf(imag(x), 0) {
		return true
	}
	return false
}

// Inf returns a complex infinity, complex(+Inf, +Inf).

// Inf 返回一个复数的无限大值，即 complex(+Inf, +Inf)。
func Inf() complex128 {
	inf := math.Inf(1)
	return complex(inf, inf)
}
