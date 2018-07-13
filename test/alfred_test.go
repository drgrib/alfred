package test

import (
	"testing"

	"github.com/drgrib/alfred"
)

func reset() {
	alfred.Rerun = 0
	alfred.Indent = "    "
	alfred.Items = []alfred.Item{}
}

func TestExample1(t *testing.T) {
	reset()

	alfred.Add(alfred.Item{
		UID:      "desktop",
		Title:    "Desktop",
		Subtitle: "~/Desktop",
		Arg:      "~/Desktop",
		Icon: &alfred.Icon{
			Type: "fileicon",
			Path: "~/Desktop",
		},
		Autocomplete: "Desktop",
		Type:         "file",
	})
	alfred.Run()
}

func TestExample2(t *testing.T) {
	reset()
	alfred.Indent = "\t"

	alfred.Add(alfred.Item{
		UID:      "desktop",
		Title:    "Desktop",
		Subtitle: "~/Desktop",
		Arg:      "~/Desktop",
		Icon: &alfred.Icon{
			Type: "fileicon",
			Path: "~/Desktop",
		},
		Autocomplete: "Desktop",
		Type:         "file",
		Valid:        alfred.Bool(false),
		Match:        "my desktop",
		QuicklookURL: "https://www.alfredapp.com/",
	})

	alfred.Run()
}

func TestExample3(t *testing.T) {
	reset()

	alfred.Add(alfred.Item{
		UID:      "desktop",
		Title:    "Desktop",
		Subtitle: "~/Desktop",
		Arg:      "~/Desktop",
		Icon: &alfred.Icon{
			Type: "fileicon",
			Path: "~/Desktop",
		},
		Autocomplete: "Desktop",
		Type:         "file",
		Valid:        alfred.Bool(false),
		Mods: map[string]alfred.Mod{
			"alt": {
				Valid:    alfred.Bool(true),
				Arg:      "alfredapp.com/powerpack",
				Subtitle: "https://www.alfredapp.com/powerpack/",
			},
			"cmd": {
				Valid:    alfred.Bool(true),
				Arg:      "alfredapp.com/powerpack/buy/",
				Subtitle: "https://www.alfredapp.com/powerpack/buy/",
			},
		},
	})

	alfred.Run()
}

func TestExample4(t *testing.T) {
	reset()

	alfred.Add(alfred.Item{
		UID:      "desktop",
		Title:    "Desktop",
		Subtitle: "~/Desktop",
		Arg:      "~/Desktop",
		Icon: &alfred.Icon{
			Type: "fileicon",
			Path: "~/Desktop",
		},
		Autocomplete: "Desktop",
		Type:         "file",
		Text: &alfred.Text{
			Copy:      "https://www.alfredapp.com/ (text here to copy)",
			Largetype: "https://www.alfredapp.com/ (text here for large type)",
		},
	})

	alfred.Run()
}

func TestExample5(t *testing.T) {
	reset()

	alfred.Rerun = 0.5

	alfred.Variables["fruit"] = "banana"
	alfred.Variables["vegetable"] = "carrot"

	alfred.Add(alfred.Item{
		UID:      "desktop",
		Title:    "Desktop",
		Subtitle: "~/Desktop",
		Arg:      "~/Desktop",
		Icon: &alfred.Icon{
			Type: "fileicon",
			Path: "~/Desktop",
		},
		Autocomplete: "Desktop",
		Type:         "file",
	})

	alfred.Run()
}

func TestExample6(t *testing.T) {
	reset()

	alfred.Variables = map[string]string{
		"fruit":     "banana",
		"vegetable": "carrot",
	}

	alfred.Add(alfred.Item{
		Variables: map[string]string{
			"fruit":     "orange",
			"vegetable": "asparagus",
		},
		UID:      "desktop",
		Title:    "Desktop",
		Subtitle: "~/Desktop",
		Arg:      "~/Desktop",
		Icon: &alfred.Icon{
			Type: "fileicon",
			Path: "~/Desktop",
		},
		Mods: map[string]alfred.Mod{
			"alt": {
				Variables: map[string]string{
					"fruit":     "apple",
					"vegetable": "radish",
				},
				Arg:      "alfredapp.com/powerpack",
				Subtitle: "https://www.alfredapp.com/powerpack/",
				Icon: &alfred.Icon{
					Type: "powericon",
					Path: "~/Desktop/power",
				},
			},
		},
	})

	alfred.Run()
}
