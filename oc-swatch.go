package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"time"
)

// RandomString ...
func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func main() {
	colors := []color.RGBA{
		color.RGBA{248, 249, 250, 255},
		color.RGBA{241, 243, 245, 255},
		color.RGBA{248, 249, 250, 255},
		color.RGBA{241, 243, 245, 255},
		color.RGBA{233, 236, 239, 255},
		color.RGBA{222, 226, 230, 255},
		color.RGBA{206, 212, 218, 255},
		color.RGBA{173, 181, 189, 255},
		color.RGBA{134, 142, 150, 255},
		color.RGBA{73, 80, 87, 255},
		color.RGBA{52, 58, 64, 255},
		color.RGBA{33, 37, 41, 255},
		color.RGBA{255, 245, 245, 255},
		color.RGBA{255, 227, 227, 255},
		color.RGBA{255, 201, 201, 255},
		color.RGBA{255, 168, 168, 255},
		color.RGBA{255, 135, 135, 255},
		color.RGBA{255, 107, 107, 255},
		color.RGBA{250, 82, 82, 255},
		color.RGBA{240, 62, 62, 255},
		color.RGBA{224, 49, 49, 255},
		color.RGBA{201, 42, 42, 255},
		color.RGBA{255, 240, 246, 255},
		color.RGBA{255, 222, 235, 255},
		color.RGBA{252, 194, 215, 255},
		color.RGBA{250, 162, 193, 255},
		color.RGBA{247, 131, 172, 255},
		color.RGBA{240, 101, 149, 255},
		color.RGBA{230, 73, 128, 255},
		color.RGBA{214, 51, 108, 255},
		color.RGBA{194, 37, 92, 255},
		color.RGBA{166, 30, 77, 255},
		color.RGBA{248, 240, 252, 255},
		color.RGBA{243, 217, 250, 255},
		color.RGBA{238, 190, 250, 255},
		color.RGBA{229, 153, 247, 255},
		color.RGBA{218, 119, 242, 255},
		color.RGBA{204, 93, 232, 255},
		color.RGBA{190, 75, 219, 255},
		color.RGBA{174, 62, 201, 255},
		color.RGBA{156, 54, 181, 255},
		color.RGBA{134, 46, 156, 255},
		color.RGBA{243, 240, 255, 255},
		color.RGBA{229, 219, 255, 255},
		color.RGBA{208, 191, 255, 255},
		color.RGBA{177, 151, 252, 255},
		color.RGBA{151, 117, 250, 255},
		color.RGBA{132, 94, 247, 255},
		color.RGBA{121, 80, 242, 255},
		color.RGBA{112, 72, 232, 255},
		color.RGBA{103, 65, 217, 255},
		color.RGBA{95, 61, 196, 255},
		color.RGBA{237, 242, 255, 255},
		color.RGBA{219, 228, 255, 255},
		color.RGBA{186, 200, 255, 255},
		color.RGBA{145, 167, 255, 255},
		color.RGBA{116, 143, 252, 255},
		color.RGBA{92, 124, 250, 255},
		color.RGBA{76, 110, 245, 255},
		color.RGBA{66, 99, 235, 255},
		color.RGBA{59, 91, 219, 255},
		color.RGBA{54, 79, 199, 255},
		color.RGBA{232, 247, 255, 255},
		color.RGBA{204, 237, 255, 255},
		color.RGBA{163, 218, 255, 255},
		color.RGBA{114, 195, 252, 255},
		color.RGBA{77, 173, 247, 255},
		color.RGBA{50, 154, 240, 255},
		color.RGBA{34, 138, 230, 255},
		color.RGBA{28, 124, 214, 255},
		color.RGBA{27, 110, 194, 255},
		color.RGBA{24, 98, 171, 255},
		color.RGBA{227, 250, 252, 255},
		color.RGBA{197, 246, 250, 255},
		color.RGBA{153, 233, 242, 255},
		color.RGBA{102, 217, 232, 255},
		color.RGBA{59, 201, 219, 255},
		color.RGBA{34, 184, 207, 255},
		color.RGBA{21, 170, 191, 255},
		color.RGBA{16, 152, 173, 255},
		color.RGBA{12, 133, 153, 255},
		color.RGBA{11, 114, 133, 255},
		color.RGBA{230, 252, 245, 255},
		color.RGBA{195, 250, 232, 255},
		color.RGBA{150, 242, 215, 255},
		color.RGBA{99, 230, 190, 255},
		color.RGBA{56, 217, 169, 255},
		color.RGBA{32, 201, 151, 255},
		color.RGBA{18, 184, 134, 255},
		color.RGBA{12, 166, 120, 255},
		color.RGBA{9, 146, 104, 255},
		color.RGBA{8, 127, 91, 255},
		color.RGBA{235, 251, 238, 255},
		color.RGBA{211, 249, 216, 255},
		color.RGBA{178, 242, 187, 255},
		color.RGBA{140, 233, 154, 255},
		color.RGBA{105, 219, 124, 255},
		color.RGBA{81, 207, 102, 255},
		color.RGBA{64, 192, 87, 255},
		color.RGBA{55, 178, 77, 255},
		color.RGBA{47, 158, 68, 255},
		color.RGBA{43, 138, 62, 255},
		color.RGBA{244, 252, 227, 255},
		color.RGBA{233, 250, 200, 255},
		color.RGBA{216, 245, 162, 255},
		color.RGBA{192, 235, 117, 255},
		color.RGBA{169, 227, 75, 255},
		color.RGBA{148, 216, 45, 255},
		color.RGBA{130, 201, 30, 255},
		color.RGBA{116, 184, 22, 255},
		color.RGBA{102, 168, 15, 255},
		color.RGBA{92, 148, 13, 255},
		color.RGBA{255, 249, 219, 255},
		color.RGBA{255, 243, 191, 255},
		color.RGBA{255, 236, 153, 255},
		color.RGBA{255, 224, 102, 255},
		color.RGBA{255, 212, 59, 255},
		color.RGBA{252, 196, 25, 255},
		color.RGBA{250, 176, 5, 255},
		color.RGBA{245, 159, 0, 255},
		color.RGBA{240, 140, 0, 255},
		color.RGBA{230, 119, 0, 255},
		color.RGBA{255, 244, 230, 255},
		color.RGBA{255, 232, 204, 255},
		color.RGBA{255, 216, 168, 255},
		color.RGBA{255, 192, 120, 255},
		color.RGBA{255, 169, 77, 255},
		color.RGBA{255, 146, 43, 255},
		color.RGBA{253, 126, 20, 255},
		color.RGBA{247, 103, 7, 255},
		color.RGBA{232, 89, 12, 255},
		color.RGBA{217, 72, 15, 255},
	}
	img := image.NewRGBA(image.Rect(0, 0, 128, 128))
	for _, color := range colors {
		draw.Draw(img, img.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)
		f, _ := os.Create(RandomString(10) + ".png")
		defer f.Close()
		png.Encode(f, img)
	}
}
