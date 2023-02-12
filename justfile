# for windows need to run this in a bash emulator - we used git for windows here.

build:
  cd wasm && GOOS=js GOARCH=wasm go build -o ../docs/static/main.wasm main.go
  cp "/c/Program Files/Go/misc/wasm/wasm_exec.js" docs/js
  
serve: build
  python -m http.server --directory docs
  
format:
  prettier --write docs/index.html