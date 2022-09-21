package profile

import (
	"testing"
)

func TestFontSetTypeface(t *testing.T) {
	expected := Font{
		Typeface: "gomonoitalic",
		Size:     12,
		DPI:      72,
	}

	font := Font{
		Typeface: "gomonobold",
		Size:     12,
		DPI:      72,
	}

	if err := font.Set("gomonoitalic"); err != nil {
		t.Fatalf("%v", err)
	}

	if font.Typeface != expected.Typeface {
		t.Errorf("Incorrect typeface - expected:%v, got:%v", expected.Typeface, font.Typeface)
	}

	if font.Size != expected.Size {
		t.Errorf("Incorrect font size - expected:%v, got:%v", expected.Size, font.Typeface)
	}

	if font.DPI != expected.DPI {
		t.Errorf("Incorrect DPI - expected:%v, got:%v", expected.DPI, font.DPI)
	}
}

func TestFontSetTypefaceWithBlank(t *testing.T) {
	expected := Font{
		Typeface: "gomonobold",
		Size:     12,
		DPI:      72,
	}

	font := Font{
		Typeface: "gomonobold",
		Size:     12,
		DPI:      72,
	}

	if err := font.Set(":12:72"); err != nil {
		t.Fatalf("%v", err)
	}

	if font.Typeface != expected.Typeface {
		t.Errorf("Incorrect typeface - expected:%v, got:%v", expected.Typeface, font.Typeface)
	}

	if font.Size != expected.Size {
		t.Errorf("Incorrect font size - expected:%v, got:%v", expected.Size, font.Typeface)
	}

	if font.DPI != expected.DPI {
		t.Errorf("Incorrect DPI - expected:%v, got:%v", expected.DPI, font.DPI)
	}
}

func TestFontSetFontSize(t *testing.T) {
	expected := Font{
		Typeface: "gomonobold",
		Size:     17.5,
		DPI:      72,
	}

	font := Font{
		Typeface: "gomonobold",
		Size:     12,
		DPI:      72,
	}

	if err := font.Set(":17.5"); err != nil {
		t.Fatalf("%v", err)
	}

	if font.Typeface != expected.Typeface {
		t.Errorf("Incorrect typeface - expected:%v, got:%v", expected.Typeface, font.Typeface)
	}

	if font.Size != expected.Size {
		t.Errorf("Incorrect font size - expected:%v, got:%v", expected.Size, font.Typeface)
	}

	if font.DPI != expected.DPI {
		t.Errorf("Incorrect DPI - expected:%v, got:%v", expected.DPI, font.DPI)
	}
}

func TestFontSetFontSizeWithBlank(t *testing.T) {
	expected := Font{
		Typeface: "gomonobold",
		Size:     12,
		DPI:      72,
	}

	font := Font{
		Typeface: "gomonobold",
		Size:     12,
		DPI:      72,
	}

	if err := font.Set("gomonobold::72"); err != nil {
		t.Fatalf("%v", err)
	}

	if font.Typeface != expected.Typeface {
		t.Errorf("Incorrect typeface - expected:%v, got:%v", expected.Typeface, font.Typeface)
	}

	if font.Size != expected.Size {
		t.Errorf("Incorrect font size - expected:%v, got:%v", expected.Size, font.Typeface)
	}

	if font.DPI != expected.DPI {
		t.Errorf("Incorrect DPI - expected:%v, got:%v", expected.DPI, font.DPI)
	}
}

func TestFontSetDPI(t *testing.T) {
	expected := Font{
		Typeface: "gomonobold",
		Size:     12,
		DPI:      100,
	}

	font := Font{
		Typeface: "gomonobold",
		Size:     12,
		DPI:      72,
	}

	if err := font.Set("::100"); err != nil {
		t.Fatalf("%v", err)
	}

	if font.Typeface != expected.Typeface {
		t.Errorf("Incorrect typeface - expected:%v, got:%v", expected.Typeface, font.Typeface)
	}

	if font.Size != expected.Size {
		t.Errorf("Incorrect font size - expected:%v, got:%v", expected.Size, font.Typeface)
	}

	if font.DPI != expected.DPI {
		t.Errorf("Incorrect DPI - expected:%v, got:%v", expected.DPI, font.DPI)
	}
}
