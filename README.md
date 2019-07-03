This is just simple research code, to determine if the current
state of TinyGo wasm can usefully present visualisation data for
DBHub.io purposes.

To compile the WebAssembly file:

    $ tinygo build -target wasm -no-debug -o docs/wasm.wasm wasm.go

To strip the custom name section from the end (reducing file size
further):

    $ wasm2wat docs/wasm.wasm -o docs/wasm.wat
    $ wat2wasm docs/wasm.wat -o docs/wasm.wasm
    $ rm -f docs/wasm.wat
