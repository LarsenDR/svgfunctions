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

//ClientData is a structure to hold point to plot in graph space.
type ClientPoint struct {
	X int `json:"x"`
	Y int `json:"y"`
	R int `json:"r"`
}

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
func SVGStart(width, height int) string {
	var startstr string
	startstr = fmt.Sprintf("<svg viewbox=\"0 0 1000 1000\"  width=\"%d\" height=\"%d\"  xmlns=\"http://www.w3.org/2000/svg\">", width, height)
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
func SVGRect(fill, stroke string, x, y, width, height int) string {
	var rectstr string
	rectstr = fmt.Sprintf("<rect fill=\"%s\" stroke=\"%s\" x=\"%d\" y=\"%d\" width=\"%d\" height=\"%d\" />",
		fill, stroke, x, y, width, height)
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
func SVGPoint(fill, stroke string, strokeWidth, x, y, r int) string {
	var pointstr string
	// pointstr = fmt.Sprintf("<g id=\"point\">\n<circle  stroke=\"%s\" stroke-width=\"%d\" fill=\"%s\" cx=\"%d\" cy=\"%d\" r=\"%d\" />\n</g>",

	pointstr = fmt.Sprintf("<circle  stroke=\"%s\" stroke-width=\"%d\" fill=\"%s\" cx=\"%d\" cy=\"%d\" r=\"%d\" />",
		stroke, strokeWidth, fill, x, y, r)
	return pointstr
}

//SVGText is a function to write a svg text commant
func SVGText(x, y, fontSize, rotdeg int, fontFamily, textAnchor, label string) string {
	var textstr string
	textstr = fmt.Sprintf("<text x=\"%d\" y=\"%d\" font-family=\"%s\" font-size=\"%dpx\" text-anchor=\"%s\" alignment-baseline=\"middle\" transform=\"rotate(%d %d %d)\"> %s </text>",
		x, y, fontFamily, fontSize, textAnchor, rotdeg, x, y, label)
	log.Printf("%s/n", textstr)
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

// LeftMargin is a set of functions to build the svg commands to draw the labels, scales and tics for the margin

//DrawMarginTesting is a function
func DrawMarginTesting(fill, stroke string, width, size int) (str string) {
	//If testing this draws the margin box
	str = SVGRect(fill, stroke, 0, 0, width, size)
	return str
}

//DrawMarginAxis is a function
//func DrawMarginAxis(axis, stroke, fontFamily string, axisStrokeWidth, ticstrokeWidth, ticunit, ticsize, ticfontpx, size, length int) (str string) {
// func DrawMarginAxis(dt Data, length int) (str string) {
// 	str = ""
// 	//axis bar  (margin units)
// 	if dt.Margs.Left.Side == "Left" {
// 		var mrg = dt.Margs.Left
// 		//draw axis line for margin
// 		str = SVGLine(mrg.Stroke, mrg.StrokeWt, mrg.Size, 0, mrg.Size, length)
// 		for i := 0; i <= length; i += mrg.Ticunit {
// 			//axis tics (margin units)
// 			str += SVGLine(mrg.Stroke, mrg.StrokeWt, (mrg.Size - mrg.Ticsize), length-i, mrg.Size, length-i)
// 			// 	//axis tic labels (margin units)
// 			str += SVGTextTicLabel(mrg.Size-(mrg.Ticsize+10), length-i, dt.FontFamily, mrg.Ticfontpx, i)
// 		}
// 	} else if dt.Margs.Bott.Side == "Bottom" {
// 		var mrg = dt.Margs.Bott
// 		//draw axis line for margin
// 		str = SVGLine(mrg.Stroke, mrg.StrokeWt, 0, 0, 0, length)
// 		for i := 0; i <= length; i += mrg.Ticunit {
// 			//axis tics (margin units)
// 			str += SVGLine(mrg.Stroke, mrg.StrokeWt, (mrg.Size - mrg.Ticsize), length-i, mrg.Size, length-i)
// 			// 	//axis tic labels (margin units)
// 			str += SVGTextTicLabel(mrg.Size-(mrg.Ticsize+10), length-i, dt.FontFamily, mrg.Ticfontpx, i)
// 		}
// 	}
// 	return str
// }

//DrawMarginLabel is a function
func DrawMarginLabel() {
	// axis side label (margin units)
}

// BottomMargin is a set of functions to build the svg commands to draw the labels, scales and tics for the margin

// RightMargin is a set of functions to build the svg commands to draw the labels, scales and tics for the margin

// TopMargin is a set of functions to build the svg commands to draw the labels, scales and tics for the margin

// Graph is a set of functions to build the svg parts of the graph space
//graph space grid (graph units)
//points in graph space (graph units)
