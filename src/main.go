package main

import (
    "os"
    "fmt"
    "flag"
    "image"
    _ "image/png"
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

    file, err = os.Create("out");
    if err != nil { panic(err); }

    for x := 0; x < 256; x++ {
        for y := 0; y < 256; y++ {
            r, g, b, _ := in.At(int(float64(x)/256.0*float64(in.Bounds().Dx())), int(float64(y)/256.0*float64(in.Bounds().Dy()))).RGBA();
            c := getValue(r, g, b);
            for i := uint8(0); i < c; i++ {
                file.Write([]byte{uint8(x), uint8(y)});
            }
        }
    }
    file.Close();
}

func getValue(r uint32, g uint32, b uint32) uint8 {
    return uint8(0.2126 * float32(r)/256 + 0.7152 * float32(g)/256 + 0.0722 * float32(b)/256);
}
