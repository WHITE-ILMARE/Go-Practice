// 生成利萨如曲线gif图,作为server1导入的外部包
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0x36, 0x78, 0x23, 0x2E}}

const (
	whiteIndex = 0 // 调色板(palette)中的第一个颜色
	blackIndex = 1 // 调色板中的下一个颜色
	selfColorIndex = 2
)

func Lissajous(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), selfColorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // ignoring encoding errors
}

//func main() {
//	lissajous(os.Stdout)
//}
