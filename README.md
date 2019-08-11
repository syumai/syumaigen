# syumaigen

* A CLI tool to generate syumai's avatar image.
* The avatar used in this command was designed by @tanakaworld.

## Features

* Randomize color generation

## Installation

```console
go get -u github.com/syumai/syumaigen/cmd
```

## Usage

```sh
# Generate image
syumaigen > syumai.png

# Show usage
syumaigen -help > syumai.png

# Upscale (default: 10) and stop randomize color generation (default: true)
syumaigen -scale=100 -random=false > syumai.png
```

### Example HTTP Server

* `go run example/server/main.go`
* This generates random avator image.

## License

MIT

## Author

syumai
