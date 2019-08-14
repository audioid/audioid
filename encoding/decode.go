// Package encoding is universal wrapper around
// Aʊdioid's codec implementations.
//
// Copyright 2019-present Audioid contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://github.com/audioid/audioid/tree/master/LICENSE
package encoding

import (
	"io"

	"github.com/audioid/audioid/encoding/flac"
	"github.com/audioid/audioid/errors"
	"github.com/audioid/audioid/metadata"
	"github.com/audioid/audioid/utils"

	"github.com/valyala/bytebufferpool"
)

// Decode given Reader into a Track.
// In current opensource release, this package supports only FLAC.
func Decode(r io.ReadSeeker) (*metadata.Track, error) {
	const detectionLength = 8
	bb := bytebufferpool.Get()
	utils.Grow(bb, detectionLength)

	_, err := io.ReadFull(r, bb.B)
	if err != nil {
		return nil, err
	}

	switch {
	case string(bb.B[:4]) == "fLaC":
		_, err := r.Seek(-detectionLength+4, io.SeekCurrent)
		if err != nil {
			return nil, errors.Wrap("could not seek back", err)
		}
		return flac.DecodeFlacUsingBuffer(r, bb)
	}

	return nil, errors.New("unknown file type")
}
