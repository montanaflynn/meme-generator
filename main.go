package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"net/http"
	"os"

	_ "./vendor/statik"

	"github.com/fogleman/gg"
)

func DownloadTemplate(file string) image.Image {
	url := fmt.Sprintf("https://imgflip.com/s/meme/%s", file)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("%s template from %s failed because of %v", file, url, err)
	}
	defer res.Body.Close()
	image, _, err := image.Decode(res.Body)
	if err != nil {
		log.Fatalf("Could not decode %s because of %v", file, err)
	}
	return image
}

func main() {
	const fontSize = 36

	meme := flag.String("meme", "fry", "the meme to use")
	text := flag.String("text", "not sure what to put here", "the text to use")
	list := flag.Bool("list", false, "list of available memes")
	flag.Parse()

	memes := make(map[string]string)
	memes["fry"] = "Futurama-Fry.jpg"
	memes["aliens"] = "Ancient-Aliens.jpg"
	memes["doge"] = "Doge.jpg"
	memes["simply"] = "One-Does-Not-Simply.jpg"
	memes["wonka"] = "Creepy-Condescending-Wonka.jpg"
	memes["grumpy"] = "Grumpy-Cat.jpg"
	memes["raptor"] = "Philosoraptor.jpg"

	if *list {
		fmt.Println("Available memes:")
		for k, _ := range memes {
			fmt.Println(k)
		}
		os.Exit(0)
	}

	path := "./meme.png"
	args := flag.Args()
	if len(args) > 0 {
		path = args[0]
	}

	img := DownloadTemplate(memes[*meme])
	r := img.Bounds()
	w := r.Dx()
	h := r.Dy()

	m := gg.NewContext(w, h)
	m.DrawImage(img, 0, 0)
	m.LoadFontFace("/Library/Fonts/Impact.ttf", fontSize)
	m.SetRGB255(255, 255, 255)
	m.DrawStringAnchored(*text, float64(w)/2, float64(h)-fontSize, .5, .5)
	m.SavePNG(path)
	fmt.Printf("Saved to %s\n", path)
}
