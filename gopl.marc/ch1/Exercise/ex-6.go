// gera animações gif de figuras de lissajous aleatórias

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

var palette = []color.Color{
	color.Black,
	color.RGBA{0x4d, 0xff, 0x4d, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0x00, 0xcc, 0xff, 0xff}}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Numero de ocilações completas do oscilador
		res     = 0.001 // resolução angular
		size    = 100   // canvas da imagem cobre de [-size ... + size]
		nFrames = 64    // número de quadros da animação
		delay   = 8     // tempo entre quadros em unidades de 10ms
	)

	idx := 1

	freq := rand.Float64() * 3.0 // Frequencia relativa do oscilador

	anim := gif.GIF{LoopCount: nFrames}

	phase := 0.0 // diferença de fase

	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(idx))
		}
		if idx < 3 {
			idx++
		} else {
			idx = 1
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTA: Ignorando erros de codificacao
}