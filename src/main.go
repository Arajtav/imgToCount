package main

import (
    "os"
    "fmt"
    "flag"
    "image"
    "image/color"
    "image/png"
    _ "image/jpeg"
)

func main() {
    flag.Parse();

    if len(flag.Args()) < 1 {
        fmt.Fprintln(os.Stderr, "You need to specify input file");
        os.Exit(1);
    }

    if len(flag.Args()) > 1 {
        fmt.Fprintln(os.Stderr, "Too many arguments");
        os.Exit(1);
    }

    file, err := os.Open(flag.Args()[0]);
    if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to open file");
        panic(err);
    }

    in, _, err := image.Decode(file);
    if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to decode image");
        panic(err);
    }
    file.Close();

    img := image.NewRGBA(image.Rect(0, 0, 256, 256));
    for x := 0; x < 256; x++ {
        for y := 0; y < 256; y++ {
            r, g, b, _ := in.At(int(float64(x)/256.0*float64(in.Bounds().Dx())), int(float64(y)/256.0*float64(in.Bounds().Dy()))).RGBA();
            c := getValue(r, g, b);
            img.Set(x, y, color.RGBA{c, c, c, 255});
        }
    }

    file, err = os.Create("out.png");
    if err != nil {
        panic(err);
    }
    png.Encode(file, img);
    file.Close();
}

func getValue(r uint32, g uint32, b uint32) uint8 {
    return uint8(0.2126 * float32(r)/256 + 0.7152 * float32(g)/256 + 0.0722 * float32(b)/256);
}
