package cobs

import (
	"bytes"
	"errors"
	"io"
)

type decoder struct {
	r    io.Reader // inner reader
	iBuf []byte    // input buffer
	iCnt int       // valid data inside iBuf
}

// NewDecoder creates a decoder instance and returns its address.
// r will be used as inner reader.
// size is used as size for the inner buffers.
// Max one decoded COBS package is returned by Read.
func NewDecoder(r io.Reader, size int) (p *decoder) {
	p = new(decoder)
	p.r = r
	p.iBuf = make([]byte, size)
	return
}

// Read returns one decoded COBS package if available.
// io.EOF is returned when inner reader reached end of input stream.
// In case of a corrupted input buffer an error is returned with n = 0.
func (p *decoder) Read(buffer []byte) (n int, e error) {
	var dataInvalid bool
start:
	if len(p.iBuf[p.iCnt:]) == 0 {
		p.iCnt = 0 // drop data
		dataInvalid = true
	}
	nCnt, e := p.r.Read(p.iBuf[p.iCnt:])
	p.iCnt += nCnt
	if e != nil && e != io.EOF {
		return // internal Read error
	}
	if e == io.EOF && p.iCnt == 0 {
		return // no more data
	}
	before, after, found := bytes.Cut(p.iBuf[:p.iCnt], []byte{0})
	if !found {
		goto start
	}
	if dataInvalid {
		p.iCnt = copy(p.iBuf, after) // remove data
		dataInvalid = false
		e = errors.New("encoded package is bigger than internal buffer, cannot read")
		return
	}
	n, e = Decode(buffer, before)
	p.iCnt = copy(p.iBuf, after)
	if e != nil && e != io.EOF {
		n = 0
	}
	return
}
