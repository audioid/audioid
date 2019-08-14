// Copyright 2019-present Audioid contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://github.com/audioid/audioid/tree/master/LICENSE

package flac

import (
	"github.com/audioid/audioid/metadata"
)

// Apply current VorbisComment to the track
func (vc *VorbisComment) Apply(t *metadata.Track) {
	for key, value := range vc.Comments {
		// ref: https://xiph.org/vorbis/doc/v-comment.html
		switch key {
		case "title":
			t.Title = value
		case "version":
			t.Version = value
		case "album":
			t.Album = value
		case "tracknumber":
			t.TrackNumber = value
		case "artist":
			t.Artist = value
		case "performer":
			t.Performer = value
		case "copyright":
			t.Copyright = value
		case "license":
			t.License = value
		case "organization":
			t.Organization = value
		case "description":
			t.Description = value
		case "genre":
			t.Genre = value
		case "date":
			t.Date = value
		case "location":
			t.Location = value
		case "isrc":
			t.ISRC = value
		default:
			if t.Comments == nil {
				t.Comments = map[string]string{}
			}
			t.Comments[key] = value
		}
	}
}
