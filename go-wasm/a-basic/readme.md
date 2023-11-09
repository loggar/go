# A-basic wasm

## go mod init

```
cd ./go-wasm/a-basic
go mod init github.com/loggar/go-wasm/a-basic
```

## Build

```
cd ./go-wasm/a-basic/cmd/wasm
GOOS=js GOARCH=wasm go build -o ./assets/json.wasm ./cmd/wasm
```

## Js Glue

```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./assets/
```

## Run server

```
go run ./cmd/server/main.go
```

## Test

Browser console:

```
formatJSON('{"website":"golangbot.com", "tutorials": {"string":"/strings/"}}');
```

Result:

```
Go Web Assembly

formatJSON('{"website":"golangbot.com", "tutorials": {"string":"/strings/"}}');
wasm_exec.js:22 input {"website":"golangbot.com", "tutorials": {"string":"/strings/"}}
'{\n  "tutorials": {\n    "string": "/strings/"\n  },\n  "website": "golangbot.com"\n}'

formatJSON();
'Invalid no of arguments passed'
```
