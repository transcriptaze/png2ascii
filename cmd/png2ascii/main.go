package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/transcriptaze/png2ascii/png2ascii"
	"github.com/transcriptaze/png2ascii/png2ascii/profile"
)

const VERSION = "v0.0.0"

type options struct {
	out        string
	format     string
	profile    string
	background png2ascii.Colour
	foreground png2ascii.Colour
	font       string
	squoosh    string
	debug      bool
}

func main() {
	fmt.Printf("png2ascii %v\n", VERSION)

	options := options{
		out:        "",
		format:     "text",
		profile:    "",
		background: png2ascii.White,
		foreground: png2ascii.Black,
		font:       "",
		squoosh:    "",
		debug:      false,
	}

	flag.StringVar(&options.out, "out", options.out, "(optional) output file. Defaults to stdout for text and mp42asc.png for PNG")
	flag.StringVar(&options.format, "format", options.format, "Format (png or text). Defaults to text")
	flag.StringVar(&options.profile, "profile", options.profile, "Profile file (defaults to none)")
	flag.Var(&options.background, "bgcolor", "Background colour. Defaults to white")
	flag.Var(&options.foreground, "fgcolor", "Foreground colour. Defaults to black")
	flag.StringVar(&options.font, "font", options.font, "(optional) font file path. Defaults to gomonobold.")
	flag.StringVar(&options.squoosh, "squoosh", "", "(optional) Prescales the image to preserve aspect ratio. Defaults to the profile 'squoosh'")
	flag.BoolVar(&options.debug, "debug", options.debug, "Displays internal conversion information")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Printf("\n   ERROR: please supply file to convert\n\n")
		os.Exit(1)
	}

	if err := exec(args[0], options); err != nil {
		fmt.Printf("\n   ERROR: %v\n\n", err)
		os.Exit(1)
	}
}

func exec(in string, options options) error {
	var codec png2ascii.Codec
	var err error

	profile := profile.DEFAULT

	if options.profile != "" {
		if err = profile.Load(options.profile); err != nil {
			return err
		}
	}

	if options.squoosh != "" {
		if err := profile.Squoosh.Set(options.squoosh); err != nil {
			return err
		}
	}

	if options.debug {
		fmt.Printf("  PROFILE\n")
		fmt.Printf("  charset:       %v\n", profile.Charset)
		fmt.Printf("  font typeface: %v\n", profile.Font.Typeface)
		fmt.Printf("       size:     %v\n", profile.Font.Size)
		fmt.Printf("       DPI:      %v\n", profile.Font.DPI)
		fmt.Printf("  squoosh:       %v\n", profile.Squoosh)
		fmt.Println()
	}

	switch options.format {
	case "text":
		if codec, err = png2ascii.NewPng2Txt(profile); err != nil {
			return err
		}

	case "png":
		if codec, err = png2ascii.NewPng2Png(profile); err != nil {
			return err
		}

	default:
		return fmt.Errorf("invalid output format (%v) - expected text or png", options.format)
	}

	info, err := os.Stat(in)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return walk(codec, in, options.out, options.debug)
	}

	if mode := info.Mode(); !mode.IsRegular() {
		return fmt.Errorf("invalid file mode (%v)", mode)
	}

	return convert(codec, in, options.out, options.debug)
}

func walk(codec png2ascii.Codec, dir string, out string, debug bool) error {
	f := func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if info, err := d.Info(); err != nil {
			return err
		} else if !info.Mode().IsRegular() {
			return nil
		} else if strings.HasPrefix(filepath.Base(path), ".") {
			return nil
		}

		return convert(codec, filepath.Join(dir, path), filepath.Join(out, path), debug)
	}

	if info, err := os.Stat(out); err != nil && !os.IsNotExist(err) {
		return err
	} else if err == nil && !info.IsDir() {
		return fmt.Errorf("%v is not a directory", out)
	} else if err := os.MkdirAll(out, 0770); err != nil {
		return err
	}

	return fs.WalkDir(os.DirFS(dir), ".", f)
}

func convert(codec png2ascii.Codec, file string, out string, debug bool) error {
	if debug {
		fmt.Printf("  ... converting %v\n", file)
	}

	if img, err := png2image(file); err != nil {
		return err
	} else {
		return codec.Convert(img, out, debug)
	}
}

func png2image(file string) (image.Image, error) {
	r, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer r.Close()

	return png.Decode(r)
}
