package main

import (
	"fmt"
	"os"

	json "github.com/goccy/go-json"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	a := ColorGroup{}
	json.Unmarshal(b, &a)
	fmt.Println()
	fmt.Printf("a=%+v", a)
}
