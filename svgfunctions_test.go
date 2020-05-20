package svgfunctions

import (
	"fmt"
	"testing"
)

//XMLStart_Test the the XMLStart function
func TestXMLStart(t *testing.T) {
	desired := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>`
	str := XMLStart()
	fmt.Printf("XMLStart test: %v\n", str)
	if str != desired {
		t.Errorf("got %v   wanted %v", str, desired)
	}
}

func TestSVGStart(t *testing.T) {
	var scdt = []struct {
		Label   string
		Xorg    int
		Yorg    int
		Width   int
		Height  int
		desired string
	}{
		{"t1", 0, 0, 1000, 1000, `<svg viewbox="0 0 1000 1000"  width="1000" height="1000"  xmlns="http://www.w3.org/2000/svg">`},
	}
	for _, tt := range scdt {
		str := SVGStart(tt.Width, tt.Height)
		fmt.Printf("SVGStart test: %v\n", str)
		if str != tt.desired {
			t.Errorf("got %v   wanted %v", str, tt.desired)

		}
	}
}

func TestSVGRect(t *testing.T) {
	var line = []struct {
		Label    string
		fill     string
		stroke   string
		strokewt int
		xorg     int
		yorg     int
		width    int
		height   int
		desired  string
	}{
		{"t1", "white", "black", 3, 0, 0, 500, 500, `<rect fill="white" stroke="black" stroke-width="3" x="0" y="0" width="500" height="500" />`},
	}
	for _, tt := range line {
		str := SVGRect(tt.fill, tt.stroke, tt.strokewt, tt.xorg, tt.yorg, tt.width, tt.height)
		fmt.Printf("SVGRect test: %v\n", str)
		if str != tt.desired {
			t.Errorf("got %v   wanted %v", str, tt.desired)

		}
	}
}

func TestSVGLine(t *testing.T) {
	var line = []struct {
		Label    string
		stroke   string
		strokewt int
		x1       int
		y1       int
		x2       int
		y2       int
		desired  string
	}{
		{"t1", "black", 3, 0, 0, 500, 500, `<line style="stroke:black;stroke-width:3" x1="0" y1="0" x2="500" y2="500" />`},
		{"t2", "#e00", 2, 500, 0, 0, 500, `<line style="stroke:#e00;stroke-width:2" x1="500" y1="0" x2="0" y2="500" />`},
	}
	for _, tt := range line {
		str := SVGLine(tt.stroke, tt.strokewt, tt.x1, tt.y1, tt.x2, tt.y2)
		fmt.Printf("SVGLine test: %v\n", str)
		if str != tt.desired {
			t.Errorf("got %v   wanted %v", str, tt.desired)

		}
	}
}

func TestSVGPoint(t *testing.T) {
	var line = []struct {
		Label    string
		fill     string
		stroke   string
		strokewt int
		x        int
		y        int
		r        int
		desired  string
	}{
		{"t1", "red", "black", 2, 400, 350, 10, `<circle  stroke="black" stroke-width="2" fill="red" cx="400" cy="350" r="10" />`},
		{"t1", "orange", "black", 2, 100, 350, 10, `<circle  stroke="black" stroke-width="2" fill="orange" cx="100" cy="350" r="10" />`},
		{"t1", "blue", "black", 2, 200, 450, 10, `<circle  stroke="black" stroke-width="2" fill="blue" cx="200" cy="450" r="10" />`},
	}
	for _, tt := range line {
		str := SVGPoint(tt.fill, tt.stroke, tt.strokewt, tt.x, tt.y, tt.r)
		fmt.Printf("SVGPoint test: %v\n", str)
		if str != tt.desired {
			t.Errorf("got %v   wanted %v", str, tt.desired)

		}
	}
}

func TestSVGEnd(t *testing.T) {
	desired := `</svg>`
	str := SVGEnd()
	fmt.Printf("SVGEnd test: %v\n", str)
	if str != desired {
		t.Errorf("got %v   wanted %v", str, desired)
	}
}
