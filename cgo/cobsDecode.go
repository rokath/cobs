// Package cobs implenments COBS (https://en.wikipedia.org/wiki/Consistent_Overhead_Byte_Stuffing).
package cobs

import (
	"errors"
)

// Decode a COBS frame to a slice of bytes.
//
// decoded := dec[:n]
func Decode(dec, cobs []byte) (n int, e error) {

	for len(cobs) > 0 {
		cnt := int(cobs[0] - 1)
		cobs = cobs[1:]
		if cnt > len(cobs) {
			e = errors.New("inconsistent COBS packet")
			return
		}
		n += copy(dec[n:], cobs[:cnt])
		cobs = cobs[cnt:]
		if cnt < 0xfe && len(cobs) > 0 {
			if n >= len(dec) {
				e = errors.New("provided decode buffer too small")
				return
			}
			dec[n] = 0
			n++
		}
	}
	return
}

//  func Encode(o, i []byte) (n int) {
//  todo
//  }
