package main

import (
	"math"
	"math/rand"
	"strconv"
	"syscall/js"
	"time"
)

const goldenRatioConjugate = 0.618033988749895

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

	// Determine the vertical size and center position of the graph
	unitSize := 3
	vertSize := highestVal * unitSize
	baseLine := displayHeight - ((displayHeight - vertSize) / 2)

	// Calculate the bar size, gap, and centering based upon the number of bars
	graphBorder := 50
	numBars := len(itemCounts)
	horizSize := displayWidth - (graphBorder * 2)
	b := float64(horizSize) / float64(numBars)
	barWidth := int(math.Round(b * 0.6))
	barGap := int(b - float64(barWidth))
	barLeft := ((graphBorder * 2) + barGap) / 2

	// Draw simple bar graph using the category data
	textGap := 5
	textSize := 20
	axisThickness := 5
	ctx.Set("strokeStyle", "black")
	ctx.Set("font", "bold "+strconv.FormatInt(int64(textSize), 10)+"px serif")
	rand.Seed(int64(time.Now().Nanosecond()))
	hue := rand.Float64()
	for label, num := range itemCounts {
		barHeight := num * unitSize
		hue += goldenRatioConjugate
		hue = math.Mod(hue, 1)
		ctx.Set("fillStyle", hsvToRgb(hue, 0.5, 0.95))
		ctx.Call("beginPath")
		ctx.Call("moveTo", barLeft, baseLine)
		ctx.Call("lineTo", barLeft+barWidth, baseLine)
		ctx.Call("lineTo", barLeft+barWidth, baseLine-barHeight)
		ctx.Call("lineTo", barLeft, baseLine-barHeight)
		ctx.Call("closePath")
		ctx.Call("fill")
		ctx.Call("stroke")
		ctx.Set("fillStyle", "black")

		// Draw the bar label horizontally centered
		textMet := ctx.Call("measureText", label)
		textWidth := textMet.Get("width").Float()
		textLeft := (float64(barWidth) - textWidth) / 2
		ctx.Call("fillText", label, barLeft+int(textLeft), baseLine+textSize+textGap+axisThickness+textGap)
		barLeft += barGap + barWidth
	}

	// TODO: Draw axis
	barLeft = ((graphBorder * 2) + barGap) / 2
	ctx.Set("lineWidth", axisThickness)
	ctx.Call("beginPath")
	ctx.Call("moveTo", graphBorder+horizSize, baseLine+axisThickness+textGap)
	ctx.Call("lineTo", barLeft-axisThickness-textGap, baseLine+axisThickness+textGap)
	ctx.Call("lineTo", barLeft-axisThickness-textGap, baseLine-int(float64(vertSize)*1.2))
	ctx.Call("stroke")

	// Add title
	title := "Marine Litter Survey - Keep Northern Ireland Beautiful"
	titleFontSize := 25
	titleMet := ctx.Call("measureText", title)
	titleWidth := titleMet.Get("width").Float()
	ctx.Set("font", "bold "+strconv.FormatInt(int64(titleFontSize), 10)+"px serif")
	ctx.Call("fillText", title, (displayWidth-int(titleWidth))/2, top+titleFontSize+20)

	// TODO: Add axis labels

	// TODO: Add units of measurement

	// TODO: Adjust the grid lines to work with the unit of measurement

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

// Ported from the JS here: https://martin.ankerl.com/2009/12/09/how-to-create-random-colors-programmatically/
func hsvToRgb(h, s, v float64) string {
	hi := math.Round(h * 6)
	f := h*6 - hi
	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)

	hiInt := int(hi)
	var r, g, b float64
	if hiInt == 0 {
		r, g, b = v, t, p
	}
	if hiInt == 1 {
		r, g, b = q, v, p
	}
	if hiInt == 2 {
		r, g, b = p, v, t
	}
	if hiInt == 3 {
		r, g, b = p, q, v
	}
	if hiInt == 4 {
		r, g, b = t, p, v
	}
	if hiInt == 5 {
		r, g, b = v, p, q
	}

	red := int(math.Round(r * 256))
	green := int(math.Round(g * 256))
	blue := int(math.Round(b * 256))
	return "rgb(" + strconv.FormatInt(int64(red), 10) + ", " + strconv.FormatInt(int64(green), 10) + ", " + strconv.FormatInt(int64(blue), 10) + ")"
}
