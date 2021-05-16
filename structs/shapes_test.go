package main

import "testing"

type Shape interface {
	Area() float64
	Perimeter() float64
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name            string
		shape           Shape
		wantedPerimeter float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, wantedPerimeter: 40.0},
		{name: "circle", shape: Circle{Radius: 10.0}, wantedPerimeter: 62.83185307179586},
		{name: "triangle not supported", shape: Triangle{Base: 0.0, Height: 0.0}, wantedPerimeter: -1.0},
	}

	for _, test := range perimeterTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Perimeter()
			if test.wantedPerimeter != got {
				t.Errorf("wanted %g but got %g given %#v", test.wantedPerimeter, got, test.shape)
			}
		})
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name       string
		shape      Shape
		wantedArea float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, wantedArea: 72.0},
		{name: "circle", shape: Circle{Radius: 10.0}, wantedArea: 314.1592653589793},
		{name: "triangle", shape: Triangle{Base: 12.0, Height: 6.0}, wantedArea: 36.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if test.wantedArea != got {
				t.Errorf("wanted %g but got %g given %#v", test.wantedArea, got, test.shape)
			}
		})
	}
}
