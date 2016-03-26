// Lissajous generates GIF animations of random Lissajous figures.
package main

import "image"
import "image/color"
import "image/gif"
import "io"
import "math"
import "math/rand"
import "os"

var palette []color.Color = []color.Color{color.White, color.Black}

const whiteColorIndex uint8 = 0
const blackColorIndex uint8 = 1

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const cycle float64 = 5                 // Number of complete x oscillator revolution
	const angularResolution float64 = 0.001 // angular resolution
	const size int = 100                    // image canvas size [-size..+size]
	const numberOfFrames int = 64           // Number of animation frames
	const delay int = 8                     // Delay between frames in 10ms unit

	var frequency float64 = rand.Float64() * 0.3 // Relative frequency of y oscillator
	var animation gif.GIF = gif.GIF{LoopCount: numberOfFrames}
	var phase float64 = 0.0 // Phase difference

	for counter := 0; counter < numberOfFrames; counter++ {
		var rect image.Rectangle = image.Rect(0, 0, 2*size+1, 2*size+1)
		var img *image.Paletted = image.NewPaletted(rect, palette)

		for cycleCounter := 0.0; cycleCounter < cycle*2.0*math.Pi; cycleCounter += angularResolution {
			var x float64 = math.Sin(cycleCounter)
			var y float64 = math.Sin(cycleCounter*frequency + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackColorIndex)
		}

		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, img)
	}

	gif.EncodeAll(out, &animation)
}
