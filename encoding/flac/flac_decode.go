// Copyright 2019-present Audioid contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://github.com/audioid/audioid/tree/master/LICENSE

package flac

import (
	"io"

	"github.com/audioid/audioid/errors"
	"github.com/audioid/audioid/metadata"
	"github.com/valyala/bytebufferpool"
)

var (
	ErrorBrokenSeeker = errors.New("broken seeker")
)

// Decode checks if f contains fLaC header, and parses metadata
// into *metadata.Track
func Decode(f io.ReadSeeker) (*metadata.Track, error) {
	bb := bytebufferpool.Get()
	track, err := DecodeUsingBuffer(f, bb)
	bytebufferpool.Put(bb)
	return track, err
}

// Decode checks if f contains fLaC header, and parses metadata
// into *metadata.Track using given byte buffer.
// You must reset bb, if it was used before.
func DecodeUsingBuffer(f io.ReadSeeker, bb *bytebufferpool.ByteBuffer) (*metadata.Track, error) {
	n, err := f.Read(bb.B)
	if n < 4 {
		return nil, errors.Wrap("invalid file: no flac header", err)
	}
	return DecodeFlacUsingBuffer(f, bb)
}

// Decode parses metadata
// into *metadata.Track using given byte buffer.
// This function DOES NOT check if f contains fLaC header.
// You must seek to 5th byte of the file before using this function.
//
// You must reset bb, if it was used before.
func DecodeFlacUsingBuffer(f io.ReadSeeker, bb *bytebufferpool.ByteBuffer) (*metadata.Track, error) {
	t := &metadata.Track{}
	for {
		block, err := parseBlock(f, bb)
		if err != nil {
			return nil, errors.Wrap("could not decode flac", err)
		}
		block.ApplyTo(t)
		if block.IsLast {
			break
		}
	}
	return t, nil
}
