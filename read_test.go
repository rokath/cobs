// Package cobs_test contains blackbox tests.
package cobs_test

import (
	"bytes"
	"github/rokath/cobs"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSingleRead concatenates encoded testData with 0-delimiters and expects that Read returns decoded single packages.
func TestSingleRead(t *testing.T) {
	enc := CreateEncodedStream()
	p := cobs.NewDecoder(bytes.NewReader(enc), testBufferSize)
	dec := make([]byte, testBufferSize)
	for _, x := range testData {
		n, e := p.Read(dec)
		assert.True(t, e == nil || e == io.EOF)
		assert.Equal(t, x.dec, dec[:n])
	}
}

// TestMultiReadPackets decodes 6 TCOBS packages with 0-delimiters and expects that Read returns 1 package decoded in one call.
// As long a package is returned the return value is nil. If no more decodable data return is io.EOF.
func TestMultiReadPackets(t *testing.T) {
	enc := []byte{
		0,
		3, 11, 22, 2, 33, 0,
		4, 11, 22, 33, 0,
		05, 11, 22, 33, 44, 0,
		0,
	}
	exp := [][]byte{
		{},
		{11, 22, 0, 33},
		{11, 22, 33},
		{11, 22, 33, 44},
		{},
	}

	p := cobs.NewDecoder(bytes.NewReader(enc), testBufferSize)
	dec := make([]byte, testBufferSize)
	for i := range exp {
		n, e := p.Read(dec)
		assert.Nil(t, e)
		assert.Equal(t, exp[i], dec[:n])
	}
}

func TestMultiReadPacketsWithOneErrorPackageInside(t *testing.T) {
	enc := []byte{
		0,
		3, 11 /*22,*/, 2, 33, 0, // error injected: missing byte
		4, 11, 22, 33, 0,
		05, 11, 22, 33, 44, 0,
		0,
	}
	exp := [][]byte{
		{},
		{}, // {11, 22, 0, 33},
		{11, 22, 33},
		{11, 22, 33, 44},
		{},
	}

	p := cobs.NewDecoder(bytes.NewReader(enc), testBufferSize)
	dec := make([]byte, testBufferSize)
	var ecnt int
	for i := range exp {
		n, e := p.Read(dec)
		if e != nil {
			ecnt++
		}
		assert.Equal(t, exp[i], dec[:n])
	}
	assert.Equal(t, 1, ecnt)
}

// When the internal buffer is smaller than the encoded buffer, no read is possible.
func TestTooSmallInnerBufferRead(t *testing.T) {
	enc := []byte{
		3, 11, 22, 2, 33, 4, 11, 22, 33, 05, 11, 22, 33, 44, 0,
		4, 11, 22, 33, 0,
	}
	p := cobs.NewDecoder(bytes.NewReader(enc), 10)
	dec := make([]byte, 100)

	// package does not fit in internal buffer and is ignored therefore
	n, e := p.Read(dec)
	assert.True(t, e != nil)
	assert.True(t, e != io.EOF)
	assert.True(t, n == 0)

	// next package is ok
	n, e = p.Read(dec)
	assert.Nil(t, e)
	assert.True(t, n == 3)
}

// When the provided buffer is smaller than the decoded buffer, no read is possible.
func TestTooSmallProvidedBufferRead(t *testing.T) {
	enc := []byte{
		3, 11, 22, 2, 33, 4, 11, 22, 33, 05, 11, 22, 33, 44, 0,
		4, 11, 22, 33, 0,
	}
	p := cobs.NewDecoder(bytes.NewReader(enc), 150)
	dec := make([]byte, 8)

	// package does not fit in external buffer
	n, e := p.Read(dec)
	assert.True(t, e != nil)
	assert.True(t, e != io.EOF)
	assert.True(t, n == 0)

	// next package is ok
	n, e = p.Read(dec)
	assert.Nil(t, e)
	assert.True(t, n == 3)
}
