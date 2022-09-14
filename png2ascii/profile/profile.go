package profile

import (
	"encoding/json"
	"os"
)

var DEFAULT = Profile{
	Charset: "#@80GCLft1i;:,.",
	Font: Font{
		Typeface: "gomonobold",
		Size:     12,
		DPI:      72,
	},
}

type Profile struct {
	Charset string `json:"charset"`
	Font    Font   `json:"font"`
}

type Font struct {
	Typeface string  `json:"typeface"`
	Size     float64 `json:"size"`
	DPI      float64 `json:"dpi"`
}

func (p *Profile) Load(file string) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	profile := struct {
		Profile *Profile `json:"profile"`
	}{
		Profile: p,
	}

	return json.Unmarshal(bytes, &profile)
}
