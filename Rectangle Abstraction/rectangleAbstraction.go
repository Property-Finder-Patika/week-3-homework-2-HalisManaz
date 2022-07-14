package main

import "fmt"

// initializing rectangle struct
type rectangle struct {
	width  float64
	height float64
}

func main() {

	r2, errR2 := NewRectangle(-2, 2) // Create invalid rectangle
	if errR2 != nil {
		fmt.Println(errR2) // Output: width and height value should be greater than zero. Your inputs --> width: -2, height: 2
	} else {
		fmt.Println(r2)
	}

	r, errR := NewRectangle(2, 2) // Create valid rectangle
	if errR != nil {
		fmt.Println(errR) // Output: width and height value should be greater than zero. Your inputs --> width: -2, height: 2
	} else {
		area := Area(r)
		circumference := Circumference(r)
		fmt.Printf("Width: %vm\nHeight: %vm\nArea: %vm2\nCircumference: %vm", r.width, r.height, area, circumference)
		// Output:
		//Height: 2m
		//Width: 2m
		//Area: 4m2
		//Circumference: 8m
	}
}

// NewRectangle functions create new rectangle struct. Also, it checks invalid input values such as negative width or height values.
func NewRectangle(w, h float64) (*rectangle, error) {
	if w <= 0 || h <= 0 {
		return nil, fmt.Errorf("widht and height value should be greater than zero. Your inputs --> width: %v, height: %v", w, h)
	} else {
		return &rectangle{width: w, height: h}, nil
	}
}

// Area function calculate rectangle area.
func Area(r *rectangle) float64 {
	return r.height * r.width
}

// Circumference function calculate rectangle circumference.
func Circumference(r *rectangle) float64 {
	return 2 * (r.height + r.width)
}
