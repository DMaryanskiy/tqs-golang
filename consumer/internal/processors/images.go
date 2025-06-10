package processors

import (
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/nfnt/resize"
)

func ResizeImage(imageURL string, width, height int) {
	resp, err := http.Get(imageURL)
	if err != nil {
		log.Fatalf("Failed to fetch image: %v", err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatalf("Failed to decode image: %v", err)
	}

	resizedImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	out, err := os.Create("resized_output.jpg")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer out.Close()

	err = jpeg.Encode(out, resizedImg, nil)
	if err != nil {
		log.Fatalf("Failed to encode resized image: %v", err)
	}

	log.Println("Image resized and saved as resized_output.jpg")
}
