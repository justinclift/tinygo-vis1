This is just simple research code, to determine if the current
state of TinyGo wasm can usefully present visualisation data for
DBHub.io purposes.

To compile the WebAssembly file:

    $ tinygo build -target wasm -no-debug -o docs/wasm.wasm wasm.go
