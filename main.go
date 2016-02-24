package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"net/http"
	"os"

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

	// Apply black stroke
	m.SetHexColor("#000")
	strokeSize := 6
	for dy := -strokeSize; dy <= strokeSize; dy++ {
		for dx := -strokeSize; dx <= strokeSize; dx++ {
			// give it rounded corners
			if dx*dx+dy*dy >= strokeSize*strokeSize {
				continue
			}
			x := float64(w/2 + dx)
			y := float64(h - fontSize + dy)
			m.DrawStringAnchored(*text, x, y, 0.5, 0.5)
		}
	}

	// Apply white fill
	m.SetHexColor("#FFF")
	m.DrawStringAnchored(*text, float64(w)/2, float64(h)-fontSize, 0.5, 0.5)
	m.SavePNG(path)

	fmt.Printf("Saved to %s\n", path)
}
