// Copyright 2019-present Audioid contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://github.com/audioid/audioid/tree/master/LICENSE

package flac

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"

	"github.com/audioid/audioid/errors"
	"github.com/audioid/audioid/metadata"
	"github.com/audioid/audioid/utils"
	"github.com/valyala/bytebufferpool"
)

// LoadStreamInfo reads stream info from f using given bb.
// Caller must reset bb before call.
//
// ref: https://xiph.org/flac/api/structFLAC____StreamMetadata__StreamInfo.html
func (block *MetadataBlock) LoadStreamInfo(f io.ReadSeeker, bb *bytebufferpool.ByteBuffer) error {
	block.Type = BlockTypeStreamInfo
	stream := &StreamInfo{}

	if err := binary.Read(f, binary.BigEndian, &stream.MinBlockSize); err != nil {
		return errors.Wrap("could not read MinBlockSize", err)
	}
	if err := binary.Read(f, binary.BigEndian, &stream.MaxBlockSize); err != nil {
		return errors.Wrap("could not read MaxBlockSize", err)
	}

	n, err := utils.ReadUint24BE(bb, f)
	if err != nil {
		return errors.Wrap("could not read MinFrameSize", err)
	}
	stream.MinFrameSize = n

	n, err = utils.ReadUint24BE(bb, f)
	if err != nil {
		return errors.Wrap("could not read MaxFrameSize", err)
	}
	stream.MaxFrameSize = n

	x := uint64(0)

	if err := binary.Read(f, binary.BigEndian, &x); err != nil {
		return errors.Wrap("could not read SampleRate, Channels, BPS and TotalSamples", err)
	}

	// 20bits
	// 1111 1111 1111 1111 1111 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000
	stream.SampleRate = uint32(x >> 44)
	// 3 bits
	// 0000 0000 0000 0000 0000 1110 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000
	stream.Channels = uint8((x<<20)>>41) + 1
	// 5 bits
	// 0000 0000 0000 0000 0000 0001 1111 0000 0000 0000 0000 0000 0000 0000 0000 0000
	// stream.BitsPerSample = uint8((x << 23) >> 36)
	stream.BitsPerSample = uint8((x<<23)>>59) + 1

	// 36 bits
	// 0000 0000 0000 0000 0000 0000 0000 1111 1111 1111 1111 1111 1111 1111 1111 1111
	stream.TotalSamples = uint64(x<<28) >> 28

	bb.Reset()
	utils.Grow(bb, 16)
	_, err = f.Read(bb.B)
	if err != nil {
		return errors.Wrap("could not read MD5Sum", err)
	}

	stream.MD5Sum = fmt.Sprintf("%x", bb.B)

	block.Data = stream
	return nil
}

func (stream *StreamInfo) Apply(t *metadata.Track) {
	if stream.TotalSamples != 0 && stream.SampleRate != 0 {
		t.Duration = time.Duration(stream.TotalSamples/uint64(stream.SampleRate)) * time.Second
	} else {
		t.Duration = -1
	}

	t.Checksum = metadata.Checksum{
		Algorithm: metadata.AlgoMD5,
		Sum:       stream.MD5Sum,
	}
}
