package main

import (
	"image"
	_ "image/jpeg"
	"log"
	"os"

	"github.com/fogleman/gg"
)

const pixels = 500
const fontSize = 36
const meme = "not sure if magick or golang"

func LoadImage(filename string) image.Image {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	image, something, err := image.Decode(file)
	if err != nil {
		log.Fatal(something, err)
	}
	return image
}

func main() {

	memes := make(map[string]string)
	memes["fry"] = "Futurama-Fry.jpg"
	memes["aliens"] = "Ancient-Aliens.jpg"
	memes["doge"] = "Doge.jpg"

	img := LoadImage("./media/" + memes["fry"])
	r := img.Bounds()
	w := r.Dx()
	h := r.Dy()

	m := gg.NewContext(w, h)
	m.DrawImage(img, 0, 0)
	m.SetRGB(255, 255, 255)
	m.LoadFontFace("/Library/Fonts/Impact.ttf", fontSize)
	m.DrawStringAnchored(meme, float64(w)/2, float64(h)-fontSize, .5, .5)
	m.SavePNG("./meme.png")
}
