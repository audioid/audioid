// Package metadata provides types for media metadata
//
// Copyright 2019-present Audioid contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://github.com/audioid/audioid/tree/master/LICENSE
package metadata

import "time"

type ChecksumAlgo uint8

const (
	AlgoUnknown ChecksumAlgo = 0
	AlgoMD5     ChecksumAlgo = 1
)

type Checksum struct {
	Algorithm ChecksumAlgo
	Sum       string
}

type Picture struct {
	// MIME MIME type string, in printable ASCII characters 0x20-0x7e.
	// The MIME type may also be --> to signify that the data part is
	// a URL of the picture instead of the picture data itself.
	MIME string
	// Description of the picture.
	Description string
	// Image dimensions.
	Width, Height uint32
	// Color depth in bits-per-pixel.
	Depth uint32
	// PaletteColors numer. 0 implies non-indexed image.
	PaletteColors uint32
	// Image data.
	Data []byte
	// IsPictureLink determines
	// if Picture.Data handles a link to the image
	IsPictureLink bool
}

// Track holds basic track information
//
// ref: https://www.xiph.org/vorbis/doc/v-comment.html
type Track struct {
	// Title is the Track/Work name
	Title string `json:"title,omitempty"`
	// Version is the version field may be used to differentiate
	// multiple versions of the same track title in a single collection. (e.g. remix info)
	Version string `json:"version,omitempty"`
	// Artist is the artist generally considered responsible for the work.
	// In popular music this is usually the performing band or singer.
	// For classical music it would be the composer. For an audio book
	// it would be the author of the original text.
	Artist string `json:"artist,omitempty"`
	// Performer is the artist(s) who performed the work. In classical music this
	// would be the conductor, orchestra, soloists. In an audio book
	// it would be the actor who did the reading. In popular music
	// this is typically the same as the ARTIST and is omitted.
	Performer string `json:"performer,omitempty"`
	// Copyright attribution, e.g., '2001 Nobody's Band' or '1999 Jack Moffitt'
	Copyright string `json:"copyright,omitempty"`
	// License information, eg, 'All Rights Reserved', 'Any Use Permitted',
	// a URL to a license such as a Creative Commons license ("www.creativecommons.org/blahblah/license.html")
	// or the EFF Open Audio License ('distributed under the terms of the Open Audio License.
	// see http://www.eff.org/IP/Open_licenses/eff_oal.html for details'), etc.
	License string `json:"license,omitempty"`
	// Contact information for the creators or distributors of the track.
	// This could be a URL, an email address, the physical address of the producing label
	Contact string `json:"contact,omitempty"`
	// Organization is the name of the organization producing the track (i.e. the 'record label')
	Organization string `json:"organization,omitempty"`
	// Album is the collection name to which this track belongs
	Album string `json:"album,omitempty"`
	// Genre is a short text indication of music genre
	Genre string `json:"genre,omitempty"`
	// TrackNumber is the track number of this piece if part of a specific larger collection or album
	TrackNumber string `json:"trackNumber,omitempty"`
	// Comments is the rest of comments about the Track,
	// which was not supported as a dedicated field
	Comments map[string]string `json:"comments,omitempty"`
	// Description is a short text description of the contents
	Description string `json:"description,omitempty"`
	// Date the track was recorded
	Date string `json:"date,omitempty"`
	// Location where track was recorded
	Location string `json:"location,omitempty"`
	// ISRC Track number; see the ISRC intro page for more information on ISRC numbers
	//
	// ref: https://isrc.ifpi.org/en/
	ISRC string `json:"isrc,omitempty"`
	// Duration returns track duration.
	// Negative duration means unknown.
	Duration time.Duration `json:"duration,omitempty"`
	// Checksum is the checksum of contents
	Checksum Checksum

	Pictures []Picture
}
