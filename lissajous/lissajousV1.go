// Lissajous generates GIF animations of random Lissajous figures.
package main

import "image"
import "image/color"
import "image/gif"
import "io"
import "math"
import "math/rand"
import "os"

var palette []color.Color = []color.Color{color.Black, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}, color.RGBA{0x94, 0x00, 0xd3, 0xff}, color.RGBA{0x1e, 0x90, 0xff, 0xff}}

var colorIndex []uint8 = []uint8{0, 1, 2, 3, 4, 5}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const cycle float64 = 5                 // Number of complete x oscillator revolution
	const angularResolution float64 = 0.001 // angular resolution
	const size int = 100                    // image canvas size [-size..+size]
	const numberOfFrames int = 64           // Number of animation frames
	const delay int = 8                     // Delay between frames in 10ms unit

	var frequency float64 = rand.Float64() * 3.0 // Relative frequency of y oscillator
	var animation gif.GIF = gif.GIF{LoopCount: numberOfFrames}
	var phase float64 = 0.0 // Phase difference
	var nextColorIndex int = 1

	for counter := 0; counter < numberOfFrames; counter++ {
		var rect image.Rectangle = image.Rect(0, 0, 2*size+1, 2*size+1)
		var img *image.Paletted = image.NewPaletted(rect, palette)

		nextColorIndex++

		if nextColorIndex >= len(colorIndex) {
			nextColorIndex = 1
		}

		for cycleCounter := 0.0; cycleCounter < cycle*2.0*math.Pi; cycleCounter += angularResolution {
			var x float64 = math.Sin(cycleCounter)
			var y float64 = math.Sin(cycleCounter*frequency + phase)

			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), colorIndex[nextColorIndex])
		}

		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, img)
	}

	gif.EncodeAll(out, &animation)
}
