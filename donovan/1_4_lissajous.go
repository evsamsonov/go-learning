package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.RGBA{0x00, 0xFF, 0x00, 0xff}, color.Black, color.White}

const (
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Кол-во полных колебаний
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения
		nFrames = 64    // Кол-во кадров анимации
		delay   = 8     // Задержка между кадрами
	)

	// Иницилизация rand текущим временем в наносекундах
	rand.Seed(time.Now().UTC().UnixNano())

	// Относительная частота колебаний Y
	freqY := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nFrames}

	// Разность фаз
	phase := 0.0

	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freqY + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
