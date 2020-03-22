package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	_, err := json.Marshal(&group)
	if err != nil {
		fmt.Println("error:", err)
	}
	// print(string(b))

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(group)
	if err != nil {
		panic(err)
	}

	// var buf2 *bytes.Buffer
	buf2 := bytes.NewBufferString("[" + buf.String() + "]")

	buf2.WriteTo(os.Stdout)
	// os.Stdout.Write(b)
}
