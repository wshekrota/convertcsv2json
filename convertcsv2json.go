package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cvtc2j(list []string) string {

	first := true
	words := []string{}
	result := "["
	for r := range list {
		if first {
			words = strings.Split(list[r], ",")
			first = false
			continue
		}

		result += "{"
		each := strings.Split(list[r], ",")
		elen := len(each)
		for i := 0; i < elen; i++ {
			result += fmt.Sprintf("%s:%s", words[i], each[i])
			if i<elen-1 {
			result += ","
			}
		}
		result += "}"
		if r != len(list)-1 {
			result += ","
		} else {
			result += "]"
		}
	}
	return string(result)
}

func main() {

// prereq input for cvtc2j is []string from a file
// so this requires you to read it line by line not as byte

	var dat = []string{}
	file, err := os.Open("./test.csv")
	if err != nil {
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		dat = append(dat, scanner.Text())
	}

	file.Close()
	fmt.Println(cvtc2j(dat))
}
