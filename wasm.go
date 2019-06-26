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
	println("Total row count: " + rowCount.String())
	colNames := db.Get("ColNames")
	println("ColNames: " + colNames.String())
	numNames := colNames.Length()
	rows := db.Get("Records")
	numRows := rows.Length()

	println("# of column names: " + strconv.FormatInt(int64(numNames), 10))
	println("# of rows: " + strconv.FormatInt(int64(numRows), 10))

	for i, n := 0, colNames.Length(); i < n; i++  {
		println("Column " + strconv.FormatInt(int64(i), 10) + ": '" + colNames.Index(i).String() + "'")
	}

	row0 := rows.Index(0)
	row0Len := row0.Length()
	println("Length of row 0: " + strconv.FormatInt(int64(row0Len), 10))

	recordSet1 := rows.Index(1)
	recordSet1Len := recordSet1.Length()
	println("Length of row 1: " + strconv.FormatInt(int64(recordSet1Len), 10))

	for i, n := 0, row0Len; i < n; i++  {
		println("Row 0 value " + strconv.FormatInt(int64(i), 10) + ": '" + row0.Index(i).Get("Value").String() + "'")
	}
}
