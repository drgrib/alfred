# Easy Alfred Workflow Script Filters in Go

[![GoDoc][godoc-icon]][godoc-link]

This is a lean but comprehensive implementation of the Alfred [Script Filter JSON Format](https://www.alfredapp.com/help/workflows/inputs/script-filter/json/) to get Alfred workflows off the ground quickly with fast, simple script filters in Go that can be easily tested and debugged in `stdout` before importing them to Alfred. 

It uses standard, familiar Go syntax and conventions as much as possible for quick and easy integration with other Go code.

## Full Alfred JSON Support
[full-json.md](full-json.md) has examples of how to produce the full range of JSON output from Alfred's [Script Filter JSON Format](https://www.alfredapp.com/help/workflows/inputs/script-filter/json/).

## A Simple Example
Let's say we want to create a simple script filter that converts a given argument to title case, lowercase, or uppercase, for which Go conveniently has built-in support.

Let's start by prototyping our logic in Go with a `case.go` file in our workflow folder. We can do this on the command line or in the editor of our choice using pure Go:

``` go
package main

import (
	"strings"

	"github.com/drgrib/alfred"
)

func addCases(arg string) {
	title := strings.Title(arg)
	alfred.Add(alfred.Item{
		Title:    title,
		Subtitle: "Title",
		Arg:      title,
		UID:      "title",
	})

	lower := strings.ToLower(arg)
	alfred.Add(alfred.Item{
		Title:    lower,
		Subtitle: "Lower",
		Arg:      lower,
		UID:      "lower",
	})

	upper := strings.ToUpper(arg)
	alfred.Add(alfred.Item{
		Title:    upper,
		Subtitle: "Upper",
		Arg:      upper,
		UID:      "upper",
	})
}

func main() {
	arg := "just a test"
	addCases(arg)
	alfred.Run()
}
```
``` json
{
    "items": [
        {
            "uid": "title",
            "title": "Just A Test",
            "subtitle": "Title",
            "arg": "Just A Test"
        },
        {
            "uid": "lower",
            "title": "just a test",
            "subtitle": "Lower",
            "arg": "just a test"
        },
        {
            "uid": "upper",
            "title": "JUST A TEST",
            "subtitle": "Upper",
            "arg": "JUST A TEST"
        }
    ]
}
```

Looks good. Now let's add `os.Args` support and test it on the command line to simulate Alfred input:

``` go
package main

import (
	"os"
	"strings"

	"github.com/drgrib/alfred"
)

// [same stuff in the middle]

func main() {
	arg := os.Args[1]
	addCases(arg)
	alfred.Run()
}
```
``` bash
go build case.go
./case "another test"
```
``` json
{
    "items": [
        {
            "uid": "title",
            "title": "Another Test",
            "subtitle": "Title",
            "arg": "Another Test"
        },
        {
            "uid": "lower",
            "title": "another test",
            "subtitle": "Lower",
            "arg": "another test"
        },
        {
            "uid": "upper",
            "title": "ANOTHER TEST",
            "subtitle": "Upper",
            "arg": "ANOTHER TEST"
        }
    ]
}
```

Right again. Alright. Let's drop this into our script filter now:

<img src="./images/1-script-filter.png" alt="script-filter">

And give it a whirl:

<img src="./images/2-test.png" alt="test">

And why not copy these to the clipboard so we can actually use them?

<img src="./images/3-clipboard.png" alt="clipboard">

With a few simple runs and a glance at the Alfred clipboard history, we can see we are ready for business:

<img src="./images/4-history.png" alt="clipboard">

Well that was easy!

[godoc-icon]: https://godoc.org/github.com/drgrib/alfred?status.svg
[godoc-link]: https://godoc.org/github.com/drgrib/alfred