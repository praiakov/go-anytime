// O comando jpeg lê uma imagem PNG da entrada-padrão
// e a escreve como uma imagem JPEG na saída-padrão
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" //registra o decodificador de PNG
	"io"
	"os"
)

func main() {
	if err := toJpeg(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJpeg(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input fomart=", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
