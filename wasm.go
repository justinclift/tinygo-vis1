package main

import (
	"strconv"
	"syscall/js"
)

var (
	canvasEl, doc js.Value
	// canvasEl, ctx, doc js.Value
)

func main() {
	width := js.Global().Get("innerWidth").Int()
	height := js.Global().Get("innerHeight").Int()
	doc = js.Global().Get("document")
	canvasEl = doc.Call("getElementById", "mycanvas")
	canvasEl.Call("setAttribute", "width", width)
	canvasEl.Call("setAttribute", "height", height)
	canvasEl.Set("tabIndex", 0) // Not sure if this is needed
	// ctx = canvasEl.Call("getContext", "2d")

	db := js.Global().Get("$scope").Get("db")
	rowCount := db.Get("RowCount")
	println("Row count: " + rowCount.String())
	colNames := db.Get("ColNames")
	println("ColNames: " + colNames.String())
	bar := colNames.Length()
	println("# of values: " + strconv.FormatInt(int64(bar), 10))
	for i, n := 0, colNames.Length(); i < n; i++  {
		println("i: " + strconv.FormatInt(int64(i), 10) + " '" + colNames.Index(i).String() + "'")
	}
}
