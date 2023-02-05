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

// CDecode decodes `i` into `o` and returns number of bytes in `o`.
// For details see https://en.wikipedia.org/wiki/Consistent_Overhead_Byte_Stuffing.
// The CDencode implementation is done in C because the aimed use case is an embedded device running C.
// This function is mainly for testing.
func CDecode(o, i []byte) (n int, e error) {
	if len(i) == 0 {
		return
	}
	in := unsafe.Pointer(&i[0])
	out := unsafe.Pointer(&o[0])
	n = int(C.COBSDecode(out, in, C.size_t(len(i))))
	return
}
