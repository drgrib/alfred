# Full Script Filter JSON Format Examples

All specifications from the Alfred [Script Filter JSON Format](https://www.alfredapp.com/help/workflows/inputs/script-filter/json/) page are shown here along with the Go code to produce them and the resulting JSON output.

Wherever it makes sense, struct members are listed in the order they are defined in the specification, so their JSON will sometimes be in a different order than the official document but will still function in Alfred.

## Standard Example

``` go
package main

import (
	"github.com/drgrib/alfred"
)

func main() {
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
```
``` json
{
    "items": [
        {
            "uid": "desktop",
            "title": "Desktop",
            "subtitle": "~/Desktop",
            "arg": "~/Desktop",
            "icon": {
                "type": "fileicon",
                "path": "~/Desktop"
            },
            "autocomplete": "Desktop",
            "type": "file"
        }
    ]
}
```

## `valid`, `match`, and `quicklookurl` Example

Addition of the `valid`, `match`, and `quicklookurl` fields, as well as a custom `alfred.Indent` value.

``` go
package main

import (
	"github.com/drgrib/alfred"
)

func main() {
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
```
``` json
{
	"items": [
		{
			"uid": "desktop",
			"title": "Desktop",
			"subtitle": "~/Desktop",
			"arg": "~/Desktop",
			"icon": {
				"type": "fileicon",
				"path": "~/Desktop"
			},
			"autocomplete": "Desktop",
			"type": "file",
			"valid": false,
			"match": "my desktop",
			"quicklookurl": "https://www.alfredapp.com/"
		}
	]
}
```

## `mods` Example
Addition of the `mods` object with `mod` sub-objects.

``` go
package main

import (
	"github.com/drgrib/alfred"
)

func main() {
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
```
``` json
{
    "items": [
        {
            "uid": "desktop",
            "title": "Desktop",
            "subtitle": "~/Desktop",
            "arg": "~/Desktop",
            "icon": {
                "type": "fileicon",
                "path": "~/Desktop"
            },
            "autocomplete": "Desktop",
            "type": "file",
            "valid": false,
            "mods": {
                "alt": {
                    "valid": true,
                    "arg": "alfredapp.com/powerpack",
                    "subtitle": "https://www.alfredapp.com/powerpack/"
                },
                "cmd": {
                    "valid": true,
                    "arg": "alfredapp.com/powerpack/buy/",
                    "subtitle": "https://www.alfredapp.com/powerpack/buy/"
                }
            }
        }
    ]
}
```

## `text` Example
Addition of the `text` object.

``` go
package main

import (
	"github.com/drgrib/alfred"
)

func main() {
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
```
``` json
{
    "items": [
        {
            "uid": "desktop",
            "title": "Desktop",
            "subtitle": "~/Desktop",
            "arg": "~/Desktop",
            "icon": {
                "type": "fileicon",
                "path": "~/Desktop"
            },
            "autocomplete": "Desktop",
            "type": "file",
            "text": {
                "copy": "https://www.alfredapp.com/ (text here to copy)",
                "largetype": "https://www.alfredapp.com/ (text here for large type)"
            }
        }
    ]
}
```

## `rerun` and `variables` Example
Addition of `rerun` value and `variables` object at the script filter level.

``` go
package main

import (
	"github.com/drgrib/alfred"
)

func main() {
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
```
``` json
{
    "rerun": 0.5,
    "variables": {
        "fruit": "banana",
        "vegetable": "carrot"
    },
    "items": [
        {
            "uid": "desktop",
            "title": "Desktop",
            "subtitle": "~/Desktop",
            "arg": "~/Desktop",
            "icon": {
                "type": "fileicon",
                "path": "~/Desktop"
            },
            "autocomplete": "Desktop",
            "type": "file"
        }
    ]
}
```

## ALL the `variables` (and `mod` `icon`) Example
Alfred 3.4.1 adds `variables` object support to `item` and `mod` objects, as well as a custom `icon` for `mod` objects. Here is an example of all of it.

``` go
package main

import (
	"github.com/drgrib/alfred"
)

func main() {
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
```
``` json
{
    "variables": {
        "fruit": "banana",
        "vegetable": "carrot"
    },
    "items": [
        {
            "variables": {
                "fruit": "orange",
                "vegetable": "asparagus"
            },
            "uid": "desktop",
            "title": "Desktop",
            "subtitle": "~/Desktop",
            "arg": "~/Desktop",
            "icon": {
                "type": "fileicon",
                "path": "~/Desktop"
            },
            "mods": {
                "alt": {
                    "variables": {
                        "fruit": "apple",
                        "vegetable": "radish"
                    },
                    "arg": "alfredapp.com/powerpack",
                    "subtitle": "https://www.alfredapp.com/powerpack/",
                    "icon": {
                        "type": "powericon",
                        "path": "~/Desktop/power"
                    }
                }
            }
        }
    ]
}
```