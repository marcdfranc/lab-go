// gera animações gif de figuras de lissajous58 aleatórias

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

var palette2 = []color.Color{color.White, color.Black}

const (
	blackIndex2 = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous2(os.Stdout)
}

func lissajous2(out io.Writer) {
	const (
		cycles  = 5     // Numero de ocilações completas do oscilador
		res     = 0.001 // resolução angular
		size    = 100   // canvas da imagem cobre de [-size ... + size]
		nFrames = 64    // número de quadros da animação
		delay   = 8     // tempo entre quadros em unidades de 10ms
	)

	freq := rand.Float64() * 3.0 // Frequencia relativa do oscilador

	anim := gif.GIF{LoopCount: nFrames}

	phase := 0.0 // diferença de fase

	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette2)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex2)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTA: Ignorando erros de codificacao
}
