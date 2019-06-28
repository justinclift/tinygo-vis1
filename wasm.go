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
	for i, j := range itemCounts {
		println(i + ": " + strconv.FormatInt(int64(j), 10))
	}

	// TODO: Sort the categories in some useful way

	// TODO: Draw simple bar graph using the category data

	// TODO: Put labels on the bar graph


	border := float64(2)
	gap := float64(2)
	left := border + gap
	top := border + gap
	displayWidth := float64(width) - border - 1
	displayHeight := float64(height) - border - 1
	// centerX := displayWidth / 2
	// centerY := displayHeight / 2

	// Clear the background
	ctx.Set("fillStyle", "white")
	ctx.Call("fillRect", 0, 0, width, height)

	// Draw grid lines
	step := math.Min(float64(width), float64(height)) / float64(30)
	ctx.Set("strokeStyle", "rgb(220, 220, 220)")
	for i := left; i < displayWidth-step; i += step {
		// Vertical dashed lines
		ctx.Call("beginPath")
		ctx.Call("moveTo", i+step, top)
		ctx.Call("lineTo", i+step, displayHeight)
		ctx.Call("stroke")
	}
	for i := top; i < displayHeight-step; i += step {
		// Horizontal dashed lines
		ctx.Call("beginPath")
		ctx.Call("moveTo", left, i+step)
		ctx.Call("lineTo", displayWidth-border, i+step)
		ctx.Call("stroke")
	}

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
