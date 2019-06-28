package main

import (
	"math"
	"strconv"
	"syscall/js"
)

var (
	canvasEl, ctx, doc js.Value
)

func main() {
	width := js.Global().Get("innerWidth").Int()
	height := js.Global().Get("innerHeight").Int()
	doc = js.Global().Get("document")
	canvasEl = doc.Call("getElementById", "mycanvas")
	canvasEl.Call("setAttribute", "width", width)
	canvasEl.Call("setAttribute", "height", height)
	canvasEl.Set("tabIndex", 0) // Not sure if this is needed
	ctx = canvasEl.Call("getContext", "2d")

	db := js.Global().Get("$scope").Get("db")
	rows := db.Get("Records")
	numRows := rows.Length()

	// Count the number of items for each category
	highestVal := 0
	itemCounts := make(map[string]int)
	var row js.Value
	for i, n := 0, numRows; i < n; i++ {
		row = rows.Index(i)
		catName := row.Index(10).Get("Value").String()
		itemCount, err := strconv.Atoi(row.Index(12).Get("Value").String())
		if err != nil {
			println(err)
		}
		c := itemCounts[catName]
		itemCounts[catName] = c + itemCount
	}

	// Determine the highest count value, so we can automatically size the graph to fit
	for _, itemCount := range itemCounts {
		if itemCount > highestVal {
			highestVal = itemCount
		}
	}

	// TODO: Sort the categories in some useful way

	border := 2
	gap := 2
	left := border + gap
	top := border + gap
	displayWidth := width - border - 1
	displayHeight := height - border - 1
	// centerX := displayWidth / 2
	// centerY := displayHeight / 2

	// Clear the background
	ctx.Set("fillStyle", "white")
	ctx.Call("fillRect", 0, 0, width, height)

	// Draw grid lines
	step := math.Min(float64(width), float64(height)) / float64(30)
	ctx.Set("strokeStyle", "rgb(220, 220, 220)")
	for i := float64(left); i < float64(displayWidth)-step; i += step {
		// Vertical dashed lines
		ctx.Call("beginPath")
		ctx.Call("moveTo", i+step, top)
		ctx.Call("lineTo", i+step, displayHeight)
		ctx.Call("stroke")
	}
	for i := float64(top); i < float64(displayHeight)-step; i += step {
		// Horizontal dashed lines
		ctx.Call("beginPath")
		ctx.Call("moveTo", left, i+step)
		ctx.Call("lineTo", displayWidth-border, i+step)
		ctx.Call("stroke")
	}

	// Determine the vertical size of the graph, and center it
	// println("top: " + strconv.FormatInt(int64(top), 10))
	// println("highestVal: " + strconv.FormatInt(int64(highestVal), 10))
	unitSize := 3
	// println("unitSize: " + strconv.FormatInt(int64(unitSize), 10))
	vertSize := highestVal * unitSize
	// println("vertSize: " + strconv.FormatInt(int64(vertSize), 10))
	baseLine := displayHeight - ((displayHeight - vertSize) / 2)
	// println("baseLine: " + strconv.FormatInt(int64(baseLine), 10))
	// println("displayHeight: " + strconv.FormatInt(int64(displayHeight), 10))

	// TODO: Determine a useful colour scheme


	// Determine the horizontal size of the graph, and center it
	barGap := 20
	barWidth := 30
	numBars := len(itemCounts)
	horizSize := (numBars * barWidth) + ((numBars - 1) * barGap)
	barLeft := (displayWidth - horizSize) / 2

	// Draw simple bar graph using the category data
	ctx.Set("strokeStyle", "black")
	for _, num := range itemCounts {
	// for label, num := range itemCounts {
		barHeight := num * unitSize
		ctx.Set("fillStyle", "blue")
		ctx.Call("beginPath")
		ctx.Call("moveTo", barLeft, baseLine)
		ctx.Call("lineTo", barLeft + barWidth, baseLine)
		// println("height: " + strconv.FormatInt(int64(height), 10))
		ctx.Call("lineTo", barLeft + barWidth, baseLine - barHeight)
		ctx.Call("lineTo", barLeft, baseLine - barHeight)
		ctx.Call("closePath")
		ctx.Call("fill")
		ctx.Call("stroke")
		barLeft += barGap + barWidth
	}

	// TODO: Put labels on the bar graph

	// Draw a border around the graph area
	ctx.Set("lineWidth", "2")
	ctx.Set("strokeStyle", "white")
	ctx.Call("beginPath")
	ctx.Call("moveTo", 0, 0)
	ctx.Call("lineTo", width, 0)
	ctx.Call("lineTo", width, height)
	ctx.Call("lineTo", 0, height)
	ctx.Call("closePath")
	ctx.Call("stroke")
	ctx.Set("lineWidth", "2")
	ctx.Set("strokeStyle", "black")
	ctx.Call("beginPath")
	ctx.Call("moveTo", border, border)
	ctx.Call("lineTo", displayWidth, border)
	ctx.Call("lineTo", displayWidth, displayHeight)
	ctx.Call("lineTo", border, displayHeight)
	ctx.Call("closePath")
	ctx.Call("stroke")
}
