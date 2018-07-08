# Alfred Workflow Script Filters in Go

[![GoDoc][godoc-icon]][godoc-link]

This is a lean but comprehensive implementation of the Alfred [Script Filter JSON Format](https://www.alfredapp.com/help/workflows/inputs/script-filter/json/) to get Alfred workflows off the ground quickly with fast, simple script filters in Go that can be easily tested and debugged in `stdout` before importing them to Alfred. 

It uses standard, familiar Go syntax and conventions as much as possible for quick and easy integration with other Go code.

## Full Alfred JSON Support
An [example.md](example.md) is populated with examples of how to use this package to produce the complete JSON output specification on the official Alfred [Script Filter JSON Format](https://www.alfredapp.com/help/workflows/inputs/script-filter/json/) page.

[godoc-icon]: https://godoc.org/github.com/drgrib/alfred?status.svg
[godoc-link]: https://godoc.org/github.com/drgrib/alfred