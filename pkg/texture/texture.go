package texture

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"reflect"

	"github.com/PJSoftware/go-chernopengl/pkg/resourcePath"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Texture struct {
	RendererID         uint32
	FilePath           string
	Width, Height, BPP int32
}

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

func New(file string) *Texture {
	t := Texture{0, file, 0, 0, 0}
	gl.GenTextures(1, &t.RendererID)
	gl.BindTexture(gl.TEXTURE_2D, t.RendererID)

	path, err := resourcePath.Texture(file)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	myImage, _, err := image.Decode(f)
	if err != nil {
		log.Fatal("error decoding image: " + err.Error())
	}
	myImageType := fmt.Sprintf("%v", reflect.TypeOf(myImage))
	log.Println(fmt.Sprintf("Image type = %s", myImageType))
	var rgbaPix []uint8
	if reflect.TypeOf(myImage).Name() != "*image.RGBA" {
		newImage := image.NewRGBA(myImage.Bounds())
		draw.Draw(newImage, myImage.Bounds(), myImage, image.Point{}, draw.Over)
		rgbaPix = newImage.Pix
	} else {
		rgbaPix = myImage.(*image.RGBA).Pix
	}

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8,
		int32(myImage.Bounds().Dx()), int32(myImage.Bounds().Dy()), 0,
		gl.RGBA, gl.UNSIGNED_BYTE,
		gl.Ptr(&rgbaPix[0]),
	)

	t.Unbind()
	return &t
}

func (t *Texture) Close() {
	gl.DeleteTextures(1, &t.RendererID)
}

func (t *Texture) Bind(slot int32) {
	gl.ActiveTexture(uint32(gl.TEXTURE0 + slot))
	gl.BindTexture(gl.TEXTURE_2D, t.RendererID)
}

func (t *Texture) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (t *Texture) GetWidth() int32  { return t.Width }
func (t *Texture) GetHeight() int32 { return t.Height }
