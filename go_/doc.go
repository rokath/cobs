// Package cobs implements COBS.
// For encoding create an Encoder and use the Write method as shown in TestWrite.
// For decoding create a Decoder and use the Read method as shown in TestSingleRead and TestMultiRead.
// It is possible to use only the functions cobs.Decode or cobs.CEncode.
// Package cobs is also a helper for testing the target C-code.
// Each C function gets a Go wrapper which ist tested in appropriate test functions.
// For some reason inside the cobs_test.go an 'import "C"' is not possible.
package cobs
