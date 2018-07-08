# Easy Alfred Workflow Script Filters in Go

[![GoDoc][godoc-icon]][godoc-link]

This is a lean but [comprehensive](full-json.md) implementation of the Alfred Script Filter JSON Format to get Alfred workflows off the ground quickly with blazingly fast script filters in Go that can be seamlessly developed inside or outside Alfred. 

It uses standard, familiar Go syntax and conventions as much as possible for rapid use by Go developers and integration with other Go code.

## Full Alfred JSON Support
[full-json.md](full-json.md) has examples of how to produce the full range of JSON output for Alfred's Script Filter JSON Format.

## A Simple Example
Let's say we want to create a simple script filter that converts a given query to title case, lower case, or upper case, for which Go conveniently has built-in support.

Let's start by prototyping our logic in Go with a `case.go` file in our workflow folder. We can do this on the command line or in the editor of our choice and easily run it inside or outside of Alfred:

``` go
package main

import (
	"strings"

	"github.com/drgrib/alfred"
)

func addCases(arg string) {
	titlecase := strings.Title(arg)
	alfred.Add(alfred.Item{
		Title:    titlecase,
		Subtitle: "Title",
		Arg:      titlecase,
		UID:      "titlecase",
	})

	lowercase := strings.ToLower(arg)
	alfred.Add(alfred.Item{
		Title:    lowercase,
		Subtitle: "Lower",
		Arg:      lowercase,
		UID:      "lowercase",
	})

	uppercase := strings.ToUpper(arg)
	alfred.Add(alfred.Item{
		Title:    uppercase,
		Subtitle: "Upper",
		Arg:      uppercase,
		UID:      "uppercase",
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