# Script Filter JSON Format Examples

_All code examples from the Alfred [Script Filter JSON Format](https://www.alfredapp.com/help/workflows/inputs/script-filter/json/) page are populated here along with the text preceding them in the original document, the Go code to produce them with this package, and the JSON resulting output._

_The `alfred.Item` struct members are listed in the order they are defined in the specification and so their rendered JSON output will be in a different order than some examples listed in the original document._

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

_Addition of the later mentioned `valid`, `match`, and `quicklookurl` fields, as well as adjusting the `alfred.Indent` value of this package._

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
_Addition of the `mods` object with `mod` sub-objects._

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
_Addition of the `text` object._

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