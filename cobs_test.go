package cobs_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"

	"github.com/rokath/cobs"

	"github.com/tj/assert"
)

func TestDecodeTestData(t *testing.T) {
	dec := make([]byte, 1024)
	for _, x := range testData {
		n, e := cobs.Decode(dec, x.enc)
		assert.Nil(t, e)
		act := dec[:n]
		assert.Equal(t, x.dec, act)
	}
}

func TestCDecodeTestData(t *testing.T) {
	dec := make([]byte, 1024)
	for _, x := range testData {
		n, e := cobs.CDecode(dec, x.enc)
		assert.Nil(t, e)
		act := dec[:n]
		assert.Equal(t, x.dec, act)
	}
}

func TestCEncodeTestData(t *testing.T) {
	enc := make([]byte, 1024)
	for _, x := range testData {
		n := cobs.CEncode(enc, x.dec)
		act := enc[:n]
		assert.Equal(t, x.enc, act)
	}
}

// TestCEncodeCDecode256 tests on generated random byte numbers for random length 0-32767.
func TestCEncodeCDecode256(t *testing.T) {
	max := 32768
	for i := 0; i < 10000; i++ {
		length := rand.Intn(max)
		datBuf := make([]byte, max)
		encBuf := make([]byte, 2*max)
		decBuf := make([]byte, 2*max)
		for i := 0; i < length; i++ {
			b := uint8(rand.Intn(256)) // 0, ..., 0xFF
			datBuf[i] = b
		}
		dat := datBuf[:length]
		fmt.Println()
		n := cobs.CEncode(encBuf, dat)
		enc := encBuf[:n]
		assert.False(t, bytes.Contains(enc, []byte{0}))
		n, e := cobs.CDecode(decBuf, enc)
		assert.True(t, e == nil)
		dec := decBuf[:n]
		assert.Equal(t, dat, dec)
	}
}

func TestCEncodeDecode256(t *testing.T) {
	max := 32768
	for i := 0; i < 10000; i++ {
		length := rand.Intn(max)
		datBuf := make([]byte, max)
		encBuf := make([]byte, 2*max)
		decBuf := make([]byte, 2*max)
		for i := 0; i < length; i++ {
			b := uint8(rand.Intn(256)) // 0, ..., 0xFF
			datBuf[i] = b
		}
		dat := datBuf[:length]
		fmt.Println()
		n := cobs.CEncode(encBuf, dat)
		enc := encBuf[:n]
		assert.False(t, bytes.Contains(enc, []byte{0}))
		n, e := cobs.Decode(decBuf, enc)
		assert.True(t, e == nil)
		dec := decBuf[:n]
		assert.Equal(t, dat, dec)
	}
}
