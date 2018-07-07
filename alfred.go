package alfred

import (
	"encoding/json"
	. "fmt"
)

// Indent specifies the indentation scheme used for the JSON in String() and Run()
var Indent = "    "

var Rerun float64
var Variables = map[string]string{}
var Items = []Item{}

type Icon struct {
	Type string `json:"type,omitempty"`
	Path string `json:"path,omitempty"`
}

type Item struct {
	UID          string `json:"uid,omitempty"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle,omitempty"`
	Arg          string `json:"arg,omitempty"`
	Icon         *Icon  `json:"icon,omitempty"`
	Autocomplete string `json:"autocomplete,omitempty"`
	Type         string `json:"type,omitempty"`
	Valid        bool   `json:"valid,omitempty"`
	Match        string `json:"match,omitempty"`
}

type output struct {
	Rerun     float64           `json:"rerun,omitempty"`
	Variables map[string]string `json:"variables,omitempty"`
	Items     []Item            `json:"items"`
}

func Add(item Item) {
	Items = append(Items, item)
}

// String returns the current JSON
func String() string {
	output := output{
		Rerun:     Rerun,
		Variables: Variables,
		Items:     Items,
	}
	var err error
	var b []byte
	if Indent == "" {
		b, err = json.Marshal(output)
	} else {
		b, err = json.MarshalIndent(output, "", Indent)
	}
	if err != nil {
		messageErr := Errorf("Error in parser. Please report this output to https://github.com/drgrib/alfred/issues: %v", err)
		panic(messageErr)
	}
	s := string(b)
	return s
}

// Run prints the result of String() to `stdout` for debugging or direct use by Alfred
func Run() {
	Println(String())
}
