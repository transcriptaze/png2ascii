![build](https://github.com/twystd/png2ascii/workflows/build/badge.svg)

# png2ascii

Renders a PNG as ASCII characters to either a text file or a PNG.

Supported operating systems:
- Linux
- MacOS
- Windows

#### mp42ascii

The `mp42ascii.sh` script in the _scripts_ folder is a rough-around-the-edges bash script that:
1. Uses _ffmpeg_ to extract all the frames of an MP4 as PNG files
2. Rerenders the PNG files as 'ASCII' PNGs using _png2ascii_
3. Scales and crops the rerendered PNGs using _ImageMagick_ 
4. Finally  reassembles the PNGs into an MP4 again (uing _ffmpeg_).

**WARNING**: it takes **hours** (and a significant amount of disk space) to process an MP4.


## Releases

| *Version* | *Description*                                                                             |
| --------- | ----------------------------------------------------------------------------------------- |
| v0.1.0    | Initial release                                                                           |

## Installation

Executables for all the supported operating systems are packaged in the [releases](https://github.com/transcriptaze/pn2ascii/releases). Installation is straightforward - download the archive, extract it to a directory of your choice and move the executable for
your platform to a convenient location.

### Building from source

Required tools:
- [Go 1.19+](https://go.dev)
- make (optional but recommended)

```
git clone https://github.com/transcriptaze/png2ascii.git
cd png2ascii
make build
```

Without using `make`:
```
git clone https://github.com/transcriptaze/png2ascii.git
cd png2ascii
go build -trimpath -o bin/ ./...
```

The above commands build the `pn2ascii` executable to the `bin` directory.


#### Dependencies

| *Dependency*                        | *Description*                          |
| ----------------------------------- | ---------------------------------------|
| golang.org/x/image                  | Go extended image processing functions |


## png2asci

Usage: ```png2ascii <options> <PNG file>```

where:
```
<PNG file> is the PNG file to convert
```

Supported options:

```
  --out <file>         File to which to write the re-rendered image Defaults to stdout for text and png2ascii.png for PNG.
  --format <text|png>  Output format. Defaults to text. 
  --profile <file>     Conversion profile file (defaults to none i.e. uses the internal default settings).
  --bgcolor            Background colour (for PNG output) Defaults to white.
  --fgcolor            Foreground colour (for PNG output) Defaults to black.
  --font <fontspect>   Sets the font to use for rendering, formatted as <typeface|filepath>:size:DPI. Defaults to gomonobold:12:72.
  --squoosh            Prescales the image to preserve aspect ratio. Defaults to the 'squoosh' setting in the profile (if defined).
  --debug              Displays internal conversion information.
```

### Profiles

Profiles are JSON files that define the settings used for re-rendering an image. The default profile (embedded in the executable)
is:
```
{
    "profile": {
        "charset": "#@80GCLft1i;:,.",
        "font": {
            "size": 12,
            "dpi": 72,
            "typeface": "gomonobold"
        }
        "squoosh": {
            "enabled": true
        }
    }
}
```