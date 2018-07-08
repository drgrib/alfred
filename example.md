# Script Filter JSON Format Examples

All specifications from the Alfred [Script Filter JSON Format](https://www.alfredapp.com/help/workflows/inputs/script-filter/json/) page are shown here with the Go code to produce them with this package and the resulting JSON output.

The `alfred.Item` struct members are listed in the order they are defined in the specification so their rendered JSON output will be in a different order than some examples in the official document.

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

Addition of the later mentioned `valid`, `match`, and `quicklookurl` fields, as well as adjusting the `alfred.Indent` value of this package.

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
		Valid:        alfred.Bool(false),
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
            "valid": false,
            "text": {
                "copy": "https://www.alfredapp.com/ (text here to copy)",
                "largetype": "https://www.alfredapp.com/ (text here for large type)"
            }
        }
    ]
}
```

## `variables` Example
Addition of `variables` object at the script filter level.

``` go
package main

import (
	"github.com/drgrib/alfred"
)

func main() {
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
		Valid:        alfred.Bool(false),
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
            "valid": false
        }
    ]
}
```