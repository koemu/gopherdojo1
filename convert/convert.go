package convert

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// CreateImageFile is convert image file
func CreateImageFile(from string, to string, toType string) error {
	fromFile, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	fromImage, _, err := image.Decode(fromFile)
	if err != nil {
		return err
	}

	toFile, err := os.Create(to)
	if err != nil {
		return err
	}
	defer toFile.Close()

	switch toType {
	case "png":
		png.Encode(toFile, fromImage)
	case "jpeg", "jpg":
		var options jpeg.Options
		options.Quality = 100
		jpeg.Encode(toFile, fromImage, &options)
	case "gif":
		var options gif.Options
		options.NumColors = 256
		gif.Encode(toFile, fromImage, &options)
	default:
		return errors.New("Unknown File Format")
	}

	return nil
}
