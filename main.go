package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	if len(os.Args) != 7 {
		fmt.Println("Usage: PutSquaresOnPicture <input_image_path> <output_image_path> <x> <y> <color> <size>")
		os.Exit(1)
	}

	inputImagePath := os.Args[1]
	outputImagePath := os.Args[2]
	x, y, size, err := validateInputs(os.Args[3], os.Args[4], os.Args[6])
	if err != nil {
		fmt.Printf("Error validating inputs: %v\n", err)
		os.Exit(1)
	}

	colorRGBA := parseColor(os.Args[5])

	err = processImage(inputImagePath, outputImagePath, x, y, colorRGBA, size)
	if err != nil {
		fmt.Printf("Error processing image: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Image with colored square saved to %s\n", outputImagePath)
}


func processImage(inputPath, outputPath string, x, y int, color_rgba color.RGBA, size int) error {
	// Read the input image
	img, err := readImage(inputPath)
	if err != nil {
		return fmt.Errorf("error reading input image: %v", err)
	}

	// Create a new draw.Image from the existing image
	drawImg := image.NewRGBA(img.Bounds())
	draw.Draw(drawImg, drawImg.Bounds(), img, image.Point{}, draw.Over)

	// Draw the colored square on the image
	drawFilledCircle(drawImg, x, y, size, color_rgba)

	// Save the modified image to the output path
	err = saveImage(outputPath, drawImg)
	if err != nil {
		return fmt.Errorf("error saving output image: %v", err)
	}

	return nil
}

func validateInputs(xStr, yStr, sizeStr string) (x, y, size int, err error) {
	// Validate x and y coordinates
	x, err = validatePositiveInt(xStr)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid x coordinate: %v", err)
	}

	y, err = validatePositiveInt(yStr)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid y coordinate: %v", err)
	}

	// Validate square size
	size, err = validatePositiveInt(sizeStr)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid square size: %v", err)
	}

	return x, y, size, nil
}

func validatePositiveInt(str string) (int, error) {
	value := parseInt(str)

	if value <= 0 {
		return 0, fmt.Errorf("value must be a positive number")
	}

	return value, nil
}

func parseColor(colorStr string) color.RGBA {
	var r, g, b, a uint8
	fmt.Sscanf(colorStr, "%02x%02x%02x%02x", &r, &g, &b, &a)
	fmt.Printf("%d, %d, %d, %d \n", r,g,b,a)
	return color.RGBA{r, g, b, a}
}

func parseInt(valStr string) int {
	var val int
	fmt.Sscanf(valStr, "%d", &val)
	return val
}

func readImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func drawSquare(img draw.Image, x, y int, color_rgba color.RGBA, size int) {
	draw.Draw(
		img,
		image.Rect(x, y, x+size, y+size),
		&image.Uniform{color_rgba},
		image.Point{},
		draw.Over)
}

func drawFilledCircle(img draw.Image, x0, y0, r int, c color.RGBA) {
	for r >= 0 {
		drawFilledCircle_AngryWayForFilling(img, x0, y0, r, c)
		r-=1;
	}
}

func drawFilledCircle_AngryWayForFilling(img draw.Image, x0, y0, r int, c color.RGBA) {
    x, y, dx, dy := r-1, 0, 1, 1
    err := dx - (r * 2)

    for x > y {
        img.Set(x0+x, y0+y, c)
        img.Set(x0+y, y0+x, c)
        img.Set(x0-y, y0+x, c)
        img.Set(x0-x, y0+y, c)
        img.Set(x0-x, y0-y, c)
        img.Set(x0-y, y0-x, c)
        img.Set(x0+y, y0-x, c)
        img.Set(x0+x, y0-y, c)

		img.Set(x0+(x-1), y0+y, c)
        img.Set(x0+(y-1), y0+x, c)
        img.Set(x0-(y-1), y0+x, c)
        img.Set(x0-(x-1), y0+y, c)
        img.Set(x0-(x-1), y0-y, c)
        img.Set(x0-(y-1), y0-x, c)
        img.Set(x0+(y-1), y0-x, c)
        img.Set(x0+(x-1), y0-y, c)

		img.Set(x0+x, y0+(y-1), c)
        img.Set(x0+y, y0+(x-1), c)
        img.Set(x0-y, y0+(x-1), c)
        img.Set(x0-x, y0+(y-1), c)
        img.Set(x0-x, y0-(y-1), c)
        img.Set(x0-y, y0-(x-1), c)
        img.Set(x0+y, y0-(x-1), c)
        img.Set(x0+x, y0-(y-1), c)
		
		img.Set(x0+(x-2), y0+y, c)
        img.Set(x0+(y-2), y0+x, c)
        img.Set(x0-(y-2), y0+x, c)
        img.Set(x0-(x-2), y0+y, c)
        img.Set(x0-(x-2), y0-y, c)
        img.Set(x0-(y-2), y0-x, c)
        img.Set(x0+(y-2), y0-x, c)
        img.Set(x0+(x-2), y0-y, c)

        if err <= 0 {
            y++
            err += dy
            dy += 2
        }
        if err > 0 {
            x--
            dx += 2
            err += dx - (r * 2)
        }
    }
}

func saveImage(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}

	return nil
}
