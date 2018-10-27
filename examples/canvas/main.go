//-----------------------------------------------------------------------------
/*

Canvas Demo

*/
//-----------------------------------------------------------------------------

package main

import (
	"fmt"

	"github.com/deadsy/gonk/nk"
)

//-----------------------------------------------------------------------------

func main() {

	atlas := nk.NewFontAtlas()
	atlas.InitDefault()
	atlas.Begin()
	font := atlas.AddDefault()
	image := atlas.Bake()
	//device_upload_atlas(&device, image, w, h);
	atlas.End()

	fmt.Printf("%v\n", atlas)

}

//-----------------------------------------------------------------------------
