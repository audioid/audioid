// Copyright 2019-present Audioid contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://github.com/audioid/audioid/tree/master/LICENSE

package flac

import (
	"encoding/binary"
	"io"
	"strings"

	"github.com/audioid/audioid/errors"
	"github.com/audioid/audioid/utils"
	"github.com/valyala/bytebufferpool"
)

func (block *MetadataBlock) LoadVorbisComment(f io.ReadSeeker, bb *bytebufferpool.ByteBuffer) error {
	block.Type = BlockTypeVorbisComment

	comment := &VorbisComment{
		Comments: map[string]string{},
	}
	// In Vorbis, the vendor field is stored separately
	// https://xiph.org/flac/api/structFLAC____StreamMetadata__VorbisComment.html

	// length is uint32 according to
	// https://xiph.org/flac/api/structFLAC____StreamMetadata__VorbisComment__Entry.html
	vendorLen := uint32(0)
	err := binary.Read(f, binary.LittleEndian, &vendorLen)
	if err != nil {
		return err
	}

	vendor, err := utils.ReadCString(bb, f, vendorLen)
	if err != nil {
		return err
	}
	comment.Vendor = vendor

	commentsLength := uint32(0)
	err = binary.Read(f, binary.LittleEndian, &commentsLength)
	if err != nil {
		return err
	}

	for i := uint32(0); i < commentsLength; i++ {
		length := uint32(0)
		err = binary.Read(f, binary.LittleEndian, &length)
		if err != nil {
			return err
		}
		s, err := utils.ReadCString(bb, f, length)
		if err != nil {
			return err
		}
		parts := strings.Split(s, "=")
		if len(parts) != 2 {
			return errors.New("Invalid vorbis comment: " + s)
		}
		// Key is case-insensitive
		// https://www.xiph.org/vorbis/doc/v-comment.html
		comment.Comments[strings.ToLower(parts[0])] = parts[1]
	}

	block.Data = comment
	return nil
}
