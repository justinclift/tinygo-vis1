'use strict';

const WASM_URL = 'wasm.wasm';

var wasm;

// // Pass mouse wheel events through to its wasm handler
// function wheelHandler(evt) {
//     wasm.exports.wheelHandler(evt.deltaY);
// }


function init() {
  const go = new Go();
  if ('instantiateStreaming' in WebAssembly) {
    WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
      wasm = obj.instance;
      go.run(wasm);

      // Set up wasm event handlers
      // document.getElementById("mycanvas").addEventListener("mousedown", clickHandler);
      // document.getElementById("mycanvas").addEventListener("keydown", keyPressHandler);
      // document.getElementById("mycanvas").addEventListener("mousemove", moveHandler);
      // document.getElementById("mycanvas").addEventListener("wheel", wheelHandler);
      //
      // Set up basic render loop
      // setInterval(function() {
      //     applyTransformation();
      // },25);
    })
  } else {
    fetch(WASM_URL).then(resp =>
      resp.arrayBuffer()
    ).then(bytes =>
      WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
        wasm = obj.instance;
        go.run(wasm);

        // // Set up wasm event handlers
        // document.getElementById("mycanvas").addEventListener("mousedown", clickHandler);
        // document.getElementById("mycanvas").addEventListener("keydown", keyPressHandler);
        // document.getElementById("mycanvas").addEventListener("mousemove", moveHandler);
        // document.getElementById("mycanvas").addEventListener("wheel", wheelHandler);

        // // Set up basic render loop
        // setInterval(function() {
        //     applyTransformation();
        // },25);
      })
    )
  }
}

init();
