// Minimal echo server, to serve a lissajour animated gif file
// Build server
//
//	     ensure you are on the server1 folder: ~/projects/go-kernighan/ch01/07-web-server/exercise_1_15
//	     Ensure the Makefile APP attribute is set: APP=exercise_1_15
//			$ make
//
// Launch server
//
//	     $ ~/projects/go-kernighan/bin/server4 &, or
//			$ make run
//
// Trigger server behavior
//
//	$ ~/projects/go-kernighan/bin/exercise_1_11
//  $ make client
//
// Remove server
//
//	Find out the port server4 is listening to
//		$ sudo lsof -i :8080
//
//	Kill the process
//		$ kill -9 <<process id>>
// todo: refactor to use the original lissajous, instead of copying and pasting it here.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	//http.HandleFunc("/", handler) // each request call handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))
		if err != nil {
			cycles = 5
		}
		lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func lissajous(out io.Writer, cycles int) {
	const (
		//cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(out, &anim)
}
