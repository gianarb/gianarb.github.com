package main

import (
	"bytes"
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		mkFile := filepath.Join(cwd, args[0])
		fmt.Println(fmt.Sprintf("parsing file %s", mkFile))
		f, err := os.Open(mkFile)
		if err != nil {
			panic(err)
		}
		s, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		markdown := goldmark.New(
			goldmark.WithExtensions(
				meta.Meta,
			),
		)

		var buf bytes.Buffer
		context := parser.NewContext()
		if err := markdown.Convert(s, &buf, parser.WithContext(context)); err != nil {
			panic(err)
		}
		metaData := meta.Get(context)

		img := metaData["heroimg"].(string)
		if img == "" {
			img = fallbackBackground
		}
		if img == "" {
			panic("image not found. Set a fallback or an heroimg param in the markdown file")
		}

		if err := generate(GenerateReq{
			Title:       metaData["title"].(string),
			Background:  img,
			Destination: filepath.Join(cwd, imgDestination),
		}); err != nil {
			panic(err)
		}
	},
}

var (
	imgDestination     string
	fallbackBackground string
	cwd                string
)

func init() {
	dir, _ := filepath.Abs(filepath.Dir("."))
	rootCmd.Flags().StringVar(&imgDestination, "img-destination", "./socialimg.png", "Where the social image will be saved")
	rootCmd.Flags().StringVar(&fallbackBackground, "fallback-background", "", "Image that will be used when hereimg is not available")
	rootCmd.Flags().StringVar(&cwd, "cwd", dir, "current working directory")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func generate(req GenerateReq) error {
	dc := gg.NewContext(1200, 628)

	backgroundImage, err := gg.LoadImage(fmt.Sprintf("%s%s", cwd, req.Background))
	if err != nil {
		return errors.Wrap(err, "load background image")
	}
	backgroundImage = imaging.Fill(backgroundImage, dc.Width(), dc.Height(), imaging.Center, imaging.Lanczos)

	dc.DrawImage(backgroundImage, 0, 0)

	margin := 20.0
	x := margin
	y := margin
	w := float64(dc.Width()) - (2.0 * margin)
	h := float64(dc.Height()) - (2.0 * margin)
	dc.SetColor(color.RGBA{0, 0, 0, 204})
	dc.DrawRectangle(x, y, w, h)
	dc.Fill()

	textShadowColor := color.Black
	fontPath := filepath.Join(cwd, "fonts", "Roboto", "Roboto-Bold.ttf")
	if err := dc.LoadFontFace(fontPath, 90); err != nil {
		return errors.Wrap(err, "load Playfair_Display")
	}
	textRightMargin := 60.0
	textTopMargin := 90.0
	x = textRightMargin
	y = textTopMargin
	maxWidth := float64(dc.Width()) - textRightMargin - textRightMargin
	dc.SetColor(textShadowColor)
	dc.DrawStringWrapped(req.Title, x+1, y+1, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	dc.SetColor(color.White)
	dc.DrawStringWrapped(req.Title, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)

	fontPath = filepath.Join(cwd, "fonts", "Roboto", "Roboto-Regular.ttf")
	if err := dc.LoadFontFace(fontPath, 60); err != nil {
		return errors.Wrap(err, "load font")
	}
	r, g, b, _ := color.White.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(200),
	}
	dc.SetColor(mutedColor)
	s := "https://gianarb.it/"
	_, textHeight := dc.MeasureString(s)
	x = 70
	y = float64(dc.Height()) - textHeight - 30
	dc.DrawString(s, x, y)

	fontPath = filepath.Join(cwd, "fonts", "Roboto", "Roboto-Regular.ttf")
	if err := dc.LoadFontFace(fontPath, 60); err != nil {
		return errors.Wrap(err, "load font")
	}
	r, g, b, _ = color.White.RGBA()
	mutedColor = color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(200),
	}
	dc.SetColor(mutedColor)
	s = "@gianarb"
	textWd, textHeight := dc.MeasureString(s)
	x = float64(dc.Width()) - textWd - 100
	y = float64(dc.Height()) - textHeight - 30
	dc.DrawString(s, x, y)

	if err := dc.SavePNG(req.Destination); err != nil {
		return errors.Wrap(err, "save png")
	}
	return nil
}

type GenerateReq struct {
	Title       string
	Background  string
	Destination string
}
