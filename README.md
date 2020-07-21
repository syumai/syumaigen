# syumaigen

![syumai.gif](assets/syumai.gif)

* A CLI tool to generate syumai's avatar image.
* The avatar used in this command was designed by [@tanakaworld](https://github.com/tanakaworld).

## Features

* Change scale of avatar image
* Randomize color generation
* Generate PNG / SVG
* Generate animated GIF

## Installation

```console
go get -u github.com/syumai/syumaigen/cmd/syumaigen
```

## Usage

```sh
# Generate image (default: PNG)
syumaigen > syumai.png

# Show usage
syumaigen -help

# Upscale (default: 10) and stop randomize color generation (default: true)
syumaigen -scale=100 -random=false > syumai.png

# generate animated GIF
syumaigen -animated > syumai.gif

# generate SVG
syumaigen -svg > syumai.svg
```

### Example HTTP Server

* `go run example/server/main.go`
* This generates random avator image.

## License

MIT

## Author

syumai
