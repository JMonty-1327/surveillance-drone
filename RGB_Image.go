package main

 import (
         "fmt"
         "image"
         "image/jpeg"
         "os"
 )

 func init() {
         // damn important or else At(), Bounds() functions will
         // caused memory pointer error!!
         image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
 }

 func main() {
         imgfile, err := os.Open("./elephant.jpg")

         if err != nil {
                 fmt.Println("elephant.jpg file not found!")
                 os.Exit(1)
         }

         defer imgfile.Close()

         // get image height and width with image/jpeg
         // change accordinly if file is png or gif

         imgCfg, _, err := image.DecodeConfig(imgfile)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         width := imgCfg.Width
         height := imgCfg.Height

         fmt.Println("Width : ", width)
         fmt.Println("Height : ", height)

         // we need to reset the io.Reader again for image.Decode() function below to work
         // otherwise we will  - panic: runtime error: invalid memory address or nil pointer dereference
         // there is no build in rewind for io.Reader, use Seek(0,0)
         imgfile.Seek(0, 0)

         // get the image
         img, _, err := image.Decode(imgfile)

         fmt.Println(img.At(10, 10).RGBA())
         for y := 0; y < height; y++ {
                 for x := 0; x < width; x++ {
                         r, g, b, a := img.At(x, y).RGBA()
						 
						 rr := uint8(r >> 8)
						 gg := uint8(g >> 8)
						 bb := uint8(b >> 8)
						 aa := uint8(a >> 8)
						 fmt.Printf("[X : %d Y : %v] R : %v, G : %v, B : %v, A : %v  \n", x, y, rr, gg, bb, aa)
                         //fmt.Printf("[X : %d Y : %v] R : %v, G : %v, B : %v, A : %v  \n", x, y, r, g, b, a)
                 }
         }

 }