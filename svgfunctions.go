package svgfunctions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

//Metadata structure is nested in to several structure

//ScreenData is one of base metadata structures
type ScreenData struct {
	Label  string
	Xorg   int
	Yorg   int
	Width  int
	Height int
}

//GraphData if the dat about the graph part of the
type GraphData struct {
	Label        string
	Space        string
	XaxisLabel   string
	XaxisUnitmax int
	XaxisUnitmin int
	YaxisLabel   string
	YaxisUnitmax int
	YaxisUnitmin int
	Grid         bool
	GridUnit     int
	GridColor    string
}

//MarginsData are a set of four side Margin structures
type MarginsData struct {
	Left  Margin
	Bott  Margin
	Right Margin
	Top   Margin
}

//Margin is a meta data structure about on side of the graph.
type Margin struct {
	Side          string
	Size          int
	AxisLine      bool
	Stroke        string
	StrokeWt      int
	Ticunit       int
	Ticsize       int
	Ticstroke     int
	Ticfontpx     int
	Ticfontoffset int
	Labelpx       int
	Labeltext     string
}

//Data is the graph general data structure
type Data struct {
	Testing                bool
	TestingBackgroundColor string
	TestingStrokeColor     string
	BackgroundColor        string
	StrokeColor            string
	FontFamily             string
	Screen                 ScreenData
	Graph                  GraphData
	Margs                  MarginsData
}

//ClientPoint is a structure to hold point to plot in graph space.
type ClientPoint struct {
	X int `json:"x"`
	Y int `json:"y"`
	R int `json:"r"`
}

//ClientPoints is a structure containing an array of ClientPoint Structures
type ClientPoints struct {
	DataVals []ClientPoint
}

//GetLayout is a function to read the graph.json settings
func GetLayout(filename string) (Data, error) {
	var data Data
	fmt.Printf("Layout: %#v\n", filename)

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	dec := json.NewDecoder(strings.NewReader(string(file)))
	err = dec.Decode(&data)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}

	// log.Printf("%#v\n", data)

	return data, err
}

//GetClientData is a function to read the datapoints.json into ClientData.
func GetClientData(filename string) (ClientPoints, error) {
	var cldt ClientPoints
	fmt.Printf("Client: Filename: %#v\n", filename)

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("Client: String: %#v\n", string(file))

	json.Unmarshal(file, &cldt)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	log.Printf("Data: %#v\n", cldt)
	return cldt, nil
}

// XMLStart is a start svg code need by SVG files
func XMLStart() string {
	str := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>`
	return str
}

// SVGEnd is a line to end a sequence of SVG commands
func SVGEnd() string {
	str := `</svg>`
	return str
}

//SVGStart is a function to write a SVG starting text for a graph.
func SVGStart(id string, width, height int) string {
	var startstr string
	startstr = fmt.Sprintf("<svg id=\"%s\" viewbox=\"0 0 1000 1000\"  width=\"%d\" height=\"%d\"  xmlns=\"http://www.w3.org/2000/svg\">", id, width, height)
	return startstr
}

//SVGGrid is a function to put a grid in the graph window
func SVGGrid(stroke string, x, y, width, height, xgrid, ygrid int) string {
	fmt.Printf("x:%v, y:%v, max x:%v, max y:%v, xgrid:%v, ygrid:%v\n", x, y, width, height, xgrid, ygrid)
	var gridstr string
	xgridunit := width / xgrid
	ygridunit := height / ygrid

	scale := float64(1)
	gridstr = fmt.Sprintf("<g transform=\"translate(%d %d) scale(%.3f %.3f)\">\n", x, y, scale, scale)

	// if d.Graph.Grid {
	for i := 0; i <= height; i += xgridunit {
		gridstr += fmt.Sprintf("%s\n", SVGLine(stroke, 1, 0, i, width, i))
	}
	for i := 0; i <= width; i += ygridunit {
		gridstr += fmt.Sprintf("%s\n", SVGLine(stroke, 1, i, 0, i, height))
	}

	gridstr += fmt.Sprintf(" %s\n", `</g>`)
	return gridstr
}

//SVGRect is that function to write a rectangle text for a graph
func SVGRect(id, fill, stroke string, strokeWt, x, y, width, height int) string {
	var rectstr string
	rectstr = fmt.Sprintf("<rect id=\"%s\" fill=\"%s\" stroke=\"%s\" stroke-width=\"%d\" x=\"%d\" y=\"%d\" width=\"%d\" height=\"%d\" />",
		id, fill, stroke, strokeWt, x, y, width, height)
	return rectstr
}

//SVGLine is a function to write a line test for a graph
func SVGLine(stroke string, strokewidth, x1, y1, x2, y2 int) string {
	var linestr string
	linestr = fmt.Sprintf("<line style=\"stroke:%s;stroke-width:%d\" x1=\"%d\" y1=\"%d\" x2=\"%d\" y2=\"%d\" />",
		stroke, strokewidth, x1, y1, x2, y2)
	return linestr
}

//SVGPoint is a function to write a point text for a graph
func SVGPoint(id, fill, stroke string, strokeWidth, x, y, r int) string {
	var pointstr string
	// pointstr = fmt.Sprintf("<g id=\"point\">\n<circle  stroke=\"%s\" stroke-width=\"%d\" fill=\"%s\" cx=\"%d\" cy=\"%d\" r=\"%d\" />\n</g>",
	pointstr = fmt.Sprintf("<circle id=\"%s\" stroke=\"%s\" stroke-width=\"%d\" fill=\"%s\" cx=\"%d\" cy=\"%d\" r=\"%d\" />",
		id, stroke, strokeWidth, fill, x, y, r)
	return pointstr
}

//SVGPath is a function to write a svg path comment
func SVGPath(id, fill, stroke string, strokewt, mx, my, c1x, c1y, c2x, c2y, c3x, c3y int) string {
	var pathstr string
	pathstr = fmt.Sprintf(
		"<path d=\"M %d %d C %d %d, %d %d, %d %d\" id=\"%s\" fill=\"%s\" stroke=\"%s\" stroke-width=\"%d\" />",
		mx, my, c1x, c1y, c2x, c2y, c3x, c3y, id, fill, stroke, strokewt)
	return pathstr
}

//SVGText is a function to write a svg text comment
func SVGText(x, y, fontSize, rotdeg int, fontFamily, textAnchor, label string) string {
	var textstr string
	textstr = fmt.Sprintf("<text x=\"%d\" y=\"%d\" font-family=\"%s\" font-size=\"%dpx\" text-anchor=\"%s\" alignment-baseline=\"middle\" transform=\"rotate(%d %d %d)\"> %s </text>",
		x, y, fontFamily, fontSize, textAnchor, rotdeg, x, y, label)
	return textstr
}

//SVGTextTicLabel is a function to make axis tic labels
// side it one of the name for the four margins.
func SVGTextTicLabel(x, y int, side, fontfamily string, fontsize, idx int) string {
	var textstr string
	if side == "Left" {
		textstr = fmt.Sprintf(
			"<text x=\"%d\" y=\"%d\" font-family=\"%s\" font-size=\"%dpx\" text-anchor=\"end\" alignment-baseline=\"middle\" > %d </text>",
			x, y, fontfamily, fontsize, idx)
	} else if side == "Bott" {
		textstr = fmt.Sprintf(
			"<text x=\"%d\" y=\"%d\" font-family=\"%s\" font-size=\"%dpx\" text-anchor=\"middle\" alignment-baseline=\"middle\" > %d </text>",
			x, y, fontfamily, fontsize, idx)
	} else if side == "Right" {
		textstr = fmt.Sprintf(
			"<text x=\"%d\" y=\"%d\" font-family=\"%s\" font-size=\"%dpx\" text-anchor=\"start\" alignment-baseline=\"middle\" > %d </text>",
			x, y, fontfamily, fontsize, idx)
	} else if side == "Top" {
		textstr = fmt.Sprintf(
			"<text x=\"%d\" y=\"%d\" font-family=\"%s\" font-size=\"%dpx\" text-anchor=\"middle\" alignment-baseline=\"middle\" > %d </text>",
			x, y, fontfamily, fontsize, idx)
	} else {

		log.Printf("SVGTextTicLabel: side == %s is unknown!", side)
	}
	//log.Printf("%s/n", textstr)
	return textstr
}

//ScaleGraphToMath is a function to convert scale values from graph to math scales
func ScaleGraphToMath(gval, gmin, gmax, mmin, mmax int64) (mval int64) {
	var grng, mrng float64

	grng = float64(gmax) - float64(gmin)
	mrng = float64(mmax) - float64(mmin)

	mval = int64((float64(gval) / grng) * mrng)

	return mval
}

//ScaleMathToGraph is a function to convert scale values from math to graph scales
func ScaleMathToGraph(mval, mmin, mmax, gmin, gmax int64) (gval int64) {
	var grng, mrng float64

	grng = float64(gmax) - float64(gmin)
	mrng = float64(mmax) - float64(mmin)

	fmt.Printf("mval %#v, mrng %#v grng %#v\n", mval, mrng, grng)
	gval = int64((float64(mval) / mrng) * grng)

	return gval
}
