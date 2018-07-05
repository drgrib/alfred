package alfred

import (
	"encoding/json"
	. "fmt"
)

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

var Items = []Item{}
var Variables = map[string]string{}

type Output struct {
	Variables map[string]string `json:"variables,omitempty"`
	Items     []Item            `json:"items"`
}

func Add(item Item) {
	Items = append(Items, item)
}

func String() string {
	output := Output{
		Variables: Variables,
		Items:     Items,
	}
	b, err := json.Marshal(output)
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
