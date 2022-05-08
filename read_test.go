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
	p := cobs.NewDecoder(bytes.NewReader(enc), testBufferSize, false)
	dec := make([]byte, testBufferSize)
	for _, x := range testData {
		n, e := p.Read(dec)
		assert.Nil(t, e)
		assert.Equal(t, x.dec, dec[:n])
	}
}

// TestMultiRead concatenates encoded testData with 0-delimiters and expects that Read returns all packages decoded in one call.
func TestMultiRead(t *testing.T) {
	enc := CreateEncodedStream()
	p := cobs.NewDecoder(bytes.NewReader(enc), testBufferSize, true)
	dec := make([]byte, testBufferSize)
	n, e := p.Read(dec)
	if e != nil {
		assert.True(t, e == io.EOF)
	}
	exp := CreateDecodedStream()
	assert.Equal(t, exp, dec[:n])
}

// TestMultiRead3Packets decodes 3 TCOBS packages with 0-delimiters and expects that Read returns all packages decoded in one call.
func TestMultiRead3Packets(t *testing.T) {
	enc := []byte{
		0, 3, 11, 22, 2, 33, 0,
		4, 11, 22, 33, 0,
		05, 11, 22, 33, 44, 0,
	}
	p := cobs.NewDecoder(bytes.NewReader(enc), testBufferSize, true)
	dec := make([]byte, testBufferSize)
	n, e := p.Read(dec)
	if e != nil {
		assert.True(t, e == io.EOF)
	}
	exp := []byte{
		11, 22, 0, 33,
		11, 22, 33,
		11, 22, 33, 44,
	}
	assert.Equal(t, exp, dec[:n])
}

// TestMultiRead3PacketsWith1stAsError decodes 3 TCOBS packages with 0-delimiters and expects that Read returns last 2 packages decoded in one call, because 1st package is wrong.
func TestMultiRead3PacketsWith1stAsError(t *testing.T) {
	enc := []byte{
		0, 3, 11, 22, 55, 2, 33, 0, // error injected
		4, 11, 22, 33, 0,
		05, 11, 22, 33, 44, 0,
	}
	p := cobs.NewDecoder(bytes.NewReader(enc), testBufferSize, true)
	dec := make([]byte, testBufferSize)
	n, e := p.Read(dec)
	assert.False(t, e == nil)
	assert.False(t, e == io.EOF)
	exp := []byte{
		//11, 22, 0, 33,
		11, 22, 33,
		11, 22, 33, 44,
	}
	assert.Equal(t, exp, dec[:n])
}

// TestMultiRead3PacketsWithLastIncomplete decodes 3 TCOBS packages with 0-delimiters and expects that Read returns all packages decoded in two calls.
func TestMultiRead3PacketsWithLastIncomplete(t *testing.T) {
	enc := []byte{
		0, 3, 11, 22, 2, 33, 0,
		4, 11, 22, 33, 0,
		05, 11, 22, 33, 44, 0,
	}
	p := cobs.NewDecoder(bytes.NewReader(enc), 16, true)
	dec := make([]byte, testBufferSize)
	n, e := p.Read(dec)
	assert.True(t, e == nil)
	exp := []byte{
		11, 22, 0, 33,
		11, 22, 33,
		//11, 22, 33, 44,
	}
	assert.Equal(t, exp, dec[:n])

	n, e = p.Read(dec)
	if e != nil {
		assert.True(t, e == io.EOF)
	}
	exp = []byte{
		//11, 22, 0, 33,
		//11, 22, 33,
		11, 22, 33, 44,
	}
	assert.Equal(t, exp, dec[:n])
}
