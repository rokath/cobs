// Package cobs implenments COBS (https://en.wikipedia.org/wiki/Consistent_Overhead_Byte_Stuffing).
package cobs

// #include <stddef.h>
// #include <stdint.h>
// #include "cobs.h"
// #cgo CFLAGS: -g -Wall
import "C"
import (
	"unsafe"
)

// CEncode encodes `i` into `o` and returns number of bytes in `o`.
// For details see https://en.wikipedia.org/wiki/Consistent_Overhead_Byte_Stuffing.
// The Encode implementation is done in C because the aimed use case is an embedded device running C.
// This function is mainly for testing.
func CEncode(o, i []byte) (n int) {
	if len(i) == 0 {
		return
	}
	in := unsafe.Pointer(&i[0])
	out := unsafe.Pointer(&o[0])
	n = int(C.COBSEncode(out, in, C.size_t(len(i))))
	return
}
