# Script Filter JSON Format Examples

_All code examples from the Alfred [Script Filter JSON Format](https://www.alfredapp.com/help/workflows/inputs/script-filter/json/) page are populated here along with the text preceding them in the original document, the Go code to produce them with this package, and the JSON resulting output._

_The `alfred.Item` struct members are listed in the order they are defined in the specification and so their rendered JSON will be in a different order than some examples listed in the original document._

Example JSON Format:

``` go
package main

import (
	"github.com/drgrib/alfred"
)

func main() {
	alfred.Add(alfred.Item{
		UID:          "desktop",
		Type:         "file",
		Title:        "Desktop",
		Subtitle:     "~/Desktop",
		Arg:          "~/Desktop",
		Autocomplete: "Desktop",
		Icon: &alfred.Icon{
			Type: "fileicon",
			Path: "~/Desktop",
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
            "type": "file"
        }
    ]
}
```

_Addition of the later mentioned `valid` and `match` fields, as well as adjusting the `alfred.Indent` feature of this package._

