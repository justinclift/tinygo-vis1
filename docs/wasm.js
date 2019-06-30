'use strict';

const WASM_URL = 'wasm.wasm';

var wasm;
var palette = 0.6;

// Handle browser resize events
function resizeEvent() {
  wasm.exports.DrawBarChart(palette);
}

function init() {
  window.onresize = resizeEvent; // Redraw the bar chart when the window is resized

  // Load and run the wasm
  const go = new Go();
  if ('instantiateStreaming' in WebAssembly) {
    WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
      wasm = obj.instance;
      go.run(wasm); // Initial setup

      wasm.exports.DrawBarChart(palette);
    })
  } else {
    fetch(WASM_URL).then(resp =>
      resp.arrayBuffer()
    ).then(bytes =>
      WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
        wasm = obj.instance;
        go.run(wasm);

        wasm.exports.DrawBarChart(palette);
      })
    )
  }
}

init();
