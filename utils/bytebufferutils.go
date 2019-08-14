// Package utils is an internal audioid package.
// May be useful for advanced cases.
// Copyright 2019-present Audioid contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://github.com/audioid/audioid/tree/master/LICENSE
package utils

import (
	"github.com/valyala/bytebufferpool"
)

func Grow(bb *bytebufferpool.ByteBuffer, n uint32) {
	if uint32(len(bb.B)) == n {
		return
	}
	if uint32(len(bb.B)) >= n || uint32(cap(bb.B)) >= n {
		bb.B = bb.B[:n]
	} else {
		bb.B = make([]byte, n)
	}
}
