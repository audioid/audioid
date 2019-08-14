// Copyright 2019-present Audioid contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at https://github.com/audioid/audioid/tree/master/LICENSE

package encoding

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/audioid/audioid/errors"
)

func TestFlacDecode(t *testing.T) {
	b, err := ioutil.ReadFile("../testdata/inputSCVAUP.flac")
	errors.Must(err)

	reader := bytes.NewReader(b)
	track, err := Decode(reader)
	if err != nil {
		t.Errorf("%+v", err)
	}

	if track.Artist != "1" {
		t.Errorf(`expected Artist to be "1", but got %q`, track.Artist)
	}

	if track.Title != "2" {
		t.Errorf(`expected Title to be "1", but got %q`, track.Title)
	}

	if x := track.Comments["replaygain_track_peak"]; x != "0.99996948" {
		t.Errorf(`expected Comments[replaygain_track_peak] to be "0.99996948", but got %q`, x)
	}

	if x := track.Comments["replaygain_track_gain"]; x != "-7.89 dB" {
		t.Errorf(`expected Comments[replaygain_track_peak] to be "-7.89 dB", but got %q`, x)
	}

	if x := track.Comments["replaygain_album_peak"]; x != "0.99996948" {
		t.Errorf(`expected Comments[replaygain_album_peak] to be "0.99996948", but got %q`, x)
	}

	if x := track.Comments["replaygain_album_gain"]; x != "-7.89 dB" {
		t.Errorf(`expected Comments[replaygain_album_peak] to be "-7.89 dB", but got %q`, x)
	}
}
