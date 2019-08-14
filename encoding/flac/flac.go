// Package flac implements FLAC codec.
//
// Copyright 2019-present Audioid contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://github.com/audioid/audioid/tree/master/LICENSE
package flac

import "fmt"

// MetadataBlock is a FLAC's polymorphic stream metadata block.
//
// ref: https://xiph.org/flac/api/structFLAC____StreamMetadata.html
type MetadataBlock struct {
	Type   BlockType
	IsLast bool
	Length uint
	Data   interface{}
}

func (meta *MetadataBlock) IsKnownType() bool {
	return meta.Type < BlockTypeReservedMin
}

// BlockType as defined in FLAC spec
// and official FLAC implementation
//
// ref: https://xiph.org/flac/api/group__flac__format.html#gac71714ba8ddbbd66d26bb78a427fac01
type BlockType uint8

const (
	BlockTypeStreamInfo    BlockType = 0
	BlockTypePadding       BlockType = 1
	BlockTypeApplication   BlockType = 2
	BlockTypeSeekTable     BlockType = 3
	BlockTypeVorbisComment BlockType = 4
	BlockTypeCueSheet      BlockType = 5
	BlockTypePicture       BlockType = 6
	// BlockTypeReservedMin used to identify reserved block type range start
	BlockTypeReservedMin BlockType = 7
	// BlockTypeReservedMax used to identify reserved block type range end
	BlockTypeReservedMax BlockType = 126
	// BlockTypeInvalid is defined invalid to avoid confusion with a frame sync code
	BlockTypeInvalid BlockType = 127
)

func (t BlockType) String() string {
	switch t {
	case BlockTypeStreamInfo:
		return "StreamInfo"
	case BlockTypePadding:
		return "Padding"
	case BlockTypeApplication:
		return "Application"
	case BlockTypeSeekTable:
		return "SeekTable"
	case BlockTypeVorbisComment:
		return "VorbisComment"
	case BlockTypeCueSheet:
		return "CueSheet"
	case BlockTypePicture:
		return "Picture"
	case BlockTypeInvalid:
		return "invalid"
	}
	if t >= BlockTypeReservedMin && t <= BlockTypeReservedMax {
		return "reserved"
	}

	return fmt.Sprintf("unknown<%d>", t)
}

// StreamInfo block
//
// ref: https://xiph.org/flac/api/structFLAC____StreamMetadata__StreamInfo.html
type StreamInfo struct {
	// MinBlockSize is the minimum block size (in samples) used in the stream.
	MinBlockSize uint16
	// MaxBlockSize is the maximum block size (in samples) used in the stream.
	// (Minimum blocksize == maximum blocksize) implies a fixed-blocksize stream.
	MaxBlockSize uint16
	// MinFrameSize is the minimum frame size (in bytes) used in the stream.
	// May be 0 to imply the value is not known.
	MinFrameSize uint32
	// MaxFrameSize is the maximum frame size (in bytes) used in the stream.
	// May be 0 to imply the value is not known.
	MaxFrameSize uint32
	// SampleRate in Hz.
	// Though 20 bits are available, the maximum sample rate
	// is limited by the structure of frame headers to 655350Hz.
	// Also, a value of 0 is invalid.
	SampleRate uint32
	// Channels is the (number of channels). FLAC supports from 1 to 8 channels.
	Channels uint8
	// BitsPerSample is the (bits per sample).
	// FLAC supports from 4 to 32 bits per sample.
	// Currently the reference encoder and decoders only support up to 24 bits per sample.
	BitsPerSample uint8
	// TotalSamples is the total number of samples in stream.
	// 'Samples' means inter-channel sample,
	// i.e. one second of 44.1Khz audio will have 44100 samples regardless of the number of channels.
	// A value of zero here means the number of total samples is unknown.
	TotalSamples uint64
	// MD5 signature of the unencoded audio data.
	// This allows the decoder to determine if an error exists in the audio data
	// even when the error does not result in an invalid bitstream.
	MD5Sum string
}

// VorbisComment defined in
//
// https://xiph.org/flac/api/structFLAC____StreamMetadata__VorbisComment.html
// ref: https://www.xiph.org/vorbis/doc/v-comment.html
//
// This structure was modified and does not implemented exactly as in standard
// to simplify working with Vorbis Comments in Go.
// We skipped NumComments, because it's available as len(vorbis.Comments).
// We replaced []string with string in form of "key=value"
// with a convinient map.
type VorbisComment struct {
	Vendor string
	// Comments are defined as a pointer to an array in C,
	// storing length as a uint32 in NumComments after Vendor,
	// and only then storing comments itself.
	Comments map[string]string
}
