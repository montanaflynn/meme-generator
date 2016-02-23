package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"

	"github.com/fogleman/gg"
)

const fontSize = 36

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
	meme := flag.String("meme", "fry", "the meme to use")
	text := flag.String("text", "not sure what to put here", "the text to use")
	flag.Parse()

	tail := flag.Args()
	if len(tail) == 0 {
		flag.Usage()
	}

	path := tail[0]

	memes := make(map[string]string)
	memes["fry"] = "Futurama-Fry.jpg"
	memes["aliens"] = "Ancient-Aliens.jpg"
	memes["doge"] = "Doge.jpg"

	img := LoadImage("./media/" + memes[*meme])
	r := img.Bounds()
	w := r.Dx()
	h := r.Dy()

	m := gg.NewContext(w, h)
	m.DrawImage(img, 0, 0)
	m.SetRGB255(255, 255, 255)
	m.SetLineWidth(10)
	m.LoadFontFace("/Library/Fonts/Impact.ttf", fontSize)
	m.DrawStringAnchored(*text, float64(w)/2, float64(h)-fontSize, .5, .5)
	m.SavePNG(path)
	fmt.Printf("Saved to %s\n", path)
}
