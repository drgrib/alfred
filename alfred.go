package alfred

import (
	"encoding/json"
	. "fmt"
)

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
	Type         string `json:"type,omitempty"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle,omitempty"`
	Arg          string `json:"arg,omitempty"`
	Autocomplete string `json:"autocomplete,omitempty"`
	Icon         *Icon  `json:"icon,omitempty"`
}

type Output struct {
	Rerun     float64           `json:"rerun,omitempty"`
	Variables map[string]string `json:"variables,omitempty"`
	Items     []Item            `json:"items"`
}

func Add(item Item) {
	Items = append(Items, item)
}

func String() string {
	output := Output{
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

func Run() {
	Println(String())
}
