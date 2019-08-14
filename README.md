# Aʊdioid

![Audioid — Fast and Reliable audio tools](.github/images/banner@2x.png)

_This project is in opensourcing stage. It's stable, but may miss some valuable features._

------

**/encoding**

This package contains implementation for supported audio formats.
You can use implementations directly, if you want minimal dependecies.

For example, we have a simple example app for you to try it in action by running `./examples/simple`:
```
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
```
