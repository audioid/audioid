package main

import (
	"log"
	"os"

	"github.com/audioid/audioid/encoding"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	file, err := os.Open("../../testdata/inputSCVAUP.flac")
	if err != nil {
		panic(err)
	}
	track, err := encoding.Decode(file)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	spew.Dump(track)
}
