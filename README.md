# pixgen

[![Go Version][go-shield]][go-url]
[![Build Status][build-shield]][build-url]
[![Go Report Card][report-shield]][report-url]
[![License][license-shield]][license-url]

**pixgen** - A simple command-line tool written in Go to generate pixel art PNG images from string-based definitions stored in a JSON file. It supports tiling multiple small images into a single grid output.

## ✨ Features

- Generate pixel art from arrays of strings (supports 8x8, 16x16, 32x32 pixels).
- Define colors using a simple character-to-color map.
- Supports transparent backgrounds and pixels.
- Automatically tiles multiple images for a single definition into a square grid.
- Configurable input JSON file path and output directory via CLI flags.
- Written in Go, produces a single executable binary.

## 💾 Installation

**(Option 1: Go Install - Recommended if Go environment is set up)**

```bash
go install [github.com/jtakakura/pixgen@latest](https://github.com/jtakakura/pixgen@latest)
```

**(Option 2: Download from Releases)**

Download the pre-compiled binary for your operating system from the Releases Page.

## 🚀 Usage

```bash
pixgen -input <path_to_json_file> [-output <output_directory>]
```

**Flags**:

- `input <path>`: (Required) Path to the input JSON file containing the image definitions.
- `output <dir>`: (Optional) Directory where the generated PNG files will be saved. Defaults to the current directory (`.`).

**Example**:

```bash
# Generate images from definitions in data.json and save them to ./output_images directory
pixgen -input data.json -output ./output_images
```

## 📄 Input JSON Format

The input JSON file must contain a single root object.

- **Keys**: Each key in the object is used as the base filename for the output PNG (e.g., `"player"` becomes `player.png`).
- **Values**: Each value must be an array of "image definitions". An image definition itself is an array of strings (`[][]string`).

**Image Definition (`[]string`)**

- An image definition must represent a square image.
- It must contain exactly N strings.
- Each string must contain exactly N characters.
- Valid values for N are 8, 16, or 32.
- All image definitions within the same array (for one output file) must have the same dimension N.

**Example** `data.json`:

```json
{
  "player": [
    [
      // Image Definition 1 (16x16)
      "................",
      "..GG......GG....",
      "..GG......GG....",
      "................",
      "....GGGGGG......",
      "....GGGGGG......",
      "................",
      "GG..........GG..",
      "GG..........GG..",
      "................",
      "..GGGGGGGGGG....",
      "..GGGGGGGGGG....",
      "................",
      "....GG....GG....",
      "....GG....GG....",
      "................"
    ],
    [
      // Image Definition 2 (16x16)
      "................",
      "..BB......BB....",
      "..BB......BB....",
      "................",
      "....BBBBBB......",
      "....BBBBBB......",
      "................",
      "BB..........BB..",
      "BB..........BB..",
      "................",
      "..BBBBBBBBBB....",
      "..BBBBBBBBBB....",
      "................",
      "....BB....BB....",
      "....BB....BB....",
      "................"
    ],
    [
      // Image Definition 3 (16x16)
      "................",
      "..RR......RR....",
      "..RR......RR....",
      "................",
      "....RRRRRR......",
      "....RRRRRR......",
      "................",
      "RR..........RR..",
      "RR..........RR..",
      "................",
      "..RRRRRRRRRR....",
      "..RRRRRRRRRR....",
      "................",
      "....RR....RR....",
      "....RR....RR....",
      "................"
    ]
  ],
  "icon": [
    [
      // Image Definition 1 (8x8)
      "........",
      ".rrrrrr.",
      ".rgggggr.",
      ".rbbbbbr.",
      ".ryyyyyr.",
      ".rppccpr.",
      ".rrrrrr.",
      "........"
    ]
  ]
}
```

Running `pixgen -input data.json -output ./out` would generate:

- `./out/player.png`: A 32x32 pixel image containing the 3 player definitions tiled in a 2x2 grid (top-left, top-right, bottom-left), with the bottom-right cell being transparent.
- `./out/icon.png`: An 8x8 pixel image based on the single icon definition.

🎨 Character Map

Each character in the definition strings maps to a specific RGBA color:

```
'.' : color.RGBA{R: 0, G: 0, B: 0, A: 0}      // transparent
'l' : color.RGBA{R: 0, G: 0, B: 0, A: 255}      // black (lowercase L)
'r' : color.RGBA{R: 255, G: 0, B: 0, A: 255}      // red
'g' : color.RGBA{R: 0, G: 255, B: 0, A: 255}      // green
'b' : color.RGBA{R: 0, G: 0, B: 255, A: 255}      // blue
'y' : color.RGBA{R: 255, G: 255, B: 0, A: 255}      // yellow
'p' : color.RGBA{R: 128, G: 0, B: 128, A: 255}      // purple
'c' : color.RGBA{R: 0, G: 255, B: 255, A: 255}      // cyan
'w' : color.RGBA{R: 255, G: 255, B: 255, A: 255}  // white
'L' : color.RGBA{R: 85, G: 85, B: 85, A: 255}       // light_black (dark gray)
'R' : color.RGBA{R: 255, G: 128, B: 128, A: 255}  // light_red (pink)
'G' : color.RGBA{R: 128, G: 255, B: 128, A: 255}  // light_green
'B' : color.RGBA{R: 128, G: 128, B: 255, A: 255}  // light_blue
'Y' : color.RGBA{R: 255, G: 255, B: 128, A: 255}  // light_yellow
'P' : color.RGBA{R: 255, G: 128, B: 255, A: 255}  // light_purple (magenta/pink)
'C' : color.RGBA{R: 128, G: 255, B: 255, A: 255}  // light_cyan
'W' : color.RGBA{R: 170, G: 170, B: 170, A: 255}  // light_white (light gray)
```

(Unrecognized characters will cause an error for that file.)

🛠️ Building from Source

1. Clone the repository:

```bash
git clone [https://github.com/jtakakura/pixgen.git](https://github.com/jtakakura/pixgen.git)
cd pixgen
```

2. Build the executable:

```bash
go build
```

This will create the `pixgen` (or `pixgen.exe` on Windows) executable in the current directory.

.
🤝 Contributing
Contributions are welcome! Please feel free to open an issue or submit a pull request.

📄 License
Unless otherwise noted, the pixgen source files are distributed under the MIT License found in the [LICENSE](./LICENSE) file.

[go-shield]: https://img.shields.io/badge/go-1.22%2B-blue.svg?style=flat-square
[go-url]: https://golang.org
[build-shield]: https://img.shields.io/github/actions/workflow/status/jtakakura/pixgen/go-ci.yml?branch=main&style=flat-square
[build-url]: https://github.com/jtakakura/pixgen/actions/workflows/go-ci.yml
[report-shield]: https://goreportcard.com/badge/github.com/jtakakura/pixgen?style=flat-square
[report-url]: https://goreportcard.com/report/github.com/jtakakura/pixgen
[license-shield]: https://img.shields.io/github/license/jtakakura/pixgen?style=flat-square
[license-url]: https://github.com/jtakakura/pixgen/blob/main/LICENSE
