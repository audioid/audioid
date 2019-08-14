// Copyright 2019-present Audioid contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://github.com/audioid/audioid/tree/master/LICENSE

package utils

import (
	"encoding/binary"
	"io"

	"github.com/audioid/audioid/errors"
	"github.com/valyala/bytebufferpool"
)

func ReadUint24BE(bb *bytebufferpool.ByteBuffer, r io.Reader) (uint32, error) {
	bb.Reset()
	if err := ReadBytes(bb, r, 24/8); err != nil {
		return 0, errors.Wrap("could not read uint24", err)
	}

	// return uint32(bb.B[0])<<16 | uint32(bb.B[1])<<8 | uint32(bb.B[2]), nil
	// return uint32(bb.B[2])<<16 | uint32(bb.B[1])<<8 | uint32(bb.B[0]), nil
	return binary.BigEndian.Uint32([]byte{bb.B[0], bb.B[1], bb.B[2], 0}), nil
}

func ReadCString(bb *bytebufferpool.ByteBuffer, r io.Reader, length uint32) (string, error) {
	bb.Reset()
	if err := ReadBytes(bb, r, length); err != nil {
		return "", err
	}

	return bb.String(), nil
}

func ReadBytes(bb *bytebufferpool.ByteBuffer, r io.Reader, length uint32) error {
	Grow(bb, length)
	_, err := r.Read(bb.B)
	if err != nil {
		return err
	}

	return nil
}

func ReadInt(bb *bytebufferpool.ByteBuffer, r io.Reader, length int) (int, error) {
	bb.Reset()
	err := ReadBytes(bb, r, uint32(length))
	if err != nil {
		return 0, err
	}
	var n int
	for _, x := range bb.B {
		n = n << 8
		n |= int(x)
	}
	return n, nil
}
