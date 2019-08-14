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

	"github.com/audioid/audioid/errors"
	"github.com/audioid/audioid/metadata"
	"github.com/audioid/audioid/utils"
	"github.com/valyala/bytebufferpool"
)

type PictureType byte

const (
	PictureTypeOther PictureType = iota
	PictureType32FileIcon
	PictureTypeOtherFileIcon
	PictureTypeCoverFront
	PictureTypeCoverBack
	PictureTypeLeafletPage
	PictureTypeMedia
	PictureTypeLeadArtist
	PictureTypeArtist
	PictureTypeConductor
	PictureTypeBand
	PictureTypeComposer
	PictureTypeLyricist
	PictureTypeRecordingLocation
	PictureTypeDuringRecording
	PictureTypeDuringPerformance
	PictureTypeScreenCapture
	PictureTypeFish
	PictureTypeIllustratoin
	PictureTypeArtistLogo
	PictureTypeStudioLogo
	PictureTypeInvalid
)

var pictureTypes = map[PictureType]string{
	0x00: "Other",
	0x01: "32x32 pixels 'file icon' (PNG only)",
	0x02: "Other file icon",
	0x03: "Cover (front)",
	0x04: "Cover (back)",
	0x05: "Leaflet page",
	0x06: "Media (e.g. lable side of CD)",
	0x07: "Lead artist/lead performer/soloist",
	0x08: "Artist/performer",
	0x09: "Conductor",
	0x0A: "Band/Orchestra",
	0x0B: "Composer",
	0x0C: "Lyricist/text writer",
	0x0D: "Recording Location",
	0x0E: "During recording",
	0x0F: "During performance",
	0x10: "Movie/video screen capture",
	0x11: "A bright coloured fish",
	0x12: "Illustration",
	0x13: "Band/artist logotype",
	0x14: "Publisher/Studio logotype",
}

func (t PictureType) String() string {
	if str, ok := pictureTypes[t]; ok {
		return str
	}
	return fmt.Sprintf("invalid<%d>", t)
}

// Picture contains the image data of an embedded picture.
//
// ref: https://www.xiph.org/flac/format.html#metadata_block_picture
type Picture struct {
	// Picture type according to the ID3v2 APIC frame:
	//
	//     0: Other
	//     1: 32x32 pixels 'file icon' (PNG only)
	//     2: Other file icon
	//     3: Cover (front)
	//     4: Cover (back)
	//     5: Leaflet page
	//     6: Media (e.g. label side of CD)
	//     7: Lead artist/lead performer/soloist
	//     8: Artist/performer
	//     9: Conductor
	//    10: Band/Orchestra
	//    11: Composer
	//    12: Lyricist/text writer
	//    13: Recording Location
	//    14: During recording
	//    15: During performance
	//    16: Movie/video screen capture
	//    17: A bright coloured fish
	//    18: Illustration
	//    19: Band/artist logotype
	//    20: Publisher/Studio logotype
	//
	// ref: http://id3.org/id3v2.4.0-frames
	Type PictureType
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
}

func (pic *Picture) Apply(t *metadata.Track) {
	t.Pictures = append(t.Pictures, metadata.Picture{
		MIME: pic.MIME,
		// Description of the picture.
		Description: pic.Description,
		// Image dimensions.
		Width:  pic.Width,
		Height: pic.Height,
		// Color depth in bits-per-pixel.
		Depth: pic.Depth,
		// PaletteColors numer. 0 implies non-indexed image.
		PaletteColors: pic.PaletteColors,
		// Image data.
		Data:          pic.Data,
		IsPictureLink: pic.MIME == "-->",
	})
}

func (block *MetadataBlock) LoadPictureBlock(f io.ReadSeeker, bb *bytebufferpool.ByteBuffer) error {
	block.Type = BlockTypePicture

	picture := &Picture{}

	rawPictureType := uint32(0)
	err := binary.Read(f, binary.BigEndian, &rawPictureType)
	if err != nil {
		return errors.Wrap("could not read picture type", err)
	}

	picture.Type = PictureType(rawPictureType)

	mimeLength, err := utils.ReadInt(bb, f, 4)
	if err != nil {
		return errors.Wrap("could not read picture MIME type length", err)
	}
	mime, err := utils.ReadCString(bb, f, uint32(mimeLength))
	if err != nil {
		return errors.Wrap("could not read picture MIME type", err)
	}
	picture.MIME = mime
	descriptionLength, err := utils.ReadInt(bb, f, 4)
	if err != nil {
		return errors.Wrap("could not read picture description length", err)
	}
	description, err := utils.ReadCString(bb, f, uint32(descriptionLength))
	if err != nil {
		return errors.Wrap("could not read picture description", err)
	}
	picture.Description = description

	n, err := utils.ReadInt(bb, f, 4)
	if err != nil {
		return errors.Wrap("could not read picture width", err)
	}

	picture.Width = uint32(n)

	n, err = utils.ReadInt(bb, f, 4)
	if err != nil {
		return errors.Wrap("could not read picture height", err)
	}
	picture.Height = uint32(n)

	n, err = utils.ReadInt(bb, f, 4)
	if err != nil {
		return errors.Wrap("could not read picture color depth", err)
	}
	picture.Depth = uint32(n)

	n, err = utils.ReadInt(bb, f, 4)
	if err != nil {
		return errors.Wrap("could not read picture color count", err)
	}
	picture.PaletteColors = uint32(n)

	pictureDataLength, err := utils.ReadInt(bb, f, 4)
	if err != nil {
		return errors.Wrap("could not read picture data length", err)
	}

	bb.Reset()
	err = utils.ReadBytes(bb, f, uint32(pictureDataLength))
	if err != nil {
		return errors.Wrap("could not read picture data", err)
	}

	picture.Data = bb.B[:]

	return nil
}
