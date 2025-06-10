package processors_test

import (
	"image"
	"image/color"
	"image/jpeg"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DMaryanskiy/tqs-golang/consumer/internal/processors"
)

func TestResizeImage(t *testing.T) {
	// Set up a mock HTTP server that serves a simple JPEG image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Set(x, y, color.RGBA{255, 0, 0, 255}) // Red
		}
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/jpeg")
		jpeg.Encode(w, img, nil)
	}))
	defer server.Close()

	// Call the function to test
	processors.ResizeImage(server.URL, 50, 50)

	// Verify that the output file exists
	_, err := os.Stat("resized_output.jpg")
	if os.IsNotExist(err) {
		t.Fatal("resized_output.jpg was not created")
	}

	// Optionally, decode and check size
	file, err := os.Open("resized_output.jpg")
	if err != nil {
		t.Fatalf("Failed to open output image: %v", err)
	}
	defer file.Close()

	resizedImg, _, err := image.Decode(file)
	if err != nil {
		t.Fatalf("Failed to decode output image: %v", err)
	}

	bounds := resizedImg.Bounds()
	if bounds.Dx() != 50 || bounds.Dy() != 50 {
		t.Errorf("Expected resized dimensions 50x50, got %dx%d", bounds.Dx(), bounds.Dy())
	}

	// Clean up
	_ = os.Remove("resized_output.jpg")
}
