package main

import (
    "image"
    "image/color"
    "image/png"
    "reflect"
    "math"
    // "math/rand"
    "os"
    "fmt"
    "github.com/skratchdot/open-golang/open"
    "github.com/ojrac/opensimplex-go"
    // "image/gif"

)
var (

	colorBlue = color.RGBA{70,130,180,255}
    colorGreen = color.RGBA{34,139,34,255}
    colorYellow = color.RGBA{255,250,205,255}
    colorWhite = color.RGBA{255,255,255,255}
    colorBrown = color.RGBA{139,69,19,255}
    zoom = 0.005 //smaller is more zoomed in
    height = 1000
    width = 1000

)

func main() {
	simp := opensimplex.NewWithSeed(2);
	simp2 := opensimplex.NewWithSeed(1);
    img := image.NewRGBA(image.Rect(0, 0, height, width))
    fmt.Println(reflect.TypeOf(img))

    for i := 0; i < height; i++ {
    	for q := 0; q < width; q++ {
    		height := simp.Eval2(zoom*float64(q),zoom*float64(i+750))
    		height2 := simp2.Eval2(zoom*4.0*float64(q),zoom*4.0*float64(i))
    		finalHeight:=height+math.Abs(height2/3)
    		if(finalHeight > .7){
		        img.Set(q, i, colorWhite)
    		}else if(finalHeight > .5){
				img.Set(q, i, colorBrown)
    		}else if(height > -0.3){
				img.Set(q, i, colorGreen)
    		}else if(height > -.45){
				img.Set(q, i, colorYellow)
    		}else{
				img.Set(q, i, colorBlue)
    		}
    	}
    }

    // Save to out.png
    f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    png.Encode(f, img)
    open.Run("out.png")

}