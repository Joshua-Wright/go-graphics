package texture

import (
	g "github.com/joshua-wright/go-graphics/graphics"
	"image"
	"image/color"
)

type Texture interface {
	GetPixel(p g.Vec2) color.NRGBA
}

type ImageTexture struct {
	Img *image.NRGBA
}

func (t ImageTexture) GetPixel(p g.Vec2) color.NRGBA {
	xi, yi := g.WindowTransformPoint(t.Img.Bounds().Dx(), t.Img.Bounds().Dy(), p, [4]g.Float{0, 1, 0, 1})
	return t.Img.NRGBAAt(t.Img.Bounds().Min.X+xi, t.Img.Bounds().Min.Y+yi)
}

func TextureFromImage(filename string) (*ImageTexture, error) {
	img, err := g.OpenImageAsNRGBA(filename)
	if err != nil {
		return nil, err
	}
	return &ImageTexture{
		Img: img,
	}, nil
}
