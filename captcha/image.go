package captcha

import (
	"captcha-zh/config"
	"gopkg.in/gographics/imagick.v3/imagick"

	"math/rand"
	"strings"
)

var (
	// random colors
	colors []string = []string{"#000000", "#b50000", "#373000", "#827482"}
)

func init() {
	// get all fonts file here
	config.Fonts = config.Child(config.PATH_CONFIG_FONT)
}

func drawSetfont(mw *imagick.MagickWand, dw *imagick.DrawingWand) {
	pw := imagick.NewPixelWand()
	defer pw.Destroy()

	pw.SetColor(colors[rand.Intn(len(colors))])

	dw.SetFont(randFont())
	dw.SetFillColor(pw)
	dw.SetFontWeight(500)
	dw.SetFontSize(33)
}

func drawMetrics(mw *imagick.MagickWand, dw *imagick.DrawingWand, dx *float64, text string) {
	rotation := float64(random(-30, 30))
	mw.AnnotateImage(dw, *dx, 35, rotation, text)
	mw.DrawImage(dw)
	// get the font metrics
	fm := mw.QueryFontMetrics(dw, text)
	if fm != nil {
		// Adjust the new x coordinate
		*dx += fm.TextWidth + 2
	}
}

func writeWord(mw *imagick.MagickWand, dw *imagick.DrawingWand, dx *float64, text string) {
	drawSetfont(nil, dw)
	drawMetrics(mw, dw, dx, text)
}

func Draw(text string, name string) {
	// init imageick
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()

	// Current coordinates of text
	var dx float64 = 20

	// Set the size of the image
	mw.SetSize(285, 50)
	mw.ReadImage(config.PATH_CONFIG_IMAGE_BG)

	// Start near the left edge
	dw.SetFontSize(40)
	// Note that we must free up the fontmetric array once we're done with it
	list := strings.Split(text, " ")
	for _, item := range list {
		writeWord(mw, dw, &dx, item)
	}
	mw.DrawImage(dw)
	// Now write the magickwand image
	mw.WriteImage(name)
}
