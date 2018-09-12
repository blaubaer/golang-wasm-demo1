# Golang wasm Demo 1

Demonstrates Webassembly created using Golang.

## Important files

* [`main.go`](main.go) - Source file of the Webassembly Golang program.
* [`docs/index.html`](docs/index.html) - Entry HTML page to trigger load of the Webassembly file.
* [`docs/wasm_exex.js`](docs/wasm_exec.js) - Support for interaction between Golang Webassembly files and JavaScript.
* [`docs/app.wasm`](docs/app.wasm) - Compiled version of [`main.go`](main.go) loaded by [`index.html`](docs/index.html) using [`wasm_exex.js`](docs/wasm_exec.js).

## Requirements

1. Latest version of browser like Chrome, Firefox or Edge
2. Golang 1.11+

## Build

```bash
GOOS=js GOARCH=wasm go build -out docs/app.wasm ./...
```

## Try it out

* Either open [`docs/index.html`](docs/index.html) locally.
* ...or simply open [the demo page](https://blaubaer.github.io/golang-wasm-demo1/)

## License

See the [LICENSE](LICENSE) file.