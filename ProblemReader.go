package main

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func readProblemFile(path string)  {
	data, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(data), "\n")
	fmt.Println(strings.Split(lines[2], ": ")[1])
}